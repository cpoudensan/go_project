package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Erreur ouverture fichier")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		if len(words) > 0 {
			lastWord := words[len(words)-1]
			fmt.Println(lastWord)
		}
	}
}
      