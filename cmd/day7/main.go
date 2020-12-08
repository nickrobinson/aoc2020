package main

import (
	"bufio"
	"os"
	"regexp"

	"github.com/nickrobinson/aoc2020/pkg/baggage"
	log "github.com/sirupsen/logrus"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/traverse"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	fp, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(fp)
	g := baggage.NewBaggageGraph()

	outerBagRe, _ := regexp.Compile(`^(\w+\s\w+) bag`)
	innerBagsRe, _ := regexp.Compile(`(\d) (\w+\s\w+) bag`)
	for scanner.Scan() {
		lineText := scanner.Text()
		outerBagColor := outerBagRe.FindAllStringSubmatch(lineText, -1)[0][1]
		g.AddNode(outerBagColor)

		innerBagColors := innerBagsRe.FindAllStringSubmatch(lineText, -1)
		if len(innerBagColors) > 0 {
			for _, c := range innerBagColors {
				log.Info(c[2])
				g.AddNode(c[2])
				log.Infof("Adding edge %s->%s", outerBagColor, c[2])
				g.AddEdge(c[2], outerBagColor)
			}
		}
	}

	bagCount := 0
	w := traverse.BreadthFirst{
		Traverse: func(g graph.Edge) bool {
			return true
		},
		Visit: func(n graph.Node) {
			log.Infof("Visiting %v", n)
			bagCount++
		},
	}
	w.Walk(g, g.NodeFor("shiny gold"), func(n graph.Node, d int) bool {
		return false
	})
	log.Infof("Found %d bags", bagCount-1)
}
