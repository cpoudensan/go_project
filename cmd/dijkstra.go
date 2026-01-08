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

