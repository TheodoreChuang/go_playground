# Simple Blockchain in Go

## Running as HTTP Server
Running the server:
`go run simple-blockchain -server=http`

View the current state:
CLI or browser at `http://localhost:8080/`

Add a new block:
`
curl --location --request POST 'localhost:8080/' \
--header 'Content-Type: text/plain' \
--data-raw '{"BPM": 60}'
`

## Running as TCP Server
Creates a network consisting of a TCP server and many other connections. Each connection can create new blocks. All block will be periodically synced with the current state.

Running the server:
`go run simple-blockchain -server=tcp`

For each additional connection:
`nc localhost 9000`

### Packages

`go get github.com/davecgh/go-spew/spew` \
Spew allows us to view structs and slices cleanly formatted in our console. This is nice to have.

`go get github.com/gorilla/mux` \
Gorilla/mux is a popular package for writing web handlers.

`go get github.com/joho/godotenv` \
Godotenv lets us read from a .env file

### Credits
https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc