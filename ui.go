package fabric

/*
	A UI is the generic interface that can be satisfied when
	generating UIs from a CDS.

	It is recommended that if the user wants to assign multiple
	sections to a UI to use the ComposeSections() function.

	This interface definition satisfies both UI and VUI objects.
*/

// UI ...
type UI interface {
	DGNode
	GetSection() Section
	IsUnique() bool  // specifies whether a UI is *strictly* unique or not (a UI will always have totality-uniqueness)
	IsVirtual() bool // specifies whether a UI is virtual or not
}

// EmptyUI can be used when a UI is needed codewise, but the CDS system will not be using
// any spatial virtualization (UI DDAGs).
type EmptyUI struct {
	AccessProcedures *ProcedureList
	Signalers        *SignalingMap
	Signals          *SignalsMap
	CDS              Section
}

// NewEmptyUI ...
func NewEmptyUI() UI {
	sm := make(SignalingMap, 0)
	s := make(SignalsMap, 0)
	p := make(ProcedureList, 0)

	return &EmptyUI{
		Signalers:        &sm,
		Signals:          &s,
		AccessProcedures: &p,
	}
}

// ID ...
func (u *EmptyUI) ID() int {
	return 0
}

// GetType ...
func (u *EmptyUI) GetType() NodeType {
	return UINode
}

// GetPriority ...
func (u *EmptyUI) GetPriority() int {
	return 0
}

// ListProcedures ...
func (u *EmptyUI) ListProcedures() ProcedureList {
	return *u.AccessProcedures
}

// ListSignals ...
func (u *EmptyUI) ListSignals() SignalsMap {
	return *u.Signals
}

// ListSignalers ...
func (u *EmptyUI) ListSignalers() SignalingMap {
	return *u.Signalers
}

// UpdateSignaling ...
func (u *EmptyUI) UpdateSignaling(sm SignalingMap, s SignalsMap) {
	*u.Signalers = sm
	*u.Signals = s
}

// Signal ...
func (u *EmptyUI) Signal(s NodeSignal) {
	sm := *u.Signalers

	for _, c := range sm {
		c <- s
	}
}

// GetSection ...
func (u *EmptyUI) GetSection() Section {
	return u.CDS
}

// IsUnique ...
func (u *EmptyUI) IsUnique() bool {
	return false
}

// IsVirtual ...
func (u *EmptyUI) IsVirtual() bool {
	return false
}

// NOTE: VUIs can be part of VUI Dependency Graphs
//	but each VUI *must* have a lifespan shorter than its dependents.
//	A VUI can have both real and virtual dependents and it can
// 	have both real and virtual dependencies. The key is that all
//	of its virtual dependencies must have lifespans shorter than
//	it's lifespan.
