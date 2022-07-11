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

func new_node(contents string, index int) node {
	return node{
		contents: contents,
		index:    index,
		children: make(map[int]struct{}),
		parents:  make(map[int]struct{}),
	}
}

func new_dag(starting_node_contents string) dag {
	return dag{
		roots:  map[int]struct{}{0: {}},
		leaves: map[int]struct{}{0: {}},
		nodes:  []node{new_node(starting_node_contents, 0)},
	}
}

func (dag dag) insert_free_node(contents string) {
	node := new_node(contents, len(dag.nodes))

	dag.nodes = append(dag.nodes, node)

	set_insert(dag.roots, node.index)
	set_insert(dag.leaves, node.index)
}

func (dag dag) insert_child(child_contents string, parent_index int) {
	child := new_node(child_contents, len(dag.nodes))

	dag.nodes = append(dag.nodes, child)
	set_insert(dag.leaves, child.index)
	delete(dag.leaves, parent_index)

	set_insert(child.parents, parent_index)
	set_insert(dag.nodes[parent_index].children, child.index)
}

func (dag dag) transitive_reduce_iter(current_index int, ancestors map[int]struct{}) {
	for child_index := range dag.nodes[current_index].children {
		for coparent_index := range dag.nodes[child_index].parents {
			if _, ok := ancestors[coparent_index]; ok {
				delete(dag.nodes[coparent_index].children, child_index)
				delete(dag.nodes[child_index].children, coparent_index)
			}
		}
	}

	set_insert(ancestors, current_index)
	for child_index := range dag.nodes[current_index].children {
		dag.transitive_reduce_iter(child_index, ancestors)
	}
}

func (dag dag) transitive_reduce() {
	for root := range dag.roots {
		dag.transitive_reduce_iter(root, make(map[int]struct{}))
	}
}
