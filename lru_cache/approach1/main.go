package main

import "log"

// DLL represents a node in the doubly linked list
type DLL struct {
	Val        int
	Prev, Next *DLL
}

// LRUCache represents the LRU cache structure
type LRUCache struct {
	Head, Tail *DLL         // Head and Tail of the doubly linked list
	NodeMap    map[int]*DLL // Mapping from value to node
	Capacity   int          // Maximum capacity of the cache
}

// NewLRUCache initializes an LRUCache
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		NodeMap:  make(map[int]*DLL),
		Capacity: capacity,
	}
}

// Insert inserts a value into the cache
func (cache *LRUCache) Insert(data int) {
	// Check if the data already exists in the cache
	if node, ok := cache.NodeMap[data]; ok {
		// Move the node to the front (head) of the DLL
		cache.moveToFront(node)
		return
	}

	// Create a new node for the data
	newNode := &DLL{Val: data}

	// Add the new node to the front (head) of the DLL
	cache.addToFront(newNode)

	// Add the new node to the map
	cache.NodeMap[data] = newNode

	// Check if the cache exceeds its capacity
	if len(cache.NodeMap) > cache.Capacity {
		// Remove the least recently used (LRU) node (tail)
		cache.removeLRU()
	}
}

// moveToFront moves an existing node to the front (head) of the DLL
func (cache *LRUCache) moveToFront(node *DLL) {
	// If the node is already the head, do nothing
	if node == cache.Head {
		return
	}

	// Unlink the node from its current position
	cache.removeNode(node)

	// Add the node to the front (head) of the DLL
	cache.addToFront(node)
}

// addToFront adds a node to the front (head) of the DLL
func (cache *LRUCache) addToFront(node *DLL) {
	node.Next = cache.Head
	node.Prev = nil

	if cache.Head != nil {
		cache.Head.Prev = node
	}
	cache.Head = node

	// If the cache is empty, set both head and tail to the new node
	if cache.Tail == nil {
		cache.Tail = node
	}
}

// removeLRU removes the least recently used node (tail) from the DLL
func (cache *LRUCache) removeLRU() {
	if cache.Tail == nil {
		return
	}

	// Remove the node from the map
	delete(cache.NodeMap, cache.Tail.Val)

	// Unlink the tail node
	cache.removeNode(cache.Tail)
}

// removeNode removes a node from the DLL
func (cache *LRUCache) removeNode(node *DLL) {
	// If the node is the head, update the head pointer
	if node == cache.Head {
		cache.Head = node.Next
	}

	// If the node is the tail, update the tail pointer
	if node == cache.Tail {
		cache.Tail = node.Prev
	}

	// Unlink the node from the DLL
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
}

// Helper function to print the cache content in order (from head to tail)
func printCache(cache *LRUCache) {
	current := cache.Head
	log.Print("Cache state: ")
	for current != nil {
		log.Printf("%d ", current.Val)
		current = current.Next
	}
	log.Println()
}

func main() {
	// Initialize the LRU Cache with capacity 3
	cache := NewLRUCache(3)

	// Scenario 1: Insert elements into the cache
	log.Println("Insert 1, 2, 3 into cache")
	cache.Insert(1) // Cache: [1]
	cache.Insert(2) // Cache: [2, 1]
	cache.Insert(3) // Cache: [3, 2, 1]
	printCache(cache)

	// Scenario 2: Access an element in the cache (this should move the element to the front)
	log.Println("Access element 2")
	cache.Insert(2) // Cache: [2, 3, 1] - Access 2, so it becomes most recently used
	printCache(cache)

	// Scenario 3: Insert another element and trigger eviction (capacity exceeded)
	log.Println("Insert element 4 (evicts least recently used element)")
	cache.Insert(4) // Cache: [4, 2, 3] - Evicts 1 (least recently used)
	printCache(cache)

	// Scenario 4: Access another element
	log.Println("Access element 3")
	cache.Insert(3) // Cache: [3, 4, 2]
	printCache(cache)

	// Scenario 5: Insert another element and trigger eviction
	log.Println("Insert element 5 (evicts least recently used element)")
	cache.Insert(5) // Cache: [5, 3, 4] - Evicts 2
	printCache(cache)

	// Scenario 6: Insert element 6 (causes eviction)
	log.Println("Insert element 6 (evicts least recently used element)")
	cache.Insert(6) // Cache: [6, 5, 3] - Evicts 4
	printCache(cache)
}
