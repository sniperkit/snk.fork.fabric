package fabric

// Temporal is what is assigned to Threads beyond the first thread
// assigned to a particular UI.
type Temporal interface {
	DGNode
	GetRoot() UI
	IsVirtual() bool
}

/*

	VIRTUAL SPAWNING (using the Virtual boolean of a Temporal node)

	Virtual spawning will work by being a special temporary node which
	can only have temporal dependencies itself. It is not like VDG nodes.
	It must be assigned to a single UI, and can only be created by temporal
	or UI threads. It must be given a position in the Temporal DAG assigned
	to a specific UI.

	IMPORTANT: The use case/purpose of virtual spawning is to avoid the use
		of cyclic dependencies. For example, if we want a node to run an
		operation and then signal the completion of that operation to its
		dependent. But, then we also want the dependent to finish its
		operation and somehow let its dependency know that it completed its
		operation. INSTEAD, we spawn a virtual node that makes our current
		nodes dependent it's dependency, and thus is allowed to be signaled
		to by our dependent node.

*/
