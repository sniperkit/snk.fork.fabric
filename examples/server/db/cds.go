package db

import (
	"math/rand"
	"time"

	"github.com/JKhawaja/fabric"
)

// List satisfies the fabric.CDS interface
type List struct {
	Root  *ElementNode
	Len   int
	Nodes fabric.NodeList
	Edges fabric.EdgeList
}

func NewList() *List {
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

// NewElementNode will create a new List element node
// NOTE: this should be defined as an access procedure of some access type
func (l *List) NewElementNode() *ElementNode {
	n := ElementNode{
		Id: l.GenNodeID(),
		L:  l,
	}

	var i interface{} = n
	in := i.(fabric.Node)

	l.Nodes = append(l.Nodes, &in)
	l.Len += 1

	return &n
}

// NewElementEdge will create a new List edge
// NOTE: this should be defined as an access procedure of some access type
func (l *List) NewElementEdge(s, d *ElementNode) {
	e := ElementEdge{
		Id:          l.GenEdgeID(),
		L:           l,
		Source:      s,
		Destination: d,
	}

	var i interface{} = e
	ie := i.(fabric.Edge)

	l.Edges = append(l.Edges, &ie)
}

// Generate an ID for a CDS Node
func (l List) GenNodeID() int {
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

// Generate an ID for a CDS Edge
func (l List) GenEdgeID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	for _, ep := range l.Edges {
		e := *ep
		if e.ID() == id {
			id = l.GenEdgeID()
		}
	}

	return id

}

func (l List) ListNodes() fabric.NodeList {
	return l.Nodes
}

func (l List) ListEdges() fabric.EdgeList {
	return l.Edges
}
