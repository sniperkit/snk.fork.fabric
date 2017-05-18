package fabric

// AccessType is the interface to define how an access procedure should
// behave; Create an Access Procedure function signature type and add
// these methods to it.
type AccessType interface {
	Name() string                       // the "class" of action (e.g. "read")
	Priority() int                      // priorities are not a necessity but can be helpful
	Commit() error                      // acidic transaction primitive, define how
	Rollback() error                    // acidic transaction primitive
	InvariantNodes(s *Section) NodeList // used to calculate which nodes will be invariant
	InvariantEdges(s *Section) EdgeList // used to calculate which edges will be invariant
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

	func (p *Procedure) Commit() {
		// commit sub-routine; when an operation completes
	}

	func (p *Procedure) Rollback() {
		// rollback sub-routine; when
	}

	func (p *Procedure) InvariantNodes(s *Section) NodeList{
		// calculate invariant nodes for a section
	}

	func (p *Procedure) InvariantEdges(s *Section) EdgeList{
		// calculate invariant edges for a section
	}

	// Now this is where we can create as many objects of Type
	// procedure as we want, and those objects will satisfy both
	// the procedure type we want and the AccessType Interface.
	// So we can use two different instantions of a procedure in
	// two different DG nodes.

*/
