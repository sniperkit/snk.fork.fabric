package dg

import (
	"github.com/JKhawaja/fabric"
)

// Virtual ...
type Virtual struct {
	Node
	Start bool
	Root  bool
	Space *fabric.DGNode
}

// ID ...
func (v Virtual) ID() int {
	return v.Id
}

// GetType ...
func (v Virtual) GetType() fabric.NodeType {
	return v.Type
}

// GetPriority ...
func (v Virtual) GetPriority() int {
	return 1
}

// ListProcedures ...
func (v Virtual) ListProcedures() fabric.ProcedureList {
	p := *v.AccessProcedures
	return p
}

// ListSignals ...
func (v Virtual) ListSignals() fabric.SignalsMap {
	return *v.Signals
}

// ListSignalers ...
func (v Virtual) ListSignalers() fabric.SignalingMap {
	return *v.Signalers
}

// UpdateSignaling ...
func (v Virtual) UpdateSignaling(sm fabric.SignalingMap, s fabric.SignalsMap) {
	*v.Signalers = sm
	*v.Signals = s
}

// Signal ...
func (v Virtual) Signal(s fabric.ProcedureSignals) {
	sm := *v.Signalers

	for _, c := range sm {
		c <- s
	}
}

// Started ...
func (v Virtual) Started() bool {
	return v.Start
}

// IsRoot ...
func (v Virtual) IsRoot() bool {
	return v.Root
}

// Subspace ...
func (v Virtual) Subspace() fabric.UI {
	space := *v.Space
	s := space.(fabric.UI)
	return s
}
