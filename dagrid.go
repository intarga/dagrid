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

func new_dag() dag {
	return dag{
		roots:  make(map[int]struct{}),
		leaves: make(map[int]struct{}),
		nodes:  make([]node, 0, 10),
	}
}

func (dag dag) insert_free_node(contents string) int {
	node := new_node(contents, len(dag.nodes))

	dag.nodes = append(dag.nodes, node)

	set_insert(dag.roots, node.index)
	set_insert(dag.leaves, node.index)

	return node.index
}

func (dag dag) insert_child(parent_index int, child_contents string) int {
	child := new_node(child_contents, len(dag.nodes))

	dag.nodes = append(dag.nodes, child)
	set_insert(dag.leaves, child.index)
	delete(dag.leaves, parent_index)

	set_insert(child.parents, parent_index)
	set_insert(dag.nodes[parent_index].children, child.index)

	return child.index
}

func (dag dag) add_edge(parent_index int, child_index int) {
	set_insert(dag.nodes[parent_index].children, child_index)
	set_insert(dag.nodes[child_index].children, parent_index)
}

func (dag dag) remove_edge(parent_index int, child_index int) {
	delete(dag.nodes[parent_index].children, child_index)
	delete(dag.nodes[child_index].children, parent_index)
}

func (dag dag) count_edges_iter(current_index int, nodes_visited *map[int]struct{}) int {
	edge_count := 0

	for child_index := range dag.nodes[current_index].children {
		edge_count += 1 + dag.count_edges_iter(child_index, nodes_visited)
	}

	set_insert(*nodes_visited, current_index)

	return edge_count
}

func (dag dag) count_edges() int {
	edge_count := 0
	nodes_visited := make(map[int]struct{})

	for root := range dag.roots {
		edge_count += dag.count_edges_iter(root, &nodes_visited)
	}

	return edge_count
}

func (dag dag) transitive_reduce_iter(current_index int, ancestors map[int]struct{}) {
	for child_index := range dag.nodes[current_index].children {
		for coparent_index := range dag.nodes[child_index].parents {
			if _, ok := ancestors[coparent_index]; ok {
				dag.remove_edge(coparent_index, child_index)
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
