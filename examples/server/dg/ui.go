package dg

import (
	"github.com/JKhawaja/fabric"
)

// Node ...
type Node struct {
	Id               int
	Type             fabric.NodeType
	Signalers        *fabric.SignalingMap
	AccessProcedures *fabric.ProcedureList
	Signals          *fabric.SignalsMap
	IsRoot           bool
	IsLeaf           bool
}

// UI ...
type UI struct {
	Node
	CDS     *fabric.Section
	Unique  bool
	Virtual bool
}

// ID ...
func (u UI) ID() int {
	return u.Id
}

// GetType ...
func (u UI) GetType() fabric.NodeType {
	return u.Type
}

// GetPriority ...
func (u UI) GetPriority() int {
	return 1
}

// ListProcedures ...
func (u UI) ListProcedures() fabric.ProcedureList {
	p := *u.AccessProcedures
	return p
}

// ListSignals ...
func (u UI) ListSignals() fabric.SignalsMap {
	return *u.Signals
}

// ListSignalers ...
func (u UI) ListSignalers() fabric.SignalingMap {
	return *u.Signalers
}

// UpdateSignaling ...
func (u UI) UpdateSignaling(sm fabric.SignalingMap, s fabric.SignalsMap) {
	*u.Signalers = sm
	*u.Signals = s
}

// Signal ...
func (u UI) Signal(s fabric.ProcedureSignals) {
	sm := *u.Signalers

	for _, c := range sm {
		c <- s
	}
}

// GetSection ...
func (u UI) GetSection() *fabric.Section {
	return u.CDS
}

// IsUnique ...
func (u UI) IsUnique() bool {
	return u.Unique
}

// IsVirtual ...
func (u UI) IsVirtual() bool {
	return u.Virtual
}
