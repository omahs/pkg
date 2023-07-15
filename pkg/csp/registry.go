package csp

import (
	"context"
	"crypto/md5"

	api "github.com/wetware/ww/api/process"
)

// HashSize is the size of the hash produced by HashFunc.
const HashSize = md5.Size

// HashFunc is the function used for hashing in the default
// executor implementation.
var HashFunc func([]byte) [HashSize]byte = md5.Sum

type Registry api.BytecodeRegistry

func (r Registry) Put(ctx context.Context, bc []byte) ([]byte, error) {
	f, release := api.BytecodeRegistry(r).Put(ctx, func(br api.BytecodeRegistry_put_Params) error {
		return br.SetBytecode(bc)
	})
	defer release()

	<-f.Done()
	res, err := f.Struct()
	if err != nil {
		return nil, err
	}
	return res.Hash()
}

func (r Registry) Get(ctx context.Context, hash []byte) ([]byte, error) {
	f, release := api.BytecodeRegistry(r).Get(ctx, func(br api.BytecodeRegistry_get_Params) error {
		return br.SetHash(hash)
	})
	defer release()

	<-f.Done()
	res, err := f.Struct()
	if err != nil {
		return nil, err
	}
	return res.Bytecode()
}

func (r Registry) Has(ctx context.Context, hash []byte) (bool, error) {
	f, release := api.BytecodeRegistry(r).Get(ctx, func(br api.BytecodeRegistry_get_Params) error {
		return br.SetHash(hash)
	})
	defer release()

	<-f.Done()
	res, err := f.Struct()
	if err != nil {
		return false, err
	}
	return res.HasBytecode(), nil
}
