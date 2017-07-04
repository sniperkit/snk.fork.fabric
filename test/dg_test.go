package fabric_test

import (
	"testing"

	"github.com/JKhawaja/fabric"
)

type Node struct {
	Id               int
	Type             fabric.NodeType
	Signalers        *fabric.SignalingMap
	AccessProcedures *fabric.ProcedureList
	Dependents       *[]fabric.DGNode
	Dependencies     *[]fabric.DGNode
	Signals          *fabric.SignalsMap
	IsRoot           bool
	IsLeaf           bool
}

type UI struct {
	Node
	CDS     *fabric.Section
	Unique  bool
	Virtual bool
}

type Temporal struct {
	Node
	UIRoot  UI
	Virtual bool
}

func (u UI) ID() int {
	return u.Id
}

func (u UI) GetType() fabric.NodeType {
	return u.Type
}

func (u UI) GetPriority() int {
	return 1
}

func (u UI) ListProcedures() fabric.ProcedureList {
	p := *u.AccessProcedures
	return p
}

func (u UI) ListSignals() fabric.SignalsMap {
	return *u.Signals
}

func (u UI) ListSignalers() fabric.SignalingMap {
	return *u.Signalers
}

func (u UI) UpdateSignaling(sm fabric.SignalingMap, s fabric.SignalsMap) {
	*u.Signalers = sm
	*u.Signals = s
}

func (u UI) Signal(s fabric.ProcedureSignals) {
	sm := *u.Signalers

	for _, c := range sm {
		c <- s
	}

}

func (u UI) GetSection() *fabric.Section {
	return u.CDS
}

func (u UI) IsUnique() bool {
	return u.Unique
}

func (u UI) IsVirtual() bool {
	return u.Virtual
}

// TEMPORAL

func (t Temporal) ID() int {
	return t.Id
}

func (t Temporal) GetType() fabric.NodeType {
	return t.Type
}

func (t Temporal) GetPriority() int {
	return 1
}

func (t Temporal) ListProcedures() fabric.ProcedureList {
	p := *t.AccessProcedures
	return p
}

func (t Temporal) ListSignals() fabric.SignalsMap {
	s := *t.Signals
	return s
}

func (t Temporal) ListSignalers() fabric.SignalingMap {
	s := *t.Signalers
	return s
}

func (t Temporal) UpdateSignaling(sm fabric.SignalingMap, s fabric.SignalsMap) {
	*t.Signalers = sm
	*t.Signals = s
}

func (t Temporal) Signal(s fabric.ProcedureSignals) {
	sm := *t.Signalers

	for _, c := range sm {
		c <- s
	}
}

func (t Temporal) IsVirtual() bool {
	return t.Virtual
}

