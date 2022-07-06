package dagrid

type node struct {
	contents string
	index    int
	children map[int]struct{}
	parents  map[int]struct{}
}

type dag struct {
	roots  map[int]struct{}
	leaves map[int]struct{}
	nodes  []node
}

func set_insert(set map[int]struct{}, elem int) {
	set[elem] = struct{}{}
}

func new_dag(starting_node node) dag {
	starting_node.index = 0
	return dag{
		roots:  map[int]struct{}{0: {}},
		leaves: map[int]struct{}{0: {}},
		nodes:  []node{starting_node},
	}
}

func (dag dag) insert_child(child node, parent node) {
	child.index = len(dag.nodes)

	dag.nodes = append(dag.nodes, child)
	set_insert(dag.leaves, child.index)
	delete(dag.leaves, parent.index)

	set_insert(child.parents, parent.index)
	set_insert(parent.children, child.index)
}
