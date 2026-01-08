package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on :8000")

	conn, err := ln.Accept() // attend 1 client
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n') // attend 1 ligne
	if err != nil {
		if err == io.EOF {
			return
		}
		io.WriteString(conn, "ERR read_error\n")
		return
	}

	fmt.Println("Received:", line)
	io.WriteString(conn, "OK\n")
}
