package baggage

import (
	"testing"
)

func TestBaggageGraph(t *testing.T) {
	g := NewBaggageGraph()
	g.AddNode("light red")
	g.AddNode("muted green")
	g.AddEdge("light red", "muted green")
	if g.DirectedGraph.HasEdgeFromTo(g.NodeFor("light red").ID(), g.NodeFor("muted green").ID()) != true {
		t.Errorf("Expected to find edge")
	}
}
