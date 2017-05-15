package main

// NOTE: this is based on a doubly-linked ring Go package:
// 		https://golang.org/src/container/list/list.go

import (
	"log"

	"github.com/JKhawaja/fabric"
)

type Element struct {
	next, prev *Element
	list       *Ring
	Value      interface{}
}

type ElementEdge struct {
	Id          int
	Source      ElementNode
	Destination ElementNode
	Imm         bool
}

func (e *ElementEdge) ID() {
	return e.ID()
}

func (e *ElementEdge) Source() ElementNode {
	return e.Source()
}

func (e *ElementEdge) Destination() ElementNode {
	return e.Destination()
}

func (e *ElementEdge) Immutable() bool {
	return e.Imm
}

// elementNode satisfies fabric Node interface
type ElementNode struct {
	Element // NOTE: if an Element struct did not contain information on edges, we could have put an Edges field in this Node struct
	Edges   []ElementEdge
	Id      int
	Imm     bool
}

func (e *ElementNode) ID() int {
	return e.Id
}

func (e *ElementNode) Immutable() bool {
	return e.Imm
}

// TODO: create edges in main function differently,
//		will need an edge object with methods

// list satisfies fabric CDS interface
type Ring struct {
	Root Element
	Len  int
}

func NewRing() *Ring {
	return &Ring{}
}

func (r *Ring) ListNodes() (fabric.NodeList, error) {
	var nl fabric.NodeList
	// TODO: traverse list and wrap each element as an elementNode,
	//		add edges to ElementNode with Ids
	// 		and return elementNode slice
	return nl, nil
}

func (r *Ring) ListEdges(nodes fabric.NodeList) (fabric.EdgeList, error) {
	var el fabric.EdgeList
	// TODO: traverse NodeList add each Nodes edge to edge slice
	//		check that edge slice does not already contain edge ID
	// 		return edge lsit.

	return el, nil
}

// TODO: Will need an initialization function that creates all
//		dependency graph permanent node assignments and verifications.

func main() {

	myRing := NewRing()

	nodes, err := myRing.ListNodes()
	if err != nil {
		log.Printf("Error while traversing CDS and creating node objects: ", err)
	}

	edges, err := myRing.ListEdges(nodes)
	if err != nil {
		log.Printf("Error while adding edges to Edges Map: ", err)
	}

	log.Println(edges)

	// TODO: now we can add nodes and edges as needed to a UI
	//		object.
}
