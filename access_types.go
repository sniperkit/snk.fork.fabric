package fabric

// NOTE: good article on function types: http://jordanorelli.com/post/42369331748/function-types-in-go-golang

// TODO: determine if this is a complete function signature
//		for an AccessProcedure.

// FIXME:
// An invariant is assigned to an Access Type that is assigned to
// an UI.

// Every Access Procedure has a set of invariants for some UI ...
// This means that an Access Procedure should be "constructed" before
// being assigned to a DGNode

type AccessProcedure func(DGNode)
type ProceduresList []AccessProcedure

func CreateProcedure(s Section) AccessProcedure {
	// TODO: creates a procedure with specified invariants

	return func(d DGNode) {}
}
