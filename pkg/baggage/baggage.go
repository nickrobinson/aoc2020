package baggage

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

type BaggageGraph struct {
	ids map[string]int64
	*simple.DirectedGraph
}

type Node struct {
	color string
	id    int64
}

func (n Node) ID() int64      { return n.id }
func (n Node) String() string { return n.color }

func NewBaggageGraph() BaggageGraph {
	return BaggageGraph{
		ids:           make(map[string]int64),
		DirectedGraph: simple.NewDirectedGraph(),
	}
}

func (g *BaggageGraph) AddNode(color string) {
	if _, exists := g.ids[color]; exists {
		return
	}

	// We know the node is not yet in the graph, so we can add it.
	u := g.DirectedGraph.NewNode()
	uid := u.ID()
	u = Node{color: color, id: uid}
	g.DirectedGraph.AddNode(u)
	g.ids[color] = uid
}

func (g *BaggageGraph) AddEdge(fromColor string, toColor string) error {
	fromNode := g.NodeFor(fromColor)
	if fromNode == nil {
		log.Errorf("Could not find node for %s", fromColor)
		return errors.New("No matching node found")
	}
	toNode := g.NodeFor(toColor)
	if toNode == nil {
		log.Errorf("Could not find node for %s", fromColor)
		return errors.New("No matching node found")
	}

	g.DirectedGraph.SetEdge(g.DirectedGraph.NewEdge(fromNode, toNode))

	return nil
}

func (g BaggageGraph) NodeFor(word string) graph.Node {
	id, ok := g.ids[word]
	if !ok {
		return nil
	}
	return g.DirectedGraph.Node(id)
}
