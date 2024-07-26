package modules

import (
	"log"
	"sync"
)

type BlockingQueue struct {
	mutex     sync.Mutex
	condition sync.Cond
	data      []interface{}
	capacity  int
}

func NewBlockingQueue(capacity int) *BlockingQueue {
	blockingQueue := new(BlockingQueue)
	blockingQueue.condition = sync.Cond{L: &blockingQueue.mutex}
	blockingQueue.capacity = capacity
	return blockingQueue
}

func (bQue *BlockingQueue) Put(item interface{}) {
	bQue.condition.L.Lock()
	defer bQue.mutex.Unlock()

	for bQue.isFull() {
		bQue.condition.Wait()
	}

	bQue.data = append(bQue.data, item)
	bQue.condition.Signal()
}

func (bQue *BlockingQueue) isFull() bool {
	return len(bQue.data) == bQue.capacity
}

func (bQue *BlockingQueue) Take() interface{} {
	bQue.condition.L.Lock()
	defer bQue.mutex.Unlock()

	for bQue.isEmpty() {
		bQue.condition.Wait()
	}

	result := bQue.data[0]
	bQue.data = bQue.data[:1]
	bQue.condition.Signal()
	return result
}

func (bQue *BlockingQueue) isEmpty() bool {
	return len(bQue.data) == 0
}

func BQueTester() {
	bQue := NewBlockingQueue(5)
	bQue.Put(1)
	bQue.Put(2)
	log.Print(bQue.data...)
	log.Print(bQue.Take())
	log.Print(bQue.data...)
	for i := 0; i < 10; i++ {
		bQue.Put(i)
		log.Print(bQue.data...)
	}
	for i := 0; i < 10; i++ {
		log.Print(bQue.Take())
		log.Print(bQue.data...)
	}
}
