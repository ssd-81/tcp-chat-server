package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	server()

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)

	}
	conn.Write([]byte("apple is good"))
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}

func server() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("server not listening; figure out what is causing the error")
	}

	// the below snippet is handling multiple connections; but right now we just want to connect a single client to the server
	// for {
	// 	conn, err := ln.Accept()
	// 	if err != nil {
	// 		// handle error
	// 	}
	// 	go handleConnection(conn)
	// }

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("connection failed; server could not connect to the client")
	}
	go handleConnection(conn)

}

func handleConnection(conn net.Conn) {
	var dataStream []byte // check if this fine with conn.Read
	conn.Read(dataStream)
	fmt.Println(dataStream)
}
