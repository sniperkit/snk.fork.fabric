package fabric

// AccessType is the interface to define how an access procedure should
// behave; Create an Access Procedure function signature type and add
// these methods to it.

// IMPORTANT: AccessTypes should always return an error value as this
//		will allow you to properly utilize the Commit() and Rollback()
//		methods.
type AccessType interface {
	Name() string                      // the "class" of action (e.g. "read")
	Priority() int                     // priorities are not a necessity but can be helpful
	Commit() error                     // acidic transaction primitive, define how
	Rollback() error                   // acidic transaction primitive
	InvariantNodes(s Section) NodeList // used to calculate which nodes will be invariant
	InvariantEdges(s Section) EdgeList // used to calculate which edges will be invariant
}

type ProcedureList []AccessType

/*

	Example:

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

	func (p *Procedure) InvariantNodes(s fabric.Section) fabric.NodeList{
		// calculate invariant nodes for a section
		var nl fabric.NodeList
		return nl
	}

	func (p *Procedure) InvariantEdges(s fabric.Section) fabric.EdgeList{
		// calculate invariant edges for a section
		var el fabric.NodeList
		return el
	}

	// Now this is where we can create as many objects of Type
	// procedure as we want, and those objects will satisfy both
	// the procedure type we want and the AccessType Interface.
	// So we can use two different instantions of a procedure in
	// two different DG nodes.

*/
