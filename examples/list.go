package main

// NOTE: this is based on a Go package:
// 		https://golang.org/src/container/list/list.go

import (
	"log"

	"github.com/JKhawaja/fabric"
)

type Element struct {
	next, prev *Element
	list       *List
	Value      interface{}
}

// elementNode satisfies fabric Node interface
type ElementNode struct {
	Element   // NOTE: if an Element struct did not contain information on edges, we could have put an Edges field in this Node struct
	Id        int
	Immutable bool
}

func (e *ElementNode) ID() int {
	return e.Id
}

// list satisfies fabric CDS interface
type List struct {
	Root Element
	Len  int
}

func NewList() *List {
	return &List{}
}

func (l *List) ListNodes() (fabric.NodeList, error) {
	var nl fabric.NodeList
	// TODO: traverse list and wrap each element as an elementNode,
	// and return elementNode slice
	return nl, nil
}

func (l *List) ListEdges(nodes fabric.NodeList) (fabric.EdgeList, error) {
	var el fabric.EdgeList
	// TODO: traverse NodeList add each Node as a Key, and its next
	// and previous elements in []int slice

	return el, nil
}

func main() {

	var myList List

	nodes, err := myList.ListNodes()
	if err != nil {
		log.Printf("Error while traversing CDS and creating node objects: ", err)
	}

	edges, err := myList.ListEdges(nodes)
	if err != nil {
		log.Printf("Error while adding edges to Edges Map: ", err)
	}

	log.Println(edges)

	// TODO: now we can add nodes and edges as needed to a UI
	//		object.
}
