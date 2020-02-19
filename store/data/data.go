package data

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/dimitarvdimitrov/sporkfs/log"
	"github.com/dimitarvdimitrov/sporkfs/store"
	"github.com/minio/highwayhash"
	"go.uber.org/zap"
)

var (
	hashKey, _ = hex.DecodeString("474c383279736a66674e48325037694c524e7a3830746e714636766f71675553")
)

type localDriver struct {
	storageRoot string
	indexM      *sync.RWMutex
	index       index // the values in the index are the relative locations to the storageRoot
}

func NewLocalDriver(location string) (*localDriver, error) {
	if err := os.MkdirAll(location, os.ModeDir|0755); err != nil {
		return nil, err
	}

	return &localDriver{
		storageRoot: location + "/",
		index:       restoreIndex(location),
		indexM:      &sync.RWMutex{},
	}, nil
}

func hashPath(path string) uint64 {
	file, err := os.Open(path)
	if err != nil {
		// retry
		file, err = os.Open(path)
		if err != nil {
			log.Error("couldn't open file to hash", zap.Error(err))
			return 0
		}
	}
	defer file.Close()
	return hashHandle(file)
}

func hashHandle(file *os.File) uint64 {
	hash, err := highwayhash.New64(hashKey)
	if err != nil {
		log.Error("couldn't start hashing file", zap.Error(err))
		return 0
	}

	_, err = io.Copy(hash, file)
	if err != nil {
		log.Error("couldn't hash file", zap.Error(err))
	}
	return hash.Sum64()
}

func (d *localDriver) Contains(id, hash uint64) bool {
	d.indexM.RLock()
	defer d.indexM.RUnlock()

	_, exists := d.index[id][hash]
	return exists
}

func (d *localDriver) ContainsAny(id uint64) bool {
	d.indexM.RLock()
	defer d.indexM.RUnlock()

	return len(d.index[id]) > 0
}

func (d *localDriver) Remove(id, hash uint64) {
	if !d.Contains(id, hash) {
		return
	}

	d.indexM.Lock()
	location := d.index[id][hash]
	delete(d.index[id], hash)
	if len(d.index[id]) == 0 {
		delete(d.index, id)
	}
	d.indexM.Unlock()

	if location != "" {
		go removeFromDisk(d.storageRoot + location)
	}
}

func removeFromDisk(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Error("couldn't remove file", zap.Error(err), zap.String("path", path))
	} else {
		log.Debug("removed file", zap.String("path", path))
	}
}

func (d *localDriver) Reader(id, hash uint64, flags int) (Reader, error) {
	d.indexM.RLock()
	location, exists := d.index[id][hash]
	d.indexM.RUnlock()
	if !exists {
		return nil, store.ErrNoSuchFile
	}

	f, err := os.OpenFile(d.storageRoot+location, flags|os.O_CREATE, store.ModeRegularFile)
	if err != nil {
		return nil, fmt.Errorf("file id=%d was in index but not on disk: %w", id, err)
	}

	return d.newSegReader(f), nil
}

func (d *localDriver) newSegReader(f *os.File) *segmentedReader {
	return &segmentedReader{
		f: f,
		onClose: func() {
			_ = f.Close()
		},
	}
}

func (d *localDriver) Open(id, hash uint64, flags int) (Reader, Writer, error) {
	if flags&(os.O_TRUNC|os.O_APPEND) == 0 {
		flags |= os.O_TRUNC
	}

	if flags&os.O_WRONLY != 0 {
		flags ^= os.O_WRONLY
	}

	if flags&os.O_CREATE != 0 {
		flags ^= os.O_CREATE
	}

	flags |= os.O_RDWR

	file, newLocation, err := d.handleForWriting(id, hash, flags)
	if err != nil {
		return nil, nil, err
	}

	writer := d.newSegWriter(id, hash, file, newLocation)
	reader := d.newSegReader(file)

	return reader, writer, nil
}

// handleForWriting duplicates the file with the given id and hash and returns an open
// file handle to the new duplicate with the provided flags. If the flags contains os.O_TRUNC
// or the hash is 0, there will be no copying - just a new empty file will be created.
// It also returns the location of the file relative to the storage root.
func (d *localDriver) handleForWriting(id, hash uint64, flags int) (*os.File, string, error) {
	if hash == 0 || flags&os.O_TRUNC != 0 {
		newLocation := generateStorageLocation(id, hash)
		newFilePath := d.storageRoot + newLocation
		f, err := os.OpenFile(newFilePath, flags|os.O_CREATE, store.ModeRegularFile)
		return f, newLocation, err
	}

	d.indexM.RLock()
	oldPath, exists := d.index[id][hash]
	d.indexM.RUnlock()
	if !exists {
		return nil, "", store.ErrNoSuchFile
	}
	oldFilePath := d.storageRoot + oldPath

	newLocation := generateStorageLocation(id, hash)
	newFilePath := d.storageRoot + newLocation
	err := duplicateFile(oldFilePath, newFilePath)
	if err != nil {
		log.Error("error while duplicating file",
			log.Id(id), log.Hash(hash),
			zap.String("old_path", oldFilePath),
			zap.String("new_path", newFilePath),
			zap.Error(err),
		)
		return nil, "", err
	}
	f, err := os.OpenFile(newFilePath, flags, store.ModeRegularFile)
	return f, newLocation, err
}

func (d *localDriver) Writer(id, hash uint64, flags int) (Writer, error) {
	if flags&(os.O_TRUNC|os.O_APPEND) == 0 {
		flags |= os.O_TRUNC
	}

	if flags&os.O_CREATE != 0 {
		flags ^= os.O_CREATE
	}

	flags |= os.O_WRONLY

	file, newLocation, err := d.handleForWriting(id, hash, flags)
	if err != nil {
		return nil, err
	}

	segWriter := d.newSegWriter(id, hash, file, newLocation)

	return segWriter, nil
}

func (d *localDriver) newSegWriter(id, oldHash uint64, file *os.File, relativeLocation string) *segmentedWriter {
	var newHash uint64
	absolutePath := file.Name()

	onClose := func() {
		_ = file.Close()

		newHash = hashPath(absolutePath)
		if newHash == oldHash {
			log.Debug("pruning file with repeated hash",
				log.Id(id),
				log.Hash(oldHash),
				zap.String("path", absolutePath),
			)
			go removeFromDisk(absolutePath)
		} else {
			d.indexM.Lock()
			defer d.indexM.Unlock()
			if d.index[id] == nil {
				d.index[id] = make(map[uint64]string)
			}
			d.index[id][newHash] = relativeLocation
		}
	}

	getHash := func() uint64 {
		return newHash
	}

	return &segmentedWriter{
		f:       file,
		sync:    syncer(file),
		onClose: onClose,
		hash:    getHash,
	}
}

func syncer(f *os.File) func() {
	return func() {
		_ = f.Sync()
	}
}

func duplicateFile(oldAbsolute, newAbsolute string) error {
	source, err := os.Create(oldAbsolute)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(newAbsolute)
	if err != nil {
		return err
	}
	defer destination.Close()
	// we sacrifice some memory to make it faster
	_, err = io.CopyBuffer(destination, source, make([]byte, 1024*1024))
	return err
}

func (d *localDriver) Size(id, hash uint64) int64 {
	d.indexM.RLock()
	f, err := os.Open(d.storageRoot + d.index[id][hash])
	d.indexM.RUnlock()
	if err != nil {
		return 0
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return 0
	}

	return info.Size()
}
