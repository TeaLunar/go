package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	for {
		reader := bufio.NewScanner(os.Stdin)
		if reader.Scan() == false {
			break
		}
		text := reader.Text()

		conn.Write([]byte(text + "\n"))

		readerConn := bufio.NewScanner(conn)
		if readerConn.Scan() == false {
			break
		}
		fmt.Println(readerConn.Text())

	}

}
