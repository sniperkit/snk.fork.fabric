package dg

import (
	"github.com/JKhawaja/fabric"
)

// Virtual ...
type Virtual struct {
	Node
	Executing bool
	Root      bool
	Space     fabric.UI
}

// NewVirtual ...
func NewVirtual(vdg *fabric.VDG, space fabric.UI, pl *fabric.ProcedureList, priority int) fabric.Virtual {
	sm1 := make(fabric.SignalingMap)
	s1 := make(fabric.SignalsMap)
	v := &Virtual{
		Node: Node{
			Id:               vdg.GenID(),
			Type:             fabric.VDGNode,
			Priority:         priority,
			AccessProcedures: pl,
			Signalers:        &sm1,
			Signals:          &s1,
		},
		Executing: false,
		Root:      false,
		Space:     space,
	}

	return v
}

// ID ...
func (v *Virtual) ID() int {
	return v.Id
}

// GetType ...
func (v *Virtual) GetType() fabric.NodeType {
	return v.Type
}

// GetPriority ...
func (v *Virtual) GetPriority() int {
	return 1
}

// ListProcedures ...
func (v *Virtual) ListProcedures() fabric.ProcedureList {
	return *v.AccessProcedures
}

// ListSignals ...
func (v *Virtual) ListSignals() fabric.SignalsMap {
	return *v.Signals
}

// ListSignalers ...
func (v *Virtual) ListSignalers() fabric.SignalingMap {
	return *v.Signalers
}

// UpdateSignaling ...
func (v *Virtual) UpdateSignaling(sm fabric.SignalingMap, s fabric.SignalsMap) {
	*v.Signalers = sm
	*v.Signals = s
}

// Signal ...
func (v *Virtual) Signal(s fabric.NodeSignal) {
	sm := *v.Signalers

	for _, c := range sm {
		c <- s
	}
}

// Start ...
func (v *Virtual) Start() {
	// set started boolean to true
	v.Executing = true

	// signal to dependents that node has started
	proc := v.ListProcedures()
	s := fabric.NodeSignal{
		AccessType: proc[0].ID(),
		Value:      fabric.Started,
		Space:      v.Subspace(),
	}
	v.Signal(s)
}

// Started ...
func (v *Virtual) Started() bool {
	return v.Executing
}

// IsRoot ...
func (v *Virtual) IsRoot() bool {
	return v.Root
}

// Subspace ...
func (v *Virtual) Subspace() fabric.UI {
	return v.Space
}
