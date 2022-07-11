package dagrid

import (
	// "dagrid"
	"testing"
)

func TestTransitiveReduce(t *testing.T) {
	dag := new_dag()

	a := dag.insert_free_node("a")

	b := dag.insert_child(a, "b")
	c := dag.insert_child(a, "c")
	d := dag.insert_child(a, "d")
	e := dag.insert_child(a, "e")

	dag.add_edge(b, d)
	dag.add_edge(c, d)
	dag.add_edge(c, e)
	dag.add_edge(d, e)

	if dag.count_edges() != 8 {
		t.Errorf("dag has %v edges, expected 8 edges", dag.count_edges())
	}

	dag.transitive_reduce()

	if dag.count_edges() != 5 {
		t.Errorf("dag has %v edges, expected 8 edges", dag.count_edges())
	}
}
