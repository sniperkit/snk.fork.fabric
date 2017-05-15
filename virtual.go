package fabric

// TODO: figure out the types of virtuals and what each virtual
//		interface will require in terms of UI or Temporal Node
//		assignments, etc. virtual nodes will need a
//		lifecycle / lifespan.

// NOTE: Virtual Dependency Graph Nodes can only have other
//		VDG nodes as dependents or dependencies
type Virtual interface {
	UI
	Dependents() []Virtual
	Dependencies() []Virtual
	ListProcedures() ProceduresList
	IsRoot() bool // specifies whether VDG node is root node or not
}

// NOTE: Virtual Dependency Graphs are always trees with a root node
//		which is a dependency for multiple UIs. It represents that those
//		UI nodes cannot execute their root spatial threads until the VDG
//		is completed. The set of UIs that are dependent on the VDG are
//		the set of UIs that at least one thread node in the VDG accesses.

//		A VDG that only affects one UI (i.e. only attaches to a single UI)
//		represents a virtual *temporal* dependency graph.

//		A VDG that affects more than one UI represents a virtual *spatial*
//		dependency graph.
