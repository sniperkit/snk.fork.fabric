package fabric

type VUI interface {
	UI
	Dependents() []VUI
}

// TODO: VUIs can be part of VUI Dependency Graphs
//		but each VUI *must* have a lifespan shorter than its dependents.
//		A VUI can have both real and virtual dependents and it can
// 		have both virtual and real dependencies. The key is that all
//		of its virtual dependencies must have lifespans shorter than
//		it's lifespan.

// TODO: Check a VUIs dependents list does not get shorter over the course
//		of its lifespan.
