# go_examples

A collection of sample projects written in the Go programming language

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

- A Go workspace environment
- MySQL

### Installing

1. Create MySQL database:

```create database novacredit```

2. Update db connection string in db.go to use a local MySQL configuration.

3. Finally, run the development server:

```
cd fileupload/main
go run handler.go
```

## Usage

Create a database row for the file meta data:

`
curl -X POST "http://localhost:8080/phase1" -H "Content-Type: application/json" -d '{"name":"Resume.pdf", "size": 74295, "contentType": "application/pdf"}'
`

Upload a file using the ID obtained from the response:

`
curl -F 'data=@/home/user/Documents/Resume.pdf' -F "id=1" http://localhost:8080/phase2
`

## Running the tests

TODO


## Deployment

TODO

## Built With

* [Go](https://golang.org) - The programming language

## Authors

* **Thomas Trebat** - *Initial work* - [tjtrebat](https://github.com/tjtrebat)

## License

This project is licensed under the GPLv2 License

## Acknowledgments

* Nova Credit Interview
