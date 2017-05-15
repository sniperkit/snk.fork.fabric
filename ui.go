package fabric

/*
	A UI is the generic interface that can be satisfied when
	generating UIs from a CDS.
*/
type UI interface {
	DGNode
	Dependents() []UI
	Dependencies() []UI
}

// FIXME: UIs are not necessarily immutable,
// 		although we could have immutable UIs if we wanted to not allow
//		structure modification access types to change that UI.
//		An UI can only be immutable if it is strictly-unique i.e.
//		iff it only accesses nodes and edges that no other UI can access.

// 		One of the main problems here is that if a UI covers an part
//		of the structure that another UI addresses, then the other UI
//		could modify the underlying CDS structure and this UI would
//		then technically be changed.

//		This is why it is better to assign immutability to CDS nodes
//		that way regardless of what UI a node or edges is in, if it is immutable
//		it can never be changed by any access procedure.
