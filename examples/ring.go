package main

// NOTE: this is based on a doubly-linked ring Go package:
// 		https://golang.org/src/container/list/list.go

import (
	"log"

	"github.com/JKhawaja/fabric"
)

// FIXME: How will we address the dynamics of the ring structure, since it is very likely
//		that nodes in the structure will be added and removed very often.

//		One possible solution: could be that we assign the entire ring to a UI, then we can
//		assign each structural node to its own VUI. This will note the fact that any given
//		structural node can be temporary relative to the overall data structure.

//		The other idea: is that ...

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

// ring satisfies fabric CDS interface
type Ring struct {
	Root  Element
	Len   int
	Nodes fabric.NodeList
	Edges fabric.EdgeList
}

func NewRing() *Ring {
	return &Ring{}
}

func (r *Ring) CreateNodes() error {
	var nl fabric.NodeList
	// TODO: traverse list and wrap each element as an elementNode,
	//		add edges to ElementNode with Ids
	// 		and return elementNode slice
	return nl, nil
}

func (r *Ring) CreateEdges() error {
	var el fabric.EdgeList
	// TODO: traverse NodeList add each Nodes edge to edge slice
	//		check that edge slice does not already contain edge ID
	// 		return edge lsit.

	return el, nil
}

func (r *Ring) ListNodes() fabric.NodeList {
	return r.Nodes
}

func (r *Ring) ListEdges() fabric.EdgeList {
	return r.Edges
}

// TODO: Will need an initialization function that creates all
//		dependency graph permanent node assignments and verifications.

func main() {

	myRing := NewRing()

	err := myRing.CreateNodes()
	if err != nil {
		log.Printf("Error while traversing CDS and creating node objects: ", err)
	}

	err = myRing.CreateEdges(nodes)
	if err != nil {
		log.Printf("Error while adding edges to Edges Map: ", err)
	}

	nodes := myRing.ListNodes()
	edges := myRing.ListEdges()

	log.Println(nodes)
	log.Println(edges)
}
