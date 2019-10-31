package list

import (
	"sync"

	"github.com/tjtrebat/messaging/channel"
)

// ChannelList a map of channels
type ChannelList struct {
	Channels map[string]*channel.Channel
	mux      sync.Mutex
}

// AddChannel adds a channel to a map
func (cl *ChannelList) AddChannel(id string) *channel.Channel {
	if _, ok := cl.Channels[id]; !ok {
		cl.mux.Lock()
		cl.Channels[id] = channel.Create(id)
		cl.mux.Unlock()
	}
	return cl.Channels[id]
}
