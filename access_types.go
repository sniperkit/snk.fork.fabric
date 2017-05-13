package fabric

// NOTE: good article on function types: http://jordanorelli.com/post/42369331748/function-types-in-go-golang

// TODO: determine if this is a complete function signature
//		for an AccessProcedure.
type AccessProcedure func(DGNode)
