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

func TestCycleCheck(t *testing.T) {
	good_dag := New_dag()

	a := good_dag.Insert_free_node("a")

	b := good_dag.Insert_child(a, "b")
	c := good_dag.Insert_child(a, "c")

	d := good_dag.Insert_child(b, "d")
	good_dag.Add_edge(c, d)

	if !good_dag.Cycle_check() {
		t.Error("cycle detected in acyclic graph")
	}

	bad_dag := New_dag()

	e := bad_dag.Insert_free_node("e")

	f := bad_dag.Insert_child(e, "f")
	g := bad_dag.Insert_child(e, "g")

	h := bad_dag.Insert_child(f, "h")
	bad_dag.Add_edge(h, g)
	bad_dag.Add_edge(g, f)

	if bad_dag.Cycle_check() {
		t.Error("no cycle detected in cyclic graph")
	}

}
