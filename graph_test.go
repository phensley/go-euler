package euler

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type N struct {
	id   int64
	name string
}

func (n *N) ID() int64 {
	return n.id
}

func node(id int64, name string) *N {
	return &N{id, name}
}

func TestPriorityQueue(t *testing.T) {
	q := newPriorityQueue(true)
	heap.Push(q, &nodeQueueItem{1, 100, nil})
	heap.Push(q, &nodeQueueItem{2, 0, nil})
	heap.Push(q, &nodeQueueItem{3, 0, nil})
	heap.Push(q, &nodeQueueItem{4, 100, nil})
	heap.Push(q, &nodeQueueItem{5, 0, nil})
	heap.Push(q, &nodeQueueItem{6, 0, nil})
	heap.Push(q, &nodeQueueItem{7, 100, nil})
	heap.Push(q, &nodeQueueItem{8, 22, nil})
	heap.Push(q, &nodeQueueItem{9, 33, nil})
	heap.Push(q, &nodeQueueItem{10, 44, nil})
	heap.Push(q, &nodeQueueItem{11, 55, nil})
	heap.Push(q, &nodeQueueItem{12, 66, nil})
	heap.Push(q, &nodeQueueItem{13, 77, nil})
	heap.Push(q, &nodeQueueItem{14, 200, nil})
	heap.Push(q, &nodeQueueItem{15, 333, nil})
	heap.Push(q, &nodeQueueItem{16, 444, nil})
	for q.Len() > 0 {
		n := heap.Pop(q)
		nn := n.(*nodeQueueItem)
		fmt.Println(nn.id, nn.cost)
	}
}

func TestGraphDijkstraShortestPath(t *testing.T) {
	g := NewGraph()

	nodes := map[string]*N{
		"s": node(1, "s"),
		"u": node(2, "u"),
		"v": node(3, "v"),
		"x": node(4, "x"),
		"y": node(5, "y"),
	}

	for _, node := range nodes {
		g.AddNode(node)
	}

	n := nodes["s"]
	g.AddEdge(n, nodes["u"], 10)
	g.AddEdge(n, nodes["x"], 5)

	n = nodes["u"]
	g.AddEdge(n, nodes["v"], 1)
	g.AddEdge(n, nodes["x"], 2)

	g.AddEdge(nodes["v"], nodes["y"], 4)

	n = nodes["x"]
	g.AddEdge(n, nodes["u"], 3)
	g.AddEdge(n, nodes["v"], 9)
	g.AddEdge(n, nodes["y"], 2)

	n = nodes["y"]
	g.AddEdge(n, nodes["s"], 7)
	g.AddEdge(n, nodes["v"], 6)

	path := g.DijkstraShortestPath(nodes["x"], nodes["v"])
	assert.Equal(t, 3, len(path))
	assert.Equal(t, "x", path[0].(*N).name)
	assert.Equal(t, "u", path[1].(*N).name)
	assert.Equal(t, "v", path[2].(*N).name)
}
