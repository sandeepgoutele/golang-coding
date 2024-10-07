package main

import (
	"github.com/brianvoe/gofakeit"
	"log"
)

var currentNum int

type SharedQueue struct {
	Queue    chan int
	Capacity int
}

func NewSharedQueue(capacity int) *SharedQueue {
	return &SharedQueue{Capacity: capacity,
		Queue: make(chan int)}
}

func (queue *SharedQueue) Produce() {
	currentNum = int(gofakeit.Int8())
	if len(queue.Queue) < queue.Capacity {
		queue.Queue <- currentNum
		return
	}

	for len(queue.Queue) < queue.Capacity {

	}
}

func (queue *SharedQueue) Consume() {
	for len(queue.Queue) == 0 {

	}

	log.Printf("Consumed: %d", <-queue.Queue)
}

func main() {
	queue := NewSharedQueue(5)
	go queue.Produce()
	go queue.Consume()
}
