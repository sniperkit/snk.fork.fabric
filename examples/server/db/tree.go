package db

import (
	"math/rand"
	"time"

	"github.com/JKhawaja/fabric"
)

// Tree ...
type Tree struct {
	Root  *ElementNode
	Nodes fabric.NodeList
	Edges fabric.EdgeList
}

// NewTree ...
func NewTree() *Tree {
	l := List{}
	n := &ElementNode{
		Id: l.GenNodeID(),
		L:  &l,
	}
	l.Root = n
	l.Len = 1

	var i interface{} = n
	in := i.(fabric.Node)
	var nl fabric.NodeList
	nl = append(nl, &in)
	l.Nodes = nl

	el := make(fabric.EdgeList, 0)
	l.Edges = el

	return &l
}

// GenNodeID ...
// Generate an ID for a CDS Node
func (t Tree) GenNodeID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	for _, np := range l.Nodes {
		n := *np
		if n.ID() == id {
			id = l.GenNodeID()
		}
	}

	return id
}

// GenEdgeID ...
// Generate an ID for a CDS Edge
func (t Tree) GenEdgeID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	for _, ep := range t.Edges {
		e := *ep
		if e.ID() == id {
			id = t.GenEdgeID()
		}
	}

	return id

}

// ListNodes ...
func (t Tree) ListNodes() fabric.NodeList {
	return t.Nodes
}

// ListEdges ...
func (t Tree) ListEdges() fabric.EdgeList {
	return t.Edges
}

// TreeNode satisfies fabric.Node interface
type TreeNode struct {
	Id    int
	Value interface{}
	Imm   bool
}

// ID ...
func (t TreeNode) ID() int {
	return t.Id
}

// Immutable ...
func (t TreeNode) Immutable() bool {
	return t.Imm
}

// TreeEdge satisfies the fabric.Edge interface
type TreeEdge struct {
	Id          int
	Source      *ElementNode
	Destination *ElementNode
	Imm         bool
}

// ID ...
func (t TreeEdge) ID() int {
	return t.Id
}

// GetSource ...
func (t TreeEdge) GetSource() *fabric.Node {
	var i interface{} = *t.Source
	in := i.(fabric.Node)
	return &in
}

// GetDestination ...
func (t TreeEdge) GetDestination() *fabric.Node {
	var i interface{} = *t.Destination
	in := i.(fabric.Node)
	return &in
}

// Immutable ...
func (t TreeEdge) Immutable() bool {
	return t.Imm
}
