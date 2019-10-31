package requestutil

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var validPath = regexp.MustCompile("^/([a-zA-Z0-9]+)/messages$")

// ChannelID returns channel id from a url
func ChannelID(path string) (string, error) {
	m := validPath.FindStringSubmatch(path)
	if m == nil {
		return "", errors.New("Invalid Channel ID")
	}
	return strings.ToLower(m[1]), nil
}

// ReadRequestBody reads request body as []byte
func ReadRequestBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// LastIDParam returns the last_id query parameter
func LastIDParam(r *http.Request) int {
	ids, ok := r.URL.Query()["last_id"]
	if !ok || len(ids) < 1 {
		return -1
	}
	id, err := strconv.Atoi(ids[0])
	if err != nil {
		return -1
	}
	return id
}
