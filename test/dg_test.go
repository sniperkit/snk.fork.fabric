package fabric_test

import (
	"testing"

	"github.com/JKhawaja/fabric"
)

// FIXME: Need to design an example for how to store dependents and dependencies
// 	as well as Signalers and Signals, and the list of access procedures outside
// 	of the struct definition.
type Node struct {
	Id   int
	Type fabric.NodeType
	// Signalers        fabric.SignalingMap
	// AccessProcedures fabric.ProcedureList
	// Dependents   []fabric.DGNode
	// Dependencies []fabric.DGNode
	// Signals          fabric.SignalsMap
	IsRoot bool
	IsLeaf bool
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
	var p fabric.ProcedureList
	return p
}

func (u UI) ListDependents() []fabric.DGNode {
	var d []fabric.DGNode
	return d
}

func (u UI) ListDependencies() []fabric.DGNode {
	var d []fabric.DGNode
	return d
}

func (u UI) ListSignals() fabric.SignalsMap {
	s := make(fabric.SignalsMap)
	return s
}

func (u UI) ListSignalers() fabric.SignalingMap {
	s := make(fabric.SignalingMap)
	return s
}

func (u UI) UpdateSignaling(sm fabric.SignalingMap, s fabric.SignalsMap) {
	//u.Signalers = sm
	//u.Signals = s
}

func (u UI) Signal(s fabric.ProcedureSignals) {
	/*
		for _, c := range u.Signalers {
			c <- s
		}
	*/
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
	var p fabric.ProcedureList
	return p
}

func (t Temporal) ListDependents() []fabric.DGNode {
	var d []fabric.DGNode
	return d
}

func (t Temporal) ListDependencies() []fabric.DGNode {
	var d []fabric.DGNode
	return d
}

func (t Temporal) ListSignals() fabric.SignalsMap {
	s := make(fabric.SignalsMap)
	return s
}

func (t Temporal) ListSignalers() fabric.SignalingMap {
	s := make(fabric.SignalingMap)
	return s
}

func (t Temporal) UpdateSignaling(sm fabric.SignalingMap, s fabric.SignalsMap) {
	//t.Signalers = sm
	//t.Signals = s
}

func (t Temporal) Signal(s fabric.ProcedureSignals) {
	/*
		for _, c := range t.Signalers {
			c <- s
		}
	*/
}

func (t Temporal) IsVirtual() bool {
	return t.Virtual
}

// TestDG: tests adding nodes and edges, leaf and root boundary checks,
// and Signalers and Signals checks as well
func TestDG(t *testing.T) {
	// Create New DG Graph (check)
	graph := fabric.NewGraph()

	// Create UI node
	u := UI{
		Node: Node{
			Id:   graph.GenID(),
			Type: fabric.UINode,
			//Signalers: make(fabric.SignalingMap),
			//Signals:   make(fabric.SignalsMap),
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
	temp := Temporal{
		Node: Node{Id: graph.GenID(),
			Type: fabric.TemporalNode,
			//Signalers: make(fabric.SignalingMap),
			//Signals:   make(fabric.SignalsMap),
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
	u2 := UI{
		Node: Node{Id: graph.GenID(),
			Type: fabric.UINode,
			//Signalers: make(fabric.SignalingMap),
			//Signals:   make(fabric.SignalsMap),
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
			if graph.IsRootBoundary(&n) != true {
				t.Fatal("Incorrectly classified a root boundary node.")
			}
		}

		if n.ID() == u2.Id {
			if graph.IsLeafBoundary(&n) != true {
				t.Fatal("Incorrectly classified a leaf boundary node.")
			}

		}
	}

	// Create VUI
	vu := UI{
		Node: Node{Id: graph.GenID(),
			Type: fabric.VUINode,
			//Signalers: make(fabric.SignalingMap),
			//Signals:   make(fabric.SignalsMap),
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
			graph.RemoveVUI(&n)
		}
	}

	for c := range graph.Top {
		if c.ID() == vu.Id {
			t.Fatalf("VUI was not removed!")
		}
	}
}

// TODO: Type() Test

// TODO: CycleDetect() Test

// TODO: TotalityUnique() Test

// TODO: Covered() Test
