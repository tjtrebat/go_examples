# go_examples

A collection of sample projects written in the Go programming language

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

A Go workspace environment

### Installing

To run the development server:

```
cd messaging/handler
go build
handler
```

Send a message to channel with http POST:
`
curl -X POST "http://localhost:8080/monkeys/messages" -H "Content-Type: application/json" -d '{"username":"tom", "message": "hello world"}'
`

Retrieve messages from channel:
`
curl -X GET "http://localhost:8080/monkeys/messages"
`

Retrieve messages from channel from timestamp:
`
curl -X GET "http://localhost:8080/monkeys/messages?last_id=<ID of last message>"
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

* GoodRx Interview
