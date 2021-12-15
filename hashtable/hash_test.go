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
