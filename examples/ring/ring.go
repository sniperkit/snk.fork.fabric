package ring

// NOTE: this is based on a doubly-linked ring Go package:
// 		https://golang.org/src/container/list/list.go

import (
	"github.com/JKhawaja/fabric"
)

type Element struct {
	next, prev *Element
	list       *Ring
	Value      interface{}
}

type ElementEdge struct {
	Id                 int
	SourceElement      ElementNode
	DestinationElement ElementNode
	Imm                bool
}

func (e *ElementEdge) ID() int {
	return e.Id
}

func (e *ElementEdge) Source() ElementNode {
	return e.SourceElement
}

func (e *ElementEdge) Destination() ElementNode {
	return e.DestinationElement
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

	r.Nodes = nl

	return nil
}

func (r *Ring) CreateEdges() error {
	var el fabric.EdgeList
	// TODO: traverse NodeList add each Nodes edge to edge slice
	//		check that edge slice does not already contain edge ID
	// 		return edge list.

	r.Edges = el

	return nil
}

func (r *Ring) ListNodes() fabric.NodeList {
	return r.Nodes
}

func (r *Ring) ListEdges() fabric.EdgeList {
	return r.Edges
}
