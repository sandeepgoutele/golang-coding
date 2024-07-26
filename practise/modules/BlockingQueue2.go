package modules

import "log"

type BlockingQueue2 struct {
	channel chan interface{}
}

func NewBlockingQueue2(cap int) *BlockingQueue2 {
	bQue := new(BlockingQueue2)
	bQue.channel = make(chan interface{}, cap)
	return bQue
}

func (bQue *BlockingQueue2) Put(item interface{}) {
	bQue.channel <- item
}

func (bQue *BlockingQueue2) Take() interface{} {
	return <-bQue.channel
}

func BQueTester2() {
	bQue := NewBlockingQueue2(5)
	bQue.Put(1)
	bQue.Put(2)
	log.Printf("%+v", bQue.channel)
	log.Print(bQue.Take())
	log.Printf("%+v", bQue.channel)
	log.Print(bQue.Take())
	log.Print(bQue.Take())
	// for i := 0; i < 10; i++ {
	// 	bQue.Put(i)
	// 	log.Print(bQue.data...)
	// }
	// for i := 0; i < 10; i++ {
	// 	log.Print(bQue.Take())
	// 	log.Print(bQue.data...)
	// }
}
