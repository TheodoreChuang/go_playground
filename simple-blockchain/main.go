package main

import (
	"flag"
	"log"
	"net"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func createGenesisBlock() {
	t := time.Now()
	genesisBlock := Block{0, t.String(), 0, "", ""}
	spew.Dump(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	createGenesisBlock()

	serverType := flag.String("server", "http", "http or tcp")
	flag.Parse()

	if *serverType == "http" {
		log.Fatal(runHTTPServer())
	}

	if *serverType == "tcp" {
		bcServer = make(chan []Block)

		tcpServer, err := net.Listen("tcp", ":"+os.Getenv("ADDR"))
		if err != nil {
			log.Fatal(err)
		}
		defer tcpServer.Close()

		// create a new connection each time there is a request
		for {
			conn, err := tcpServer.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go handleConn(conn)
		}
	}
}
