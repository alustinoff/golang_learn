package main

import (
	"io"
	"log"
	"net"
)

func echo (conn net.Conn) {
	defer conn.Close()
	b := make([]byte, 512)

	for {
		size, err := conn.Read(b[0:])

		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("unexpected error")
			break
		}
		log.Printf("received %d bytes: %s\n", size, string(b))
		log.Println("Writing data")

		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {  											// infinite loop
		conn, err := listener.Accept()              // block execution as it awaits client connections
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go echo(conn)
	}
}
