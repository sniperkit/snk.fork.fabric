package dg

import (
	"github.com/JKhawaja/fabric"
)

// Node ...
type Node struct {
	Id               int
	Priority         int
	Type             fabric.NodeType
	AccessProcedures *fabric.ProcedureList
	Signalers        *fabric.SignalingMap
	Signals          *fabric.SignalsMap
}

// UI ...
type UI struct {
	Node
	CDS     fabric.Section
	Unique  bool
	Virtual bool
}

// NewUI will return a Fabric UI object
func NewUI(g *fabric.Graph, s fabric.Section) fabric.UI {
	sm1 := make(fabric.SignalingMap)
	s1 := make(fabric.SignalsMap)

	ui := &UI{
		Node: Node{
			Id:        g.GenID(),
			Type:      fabric.UINode,
			Signalers: &sm1,
			Signals:   &s1,
		},
		CDS:     s,
		Virtual: false,
	}

	return ui
}

// NewVUI will return a virtual Fabric UI object
func NewVUI(g *fabric.Graph, s fabric.Section) fabric.UI {
	sm1 := make(fabric.SignalingMap)
	s1 := make(fabric.SignalsMap)

	vui := &UI{
		Node: Node{
			Id:        g.GenID(),
			Type:      fabric.UINode,
			Signalers: &sm1,
			Signals:   &s1,
		},
		CDS:     s,
		Virtual: true,
	}

	return vui
}

// ID ...
func (u *UI) ID() int {
	return u.Id
}

// GetType ...
func (u *UI) GetType() fabric.NodeType {
	return u.Type
}

// GetPriority ...
func (u *UI) GetPriority() int {
	return u.Priority
}

// ListProcedures ...
func (u *UI) ListProcedures() fabric.ProcedureList {
	p := *u.AccessProcedures
	return p
}

// ListSignals ...
func (u *UI) ListSignals() fabric.SignalsMap {
	return *u.Signals
}

// ListSignalers ...
func (u *UI) ListSignalers() fabric.SignalingMap {
	return *u.Signalers
}

// UpdateSignaling ...
func (u *UI) UpdateSignaling(sm fabric.SignalingMap, s fabric.SignalsMap) {
	*u.Signalers = sm
	*u.Signals = s
}

// Signal ...
func (u *UI) Signal(s fabric.NodeSignal) {
	sm := *u.Signalers

	for _, c := range sm {
		c <- s
	}
}

// GetSection ...
func (u *UI) GetSection() fabric.Section {
	return u.CDS
}

// IsUnique ...
func (u *UI) IsUnique() bool {
	return u.Unique
}

// IsVirtual ...
func (u *UI) IsVirtual() bool {
	return u.Virtual
}
