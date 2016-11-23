package euler

// IntNode ...
type IntNode struct {
	id    int64
	value int
}

// NewIntNode ...
func NewIntNode(id int64, value int) *IntNode {
	return &IntNode{id, value}
}

// ID ...
func (n *IntNode) ID() int64 {
	return n.id
}

// Value ...
func (n *IntNode) Value() int {
	return n.value
}

// IntMatrixToGraph converts a matrix to a graph
func IntMatrixToGraph(m [][]int, directions [][]int) *Graph {
	rows, cols := len(m), len(m[0])
	g := NewGraph()
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			id := int64((r * cols) + c)
			n := &IntNode{id, m[r][c]}

			// Link all the neighboring nodes in allowed directions
			for _, d := range directions {
				rr := r + d[0]
				cc := c + d[1]
				if rr >= 0 && rr < rows && cc >= 0 && cc < cols {
					id = int64((rr * cols) + cc)
					wt := m[rr][cc]
					g.AddEdge(n, &IntNode{id, wt}, float64(wt))
				}
			}
		}
	}
	return g
}
