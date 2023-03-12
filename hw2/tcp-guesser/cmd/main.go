package main

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"strings"
	"tcp-guesser/internal/guesses"
)

const (
	addr = "127.0.0.1:8080"
)

func main() {
	log.Printf("starting listening at tcp://%v ...\n", addr)
	lnr, err := net.Listen("tcp", addr)
	defer lnr.Close()
	if err != nil {
		log.Fatalf("can't create tcp listener: %v", err)
	}
	log.Println("start listening tcp port")

	for {
		log.Println("accepting port...")
		conn, err := lnr.Accept()
		defer conn.Close()
		if err != nil {
			log.Fatalf("can't accept port: %v", err)
		}
		log.Println("accept port")
		loop(conn)
	}
}

func loop(conn net.Conn) {
	for {
		log.Print("waiting for message...")
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("can't receive message: %v", err)
			return
		}

		log.Printf("message received: %v", string(message))

		newMessage := handleMessage(message[:len(message)-1])

		log.Printf("sending message: %v ...", newMessage)
		_, err = conn.Write([]byte(newMessage + "\n"))
		if err != nil {
			log.Printf("can't send message: %v", err)
			return
		}
		log.Print("succesfuly send message")
	}
}

type guesser interface {
	Guess(int) int
	Reset()
}

func handleMessage(msg string) string {
	cmd := strings.Split(msg, " ")
	if len(cmd) != 2 {
		return "Wrong format. Use GUESS <number>"
	}

	guess, err := strconv.ParseInt(cmd[1], 10, 64)
	if err != nil {
		return "Can't parse given number"
	}

	var g guesser = guesses.New()
	switch g.Guess(int(guess)) {
	case 1:
		return "LESS"
	case -1:
		return "MORE"
	default:
		g.Reset()
		return "EQUAL"
	}
}
