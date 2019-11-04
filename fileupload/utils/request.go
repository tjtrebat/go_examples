package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// FileID returns the id of the file in the request
func FileID(r *http.Request) (int, error) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		return -1, err
	}
	return id, nil
}

var validPath = regexp.MustCompile("^/data/([0-9]+)$")

// FileIDFromPath returns channel id from a url
func FileIDFromPath(path string) (int, error) {
	m := validPath.FindStringSubmatch(path)
	if m == nil {
		return -1, errors.New("Invalid File ID")
	}
	return strconv.Atoi(m[1])
}

// RetrieveFile retrieves the file in the request
func RetrieveFile(r *http.Request) (multipart.File, *multipart.FileHeader, error) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("data")
	if err != nil {
		fmt.Printf("An error occurred: %v", file)
		return nil, nil, err
	}
	defer file.Close()
	return file, handler, nil
}

// UploadTempFile uploads a file to a temporary location
func UploadTempFile(f multipart.File, h *multipart.FileHeader) (string, error) {
	temp, err := createTempFile(suffix(h.Filename))
	if temp != nil {
		defer temp.Close()
		name, err := writeTempFile(temp, f)
		if err == nil {
			return name, nil
		}
	}
	return "", err
}

func suffix(name string) string {
	if i := strings.LastIndex(name, "."); i > -1 {
		return name[i:]
	}
	return ""
}

func createTempFile(e string) (*os.File, error) {
	tempFile, err := ioutil.TempFile("C:/temp", "upload-*"+e)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return nil, err
	}
	return tempFile, nil
}

func writeTempFile(t *os.File, f multipart.File) (string, error) {
	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return "", err
	}
	t.Write(fileBytes)
	return t.Name(), nil
}

// ReadRequestBody reads request body as []byte
func ReadRequestBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return nil, err
	}
	return body, nil
}
