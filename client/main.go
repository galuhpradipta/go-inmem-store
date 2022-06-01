package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// cmdString, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// }

	// cmdString = strings.TrimSuffix(cmdString, "\n")
	// cmd := exec.Command(cmdString)
	// cmd.Stderr = os.Stderr
	// cmd.Stdout = os.Stdout
	// cmd.Run()

	for {
		fmt.Print(">>> ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		cmdString = strings.TrimSuffix(cmdString, "\n")
		cmd := exec.Command(cmdString)
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		err = cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	// strEcho := "GET"
	// servAddr := "localhost:6370"
	// tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	// if err != nil {
	// 	println("ResolveTCPAddr failed:", err.Error())
	// 	os.Exit(1)
	// }

	// conn, err := net.DialTCP("tcp", nil, tcpAddr)
	// if err != nil {
	// 	println("Dial failed:", err.Error())
	// 	os.Exit(1)
	// }

	// _, err = conn.Write([]byte(strEcho))
	// if err != nil {
	// 	println("Write to server failed:", err.Error())
	// 	os.Exit(1)
	// }

	// println("write to server =", strEcho)

	// reply := make([]byte, 1024)

	// _, err = conn.Read(reply)
	// if err != nil {
	// 	println("Write to server failed:", err.Error())
	// 	os.Exit(1)
	// }

	// println("reply from server=", string(reply))

	// conn.Close()
}
