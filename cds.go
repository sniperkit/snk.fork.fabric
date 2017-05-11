package fabric

// TODO: We still need a way to have the information for the ENTIRE CDS in order to create
// 		objects that are section types of a CDS.

// TODO: create a map between elements of a CDS and Nodes (?)

// So we will want extensional lists of nodes per UI,
// and thus, we could potentially wrap our original CDS elements to be Nodes
// and wrap our CDS itself to match the CDS interface.

// Example of a CDS: https://golang.org/src/container/list/list.go

type Node interface {
	ID() int // assigns an ID to a node
}

type CDS interface {
	ListNodes() []Node
	ListEdges() map[Node][]Node
}
