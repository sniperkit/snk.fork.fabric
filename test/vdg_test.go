package fabric_test

import (
	"testing"

	"github.com/JKhawaja/fabric"
)

type Virtual struct {
	Node
	Space fabric.UI
	Root  bool
}

func (v Virtual) ID() int {
	return v.Id
}

func (v Virtual) GetType() fabric.NodeType {
	return v.Type
}

func (v Virtual) GetPriority() int {
	return 1
}

func (v Virtual) ListProcedures() fabric.ProcedureList {
	p := *v.AccessProcedures
	return p
}

func (v Virtual) ListSignals() fabric.SignalsMap {
	return *v.Signals
}

func (v Virtual) ListSignalers() fabric.SignalingMap {
	return *v.Signalers
}

func (v Virtual) UpdateSignaling(sm fabric.SignalingMap, s fabric.SignalsMap) {
	*v.Signalers = sm
	*v.Signals = s
}

func (v Virtual) Signal(s fabric.ProcedureSignals) {
	sm := *v.Signalers

	for _, c := range sm {
		c <- s
	}
}

func (v Virtual) IsRoot() bool {
	return v.Root
}

func (v Virtual) Subspace() fabric.UI {
	return v.Space
}

func TestVDG(t *testing.T) {
	// Create Graph
	graph := fabric.NewGraph()
	// Add UI to Graph
	sm1 := make(fabric.SignalingMap)
	s1 := make(fabric.SignalsMap)
	u := UI{
		Node: Node{
			Id:        graph.GenID(),
			Type:      fabric.UINode,
			Signalers: &sm1,
			Signals:   &s1,
		},
		Virtual: false,
	}

	_, err := graph.AddRealNode(u)
	if err != nil {
		t.Fatalf("Could not add Temporal node to graph: %v", err)
	}

	// Create VDG
	vdg, err := fabric.NewVDG(graph)
	if err != nil {
		t.Fatalf("Could not create VDG and add to graph: %v", err)
	}

	// Add first Virtual node to VDG
	sm2 := make(fabric.SignalingMap)
	s2 := make(fabric.SignalsMap)
	var space fabric.UI
	for n := range graph.Top {
		if n.ID() == u.Id {
			space = n.(fabric.UI)
		}
	}
	v := Virtual{
		Node: Node{
			Id:        vdg.GenID(),
			Type:      fabric.UINode,
			Signalers: &sm2,
			Signals:   &s2,
		},
		Space: space,
		Root:  true,
	}
	vp, err := vdg.AddVirtualNode(v)
	if err != nil {
		t.Fatalf("Could not first add Virtual node to VDG: %v", err)
	}

	for n := range vdg.Top {
		if n.ID() == v.Id {
			t.Logf("First virtual node is created")
		}
	}

	// Add second Virtual node to graph
	v2 := Virtual{
		Node: Node{
			Id:        vdg.GenID(),
			Type:      fabric.UINode,
			Signalers: &sm2,
			Signals:   &s2,
		},
		Space: space,
		Root:  false,
	}
	vp2, err := vdg.AddVirtualNode(v2)
	if err != nil {
		t.Fatalf("Could not add second Virtual node to VDG: %v", err)
	}

	for n := range vdg.Top {
		if n.ID() == v2.Id {
			t.Logf("Second virtual node is created")
		}
	}

	// Create edge between two virtual nodes
	for n := range vdg.Top {
		if n.ID() == v2.Id {
			vdg.AddVirtualEdge(v.Id, &n)
		}
	}

	for n, l := range vdg.Top {
		if n.ID() == v.Id {
			li := l[0]
			node := *li
			for n2 := range vdg.Top {
				if n.ID() == v2.Id {
					if node.ID() != n2.ID() {
						t.Fatal("Edge was not properly added")
					}
				}
			}
		}
	}

	// Remove the second virtual node
	err = vdg.RemoveVirtualNode(vp2)
	if err != nil {
		t.Fatalf("Could not remove second Virtual node from VDG: %v", err)
	}

	for n := range vdg.Top {
		if n.ID() == v2.Id {
			t.Fatal("Second Virtual node was not removed")
		}
	}

	for n := range vdg.Top {
		if n.ID() == v.Id {
			t.Logf("First virtual node space: %v", n.Subspace())
		}
	}

	// Remove the first virtual node
	err = vdg.RemoveVirtualNode(vp)
	if err != nil {
		t.Fatalf("Could not remove first Virtual node from VDG: %v", err)
	}

	for n := range vdg.Top {
		if n.ID() == v2.Id {
			t.Fatal("First Virtual node was not removed")
		}
	}
}
