package ekfile

import "fmt"

type Store interface {
	Retrieve(mctx libkb.MetaContext)
	Erase(mctx libkb.MetaContext)
}

type EKStore struct {
	directory string
}
