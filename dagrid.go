package dagrid

type node struct {
	contents string
	index    int
	children map[int]struct{}
	parents  map[int]struct{}
}

type edge struct {
	index int
	start int
	end   int
}

type dag struct {
	roots  map[int]struct{}
	leaves map[int]struct{}
	nodes  []node
	edges  []edge
}

func set_insert(set map[int]struct{}, elem int) {
	set[elem] = struct{}{}
}

// func (dag dag) insertNode(node node) {
// 	index := len(dag.nodes)
// 	node.index = index

// 	dag.nodes = append(dag.nodes, node)
// }

func new_dag(starting_node node) dag {
	starting_node.index = 0
	return dag{
		roots:  make(map[int]struct{}),
		leaves: make(map[int]struct{}),
		nodes:  []node{starting_node},
		edges:  make([]edge, 10),
	}
}

func (dag dag) insert_child(child node, parent node) {
	child.index = len(dag.nodes)

	dag.nodes = append(dag.nodes, child)
	set_insert(dag.leaves, child.index)
	delete(dag.leaves, parent.index)
}
