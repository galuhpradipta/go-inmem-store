package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		log.Println("please provide port number")
		return
	}

	port := arguments[1]
	listener, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleClientRequest(con)
	}
}

func handleClientRequest(con net.Conn) {
	defer con.Close()

	clientReader := bufio.NewReader(con)

	for {
		clientRequest, err := clientReader.ReadString('\n')

		switch err {
		case nil:
			clientRequest := strings.TrimSpace(clientRequest)
			if clientRequest == ":QUIT" {
				log.Println("client requested server to close the connection so closing")
				return
			} else {
				log.Println(clientRequest)
			}
		case io.EOF:
			log.Println("client closed the connection by terminating the process")
			return
		default:
			log.Printf("error: %v\n", err)
			return
		}

		// Responding to the client request
		if _, err = con.Write([]byte("GOT IT!\n")); err != nil {
			log.Printf("failed to respond to client: %v\n", err)
		}
	}
}
