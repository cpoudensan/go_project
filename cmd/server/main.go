package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return // client fermé
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 3 {
			io.WriteString(conn, "ERR bad_request\n")
			continue
		}

		if strings.ToUpper(parts[0]) != "ROUTE" {
			io.WriteString(conn, "ERR unknown_cmd\n")
			continue
		}

		from := parts[1]
		to := parts[2]

		path, dist, err := Dijkstra(Graph, from, to)
		if err != nil {
			if err == ErrUnknownCity {
				io.WriteString(conn, "ERR unknown_city\n")
			} else if err == ErrNoRoute {
				io.WriteString(conn, "ERR no_route\n")
			} else {
				io.WriteString(conn, "ERR internal\n")
			}
			continue
		}

		io.WriteString(conn, fmt.Sprintf("OK dist=%d path=%s\n", dist, strings.Join(path, "->")))
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on :8000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		go handleClient(conn)
	}
}
