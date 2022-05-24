package main

import (
	"fmt"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "3333"
	connType = "tcp"
)

func main() {
	// listen to port
	l, err := net.Listen(connType, connHost+":"+connPort)
	// error occurred while listening
	if err != nil {
		fmt.Println("Error while listening: " + err.Error())
		os.Exit(1)
	}

	defer l.Close()

	fmt.Println("Listening at " + connHost + ":" + connPort)
	// while loop to listen
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error while accepting connection: " + err.Error())
			os.Exit(1)
		}
		fmt.Println("Accepted connection")
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	buf := make([]byte, 2048)
	len, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error while handling request: " + err.Error())
	}
	fmt.Printf("Read %v bytes\n", len)
	conn.Write(buf)
	fmt.Printf("Responded\n")
	conn.Close()
}
