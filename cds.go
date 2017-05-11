package fabric

// Example of a CDS: https://golang.org/src/container/list/list.go

// wrap data structure elements to become generic CDS Nodes
type Node interface {
	ID() int    // returns node id
	NewID() int // assigns a new id to node
}

// add these methods to data structure objects to use as CDS
type CDS interface {
	ListNodes() []Node
	ListEdges() map[Node][]Node
}
