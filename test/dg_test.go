package fabric_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/JKhawaja/fabric"
)

type Node struct {
	Id               int
	Type             fabric.NodeType
	Signalers        *fabric.SignalingMap
	AccessProcedures *fabric.ProcedureList
	Signals          *fabric.SignalsMap
	IsRoot           bool
	IsLeaf           bool
}

type UI struct {
	Node
	CDS     fabric.Section
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

func (u UI) Signal(s fabric.NodeSignal) {
	sm := *u.Signalers

	for _, c := range sm {
		c <- s
	}
}

func (u UI) GetSection() fabric.Section {
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

func (t Temporal) Signal(s fabric.NodeSignal) {
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

	// Create UI node
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

	// Add UI node to graph
	_, err := graph.AddRealNode(u)
	if err != nil {
		t.Fatalf("Could not add UI node to graph: %v", err)
	}

	// Check that UI node is both a leaf and root boundary
	for n := range graph.Top {
		if n.ID() == u.Id {
			if graph.IsLeafBoundary(n) != true {
				t.Fatal("Incorrectly classified a leaf boundary node.")
			}
			if graph.IsRootBoundary(n) != true {
				t.Fatal("Incorrectly classified a root boundary node.")
			}
		}
	}

	// Create Temporal node (for first UI) and add to graph (check)
	sm2 := make(fabric.SignalingMap)
	s2 := make(fabric.SignalsMap)
	temp := Temporal{
		Node: Node{Id: graph.GenID(),
			Type:      fabric.TemporalNode,
			Signalers: &sm2,
			Signals:   &s2,
		},
		UIRoot: u,
	}

	tempp, err := graph.AddRealNode(temp)
	if err != nil {
		t.Fatalf("Could not add Temporal node to graph: %v", err)
	}

	// Add edge
	graph.AddRealEdge(u.ID(), tempp)

	// Create second UI node, and add to graph
	sm3 := make(fabric.SignalingMap)
	s3 := make(fabric.SignalsMap)
	u2 := UI{
		Node: Node{Id: graph.GenID(),
			Type:      fabric.UINode,
			Signalers: &sm3,
			Signals:   &s3,
		},
	}

	up2, err := graph.AddRealNode(u2)
	if err != nil {
		t.Fatalf("Could not add UI node to graph: %v", err)
	}

	// Add edge from first UI node to second UI node
	graph.AddRealEdge(u.ID(), up2)

	// Do some Leaf and Root Boundary checks again
	for n := range graph.Top {
		if n.ID() == u.Id {
			if !graph.IsRootBoundary(n) {
				t.Fatal("Incorrectly classified a root boundary node.")
			}
		}

		if n.ID() == u2.Id {
			if !graph.IsLeafBoundary(n) {
				t.Fatal("Incorrectly classified a leaf boundary node.")
			}
		}
	}

	// Create VUI
	sm4 := make(fabric.SignalingMap)
	s4 := make(fabric.SignalsMap)
	vu := UI{
		Node: Node{Id: graph.GenID(),
			Type:      fabric.VUINode,
			Signalers: &sm4,
			Signals:   &s4,
		},
		Virtual: true,
	}

	// Add VUI to graph
	vup, err := graph.AddVUI(vu)
	if err != nil {
		t.Fatalf("Could not add VUI node to graph: %v", err)
	}

	// Quick Signalers and Signals check (verbose test logs)
	for n := range graph.Top {
		if n.ID() == u.Id {
			signals := n.ListSignals()
			t.Logf("UI-1 Signals: %v", signals)
			signalers := n.ListSignalers()
			t.Logf("UI-1 Signalers: %v", signalers)
		}

		if n.ID() == temp.Id {
			signals := n.ListSignals()
			t.Logf("Temporal Signals: %v", signals)
			signalers := n.ListSignalers()
			t.Logf("Temporal Signalers: %v", signalers)
		}
	}

	// Remove VUI from graph (check)
	err = graph.RemoveVUI(vup)
	if err != nil {
		t.Fatalf("Could not remove VUI node from graph: %v", err)
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
	u1 := UI{
		Node: Node{
			Id:        graph.GenID(),
			Type:      fabric.UINode,
			Signalers: &sm1,
			Signals:   &s1,
		},
		Virtual: false,
	}

	u1p, err := graph.AddRealNode(u1)
	if err != nil {
		t.Fatalf("Could not add UI node to graph: %v", err)
	}

	// Create second UI node
	sm2 := make(fabric.SignalingMap)
	s2 := make(fabric.SignalsMap)
	u2 := UI{
		Node: Node{
			Id:        graph.GenID(),
			Type:      fabric.UINode,
			Signalers: &sm2,
			Signals:   &s2,
		},
		Virtual: false,
	}

	u2p, err := graph.AddRealNode(u2)
	if err != nil {
		t.Fatalf("Could not add UI node to graph: %v", err)
	}

	// Create third UI node
	sm3 := make(fabric.SignalingMap)
	s3 := make(fabric.SignalsMap)
	u3 := UI{
		Node: Node{
			Id:        graph.GenID(),
			Type:      fabric.UINode,
			Signalers: &sm3,
			Signals:   &s3,
		},
		Virtual: false,
	}

	u3p, err := graph.AddRealNode(u3)
	if err != nil {
		t.Fatalf("Could not add UI node to graph: %v", err)
	}

	// Add Edges (creates a cycle)
	graph.AddRealEdge(u1.ID(), u2p)
	graph.AddRealEdge(u2.ID(), u3p)
	graph.AddRealEdge(u3.ID(), u1p)

	if !graph.CycleDetect() {
		t.Fatalf("Did not detect cycle in the graph")
	}
}

/* CDS Testing */

// ElementNode satisfies fabric.Node interface
type ElementNode struct {
	Id    int
	Value interface{}
	L     *List
	Imm   bool
}

func (e ElementNode) ID() int {
	return e.Id
}

func (e ElementNode) Immutable() bool {
	return e.Imm
}

// ElementEdge satisfies the fabric.Edge interface
type ElementEdge struct {
	Id          int
	L           *List
	Source      *ElementNode
	Destination *ElementNode
	Imm         bool
}

func (e ElementEdge) ID() int {
	return e.Id
}

func (e ElementEdge) GetSource() fabric.Node {
	var i interface{} = *e.Source
	in := i.(fabric.Node)
	return in
}

func (e ElementEdge) GetDestination() fabric.Node {
	var i interface{} = *e.Destination
	in := i.(fabric.Node)
	return in
}

func (e ElementEdge) Immutable() bool {
	return e.Imm
}

// List satisfies the fabric.CDS interface
type List struct {
	Root  *ElementNode
	Len   int
	Nodes fabric.NodeList
	Edges fabric.EdgeList
}

func NewList() *List {
	l := List{}
	n := &ElementNode{
		Id: l.GenNodeID(),
		L:  &l,
	}
	l.Root = n
	l.Len = 1

	var i interface{} = n
	in := i.(fabric.Node)

	var nl fabric.NodeList
	nl = append(nl, in)
	l.Nodes = nl

	el := make(fabric.EdgeList, 0)
	l.Edges = el

	return &l
}

// NewElementNode will create a new List element node
// NOTE: this should be defined as an access procedure of some access type
func (l *List) NewElementNode() *ElementNode {
	n := ElementNode{
		Id: l.GenNodeID(),
		L:  l,
	}

	var i interface{} = n
	in := i.(fabric.Node)

	l.Nodes = append(l.Nodes, in)
	l.Len += 1

	return &n
}

// NewElementEdge will create a new List edge
// NOTE: this should be defined as an access procedure of some access type
func (l *List) NewElementEdge(s, d *ElementNode) {
	e := ElementEdge{
		Id:          l.GenEdgeID(),
		L:           l,
		Source:      s,
		Destination: d,
	}

	var i interface{} = e
	ie := i.(fabric.Edge)

	l.Edges = append(l.Edges, ie)
}

// Generate an ID for a CDS Node
func (l List) GenNodeID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	for _, n := range l.Nodes {
		if n.ID() == id {
			id = l.GenNodeID()
		}
	}

	return id
}

// Generate an ID for a CDS Edge
func (l List) GenEdgeID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	for _, e := range l.Edges {
		if e.ID() == id {
			id = l.GenEdgeID()
		}
	}

	return id

}

func (l List) ListNodes() fabric.NodeList {
	return l.Nodes
}

func (l List) ListEdges() fabric.EdgeList {
	return l.Edges
}

func TestTotalityUnique(t *testing.T) {
	// create CDS
	list := NewList()
	n1 := list.Root
	n2 := list.NewElementNode()
	n3 := list.NewElementNode()
	n4 := list.NewElementNode()
	list.NewElementEdge(n1, n2)
	list.NewElementEdge(n2, n3)
	list.NewElementEdge(n3, n4)
	l := *list
	var il interface{} = l
	li := il.(fabric.CDS)

	// Create graph (and add CDS to graph)
	graph := fabric.NewGraph()
	graph.DS = li

	// create section
	branch := fabric.NewBranch(list.Nodes[0], li)
	var ib interface{} = branch
	b := ib.(fabric.Section)

	// Create UI nodes
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
		CDS:     b, // both UI nodes will address the same CDS section
	}

	sm2 := make(fabric.SignalingMap)
	s2 := make(fabric.SignalsMap)
	u2 := UI{
		Node: Node{
			Id:        graph.GenID(),
			Type:      fabric.UINode,
			Signalers: &sm2,
			Signals:   &s2,
		},
		Virtual: false,
		CDS:     b, // both UI nodes will address the same CDS section
	}

	// Add UI nodes to graph
	_, err := graph.AddRealNode(u)
	if err != nil {
		t.Fatalf("Could not add UI node to graph: %v", err)
	}

	_, err = graph.AddRealNode(u2)
	if err != nil {
		t.Fatalf("Could not add UI node to graph: %v", err)
	}

	// check that a TotalityUnique verification fails
	if graph.TotalityUnique() {
		t.Fatal("Incorrectly classified graph has totality unique.")
	}
}

func TestCovered(t *testing.T) {
	// Create CDS
	list := NewList()
	n1 := list.Root
	n2 := list.NewElementNode()
	n3 := list.NewElementNode()
	n4 := list.NewElementNode()
	n5 := list.NewElementNode()
	n6 := list.NewElementNode()
	list.NewElementEdge(n1, n2)
	list.NewElementEdge(n2, n3)
	list.NewElementEdge(n3, n4)
	list.NewElementEdge(n5, n6)
	l := *list
	var il interface{} = l
	li := il.(fabric.CDS)

	// Create graph (and add CDS reference to graph)
	graph := fabric.NewGraph()
	graph.DS = li

	// Create a section
	branch := fabric.NewBranch(list.Nodes[0], li)
	var ib interface{} = branch
	b := ib.(fabric.Section)

	// Create UI node
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
		CDS:     b,
	}

	// Add UI node to graph
	_, err := graph.AddRealNode(u)
	if err != nil {
		t.Fatalf("Could not add UI node to graph: %v", err)
	}

	// check that a CDS covered verification fails
	if graph.Covered() {
		t.Fatal("Incorrectly classified graph as covering entire CDS")
	}
}
