package fabric

// NOTE: Virtual Dependency Graph Nodes can only have other
//		VDG nodes as dependents or dependencies
type Virtual interface {
	DGNode
	IsRoot() bool // specifies whether VDG node is root node or not
	Subspace() UI // the (V)UI that the Virtual node is associated to
}

// NOTE: Virtual Dependency Graphs are always trees with a root node
//		which is associated with multiple (V)UIs.

//		The set of (V)UIs that are associated with the VDG are the
//		set of (V)UIs that at least one node in the VDG accesses.

//		A VDG that only affects one (V)UI represents a virtual
//		*temporal* dependency graph.

//		A VDG that affects more than one (V)UI represents a virtual
//		*spatial* dependency graph.
type VDG interface {
	ID() int
	Root() Virtual
	Space() []UI // Space() lists all (V)UIs that the VDG is associated to
}

// IMPORTANT: VDGs will run independent of all other dependency graphs
//		in other words, a VDG does not have any dependents or dependencies
// 		on any (V)UIs! It simply has associations to (V)UIs. The purpose
// 		of a VDG is to order temporary threads (even if they are
//		associated with different UIs).

/*
	// Example:

	// VDG definitions should ALWAYS encapsulate a graph.
	type MyVDG struct{
		*fabric.Graph
		Id int
		root fabric.Virtual
		Dependents []fabric.UI
	}

	func NewVDG() *MyVDG {
		return &MyVDG{}
	}

	func(m *MyVDG) ID() int{
		return m.Id
	}

	func (m *MyVDG) Root() fabric.Virtual {
		return m.root
	}

	func (m *MyVDG) Space() []fabric.UI{
		return m.Dependents
	}
*/
