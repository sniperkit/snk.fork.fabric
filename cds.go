package fabric

// Node is used to wrap data structure elements to become generic CDS Nodes
type Node interface {
	ID() int // returns node id
	Immutable() bool
}

// NodeList is a list of references to each node in the CDS
type NodeList []Node

// NOTE: for undirected edges, choice of source and destination nodes
//		are up to the developer.

// Edge ...
type Edge interface {
	ID() int // returns edge id
	GetSource() Node
	GetDestination() Node
	Immutable() bool
}

// EdgeList is a list of references to each edge in the CDS
type EdgeList []Edge

// CDS is the interface definition that must be satisfied for the global shared data structure
type CDS interface {
	GenNodeID() int
	GenEdgeID() int
	ListNodes() NodeList // a simple `return MyCDS.Nodes` will suffice here; once a NodeList has been created
	ListEdges() EdgeList // a simple `return MyCDS.Edges` will suffice here; once an EdgesList has been created
}
