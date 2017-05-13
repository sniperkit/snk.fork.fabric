package main

// NOTE: this is based on a Go package found here: https://golang.org/src/container/list/list.go

import (
	"log"

	"github.com/JKhawaja/fabric"
)

type element struct {
	next, prev *Element
	list       *List
	Value      interface{}
}

// elementNode satisfies fabric Node interface
type elementNode struct {
	element
	Id int
}

func (e *elementNode) ID() int {
	return e.Id
}

// list satisfies fabric CDS interface
type list struct {
	Root element
	Len  int
}

func NewList() *list {
	return &list{}
}

func (l *list) ListNodes() (fabric.NodeList, error) {
	// TODO: traverse list and wrap each element as an elementNode, and return elementNode slice
}

func (l *list) ListEdges(nodes fabric.NodeList) (fabric.EdgesMap, error) {
	// TODO: traverse NodeList add each Node as a Key, and its next and previous elements in []int slice
}

func main() {

	var myList list

	nodes, err := myList.ListNodes()
	if err != nil {
		log.Printf("Error while traversing CDS and creating node objects: ", err)
	}

	edges, err := myList.ListEdges(nodes)
	if err != nil {
		log.Printf("Error while adding edges to Edges Map: ", err)
	}

	// TODO: now we can add nodes and edges as needed to a UI object

}
