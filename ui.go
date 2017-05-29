package fabric

/*
	A UI is the generic interface that can be satisfied when
	generating UIs from a CDS.

	It is recommended that if the user wants to assign multiple
	sections to a UI to use the ComposeSections() function.

	This interface definition satisfies both UI and VUI objects.
*/
type UI interface {
	DGNode
	GetSection() *Section
	IsUnique() bool  // specifies whether a UI is *strictly* unique or not (a UI will always have totality-uniqueness)
	IsVirtual() bool // specifies whether a UI is virtual or not
}

// NOTE: VUIs can be part of VUI Dependency Graphs
//		but each VUI *must* have a lifespan shorter than its dependents.
//		A VUI can have both real and virtual dependents and it can
// 		have both real and virtual dependencies. The key is that all
//		of its virtual dependencies must have lifespans shorter than
//		it's lifespan.
