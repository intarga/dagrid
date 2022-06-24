package dagrid

type node struct {
	contents string
	index    int
	children []int
	parents  []int
}

type edge struct {
	index int
	start int
	end   int
}

type dag struct {
	roots  []int
	leaves []int
	nodes  []node
	edges  []edge
}

// func (dag dag) insertNode(node node) {
// 	index := len(dag.nodes)
// 	node.index = index

// 	dag.nodes = append(dag.nodes, node)
// }

func new_dag(starting_node node) dag {
	starting_node.index = 0
	return dag{
		roots:  []int{0},
		leaves: []int{0},
		nodes:  []node{starting_node},
		edges:  make([]edge, 10),
	}
}

func (dag dag) insert_child(child node, parent node) {
	child.index = len(dag.nodes)

	dag.nodes = append(dag.nodes, child)
	dag.leaves = append(dag.leaves, child.index)
	//TODO
}
