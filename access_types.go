package fabric

// Every Access Procedure has a set of invariants for some UI ...
// This means that an Access Procedure should be "constructed" before
// being assigned to a DGNode

// TODO: Access Types can have priorities and can be used to create
//		partial orderings for DGs.
//		Access Types can be a generic thing like "write" type and have
//		a list of access procedures which are "write" procedures. (!!)

type AccessProcedure func(DGNode)
type ProceduresList []AccessProcedure

func CreateProcedure(s Section) AccessProcedure {
	// TODO: creates a procedure with specified invariants

	return func(d DGNode) {}
}
