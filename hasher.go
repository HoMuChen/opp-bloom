package oppbloom

import (
        "hash/fnv"
)

type Hasher interface {
        Sum64(key []byte)       uint64
}

type FNVHasher struct {}

func (this *FNVHasher) Sum64(key []byte) uint64 {
        hash := fnv.New64a()
        hash.Write(key)
        return hash.Sum64()
}
