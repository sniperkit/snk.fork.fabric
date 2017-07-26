package fabric

// AccessType is the interface to define how an access procedure should
// behave; Create an Access Procedure function signature type and add
// these methods to it.
// IMPORTANT: In order to best utilize the commit and rollback features
// the Access Type should have an error return value.
type AccessType interface {
	// Class() string                             // the "class" of action (e.g. "read")
	ID() int                                   // integer id assigned to Access Type
	Priority() int                             // priorities are not a necessity but can be helpful for ordering algorithms for posets
	Commit(*DGNode) error                      // takes a DGNode to signal for ...
	Rollback(RestoreNodes, RestoreEdges) error // takes a list of all CDS nodes and edges that have been operated on
	// InvariantNode(*Node) bool                  // used to calculate if a CDS node should remain invariant
	// InvariantEdge(*Edge) bool                  // used to calculate if a CDS edge should remain invariant
}

// RestoreNodes is a list of Node values that can be used to overwrite existing
// Node values after an operation failure.
type RestoreNodes []Node

// RestoreEdges is a list of Edge values that can be used to overwrite existing
// Edge values after an operation failure.
type RestoreEdges []Edge

// IDEA:
// 	create a function that accepts a Section as an argument
// 	and uses InvariantNode() and InvariantEdge() on all objects in the section
// 	then returns a list of invariant nodes and edges for an Access Procedure
// 	on a section.
// 	`InvariantSets(s Section) (NodeList, EdgeList)`

// ProcedureList ...
type ProcedureList []AccessType
