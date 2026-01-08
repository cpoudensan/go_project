import "math"

func dijkstra(
	graph map[string]map[string]int,
	start string,
	target string,
) (int, []string, bool) {

	dist := make(map[string]int)
	prev := make(map[string]string)
	visited := make(map[string]bool)

	// initialisation
	for node := range graph {
		dist[node] = math.MaxInt
	}
	if _, ok := graph[start]; !ok {
		return 0, nil, false
	}
	dist[start] = 0

	for {
		// trouver le noeud non visité avec la plus petite distance
		cur := ""
		best := math.MaxInt
		for node := range dist {
			if !visited[node] && dist[node] < best {
				best = dist[node]
				cur = node
			}
		}

		if cur == "" {
			break
		}
		if cur == target {
			break
		}

		visited[cur] = true

		// relaxation
		for neigh, weight := range graph[cur] {
			if visited[neigh] {
				continue
			}
			if dist[cur]+weight < dist[neigh] {
				dist[neigh] = dist[cur] + weight
				prev[neigh] = cur
			}
		}
	}

	// pas atteignable
	if dist[target] == math.MaxInt {
		return 0, nil, false
	}

	// reconstruire le chemin
	path := []string{}
	for at := target; at != ""; at = prev[at] {
		path = append(path, at)
		if at == start {
			break
		}
	}

	// reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return dist[target], path, true
}
