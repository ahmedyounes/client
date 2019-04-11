package ekfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/keybase/client/go/libkb"
)

type Store interface {
	Store(mctx libkb.MetaContext, value []byte, additionalKeyMaterial []byte)
	Retrieve(mctx libkb.MetaContext, additionalKeyMaterial []byte)
	Erase(mctx libkb.MetaContext)
}

type EKStore struct {
	directory            string
	prefix               string
	version              Version
	noiseFileLengthBytes int
}

type Version int

const LatestVersion = 1
const NoiseFileLengthBytes = 1024 * 1024 * 2

func NewEKStore(mctx libkb.MetaContext, directory string, prefix string) (*EKStore, error) {
	err := os.MkdirAll(directory, 0700)
	if err != nil {
		return nil, err
	}
	return &EKStore{
		directory:            directory,
		prefix:               prefix,
		version:              LatestVersion,
		noiseFileLengthBytes: NoiseFileLengthBytes,
	}, nil
}

func (s *EKStore) Store(mctx libkb.MetaContext, value []byte, additionalKeyMaterial []byte) error {
	noiseHandle, err := ioutil.TempFile(s.directory, "")
	if err != nil {
		return err
	}

	storeHandle, err := ioutil.TempFile(s.directory, "")
	if err != nil {
		return err
	}

	handle, err := os.Open(s.noiseFilename())
	if err != nil {
		return err
	}
	defer handle.Close()
	return nil
}

// func (s *EKStore) keygen() ([]byte noise, []byte key) {
// 	// noise := [s.NoiseFileLengthBytes]byte
// 	// cryptorand.Read(noise)
// }

func (s *EKStore) noiseFilename() string {
	return filepath.Join(s.directory, fmt.Sprintf("%s.noise%s", s.prefix, s.version))
}

func (s *EKStore) storeFilename() string {
	return filepath.Join(s.directory, fmt.Sprintf("%s.store%s", s.prefix, s.version))
}

func (s *EKStore) computeKey(storeKey, additionalKeyMaterial []byte) ([]byte, error) {
	return nil, nil
}
