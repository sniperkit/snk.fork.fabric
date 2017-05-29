package fabric

// wrap data structure elements to become generic CDS Nodes
type Node interface {
	ID() int // returns node id
	Immutable() bool
}

// A list of references to each node in the CDS
type NodeList []*Node

// NOTE: for undirected edges, choice of source and destination nodes
//		are up to the developer.
type Edge interface {
	ID() int // returns edge id
	GetSource() *Node
	GetDestination() *Node
	Immutable() bool
}

// A list of references to each edge in the CDS
type EdgeList []*Edge

// add these methods to data structure objects to use as CDS
type CDS interface {
	GenNodeID(*Node) int
	GenEdgeID(*Edge) int
	ListNodes() NodeList // a simple `return MyCDS.Nodes` will suffice here; once a NodeList has been created
	ListEdges() EdgeList // a simple `return MyCDS.Edges` will suffice here; once an EdgesList has been created
}
