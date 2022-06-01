package main

import (
	"fmt"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "6370"
	TYPE = "tcp"
)

type Database struct {
	// map of key/value pairs
	data map[string]string
}

func main() {
	db := Database{data: make(map[string]string)}
	db.Set("person_name", "John")
	db.Set("person", "Galuh")
	data := db.Get("person")
	fmt.Println(data)

	l, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	fmt.Println("listening on " + HOST + ":" + PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go db.handleRequest(conn)
	}

}

func (db *Database) handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	strLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("failed to read request", err)
	}

	msg := string(buf[:strLen])
	switch msg {
	case "SET":
		db.Set("name", "galuh")
		conn.Write([]byte("set success"))
	case "GET":
		data := db.Get("name")
		conn.Write([]byte(data))
	default:
		conn.Write([]byte("Command not found"))
	}

	conn.Close()
}

func (db *Database) Set(key, value string) {
	db.data[key] = value
}

// create retrieve function
func (db *Database) Get(key string) string {
	return db.data[key]
}
