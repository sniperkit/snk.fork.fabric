package ring

// Based on: https://golang.org/src/container/list/list.go

import (
	"github.com/JKhawaja/fabric"
)

// FIXME: get rid of this Element struct, and ONLY have node and edge definitions
type Element struct {
	next, prev *Element
	list       *Ring
	Value      interface{}
}

// ElementNode satisfies fabric.Node interface
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
	Id          int
	Source      *ElementNode
	Destination *ElementNode
	Imm         bool
}

func (e *ElementEdge) ID() int {
	return e.Id
}

func (e *ElementEdge) GetSource() *ElementNode {
	return e.Source
}

func (e *ElementEdge) GetDestination() *ElementNode {
	return e.Destination
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
	var e ElementNode
	// TODO: return a Ring with:
	//			single-element set as Root Node
	//			Length = 1
	//			NodeList containing only reference to Root node
	//			empty Edgelist
	return &Ring{
		Root: &e,
	}
}

func (r *Ring) ListNodes() fabric.NodeList {
	return r.Nodes
}

func (r *Ring) ListEdges() fabric.EdgeList {
	return r.Edges
}
