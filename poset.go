package fabric

// Poset is an object that wraps a dependency graph
type Poset interface {
	// Graph should return a pointer to the graph that our POSET object is "wrapping"
	Graph() *Graph
	// GenerateGraph should take a list of nodes and order them according to the Order() method and return a new Graph
	GenerateGraph([]DGNode) *Graph
	// Order should be a method that determines what dependents and what
	// dependencies to assign a node in the wrapped Graph i.e. it determines
	// what edges to make for the node in the Graph.
	Order(DGNode)
}

// VPoset is an object that wraps a virtual dependency graph
type VPoset interface {
	// VDG should return a pointer to the VDG that our VPOSET object is "wrapping"
	VDG() *VDG
	// GenerateGraph should take a list of nodes and order them according to the Order() method and return a new VDG
	GenerateGraph([]Virtual) *VDG
	// Order should be a method that determines what dependents and what
	// dependencies to assign a node in the wrapped VDG i.e. it determines
	// what edges to make for the node in the VDG.
	Order(Virtual)
}

// EXAMPLE: Access Type Priority Ordering
//		if a DGNode has an Access type with priority lower than
//		all other Access Types in another DGNode, then it automatically
//		becomes a dependency of that node.
