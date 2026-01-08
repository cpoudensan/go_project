package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Lire les réponses du serveur en continu
	go func() {
		r := bufio.NewReader(conn)
		for {
			line, err := r.ReadString('\n')
			if err != nil {
				return
			}
			fmt.Print("Server: ", line)
		}
	}()

	fmt.Println("Tape une ligne puis Entrée :")
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		// Envoie la ligne au serveur (avec \n)
		fmt.Fprintln(conn, sc.Text())
	}
}
