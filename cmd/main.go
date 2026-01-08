package main

import "fmt"

func main() {
	path, dist, err := Dijkstra(graph, "Paris", "Lyon")

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Println("Chemin :", path)
	fmt.Println("Distance :", dist, "km")
}