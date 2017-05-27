package ring

import (
	"github.com/JKhawaja/fabric"
)

/* Single-UI Dependency Graph code */

// TODO: This code is about defining what a UI and Temporal object would look like
// these are the objects that are assigned to goroutines/threads in any system
// that we build and use this ring package.

type DGNode struct {
	Id               int
	Signalers        []chan fabric.Signal
	AccessProcedures fabric.ProcedureList
	Signals          []<-chan fabric.Signal
}

type RingUI struct {
	DGNode
	Section fabric.Section
	Unique  bool
	Virtual bool
}

// TODO: Implement the appropriate methods for Temporal Nodes
//		which is all DGNode methods + GetRoot() and IsVirtual()
type RingTemporal struct {
	DGNode
	Root    RingUI
	Virtual bool
}

func NewRingUI(pl fabric.ProcedureList) RingUI {
	// TODO: have to decide how to set up a Ring UI ...
	// Inputs:
	//		Section
	//		Unique & Virtual boolean
	//		ProceduresList

	dg := DGNode{
		AccessProcedures: pl,
	}

	return RingUI{
		DGNode: dg,
	}
}

func (r *RingUI) ID() int {
	return r.Id
}

// TODO: decide if we should determine signals and signal channels from the graph?
func (r *RingUI) ListSignalingChannels(g *fabric.Graph) fabric.SignalingMap {
	sm := make(fabric.SignalingMap)

	var i interface{} = r

	deps := g.Dependents(i.(fabric.DGNode))
	for _, d := range deps {
		c := make(chan fabric.Signal)
		sm[d.ID()] = c
	}

	return sm
}

func (r *RingUI) GetPriority() int {
	// EXAMPLE: could calculate priority based on
	// priorities assigned to procedures in procedures list ...
	// Priority is used to
	p := len(r.AccessProcedures)
	return p
}

func (r *RingUI) ListProcedures() fabric.ProcedureList {
	return r.AccessProcedures
}

// NOTE: the dependents for a UI node can be determined
// 		by the fabric.Graph method: Dependents()
func (r *RingUI) ListDependents(g *fabric.Graph) []fabric.DGNode {
	var i interface{} = r
	return g.Dependents(i.(fabric.DGNode))
}

// NOTE: the dependencies for a UI node can be determined
// 		by the fabric.Graph method: Dependencies()
func (r *RingUI) ListDependencies(g *fabric.Graph) []fabric.DGNode {
	var i interface{} = r
	return g.Dependencies(i.(fabric.DGNode))
}

func (r *RingUI) ListSignals() []<-chan fabric.Signal {
	return r.Signals
}

func (r *RingUI) GetSection() fabric.Section {
	return r.Section
}

func (r *RingUI) IsUnique() bool {
	return r.Unique
}

func (r *RingUI) IsVirtual() bool {
	return r.Virtual
}

// Auto-generated
func (r *RingUI) Type() fabric.NodeType {

	var i interface{} = r
	j, ok := i.(fabric.UI)
	if ok {
		if j.IsVirtual() {
			return fabric.VUINode
		}
		return fabric.UINode
	}
	k, ok := i.(fabric.Temporal)
	if ok {
		if k.IsVirtual() {
			return fabric.VirtualTemporalNode
		}
		return fabric.TemporalNode
	}
	_, ok = i.(fabric.Virtual)
	if ok {
		return fabric.VirtualNode
	}

	return fabric.Unknown
}

// Auto-generated
func (r *RingUI) IsRootBoundary(g *fabric.Graph) bool {
	if len(r.ListDependents(g)) == 0 {
		return true
	}

	return false
}

// Auto-generated
func (r *RingUI) IsLeafBoundary(g *fabric.Graph) bool {
	if len(r.ListDependencies(g)) == 0 {
		return true
	}

	return false
}

// Auto-generated (not in a fabric interface definition)
func (r *RingUI) SignalArrayLength(g *fabric.Graph) int {
	return len(r.ListDependents(g))
}
