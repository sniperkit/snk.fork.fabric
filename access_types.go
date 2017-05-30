package fabric

// AccessType is the interface to define how an access procedure should
// behave; Create an Access Procedure function signature type and add
// these methods to it.

// IMPORTANT: AccessTypes should always return an error value as this
// 	will allow you to properly utilize the Commit() and Rollback()
// 	methods.
type AccessType interface {
	Name() string             // the "class" of action (e.g. "read")
	Priority() int            // priorities are not a necessity but can be helpful
	Commit() error            // acidic transaction primitive, define how
	Rollback() error          // acidic transaction primitive
	InvariantNode(*Node) bool // used to calculate if a CDS node should remain invariant
	InvariantEdge(*Edge) bool // used to calculate if a CDS edge should remain invariant
}

// IDEA:
// 	create a function that accepts a Section as an argument
// 	and uses InvariantNode() and InvariantEdge() on all objects in the section
// 	then returns a list of invariant nodes and edges for an Access Procedure
// 	on a section.
// 	`InvariantSets(s Section) (NodeList, EdgeList)`

type ProcedureList []AccessType

/*

	EXAMPLE:

	type Procedure func(node) error

	func (p *Procedure) Name() string {
		// return class name
	}

	func (p *Procedure) Priority() int {
		// calculate priority
	}

	func (p *Procedure) Commit() error {
		// commit sub-routine; when an operation completes
		return nil
	}

	func (p *Procedure) Rollback() error {
		// rollback sub-routine; when
		return nil
	}

	func (p *Procedure) InvariantNodes(n fabric.Node) bool{
		var b bool
		return b
	}

	func (p *Procedure) InvariantEdges(e fabric.Edge) bool{
		var b bool
		return b
	}

	// Now this is where we can create as many objects of Type
	// procedure as we want, and those objects will satisfy both
	// the procedure type we want and the AccessType Interface.
	// So we can use two different instantiations of a procedure in
	// two different DG nodes.

*/
