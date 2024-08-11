// Reference: https://medium.com/@parvjn616/building-a-message-broker-in-go-a-beginners-guide-e4d8be2359c

package main

import (
	"log"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit"
)

type Message struct {
	Topic   string
	Payload interface{}
}

type Subscriber struct {
	Channel     chan interface{}
	Unsubscribe chan bool
}

type Broker struct {
	subscriber map[string][]*Subscriber
	mutax      sync.Mutex
}

func NewBroker() *Broker {
	return &Broker{subscriber: make(map[string][]*Subscriber)}
}

func (broker *Broker) Subscribe(topic string) *Subscriber {
	broker.mutax.Lock()
	defer broker.mutax.Unlock()
	subscriber := &Subscriber{
		Channel:     make(chan interface{}, 1),
		Unsubscribe: make(chan bool),
	}
	broker.subscriber[topic] = append(broker.subscriber[topic], subscriber)
	return subscriber
}

func (broker *Broker) Unsubscribe(topic string, subscriber *Subscriber) {
	broker.mutax.Lock()
	defer broker.mutax.Unlock()
	if subscribers, found := broker.subscriber[topic]; found {
		for idx, sub := range subscribers {
			if sub == subscriber {
				close(sub.Channel)
				broker.subscriber[topic] = append(subscribers[:idx], subscribers[idx+1:]...)
				return
			}
		}
	}
}

func (broker *Broker) Publish(message *Message) {
	broker.mutax.Lock()
	defer broker.mutax.Unlock()
	if subscribers, found := broker.subscriber[message.Topic]; found {
		for _, subs := range subscribers {
			select {
			case subs.Channel <- message.Payload:
			case <-time.After(time.Second):
				log.Printf("Subscriber is slow, unsubscribing from topi: %s\n", message.Topic)
				broker.Unsubscribe(message.Topic, subs)
			}
		}
	}
}

func main() {
	broker := NewBroker()
	topic := "example_topic"
	subscriber := broker.Subscribe(topic)

	go func() {
		for {
			select {
			case msg, ok := <-subscriber.Channel:
				if !ok {
					log.Println("Subscriber channel closed.")
					return
				}
				log.Printf("Received msg: %v", msg)
			case <-subscriber.Unsubscribe:
				log.Printf("Subscriber :%v unsubscribed.", subscriber)
				return
			}
		}
	}()

	gofakeit.Seed(0)
	var wg sync.WaitGroup
	for idx := 0; idx < 20; idx++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			broker.Publish(&Message{Topic: topic, Payload: gofakeit.Sentence(5)})
		}()
	}
	wg.Wait()

	time.Sleep(2 * time.Second)
	broker.Unsubscribe(topic, subscriber)

	broker.Publish(&Message{Topic: topic, Payload: "This message won't be received."})

	time.Sleep(time.Second)
}
