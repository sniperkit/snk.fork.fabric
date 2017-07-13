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

type ProcedureList []AccessType

/*

	EXAMPLE:

	type Procedure func(node) error

	func (p *Procedure) Class() string {
		// return class name
		return "Read"
	}

	func (p *Procedure) Priority() int {
		// calculate priority
	}

	func (p *Procedure) Commit(np *fabric.DGNode) error {
		// commit sub-routine; when an operation completes
		n.
		return nil
	}

	func (p *Procedure) Rollback(np *fabric.DGNode) error {
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

/*
	EXAMPLE:

	Inside a Thread:

	go func() {
		// Create an access Procedure variable

		var myProcedure ProcedureType = func(node) (node, error){
			// do stuff ...
		}

		// Then we can call methods with that procedure
		mrProcedure.Commit()
	}()


	Outside of a Thread:

	func MyProcedure(*node, int) error {
		// do stuff ...
	}

	go func() {
		// convert to access type object
		p := ProcedureType(MyProcedure)
		// convert to fiber interface object
		i := fiber.AcccessType(p)

		// call access type method
		i.Commit()
		i.Rollback()
	}()
*/
