package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Println("Please provide host and port.")
		return
	}

	host := arguments[1]
	port := arguments[2]
	c, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("mockRedis$> ")
		command, _ := reader.ReadString('\n')
		fmt.Fprintf(c, command+"\n")

		buf := make([]byte, 1024)
		_, err := c.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(">>: " + string(buf))
	}
}
