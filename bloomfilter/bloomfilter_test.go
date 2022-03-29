package bloomfilter

import (
	"log"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := NewBloomFilter()
	bf.Add("test13")
	if bf.Contain("test13"){
		log.Println("exist")
	}else {
		log.Println("no exist")
	}
}