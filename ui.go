package fabric

/*
	A UI is the generic interface that can be satisfied when
	generating UIs from a CDS.
*/
type UI interface {
	CDS // CDS Access
	NodeCount() int
	EdgeCount() int
}
