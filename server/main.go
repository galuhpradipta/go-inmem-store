package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/galuhpradipta/go-inmem-db/server/store"
)

const (
	commandSet    = "SET"
	commandDump   = "DUMP"
	commandRename = "RENAME"

	errNotFound        = "not found"
	errWrongCommand    = "wrong command"
	errKeyAlreadyExist = "key already exists"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	store := store.NewDB()

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		command, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		switch {
		case strings.HasPrefix(command, commandSet):
			splitted := strings.Split(command, " ")
			if len(splitted) != 3 {
				writeRes(c, errWrongCommand)
				break
			}

			key := sanitizeStr(splitted[1])
			if store.Get(key) != "" {
				writeRes(c, errKeyAlreadyExist)
				break
			}

			value := sanitizeStr(splitted[2])
			store.Set(key, value)
			writeRes(c, "OK")
		case strings.HasPrefix(command, commandDump):
			splitted := strings.Split(command, " ")
			if len(splitted) != 2 {
				writeRes(c, errWrongCommand)
				break
			}

			key := sanitizeStr(splitted[1])
			value := store.Get(key)
			if value == "" {
				writeRes(c, errNotFound)
				break
			}

			writeRes(c, value)
		case strings.HasPrefix(command, commandRename):
			splitted := strings.Split(command, " ")
			if len(splitted) != 3 {
				writeRes(c, errWrongCommand)
				break
			}

			oldKey := sanitizeStr(splitted[1])
			value := store.Get(oldKey)
			if value == "" {
				writeRes(c, errNotFound)
				break
			}

			store.Delete(oldKey)
			newKey := sanitizeStr(splitted[2])
			store.Set(newKey, value)
			writeRes(c, "OK")
		default:
			writeRes(c, errWrongCommand)
		}

	}
}

func writeRes(c net.Conn, msg string) {
	msg = msg + "\n"
	c.Write([]byte(msg))
}

// sanitize string input from newline character
func sanitizeStr(s string) string {
	return strings.TrimSuffix(s, "\n")
}
