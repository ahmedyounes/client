package ekfile

import "github.com/keybase/client/go/libkb"

type Store interface {
	Retrieve(mctx libkb.MetaContext)
	Erase(mctx libkb.MetaContext)
}

type EKStore struct {
	directory string
	prefix    string
}

func NewEKStore(directory, prefix string, value []byte) (*EKStore, error) {
	return &EKStore{
		directory: directory,
		prefix:    prefix,
	}, nil
}
