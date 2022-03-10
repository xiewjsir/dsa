package hashtable

import (
	"fmt"
	"testing"
)

func TestXxHash(t *testing.T) {
	keys := []string{"hi", "my", "friend", "I", "love", "you", "my", "apple"}
	for _, key := range keys {
		fmt.Printf("xxhash('%s')=%d\n", key, hashAlgorithm([]byte(key)))
	}
}

func TestHashMap(t *testing.T) {
	hashMap := NewHashMap(16)
	for i := 0; i < 35; i++ {
		hashMap.Put(fmt.Sprintf("%d", i), fmt.Sprintf("v%d", i))
	}

	hashMap.Range()
}
