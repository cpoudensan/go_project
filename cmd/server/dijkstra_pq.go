package main

import (
	"container/heap"
	"errors"
	"math"
	"sync"
)

/*
Erreurs retournées par Dijkstra
*/
var ErrUnknownCity = errors.New("unknown city")
var ErrNoRoute = errors.New("no route possible")

/*
Cache global (partagé entre goroutines)
- clé : "start|target"
- valeur : résultat (path, dist, err)
*/
type RouteResult struct {
	Path []string
	Dist int
	Err  error
}

var (
	cacheMu    sync.RWMutex
	routeCache = make(map[string]RouteResult)
)

// ---- Priority Queue (min-heap) ----

type pqItem struct {
	city string
	dist int
}

type priorityQueue []pqItem

func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist // min-heap: plus petite distance d'abord
}
func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(x any) {
	*pq = append(*pq, x.(pqItem))
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	it := old[n-1]
	*pq = old[:n-1]
	return it
}

/*
Dijkstra avec :
- Priority Queue (heap) => plus rapide que le scan O(V) à chaque étape
- Cache intégré => si même requête (start,target), réponse immédiate

Signature identique à ton ancienne fonction:
([]string, int, error)
*/
func Dijkstra_pq(graph map[string]map[string]int, start string, target string) ([]string, int, error) {
	// ---- 0) Cache lookup ----
	key := start + "|" + target

	cacheMu.RLock()
	cached, ok := routeCache[key]
	cacheMu.RUnlock()

	if ok {
		// On renvoie des copies pour éviter toute surprise si quelqu'un modifie un slice
		if cached.Path != nil {
			pathCopy := make([]string, len(cached.Path))
			copy(pathCopy, cached.Path)
			return pathCopy, cached.Dist, cached.Err
		}
		return nil, cached.Dist, cached.Err
	}

	// ---- 1) Vérifs de base ----
	if _, ok := graph[start]; !ok {
		// stocker dans le cache
		cacheMu.Lock()
		routeCache[key] = RouteResult{Path: nil, Dist: 0, Err: ErrUnknownCity}
		cacheMu.Unlock()
		return nil, 0, ErrUnknownCity
	}
	if _, ok := graph[target]; !ok {
		cacheMu.Lock()
		routeCache[key] = RouteResult{Path: nil, Dist: 0, Err: ErrUnknownCity}
		cacheMu.Unlock()
		return nil, 0, ErrUnknownCity
	}
	if start == target {
		resPath := []string{start}
		cacheMu.Lock()
		routeCache[key] = RouteResult{Path: resPath, Dist: 0, Err: nil}
		cacheMu.Unlock()

		pathCopy := make([]string, len(resPath))
		copy(pathCopy, resPath)
		return pathCopy, 0, nil
	}

	// ---- 2) Structures Dijkstra ----
	dist := make(map[string]int, len(graph))
	prev := make(map[string]string, len(graph))
	visited := make(map[string]bool, len(graph))

	inf := math.MaxInt / 4
	for city := range graph {
		dist[city] = inf
	}
	dist[start] = 0

	// ---- 3) Init heap ----
	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, pqItem{city: start, dist: 0})

	// ---- 4) Boucle principale ----
	for pq.Len() > 0 {
		it := heap.Pop(pq).(pqItem)
		u := it.city

		// Si entrée périmée (on a déjà trouvé mieux), ignore
		if it.dist != dist[u] {
			continue
		}
		// Si déjà visitée, ignore
		if visited[u] {
			continue
		}

		// Si on atteint target, on peut stopper
		if u == target {
			break
		}

		visited[u] = true

		// Relaxer les voisins
		for v, w := range graph[u] {
			if visited[v] {
				continue
			}
			alt := dist[u] + w
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				heap.Push(pq, pqItem{city: v, dist: alt})
			}
		}
	}

	// ---- 5) Pas de route ? ----
	if dist[target] >= inf {
		cacheMu.Lock()
		routeCache[key] = RouteResult{Path: nil, Dist: 0, Err: ErrNoRoute}
		cacheMu.Unlock()
		return nil, 0, ErrNoRoute
	}

	// ---- 6) Reconstruire le chemin ----
	path := []string{}
	cur := target
	path = append(path, cur)

	for cur != start {
		p, ok := prev[cur]
		if !ok {
			cacheMu.Lock()
			routeCache[key] = RouteResult{Path: nil, Dist: 0, Err: ErrNoRoute}
			cacheMu.Unlock()
			return nil, 0, ErrNoRoute
		}
		cur = p
		path = append(path, cur)
	}

	// Reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	// ---- 7) Stocker résultat dans le cache ----
	// (on stocke une copie du slice)
	stored := make([]string, len(path))
	copy(stored, path)

	cacheMu.Lock()
	routeCache[key] = RouteResult{Path: stored, Dist: dist[target], Err: nil}
	cacheMu.Unlock()

	// Renvoie une copie
	out := make([]string, len(path))
	copy(out, path)
	return out, dist[target], nil
}
