package oppbloom

import (
        "bytes"
)

type OppBloom struct {
        size    uint64
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
        if size <= 0 {
                panic("oppbloom: size must be greater than 0")
        }

        oppbloom := &OppBloom{
                size:   uint64(size),
                array:  make([][]byte, size),
                hasher: &FNVHasher{},
        }

        for _, option := range options {
                option(oppbloom)
        }

        return oppbloom
}

func (this *OppBloom) Add(key []byte) {
        index := this.hasher.Sum64(key) % this.size

        this.array[index] = append([]byte(nil), key...)
}

func (this *OppBloom) Contain(key []byte) bool {
        index := this.hasher.Sum64(key) % this.size

        if k := this.array[index]; k != nil && bytes.Equal(k, key) {
                return true
        }

        return false
}

func (this *OppBloom) Cap() int {
        return int(this.size)
}
