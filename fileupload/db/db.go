package db

import (
	"database/sql"
	"fmt"
	"log"

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

// FileMetaData used for db insert operation
type FileMetaData struct {
	ID          int64
	Name        string `json:"name"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	Timestamp   int64
}

// Insert a new row in the db
func Insert(f *FileMetaData) (*FileMetaData, error) {
	stmt, err := db.Prepare("INSERT INTO uploads (name, size, content_type, timestamp) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(f.Name, f.Size, f.ContentType, f.Timestamp)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId() returns null!")
	} else {
		f.ID = id
	}
	return f, nil
}

// Update a location in db
func Update(id int, location string) (sql.Result, error) {
	stmt, err := db.Prepare("UPDATE uploads SET location=? WHERE id=?")
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(location, id)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return nil, err
	}
	return result, nil
}
