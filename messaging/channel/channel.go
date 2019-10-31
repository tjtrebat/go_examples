package channel

import (
	"sync"
)

// Channel used to store user messages
type Channel struct {
	id       string
	messages chan *Message
	mux      sync.Mutex
}

// Create used for creating a new channel
func Create(id string) *Channel {
	return &Channel{id: id, messages: make(chan *Message, 1)}
}

// SaveMessage used for saving a message to a channel
func (c *Channel) SaveMessage(m *Message) {
	c.mux.Lock()
	c.messages <- m
	c.mux.Unlock()
}

// LoadMessages used for loading messages from a channel
func (c *Channel) LoadMessages(t int) []*Message {
	messages := make([]*Message, 0)
	for {
		select {
		case m := <-c.messages:
			if m.ID > t {
				messages = append(messages, m)
			}
		default:
			return messages
		}
	}
}

func (c *Channel) String() string {
	return c.id
}
