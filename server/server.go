package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/galuhpradipta/go-inmem-db/server/services"
	"github.com/galuhpradipta/go-inmem-db/server/store"
)

const (
	cmdQuit = "QUIT"
	cmdHelp = "HELP"

	msgCloseConn = "client closing connection"
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

	store := store.NewStore()
	svc := services.NewInMemHandler(store)

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleClientRequest(con, svc)
	}
}

func handleClientRequest(con net.Conn, svc services.InMemHandler) {
	defer con.Close()

	clientReader := bufio.NewReader(con)
	for {
		cmd, err := clientReader.ReadString('\n')
		cmd = strings.TrimSuffix(cmd, ";\n")
		switch err {
		case nil:
			if cmd == cmdQuit {
				log.Println(msgCloseConn)
				return
			}
		case io.EOF:
			log.Println(msgCloseConn)
			return
		default:
			log.Printf("error: %v\n", err)
			return
		}

		res, err := svc.Run(cmd)
		if err != nil {
			res = err.Error()
		}

		if _, err = con.Write(formatCmd(res)); err != nil {
			log.Printf("failed to respond to client: %v\n", err)
		}
	}
}

func formatCmd(cmd string) []byte {
	return []byte(cmd + ";")
}
