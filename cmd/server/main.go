package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)

	for {
		line, err := r.ReadString('\n') // attend une ligne
		if err != nil {
			return
		}
		fmt.Println("Received:", line)
		io.WriteString(conn, "OK\n")
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on :8000")

	for {
		conn, err := ln.Accept() // attend un client
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		go handleClient(conn)
	}
}
