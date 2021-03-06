package main

import (
	"github.com/bakerolls/adventofcode-2018/day-07/node"
)

type worker struct {
	steps int
	node  *node.Node
}

func (w worker) done(done map[rune]bool) bool {
	return w.node == nil
}

type workers []worker

func (ws workers) done(done map[rune]bool) bool {
	for _, w := range ws {
		// fmt.Println("worker", i, "done", w.done(done), "node", w.node)
		if !w.done(done) {
			return false
		}
	}
	return true
}
