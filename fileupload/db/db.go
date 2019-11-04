package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	// MySql driver
	_ "github.com/go-sql-driver/mysql"
)

var db, err = sql.Open("mysql", "tom:password@/novacredit")

// InitDB initializes tables for the novacredit db
func InitDB() {
	q := `
		create table if not exists uploads (
			id 						int primary key auto_increment,
			name 					varchar(500) not null,
			size 					int not null,
			content_type 	varchar(500) not null,
			timestamp			int not null,
			location 			varchar(500)
		);
	`
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

// FileMetaData used for db insert/update operations
type FileMetaData struct {
	ID          int64
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	ContentType string `json:"contentType"`
	Location    string
}

// Insert a new row in the db
func Insert(f *FileMetaData) (*FileMetaData, error) {
	stmt, err := db.Prepare("INSERT INTO uploads (name, size, content_type, timestamp) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(f.Name, f.Size, f.ContentType, time.Now().Unix())
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return nil, err
	}
	f.ID, _ = result.LastInsertId()
	return f, nil
}

// Update a location in db
func Update(f *FileMetaData) (sql.Result, error) {
	stmt, err := db.Prepare("UPDATE uploads SET name=?, size=?, content_type=?, location=? WHERE id=?")
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(f.Name, f.Size, f.ContentType, f.Location, f.ID)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return nil, err
	}
	return result, nil
}

// Select a row from the db
func Select(id int) (*FileMetaData, error) {
	rows, err := db.Query("SELECT name, size, content_type, location FROM uploads WHERE id=?", id)
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.New("File with specified id does not exist")
	}
	f := FileMetaData{}
	if err := rows.Scan(&f.Name, &f.Size, &f.ContentType, &f.Location); err != nil {
		fmt.Printf("An error occurred: %v", err)
		return nil, err
	}
	return &f, nil
}
