package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tjtrebat/messaging/channel"
	"github.com/tjtrebat/messaging/list"
	"github.com/tjtrebat/messaging/requestutil"
)

// ViewChannelResponse response for user messages
type ViewChannelResponse struct {
	Messages []*channel.Message `json:"messages"`
}

// UpdateChannelResponse response for message id
type UpdateChannelResponse struct {
	ID int `json:"id"`
}

var channels = list.ChannelList{Channels: make(map[string]*channel.Channel)}

func handler(w http.ResponseWriter, r *http.Request) {
	id, err := requestutil.ChannelID(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if r.Method == http.MethodGet {
		viewChannel(w, r, id)
	} else if r.Method == http.MethodPost {
		updateChannel(w, r, id)
	}
}

func viewChannel(w http.ResponseWriter, r *http.Request, id string) {
	c, ok := channels.Channels[id]
	if !ok {
		http.Error(w, "Channel Not Found", http.StatusNotFound)
		return
	}
	writeResponse(w, &ViewChannelResponse{c.LoadMessages(requestutil.LastIDParam(r))})
}

func updateChannel(w http.ResponseWriter, r *http.Request, id string) {
	m, err := channel.ParseMessage(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	addChannelMessage(id, m)
	writeResponse(w, &UpdateChannelResponse{m.ID})
}

func writeResponse(w http.ResponseWriter, r interface{}) {
	resp, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func addChannelMessage(id string, m *channel.Message) {
	channel := channels.AddChannel(id)
	go channel.SaveMessage(m)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
