package euler

import (
	"container/heap"
	"fmt"
	"strings"
)

// Graph ..
type Graph struct {
	adjacency map[int64]map[int64]float64
	nodes     map[int64]Node
}

// Node ..
type Node interface {
	ID() int64
}

// NewGraph ...
func NewGraph() *Graph {
	return &Graph{
		make(map[int64]map[int64]float64),
		make(map[int64]Node),
	}
}

// Dump ...
func (g *Graph) Dump(printer func(Node) string) string {
	r := []string{}
	for fromID, edges := range g.adjacency {
		j := 0
		m := fmt.Sprintf("%s: [", printer(g.nodes[fromID]))
		for toID := range edges {
			if j > 0 {
				m += ", "
			}
			m += printer(g.nodes[toID])
			j++
		}
		m += "],"
		r = append(r, m)
	}
	return "{\n" + strings.Join(r, "\n") + "\n}"
}

// AddNode ...
func (g *Graph) AddNode(n Node) {
	id := n.ID()
	if g.adjacency[id] != nil {
		return
	}
	g.adjacency[id] = make(map[int64]float64)
	g.nodes[id] = n
}

// AddEdge ...
func (g *Graph) AddEdge(a, b Node, weight float64) {
	g.AddNode(a)
	g.AddNode(b)
	aID := a.ID()
	bID := b.ID()
	g.adjacency[aID][bID] = weight
}

type nodePath struct {
	id   int64
	path *nodePath
}

type nodeQueueItem struct {
	id   int64
	cost float64
	path *nodePath
}

type nodePriorityQueue struct {
	min   bool
	queue []*nodeQueueItem
}

func newPriorityQueue(min bool) *nodePriorityQueue {
	return &nodePriorityQueue{min, []*nodeQueueItem{}}
}

func (q nodePriorityQueue) Len() int {
	return len(q.queue)
}

func (q nodePriorityQueue) Less(i, j int) bool {
	if q.min {
		return q.queue[i].cost < q.queue[j].cost
	}
	return q.queue[i].cost > q.queue[j].cost
}

func (q nodePriorityQueue) Swap(i, j int) {
	qq := q.queue
	qq[i], qq[j] = qq[j], qq[i]
}

func (q *nodePriorityQueue) Push(x interface{}) {
	elem := x.(*nodeQueueItem)
	q.queue = append(q.queue, elem)
}

func (q *nodePriorityQueue) Pop() interface{} {
	n := len(q.queue)
	elem := q.queue[n-1]
	q.queue = q.queue[0 : n-1]
	return elem
}

func resolvePath(path *nodePath) []int64 {
	r := []int64{}
	for path != nil {
		r = append([]int64{path.id}, r...)
		path = path.path
	}
	return r
}

func dumpQueue(q *nodePriorityQueue) {
	for _, n := range q.queue {
		fmt.Printf("%d[%.0f] ", n.id, n.cost)
	}
	fmt.Println()
}

// DijkstraShortestPath returns the shortest (lowest-cost) path between
// the start and end nodes in the graph.
func (g *Graph) DijkstraShortestPath(start, end Node) []Node {
	queue := newPriorityQueue(true)
	queue.Push(&nodeQueueItem{start.ID(), 0, nil})
	seen := make(map[int64]*struct{})
	res := []Node{}
	for queue.Len() > 0 {
		curr := heap.Pop(queue).(*nodeQueueItem)
		id := curr.id
		if seen[id] != nil {
			continue
		}

		seen[id] = &struct{}{}
		path := curr.path

		// Reached the end node. Resolve and return the result.
		if id == end.ID() {
			for _, nodeID := range append(resolvePath(path), end.ID()) {
				res = append(res, g.nodes[nodeID])
			}
			break
		}

		path = &nodePath{curr.id, curr.path}
		for destID, destCost := range g.adjacency[id] {
			if seen[destID] == nil {
				heap.Push(queue, &nodeQueueItem{destID, curr.cost + destCost, path})
			}
		}
	}
	return res
}
