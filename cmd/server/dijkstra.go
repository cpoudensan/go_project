package main

import (
	"errors"
	"math"
)

var ErrUnknownCity = errors.New("unknown city")
var ErrNoRoute = errors.New("no route possible")
func Dijkstra(graph map[string]map[string]int, start string, target string) ([]string, int, error) {
	if _, ok := graph[start]; ok == false {
		return nil, 0, ErrUnknownCity
	}
	if _, ok := graph[target]; ok == false {
		return nil, 0, ErrUnknownCity
	}
	if start == target {
		return []string{start}, 0, nil
	}
	dist := make(map[string]int)
	prev := make(map[string]string)
	visited := make(map[string]bool)

	// initialiser toutes les distances à "infini" sauf le départ
	for city := range graph {
		dist[city] = math.MaxInt / 4
	}
	dist[start] = 0

	for {
		// trouver la ville non visitée la plus proche
		u := ""
		best := math.MaxInt / 4

		for city := range graph {
			if !visited[city] && dist[city] < best {
				best = dist[city]
				u = city
			}
		}

		// plus rien atteignable
		if u == "" {
			return nil, 0, ErrNoRoute
		}

		// on a atteint target
		if u == target {
			break
		}
         //marquer qu'on a bien visité la ville
		visited[u] = true

		// check si on peut améliorer la distance pour aller vers un voisin
		for v, w := range graph[u] {
			if visited[v] {
				continue
			}
			alt := dist[u] + w
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}

	// reconstruire le chemin (à l'envers)
	path := []string{}
	cur := target
	path = append(path, cur)

	for cur != start {
		p, ok := prev[cur]
		if !ok {
			return nil, 0, ErrNoRoute
		}
		cur = p
		path = append(path, cur)
	}

	// on renverse pour avoir le bon chemin
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path, dist[target], nil
}