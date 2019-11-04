package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tjtrebat/fileupload/db"
	"github.com/tjtrebat/fileupload/utils"
)

func insertMetaData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		f, err := parseFileMetaData(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		db.Insert(f)
		fmt.Printf("name: %s, size: %d, contentType: %s", f.Name, f.Size, f.ContentType)
		writeResponse(w, f)
	}
}

func parseFileMetaData(r *http.Request) (*db.FileMetaData, error) {
	body, err := utils.ReadRequestBody(r)
	if err == nil {
		f, err := unmarshal(body)
		if err == nil {
			return f, nil
		}
	}
	return nil, err
}

func unmarshal(body []byte) (*db.FileMetaData, error) {
	f := db.FileMetaData{Timestamp: time.Now().Unix()}
	err := json.Unmarshal(body, &f)
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		return nil, err
	}
	return &f, nil
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

func uploadFile(w http.ResponseWriter, r *http.Request) {
	id, err := utils.FileID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	file, handler, err := utils.RetrieveFile(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	temp, err := utils.UploadTempFile(file, handler)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = db.Update(id, temp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Successfully Uploaded File: %s\n", temp)
}

func main() {
	db.InitDB()
	http.HandleFunc("/phase1", insertMetaData)
	http.HandleFunc("/phase2", uploadFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
