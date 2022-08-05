package dagrid

type node struct {
	Contents string
	Index    int
	Children map[int]struct{}
	Parents  map[int]struct{}
}

type Dag struct {
	Roots       map[int]struct{}
	Leaves      map[int]struct{}
	Nodes       []node
	IndexLookup map[string]int
}

func set_insert(set map[int]struct{}, elem int) {
	set[elem] = struct{}{}
}

func new_node(contents string, index int) node {
	return node{
		Contents: contents,
		Index:    index,
		Children: make(map[int]struct{}),
		Parents:  make(map[int]struct{}),
	}
}

func New_dag() Dag {
	return Dag{
		Roots:       make(map[int]struct{}),
		Leaves:      make(map[int]struct{}),
		Nodes:       make([]node, 0, 10),
		IndexLookup: make(map[string]int),
	}
}

func (dag *Dag) Insert_free_node(contents string) int {
	node := new_node(contents, len(dag.Nodes))

	dag.Nodes = append(dag.Nodes, node)

	set_insert(dag.Roots, node.Index)
	set_insert(dag.Leaves, node.Index)

	dag.IndexLookup[node.Contents] = node.Index

	return node.Index
}

func (dag *Dag) Insert_child(parent_index int, child_contents string) int {
	child := new_node(child_contents, len(dag.Nodes))

	dag.Nodes = append(dag.Nodes, child)
	set_insert(dag.Leaves, child.Index)
	delete(dag.Leaves, parent_index)

	set_insert(child.Parents, parent_index)
	set_insert(dag.Nodes[parent_index].Children, child.Index)

	dag.IndexLookup[child.Contents] = child.Index

	return child.Index
}

func (dag *Dag) Add_edge(parent_index int, child_index int) {
	set_insert(dag.Nodes[parent_index].Children, child_index)
	set_insert(dag.Nodes[child_index].Parents, parent_index)

	delete(dag.Leaves, parent_index)
	delete(dag.Roots, child_index)
}

func (dag *Dag) Remove_edge(parent_index int, child_index int) {
	delete(dag.Nodes[parent_index].Children, child_index)
	delete(dag.Nodes[child_index].Parents, parent_index)

	if len(dag.Nodes[parent_index].Children) == 0 {
		set_insert(dag.Leaves, parent_index)
	}
	if len(dag.Nodes[child_index].Parents) == 0 {
		set_insert(dag.Roots, child_index)
	}
}

func (dag *Dag) count_edges_iter(current_index int, nodes_visited map[int]struct{}) int {
	edge_count := 0

	for child_index := range dag.Nodes[current_index].Children {
		edge_count += 1
		if _, ok := nodes_visited[child_index]; !ok {
			edge_count += dag.count_edges_iter(child_index, nodes_visited)
		}
	}

	set_insert(nodes_visited, current_index)

	return edge_count
}

func (dag *Dag) Count_edges() int {
	edge_count := 0
	nodes_visited := make(map[int]struct{})

	for root := range dag.Roots {
		edge_count += dag.count_edges_iter(root, nodes_visited)
	}

	return edge_count
}

func (dag *Dag) transitive_reduce_iter(current_index int, ancestors map[int]struct{}) {
	for child_index := range dag.Nodes[current_index].Children {
		for coparent_index := range dag.Nodes[child_index].Parents {
			if _, ok := ancestors[coparent_index]; ok {
				dag.Remove_edge(coparent_index, child_index)
			}
		}
	}

	set_insert(ancestors, current_index)
	for child_index := range dag.Nodes[current_index].Children {
		dag.transitive_reduce_iter(child_index, ancestors)
	}
	delete(ancestors, current_index)
}

func (dag *Dag) Transitive_reduce() {
	for root := range dag.Roots {
		dag.transitive_reduce_iter(root, make(map[int]struct{}))
	}
}

func (dag *Dag) cycle_check_iter(current_index int, ancestors map[int]struct{}) (ok bool) {
	_, ok = ancestors[current_index]
	if ok {
		return false
	}

	set_insert(ancestors, current_index)

	for child_index := range dag.Nodes[current_index].Children {
		if !dag.cycle_check_iter(child_index, ancestors) {
			return false
		}
	}

	delete(ancestors, current_index)

	return true
}

func (dag *Dag) Cycle_check() (ok bool) {
	ancestors := make(map[int]struct{})

	for root := range dag.Roots {
		if !dag.cycle_check_iter(root, ancestors) {
			return false
		}
	}

	return true
}