// TestDG: tests adding nodes and edges, leaf and root boundary checks,
// and Signalers and Signals checks as well
func TestDG(t *testing.T) {
	// Create New DG Graph (check)
	graph := fabric.NewGraph()
	sm1 := make(fabric.SignalingMap)
	s1 := make(fabric.SignalsMap)
	var d11 []fabric.DGNode
	var d12 []fabric.DGNode

	// Create UI node
	u := UI{
		Node: Node{
			Id:           graph.GenID(),
			Type:         fabric.UINode,
			Signalers:    &sm1,
			Signals:      &s1,
			Dependents:   &d11,
			Dependencies: &d12,
		},
		Virtual: false,
	}

	// Add UI node to graph (check)
	err := graph.AddRealNode(u)
	if err != nil {
		t.Fatalf("Could not add UI node to graph: %v", err)
	}

	// Check that UI node is both a leaf and root boundary
	for n := range graph.Top {
		if n.ID() == u.Id {
			if graph.IsLeafBoundary(&n) != true {
				t.Fatal("Incorrectly classified a leaf boundary node.")
			}
			if graph.IsRootBoundary(&n) != true {
				t.Fatal("Incorrectly classified a root boundary node.")
			}
		}
	}

	// Create Temporal node (for first UI) and add to graph (check)
	sm2 := make(fabric.SignalingMap)
	s2 := make(fabric.SignalsMap)
	var d21 []fabric.DGNode
	var d22 []fabric.DGNode
	temp := Temporal{
		Node: Node{Id: graph.GenID(),
			Type:         fabric.TemporalNode,
			Signalers:    &sm2,
			Signals:      &s2,
			Dependents:   &d21,
			Dependencies: &d22,
		},
		UIRoot: u,
	}

	var t1 interface{} = temp
	err = graph.AddRealNode(t1.(fabric.DGNode))
	if err != nil {
		t.Fatalf("Could not add Temporal node to graph: %v", err)
	}

	// Add edge
	for n := range graph.Top {
		if n.ID() == temp.Id {
			graph.AddRealEdge(u.ID(), &n)
		}
	}

	// Create second UI node, and add to graph
	// Add Edge from first UI node to second UInode
	sm3 := make(fabric.SignalingMap)
	s3 := make(fabric.SignalsMap)
	var d31 []fabric.DGNode
	var d32 []fabric.DGNode
	u2 := UI{
		Node: Node{Id: graph.GenID(),
			Type:         fabric.UINode,
			Signalers:    &sm3,
			Signals:      &s3,
			Dependents:   &d31,
			Dependencies: &d32,
		},
	}

	err = graph.AddRealNode(u2)
	if err != nil {
		t.Fatalf("Could not add UI node to graph: %v", err)
	}

	// Add edge from first UI node to second UI node
	for n := range graph.Top {
		if n.ID() == u2.Id {
			graph.AddRealEdge(u.ID(), &n)
		}
	}

	// Do Leaf and Root Boundary checks again
	for n := range graph.Top {
		if n.ID() == u.Id {
			if !graph.IsRootBoundary(&n) {
				t.Fatal("Incorrectly classified a root boundary node.")
			}
		}

		if n.ID() == u2.Id {
			if !graph.IsLeafBoundary(&n) {
				t.Fatal("Incorrectly classified a leaf boundary node.")
			}
		}
	}

	// Create VUI
	sm4 := make(fabric.SignalingMap)
	s4 := make(fabric.SignalsMap)
	var d41 []fabric.DGNode
	var d42 []fabric.DGNode
	vu := UI{
		Node: Node{Id: graph.GenID(),
			Type:         fabric.VUINode,
			Signalers:    &sm4,
			Signals:      &s4,
			Dependents:   &d41,
			Dependencies: &d42,
		},
		Virtual: true,
	}

	// Add VUI to graph
	err = graph.AddVUI(vu)
	if err != nil {
		t.Fatalf("Could not add VUI node to graph: %v", err)
	}

	// TODO: Check Signalers and Signals for all Nodes

	// Remove VUI from graph (check)
	for n := range graph.Top {
		if n.ID() == vu.Id {
			graph.RemoveVUI(n)
			break
		}
	}

	for c := range graph.Top {
		if c.ID() == vu.Id {
			t.Fatalf("VUI was not removed!")
		}
	}
}

func TestCycleDetect(t *testing.T) {
	// Create New DG Graph (check)
	graph := fabric.NewGraph()

	// Create first UI node
	sm1 := make(fabric.SignalingMap)
	s1 := make(fabric.SignalsMap)
	var d11 []fabric.DGNode
	var d12 []fabric.DGNode
	u1 := UI{
		Node: Node{
			Id:           graph.GenID(),
			Type:         fabric.UINode,
			Signalers:    &sm1,
			Signals:      &s1,
			Dependents:   &d11,
			Dependencies: &d12,
		},
		Virtual: false,
	}

	err := graph.AddRealNode(u1)
	if err != nil {
		t.Fatalf("Could not add UI node to graph: %v", err)
	}

	// Create second UI node
	sm2 := make(fabric.SignalingMap)
	s2 := make(fabric.SignalsMap)
	var d21 []fabric.DGNode
	var d22 []fabric.DGNode
	u2 := UI{
		Node: Node{
			Id:           graph.GenID(),
			Type:         fabric.UINode,
			Signalers:    &sm2,
			Signals:      &s2,
			Dependents:   &d21,
			Dependencies: &d22,
		},
		Virtual: false,
	}

	err = graph.AddRealNode(u2)
	if err != nil {
		t.Fatalf("Could not add UI node to graph: %v", err)
	}

	// Create third UI node
	sm3 := make(fabric.SignalingMap)
	s3 := make(fabric.SignalsMap)
	var d31 []fabric.DGNode
	var d32 []fabric.DGNode
	u3 := UI{
		Node: Node{
			Id:           graph.GenID(),
			Type:         fabric.UINode,
			Signalers:    &sm3,
			Signals:      &s3,
			Dependents:   &d31,
			Dependencies: &d32,
		},
		Virtual: false,
	}

	// Add UI node to graph (check)
	err = graph.AddRealNode(u3)
	if err != nil {
		t.Fatalf("Could not add UI node to graph: %v", err)
	}

	// Add Edges (creates a cycle)
	for n := range graph.Top {
		// from node 1 to node 2
		if n.ID() == u2.Id {
			graph.AddRealEdge(u1.ID(), &n)
		}

		// from node 2 to node 3
		if n.ID() == u3.Id {
			graph.AddRealEdge(u2.ID(), &n)
		}

		// from node 3 to node 1
		if n.ID() == u1.Id {
			graph.AddRealEdge(u3.ID(), &n)
		}
	}

	if !graph.CycleDetect() {
		t.Fatalf("Did not detect cycle in the graph")
	}
}

// TODO: create CDS
// TODO: TotalityUnique() Test
// TODO: Covered() Test
