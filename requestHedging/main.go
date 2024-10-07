package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var cache = make(map[string]string)
var inMemoryCache = make(map[string]string)
var keyLocks sync.Map
var dbWaitGroup = sync.Map{}

func dbQuery(key string) string {
	log.Printf("Querying DB for key: %s", key)
	time.Sleep(2 * time.Second) // Simulate a slow DB response
	return fmt.Sprintf("DBValue:%s", key)
}

func getKeyLock(key string) *sync.Mutex {
	val, _ := keyLocks.LoadOrStore(key, &sync.Map{})
	return val.(*sync.Mutex)
}

func checkCache(key string) (string, bool) {
	keyLock := getKeyLock(key)
	keyLock.Lock()
	defer keyLock.Unlock()

	value, existed := cache[key]
	return value, existed
}

func updateCache(key string, value string) {
	keyLock := getKeyLock(key)
	keyLock.Lock()
	defer keyLock.Unlock()

	cache[key] = value
}

func apiServer(key string) string {
	keyLock := getKeyLock(key)
	keyLock.Lock()
	defer keyLock.Unlock()

	if value, exists := inMemoryCache[key]; exists {
		log.Printf("Serving from in memory cache for key: %s", key)
		return value
	}

	if value, exists := cache[key]; exists {
		log.Printf("Serving from cache server for key: %s", key)
		return value
	}
	_, inProgress := dbWaitGroup.LoadOrStore(key, &sync.WaitGroup{})
	waitGroup, _ := dbWaitGroup.LoadOrStore(key, &sync.WaitGroup{})
	wg := waitGroup.(*sync.WaitGroup)

	if !inProgress {
		wg.Add(1)
		go func() {
			defer wg.Done()
			valueFromDB := dbQuery(key)

			inMemoryCache[key] = valueFromDB
			updateCache(key, valueFromDB)
		}()
	}

	wg.Wait()

	value := inMemoryCache[key]
	log.Printf("Serving from in-memory cache (after DB query) for key: %s", key)

	dbWaitGroup.Delete(key)

	return value
}

func simulateClientRequest(clientID int, key string, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Client %d requesting key: %s", clientID, key)
	value := apiServer(key)
	log.Printf("Client %d received value: %s", clientID, value)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	keys := []string{
		"key1", "key2", "key3", "key1", "key3", "key2", // Initial requests
		"key4", "key1", "key5", "key2", "key6", "key3", // More requests for different keys
	}

	for i, key := range keys {
		wg.Add(1)
		go simulateClientRequest(i, key, &wg)

		if i == 5 {
			time.Sleep(2 * time.Second) // Short delay to simulate staggered requests
		}
	}

	wg.Wait()
	log.Println("All clients have received their responses.")
}
