package fabric

// NOTE:
//		It is best to think of Posets as objects that you use to
//		generate graphs from. Instead of specifying every node and edge
//		a poset has some ordering algorithm that auto-orders a list of
//		nodes by creating a graph with the nodes and appropriate edges.

type Poset interface {
	ListNodes() []DGNode
	GenerateGraph([]DGNode) *Graph
	// Order should be a method that determines what dependents and what
	// dependencies to assign a node in a specified graph. It will also
	// add the node to its dependencies dependents list, and the same for
	// its dependents dependencies lists.
	Order(*Graph, DGNode) *Graph
}

// EXAMPLE: Access Type Priority Ordering
//		if a DGNode has an Access type with priority lower than
//		all other Access Types in another DGNode, then it automatically
//		becomes a dependency of that node.
