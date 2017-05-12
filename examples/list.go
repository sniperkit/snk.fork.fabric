package main

// NOTE: this is based on a Go package found here: https://golang.org/src/container/list/list.go

import "github.com/JKhawaja/fabric"

type element struct {
	next, prev *Element
	list       *List
	Value      interface{}
}

// elementNode satisfies fabric Node interface
type elementNode struct {
	element
	Id int
}

func (e *elementNode) ID() int {
	return e.Id
}

// list satisfies fabric CDS interface
type list struct {
	root element
	len  int
}

func NewList() *list {

}

func (l *list) ListNodes() fabric.NodeList {

}

func (l *list) ListEdges() fabric.EdgesMap {

}

func main() {

}
