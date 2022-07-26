package dagrid

import "testing"

func TestTransitiveReduce(t *testing.T) {
	dag := New_dag()

	a := dag.Insert_free_node("a")

	b := dag.Insert_child(a, "b")
	c := dag.Insert_child(a, "c")
	d := dag.Insert_child(a, "d")
	e := dag.Insert_child(a, "e")

	dag.Add_edge(b, d)
	dag.Add_edge(c, d)
	dag.Add_edge(c, e)
	dag.Add_edge(d, e)

	if dag.Count_edges() != 8 {
		t.Errorf("dag has %v edges, expected 8 edges", dag.Count_edges())
	}

	dag.Transitive_reduce()

	if dag.Count_edges() != 5 {
		t.Errorf("dag has %v edges, expected 5 edges", dag.Count_edges())
	}
}
