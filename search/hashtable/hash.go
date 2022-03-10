package hashtable

import (
	"fmt"
	"log"
	"math"
	"sync"

	"github.com/OneOfOne/xxhash"
)

const expandFactor = 0.65 //扩容因子

type keyPairs struct {
	key   string
	value interface{}
	next  *keyPairs
}

type HashMap struct {
	array        []*keyPairs
	capacity     int //容量
	len          int //长度
	capacityMask int // 掩码，等于 capacity-1
	sync.Mutex
}

func NewHashMap(capacity int) *HashMap {
	// 默认大小为 16
	defaultCapacity := 1 << 4
	if capacity < defaultCapacity {
		capacity = defaultCapacity
	} else {
		// 否则，实际大小为>= capacity 的第一个 2^k
		capacity = 1 << int(math.Ceil(math.Log2(float64(capacity))))
	}

	m := new(HashMap)
	m.array = make([]*keyPairs, capacity, capacity)
	m.capacity = capacity
	m.capacityMask = capacity - 1

	return m
}

var hashAlgorithm = func(key []byte) uint64 {
	h := xxhash.New64()
	_, err := h.Write(key)
	if err != nil {
		log.Println(err)
	}

	return h.Sum64()
}

func (m *HashMap) Len() int {
	return m.len
}

// HashIndex 对键进行哈希求值，并计算下标
func (m *HashMap) HashIndex(key string, mask int) int {
	hash := hashAlgorithm([]byte(key))
	//求下标
	index := hash & uint64(mask)
	return int(index)
}

func (m *HashMap) Put(key string, value interface{}) {
	m.Lock()
	defer m.Unlock()

	index := m.HashIndex(key, m.capacityMask)

	element := m.array[index]
	if element == nil {
		m.array[index] = &keyPairs{
			key:   key,
			value: value,
		}
	} else {
		var lastPairs *keyPairs

		for element != nil {
			if element.key == key {
				element.value = value
				return
			}

			lastPairs = element
			element = element.next
		}

		lastPairs.next = &keyPairs{
			key:   key,
			value: value,
		}
	}

	newLen := m.len + 1
	if float64(newLen)/float64(m.capacity) > expandFactor {
		newM := new(HashMap)
		newM.capacity = 2 * m.capacity
		newM.capacityMask = newM.capacity - 1
		newM.array = make([]*keyPairs, newM.capacity, newM.capacity)
		for _, pairs := range m.array {
			for pairs != nil {
				newM.Put(pairs.key, pairs.value)
				pairs = pairs.next
			}
		}

		m.array = newM.array
		m.capacity = newM.capacity
		m.capacityMask = newM.capacityMask
	}
	m.len = newLen
}

func (m *HashMap) Get(key string) (value interface{}, ok bool) {
	m.Lock()
	defer m.Unlock()

	index := m.HashIndex(key, m.capacityMask)

	element := m.array[index]
	if element == nil {
		return
	}

	for element != nil {
		if element.key == key {
			return element.value, true
		}

		element = element.next
	}

	return
}

func (m *HashMap) Delete(key string) {
	m.Lock()
	defer m.Unlock()

	index := m.HashIndex(key, m.capacityMask)

	element := m.array[index]

	if element == nil {
		return
	}

	if element.key == key {
		m.array[index] = element.next
		m.len--
	} else {
		nextElement := element.next
		for nextElement != nil {
			if nextElement.key == key {
				element.next = nextElement.next
				m.len--
				return
			}

			element = nextElement
			nextElement = nextElement.next
		}
	}

}

func (m *HashMap) Range() {
	m.Lock()
	defer m.Unlock()

	for _, pairs := range m.array {
		for pairs != nil {
			fmt.Println(pairs.key, pairs.value)
			pairs = pairs.next
		}
	}

	fmt.Println()
}
