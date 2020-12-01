package oppbloom

import (
        "bytes"
)

type OppBloom struct {
        size    int
        array   [][]byte
        hasher  Hasher
}

type Option func(*OppBloom)

func WithHasher(hasher Hasher) Option {
        return func(o *OppBloom) {
                o.hasher = hasher
        }
}

func New(size int, options ...Option) *OppBloom {
        oppbloom := &OppBloom{
                size:   size,
                array:  make([][]byte, size),
                hasher: FNVHasher{},
        }

        for _, option := range options {
                option(oppbloom)
        }

        return oppbloom
}

func (this *OppBloom) Add(key []byte) {
        index := this.hasher.Sum64(key) % uint64(this.size)

        this.array[index] = key
}

func (this *OppBloom) Contain(key []byte) bool {
        index := this.hasher.Sum64(key) % uint64(this.size)

        if k := this.array[index]; k != nil && bytes.Equal(k, key) {
                return true
        }

        return false
}

func (this *OppBloom) Cap() int {
        return this.size
}
