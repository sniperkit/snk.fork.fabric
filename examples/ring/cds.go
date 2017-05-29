package ring

// NOTE: this is based on a doubly-linked ring Go package:
// 		https://golang.org/src/container/list/list.go

// NOTE: The purpose of this package is to show how you can
// convert an existing data structure package into a fabric-based
// package.

import (
	"github.com/JKhawaja/fabric"
)

// TODO: get rid of this Element struct, and ONLY have node and edge definitions
type Element struct {
	next, prev *Element
	list       *Ring
	Value      interface{}
}

// ElementNode satisfies fabric.Node interface
// NOTE: if an Element struct did not already contain information on edges,
// we could have put an Edges field in this Node struct
type ElementNode struct {
	Element
	Edges []*ElementEdge
	Id    int
	Imm   bool
}

func (e *ElementNode) ID() int {
	return e.Id
}

func (e *ElementNode) Immutable() bool {
	return e.Imm
}

// ElementEdge satisfies the fabric.Edge interface
type ElementEdge struct {
	Id                 int
	SourceElement      *ElementNode
	DestinationElement *ElementNode
	Imm                bool
}

func (e *ElementEdge) ID() int {
	return e.Id
}

func (e *ElementEdge) Source() *ElementNode {
	return e.SourceElement
}

func (e *ElementEdge) Destination() *ElementNode {
	return e.DestinationElement
}

func (e *ElementEdge) Immutable() bool {
	return e.Imm
}

// Ring satisfies the fabric.CDS interface
type Ring struct {
	Root  *ElementNode
	Len   int
	Nodes fabric.NodeList
	Edges fabric.EdgeList
}

func NewRing() *Ring {
	return &Ring{}
}

func (r *Ring) CreateNodes() (fabric.NodeList, error) {
	var nl fabric.NodeList
	// TODO: traverse list and wrap each element as an elementNode,
	// 		use DFS recursion to traverse nodes ...
	//		add edges to ElementNode with Ids
	// 		and return elementNode slice

	return nl, nil
}

func (r *Ring) CreateEdges() (fabric.EdgeList, error) {
	var el fabric.EdgeList
	// TODO: traverse NodeList add each Nodes edge to edge slice
	//		check that edge slice does not already contain edge ID
	// 		return edge list.

	return el, nil
}

func (r *Ring) ListNodes() fabric.NodeList {
	return r.Nodes
}

func (r *Ring) ListEdges() fabric.EdgeList {
	return r.Edges
}

func (r *Ring) Length() int {
	var l int
	// TODO: dfs traverse and add up number of nodes
	return l
}

func (r *Ring) Traverse() {

}
