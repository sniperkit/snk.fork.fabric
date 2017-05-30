package ring

// Based on: https://golang.org/src/container/list/list.go

import (
	"math/rand"
	"time"

	"github.com/JKhawaja/fabric"
)

// ElementNode satisfies fabric.Node interface
type ElementNode struct {
	Id    int
	Value interface{}
	List  *Ring
	// `Next` is the edge that the element is the source for
	// `Prev` is the edge that the element is the destination for
	Next, Prev *ElementEdge
	Imm        bool
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
	// Create fabric.Node interface
	var e ElementNode
	var ei interface{} = e
	ein := ei.(fabric.Node)

	var nl fabric.NodeList
	nl[0] = &ein

	var el fabric.EdgeList

	return &Ring{
		Root:  &e,
		Len:   1,
		Nodes: nl,
		Edges: el,
	}
}

// Generate an ID for a CDS Node
func (r *Ring) GenNodeID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	for _, np := range r.Nodes {
		n := *np
		if n.ID() == id {
			id = r.GenNodeID()
		}
	}

	return id
}

// Generate an ID for a CDS Edge
func (r *Ring) GenEdgeID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	for _, ep := range r.Edges {
		e := *ep
		if e.ID() == id {
			id = r.GenEdgeID()
		}
	}

	return id

}

func (r *Ring) ListNodes() fabric.NodeList {
	return r.Nodes
}

func (r *Ring) ListEdges() fabric.EdgeList {
	return r.Edges
}
