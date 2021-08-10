package ch04

import (
	"bufio"
	"io"
	"unicode"
)

type RuneGraph map[rune]map[rune]bool

func CreateRuneGraph(in io.Reader) (RuneGraph, error) {
	graph := make(RuneGraph)
	reader := bufio.NewReader(in)
	var from rune
	for {
		if from == 0 {
			r, n, err := reader.ReadRune()
			if err == io.EOF {
				break
			}
			if err != nil {
				return graph, err
			}
			if r == unicode.ReplacementChar && n == 1 {
				continue
			}
			from = r
		}
		to, n, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return graph, err
		}
		if to == unicode.ReplacementChar && n == 1 {
			continue
		}
		graph.addEdge(from, to)
		from = to
	}
	return graph, nil
}

func (g *RuneGraph) addEdge(from, to rune) {
	edges := (*g)[from]
	if edges == nil {
		edges = make(map[rune]bool)
		(*g)[from] = edges
	}
	edges[to] = true
}
