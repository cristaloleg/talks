package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// START_DEFS OMIT
type (
	Post    string
	Message string
	Feed    chan<- Message
)

type Subscriber interface {
	Socket() Feed
}

type Broker struct {
	sync.Mutex
	connections map[Post][]Feed
}

type Client struct {
	name  string
	feed  chan Message
	sleep chan time.Duration
	done  chan struct{}
}

// END_DEFS OMIT

func NewBroker() *Broker {
	return &Broker{
		connections: make(map[Post][]Feed),
	}
}

// START_PUBLISH OMIT
func (b *Broker) Publish(post Post) {
	b.Lock()

	if _, ok := b.connections[post]; !ok {
		b.connections[post] = make([]Feed, 0)
	}

	b.Unlock()
}

// END_PUBLISH OMIT

// START_SUBSCRIBE OMIT
func (b *Broker) Subscribe(post Post, client Subscriber) (err error) {
	b.Lock()
	defer b.Unlock()

	queue, ok := b.connections[post]
	if !ok {
		return errors.New("no such post")
	}
	b.connections[post] = append(queue, client.Socket())

	return
}

// END_SUBSCRIBE OMIT

func (b *Broker) Close(post Post) {
	b.Lock()

	delete(b.connections, post)

	b.Unlock()
}

// START_NOTIFY OMIT
func (b *Broker) Notify(post Post, message Message) {
	b.Lock()
	defer b.Unlock()

	queue, ok := b.connections[post]
	if !ok {
		return
	}
	for _, q := range queue {
		go func(q Feed) {
			q <- message
		}(q)
	}
}

// END_NOTIFY OMIT

// START_NEWCLIENT OMIT
func NewClient(name string) *Client {
	return &Client{
		name:  name,
		feed:  make(chan Message),
		sleep: make(chan time.Duration, 1),
		done:  make(chan struct{}, 1),
	}
}

// END_NEWCLIENT OMIT

// START_IMPL OMIT
func (c *Client) Socket() Feed {
	return c.feed
}

// END_IMPL OMIT

// START_SLEEP OMIT
func (c *Client) Sleep(delay time.Duration) {
	c.sleep <- delay
}

// END_SLEEP OMIT

// START_BYE OMIT
func (c *Client) Disconnect() {
	c.done <- struct{}{}
}

// END_BYE OMIT

// START_LISTEN OMIT
func (c *Client) Listen() {
	go func() {
		for {
			select {
			case msg := <-c.feed:
				fmt.Println(msg)

			case delay := <-c.sleep:
				time.Sleep(delay)

			case <-c.done:
				return

			default:
				time.Sleep(451 * time.Millisecond)
			}
		}
	}()
}

// END_IMPL OMIT

func main() {
	var title Post = "wow, such title, very informative, much interesting"
	client1 := NewClient("gopher")
	client1.Listen()

	client2 := NewClient("someone")
	client2.Listen()

	broker := NewBroker()
	broker.Publish(title)

	pinger := func(title Post, interval time.Duration) {
		for {
			broker.Notify(title, "ping...")
			time.Sleep(interval)
		}
	}
	go pinger(title, time.Second*2)

	broker.Subscribe(title, client1)
	broker.Subscribe(title, client2)

	broker.Notify(title, "Hi, everyone!")

	client2.Sleep(time.Second * 5)

	broker.Notify(title, "is anyone here?")

	time.Sleep(time.Second * 20)
}
