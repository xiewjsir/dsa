/*
https://blog.csdn.net/weixin_29113371/article/details/112432406
*/
package bloomfilter

import (
	"github.com/bits-and-blooms/bitset"
)

const DEFAULT_SIZE = 1 << 26

var seeds = []uint{7, 11, 13, 31, 37}

type BloomFilter struct {
	set *bitset.BitSet
	fns [5]SimpleHash
}

type SimpleHash struct {
	cap  uint
	seed uint
}

func (s *SimpleHash) hash(value string) uint {
	var result uint = 0
	for i := 0; i < len(value); i++ {
		result = result*s.seed + uint(value[i])
	}

	return (s.cap - 1) & result
}

func NewBloomFilter() *BloomFilter {
	bf := new(BloomFilter)
	for i := 0; i < len(bf.fns); i++ {
		bf.fns[i] = SimpleHash{DEFAULT_SIZE, seeds[i]}
	}
	bf.set = bitset.New(DEFAULT_SIZE)
	return bf
}

func (bf *BloomFilter) Add(value string) {
	for _, fn := range bf.fns {
		bf.set.Set(fn.hash(value))
	}
}

func (bf *BloomFilter) Contain(value string) bool {
	if value == "" {
		return false
	}
	
	res := true
	for _, fn := range bf.fns {
		res = res && bf.set.Test(fn.hash(value))
	}
	
	return res
}
