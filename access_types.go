package fabric

// NOTE: good article on function types: http://jordanorelli.com/post/42369331748/function-types-in-go-golang

// TODO: determine if this is a complete function signature
//		for an AccessProcedure.

// TODO: Determine where we are going to specify invariants for an
//		access procedure. We could probably do it in some sort of
//		structure that is assigned to a depndency graph node.
//		As every dependency graph node has a set of access procedures
//		that it is allowed to perform.
type AccessProcedure func(DGNode)
type ProceduresList []AccessProcedure
