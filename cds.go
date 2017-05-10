package fabric

// TODO: How to generically define a CDS
//		A set of Nodes (each node holds typed data)
//		A set of edges (often pointers e.g. linked lists)
// TODO: Methods for virtualizing a CDS into UIs
//		Partitioning Rules (by regexp for KVs, section cuts (for linear ordered CDS graphs), etc.)
//		Node Subset Definitions (for UIs that will have overlap, generically defining subsets of CDS nodes work best)
//		Sub-graph definitions
// TODO: How to assign CDS sub-graphs to a UI

// CDS Subgraph
type Subgraph interface {
	NodeCount() int
	EdgeCount() int
}
