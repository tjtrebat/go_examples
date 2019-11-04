package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	f := db.FileMetaData{}
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
	_, err = db.Update(
		&db.FileMetaData{
			ID:          int64(id),
			Name:        handler.Filename,
			Size:        handler.Size,
			ContentType: handler.Header["Content-Type"][0],
			Location:    temp,
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Successfully Uploaded File: %s\n", temp)
}

func viewFile(w http.ResponseWriter, r *http.Request) {
	id, err := utils.FileIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	f, err := db.Select(id)
	http.ServeFile(w, r, f.Location)
}

func main() {
	db.InitDB()
	http.HandleFunc("/phase1", insertMetaData)
	http.HandleFunc("/phase2", uploadFile)
	http.HandleFunc("/", viewFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
