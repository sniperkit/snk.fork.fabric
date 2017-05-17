package fabric

/*
	A UI is the generic interface that can be satisfied when
	generating UIs from a CDS.

	It is recommended that if the user wants to assign multiple
	sections to a UI to use the ComposeSections() function.
*/
type UI interface {
	DGNode
	Sections() Section
	Unique() bool // specifies whether a UI is strictly unique or not
}
