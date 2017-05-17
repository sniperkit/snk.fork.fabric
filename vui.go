package fabric

type VUI interface {
	DGNode
	Section() Section
	CheckDependents() bool
}

// TODO: VUIs can be part of VUI Dependency Graphs
//		but each VUI *must* have a lifespan shorter than its dependents.
//		A VUI can have both real and virtual dependents and it can
// 		have both real and virtual dependencies. The key is that all
//		of its virtual dependencies must have lifespans shorter than
//		it's lifespan.

// TODO: Check a VUIs dependents list does not get shorter over the course
//		of its lifespan. If a dependent has ended its lifecycle before
//		a VUI finishes it's lifespan then the VUI needs to abort its
//		operation.
