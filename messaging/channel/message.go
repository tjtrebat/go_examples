package channel

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/tjtrebat/messaging/requestutil"
)

// Message contains message text from user
type Message struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

// ParseMessage parses a Message from an http.Request
func ParseMessage(r *http.Request) (*Message, error) {
	body, err := requestutil.ReadRequestBody(r)
	if err != nil {
		return nil, errors.New("Error in request body")
	}
	return unmarshal(body), nil
}

func unmarshal(body []byte) *Message {
	m := Message{ID: int(time.Now().Unix())}
	json.Unmarshal(body, &m)
	return &m
}

func (m *Message) String() string {
	return fmt.Sprintf("Message: %s, Sent: %d", m.Message, m.ID)
}
