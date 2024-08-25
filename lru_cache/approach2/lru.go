package main

import "log"

type DLL struct {
	Key   int
	Value int
	Prev  *DLL
	Next  *DLL
}

type LRUCache struct {
	Head     *DLL
	Tail     *DLL
	Capacity int
	Cache    map[int]*DLL
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Capacity: capacity,
		Cache:    make(map[int]*DLL),
	}
}

func (lru *LRUCache) moveToFront(node *DLL) {
	if node == lru.Head {
		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if node == lru.Tail {
		lru.Tail = node.Prev
	}

	node.Prev = nil
	node.Next = lru.Head

	if lru.Head != nil {
		lru.Head.Prev = node
	}
	lru.Head = node
}

func (lru *LRUCache) addToFront(node *DLL) {
	if lru.Head == nil {
		lru.Head, lru.Tail = node, node
		return
	}

	node.Next = lru.Head
	lru.Head.Prev = node
	lru.Head = node
}

func (lru *LRUCache) evictLRU() {
	if lru.Tail == nil {
		return
	}

	delete(lru.Cache, lru.Tail.Key)

	if lru.Tail.Prev != nil {
		lru.Tail.Prev.Next = nil
	} else {
		lru.Head = nil
	}
	lru.Tail = lru.Tail.Prev
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.Cache[key]; ok {
		lru.moveToFront(node)
		return node.Value
	}
	log.Printf("Key: %d does not exists.", key)
	return -1
}

func (lru *LRUCache) Put(key, value int) {
	if node, ok := lru.Cache[key]; ok {
		node.Value = value
		lru.moveToFront(node)
		return
	}

	node := &DLL{Key: key, Value: value}
	lru.Cache[key] = node
	lru.addToFront(node)
	if len(lru.Cache) > lru.Capacity {
		lru.evictLRU()
	}
}

func main() {
	// Create an LRUCache with capacity 3
	cache := NewLRUCache(3)

	// Put some key-value pairs into the cache
	cache.Put(1, 10) // Cache: [1]
	cache.Put(2, 20) // Cache: [2, 1]
	cache.Put(3, 30) // Cache: [3, 2, 1]

	// Access elements in the cache
	log.Printf("%d", cache.Get(1)) // Output: 10, Cache: [1, 3, 2]

	// Insert a new element and trigger eviction
	cache.Put(4, 40)               // Cache: [4, 1, 3], Evicts key 2
	log.Printf("%d", cache.Get(2)) // Output: -1 (not found)

	// Access element 3
	log.Printf("%d", cache.Get(3)) // Output: 30, Cache: [3, 4, 1]

	// Insert another element and trigger eviction
	cache.Put(5, 50)               // Cache: [5, 3, 4], Evicts key 1
	log.Printf("%d", cache.Get(1)) // Output: -1 (not found)
	log.Printf("%d", cache.Get(4)) // Output: 40
	log.Printf("%d", cache.Get(5)) // Output: 50
}
