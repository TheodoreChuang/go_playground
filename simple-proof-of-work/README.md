# Simple Proof of Work Blockchain in Go

## Running as HTTP Server
Running the server:
`go run main.go`

View the current state:
`
curl --request GET 'localhost:8080/'
`

Add a new block:
`
curl --request POST 'localhost:8080/' \
--header 'Content-Type: text/plain' \
--data-raw '{"BPM": 60}'
`


### Credits
https://medium.com/@mycoralhealth/code-your-own-blockchain-mining-algorithm-in-go-82c6a71aba1f