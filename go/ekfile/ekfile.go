package ekfile

import (
	cryptorand "crypto/rand"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/keybase/client/go/libkb"
	"golang.org/x/crypto/hkdf"
	"golang.org/x/crypto/nacl/secretbox"
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
	return &EKStore{
		directory:            directory,
		prefix:               prefix,
		version:              LatestVersion,
		noiseFileLengthBytes: NoiseFileLengthBytes,
	}, nil
}

func (s *EKStore) Store(mctx libkb.MetaContext, value []byte, additionalKeyMaterial []byte) error {
	err := os.MkdirAll(directory, 0700)
	if err != nil {
		return nil, err
	}
	noiseTemp, err := ioutil.TempFile(s.directory, "")
	if err != nil {
		return err
	}
	defer libkb.ShredFile(noiseTemp.Name())

	err = os.Chmod(noiseTemp, 0600)
	if err != nil {
		return err
	}

	storeTemp, err := ioutil.TempFile(s.directory, "")
	if err != nil {
		return err
	}
	defer libkb.ShredFile(storeTemp.Name())

	err = os.Chmod(storeTemp, 0600)
	if err != nil {
		return err
	}

	noise, err := s.noiseGen()
	if err != nil {
		return err
	}

	key, err := s.keyGen(noise, additionalKeyMaterial)
	if err != nil {
		return err
	}

	box, nonce, err := s.seal(key, value)
	if err != nil {
		return err
	}

	_, err = noiseTemp.Write(noise)
	if err != nil {
		return err
	}

	_, err = storeTemp.Write(encryptedValue)
	if err != nil {
		return err
	}

	err = os.Rename(noiseTemp.Name(), s.noiseFilename())
	if err != nil {
		return err
	}

	err = os.Rename(storeTemp.Name(), s.storeFilename())
	if err != nil {
		return err
	}

	return nil
}

func (s *EKStore) noiseGen() ([]byte, error) {
	noise := make([]byte, s.noiseFileLengthBytes)
	_, err = cryptorand.Read(noise)
	if err != nil {
		return nil, err
	}
	return noise, nil
}

func (s *EKStore) keyGen(noise []byte, additionalKeyMaterial []byte) ([]byte, error) {
	keyMaterial := append(noise, additionalKeyMaterial...)
	key := make([]byte, 32)
	_, err = hkdf.New(sha256.New, keyMaterial, nil, s.combinedKeyDerivationContext()).Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func (s *EKStore) seal(key []byte, value []byte) ([]byte, error) {
	nonce := make([]byte, 24)
	_, err := cryptorand.Read(nonce)
	if err != nil {
		return nil, err
	}
	var nonceArray [24]byte
	copy(nonceArray[:], nonce)
	var keyArray [32]byte
	copy(keyArray[:], key)
	var out []byte
	boxed := secretbox.Seal(out, value, &nonceArray, &keyArray)
	return boxed, nil
}

func (s *EKStore) derivationContext() []byte {
	return []byte(fmt.Sprintf("Keybase-Derived-LKS-NaCl-SecretBox-%s", s.version))
}

func (s *EKStore) noiseFilename() string {
	return filepath.Join(s.directory, fmt.Sprintf("%s.noise%s", s.prefix, s.version))
}

func (s *EKStore) storeFilename() string {
	return filepath.Join(s.directory, fmt.Sprintf("%s.store%s", s.prefix, s.version))
}
