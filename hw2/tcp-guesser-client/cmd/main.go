package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Can't connect to socket: %v", err)
	}
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("can't read from stdin: %v", err)
		}
		fmt.Fprintf(conn, "GUESS "+text+"\n")
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalf("can't recieve answer: %v", err)
		}
		fmt.Print("Message from server: " + message)
	}

}
