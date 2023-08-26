package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

func main() {

	fmt.Println("Launching server...")

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	readerStrings := bufio.NewScanner(conn)

	for readerStrings.Scan() {

		message := readerStrings.Text()

		fmt.Println("Message Received:", string(message))

		f, err := strconv.ParseFloat(message, 64)
		if err != nil {
			newFstr := fmt.Sprintf("Incorrect request: %s", message)
			conn.Write([]byte(newFstr + "\n"))
			continue
		}

		newFstr := fmt.Sprintf("%f", f*2)

		conn.Write([]byte(newFstr + "\n"))
	}
}
