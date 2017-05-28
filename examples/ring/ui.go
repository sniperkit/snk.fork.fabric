package ring

import (
	"github.com/JKhawaja/fabric"
)

type RingUI struct {
	DGNode
	Section fabric.Section
	Unique  bool
	Virtual bool
}

// NOTE: Our arguments ...
//		- Graph that node will be a part of
//		- CDS section that node will have access too
//		- Strict-Uniqueness boolean, and VUI boolean
//		- List of usable access procedures
// AUTOGENERATED: function signature can be autogenerated as well ...
func NewRingUI(g *fabric.Graph, s fabric.Section, u, v bool, pl fabric.ProcedureList) (*RingUI, error) {

	// FIRST: Get ID (init)
	var R RingUI // Auto-generated
	// `id := g.GenID()` can be autogenerated as well
	R.Id = g.GenID()
	R.Type = g.Type(R)
	R.AccessProcedures = pl

	// SECOND: Add to graph
	err := g.AddRealNode(R) // Auto-generated
	if err != nil {
		return &R, err
	}

	// THIRD: Set up data
	R.Signalers = g.CreateSignalers(R)
	R.Signals = g.Signals(R)
	R.IsRoot = g.IsRootBoundary(R)
	R.IsLeaf = g.IsLeafBoundary(R)
	R.Dependents = g.Dependents(R)
	R.Dependencies = g.Dependencies(R)
	R.Section = s
	R.Unique = u
	R.Virtual = v

	return &R, nil
}

func (r RingUI) ID() int {
	return r.Id
}

func (r RingUI) GetType() fabric.NodeType {
	return r.Type
}

func (r RingUI) GetPriority() int {
	// EXAMPLE: could calculate priority based on
	// priorities assigned to procedures in procedures list ...
	// Priority is used to
	p := len(r.AccessProcedures)
	return p
}

func (r RingUI) ListProcedures() fabric.ProcedureList {
	return r.AccessProcedures
}

func (r RingUI) ListDependents() []fabric.DGNode {
	return r.Dependents
}

func (r RingUI) ListDependencies() []fabric.DGNode {
	return r.Dependencies
}

func (r RingUI) ListSignals() fabric.SignalsMap {
	return r.Signals
}

func (r RingUI) ListSignalers() fabric.SignalingMap {
	return r.Signalers
}

func (r RingUI) Signal(s fabric.Signal) {
	for _, c := range r.Signalers {
		c <- s
	}
}

func (r RingUI) GetSection() fabric.Section {
	return r.Section
}

func (r RingUI) IsUnique() bool {
	return r.Unique
}

func (r RingUI) IsVirtual() bool {
	return r.Virtual
}
