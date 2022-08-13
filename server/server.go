package main

import (
	"io"
	"log"
	"net"
	"time"
)

// ref1: https://hackthedeveloper.com/golang-http-using-tcp-go-net-package/
// ref2: https://www.developer.com/languages/intro-socket-programming-go/

func main() {
	// write server program to handle concurent client connection

	// The Golang Net Package has net.Listen function that takes the name of connection
	// type and the required port and enables us to make a TCP Server.
	listener, err := net.Listen("tcp", "localhost:8000")

	// Next, we check for any kind of error from the net.Listen and print it.
	// Types of Errors might be that the port is occupied or unable to connect.
	if err != nil {
		log.Fatal(err)
	}

	// Then, we use the deferred close statement that closes the connection and thus the resource (port) is
	// taken from the program.
	defer listener.Close()

	// After handling all kinds of errors and closing the connection, it’s time to accept
	// the connection. We will use an infinite loop that will accept the connection and
	// handle the connection.
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// handleConn(conn)
		go handleConn(conn)
	}
}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
