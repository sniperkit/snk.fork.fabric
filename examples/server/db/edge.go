package db

import (
	"github.com/JKhawaja/fabric"
)

// ElementEdge satisfies the fabric.Edge interface
type ElementEdge struct {
	Id          int
	L           *List
	Source      *ElementNode
	Destination *ElementNode
	Imm         bool
}

func (e ElementEdge) ID() int {
	return e.Id
}

func (e ElementEdge) GetSource() *fabric.Node {
	var i interface{} = *e.Source
	in := i.(fabric.Node)
	return &in
}

func (e ElementEdge) GetDestination() *fabric.Node {
	var i interface{} = *e.Destination
	in := i.(fabric.Node)
	return &in
}

func (e ElementEdge) Immutable() bool {
	return e.Imm
}
