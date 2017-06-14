package fabric

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// RECOMMENDATION: every thread should have a proper reaction (which may be a non-reaction) to each signal value for each access
// 	procedure in each dependency node.
// EXAMPLE: an example reaction to an abort signal could be "Abort Chain/Tree" where the dependents
// 	and their dependents, etc. all abort their operations if a signal value from a dependency node
// 	is an 'Abort' signal.
type Signal int

const (
	Waiting Signal = iota
	Started
	Completed
	Aborted
	AbortRetry   // EXAMPLE: could use exponential backoff checks on retries for AbortRetry signals from dependencies ...
	PartialAbort // (used to specify if an operation partially-completed before aborting)
)

// ProcedureSignals is used to map a signal to the access type that caused the signal
// NOTE: the string key should be equivalent to the Class() method return value for that AccessType
// EXAMPLE: a system design calls for a single thread having multiple access procedures,
// 	only some of which induce a dependent to invoke a responsive operation, then to know which
// 	procedure a signal is from you can use this map.
type ProcedureSignals map[string]Signal

type NodeType int

const (
	UINode NodeType = iota
	TemporalNode
	VirtualTemporalNode
	VUINode
	VirtualNode
	Unknown
)

// SignalingMap is a map of dependent node ids to a set of
// access procedures and their current signal states.
type SignalingMap map[int]chan ProcedureSignals

// SignalsMap is a map of dependency node ids to a set of
// their access procedures and their current signal states.
type SignalsMap map[int]<-chan ProcedureSignals

// Dependency Graph Node
// every DGNode has an id, a Type, a state, and a set of Access Procedures
// NOTE: This will require assigning signals to their appropriate nodes
//		when setting up a dependency graph.
type DGNode interface {
	ID() int           // must be unique from all other DGNodes in our graph
	GetType() NodeType // specifies whether node is UI, VUI, etc.
	GetPriority() int  // not necessary, but can be useful
	ListProcedures() ProcedureList
	ListDependents() []DGNode
	ListDependencies() []DGNode
	ListSignalers() SignalingMap // lists signaler channels for signaling dependents
	ListSignals() SignalsMap     // lists signal channels from *dependencies*
	Signal(ProcedureSignals)     // used to send the same signal to all dependents in signalers list
}

// Graph can be either UI DDAG, Temporal DAG or VDG
type Graph struct {
	DS    *CDS // reference to CDS that the dependency graph is for
	Nodes []DGNode
	// FIXME: Edges should be a reference to a node and a list of references to other nodes
	Edges map[DGNode][]DGNode // each node (id) has a list of node ids that it points too
}

// NewGraph creates a new empty graph
func NewGraph() *Graph {
	var nodes []DGNode
	return &Graph{
		Nodes: nodes,
		Edges: make(map[DGNode][]DGNode),
	}
}

func (g *Graph) GenID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	for _, n := range g.Nodes {
		if n.ID() == id {
			id = g.GenID()
		}
	}
	return id
}

func (g *Graph) IsLeafBoundary(n DGNode) bool {
	if len(g.Dependents(n)) == 0 {
		return true
	}

	return false
}

func (g *Graph) IsRootBoundary(n DGNode) bool {
	if len(g.Dependencies(n)) == 0 {
		return true
	}

	return false
}

func (g *Graph) CreateSignalers(n DGNode) SignalingMap {
	sm := make(SignalingMap)

	deps := g.Dependents(n)
	for _, d := range deps {
		c := make(chan ProcedureSignals)
		sm[d.ID()] = c
	}

	return sm
}

func (g *Graph) Signals(n DGNode) SignalsMap {
	sm := make(SignalsMap)

	deps := g.Dependencies(n)
	for _, d := range deps {
		channels := d.ListSignalers()
		ch := channels[n.ID()]
		sm[n.ID()] = ch
	}

	return sm
}

// This should only be used for adding nodes to a graph
// to intialize the graph.
func (g *Graph) AddRealNode(node DGNode) error {
	if !contains(g.Nodes, node) {
		g.Nodes = append(g.Nodes, node)
	} else {
		return fmt.Errorf("Node already exists in Dependency Graph.")
	}
	return nil
}

func (g *Graph) AddRealEdge(source, dest DGNode) {
	if _, ok := g.Edges[source]; !ok {
		g.Edges[source] = []DGNode{dest}
	} else {
		s := g.Edges[source]
		s = append(s, dest)
		g.Edges[source] = s
	}
}

// CycleDetect will check whether a graph has cycles or not
func (g *Graph) CycleDetect() bool {
	var seen []DGNode
	var done []DGNode

	for _, v := range g.Nodes {
		if !contains(done, v) {
			result, d := g.cycleDfs(v, seen, done)
			done = d
			if result {
				return true
			}
		}
	}
	return false
}

// GetAdjacents will return the list of nodes a supplied node points too
func (g *Graph) GetAdjacents(node DGNode) []DGNode {
	return g.Edges[node]
}

// Recursive Depth-First-Search; used for Cycle Detection
func (g *Graph) cycleDfs(start DGNode, seen, done []DGNode) (bool, []DGNode) {
	seen = append(seen, start)
	adj := g.Edges[start]
	for _, v := range adj {
		if contains(done, v) {
			continue
		}

		if contains(seen, v) {
			return true, done
		}

		if result, done := g.cycleDfs(v, seen, done); result {
			return true, done
		}
	}
	seen = seen[:len(seen)-1]
	done = append(done, start)
	return false, done
}

// Totality-Uniqueness check for the UI nodes of a graph...
// should only be called once when creating the UI dependency graph;
// can be called with the creation of each UI if needed for
// more "real-time" verification.
func (g *Graph) TotalityUnique() bool {
	// grab all UI nodes
	uiSlice := make([]DGNode, 0)
	for _, v := range g.Nodes {
		if v.GetType() == UINode {
			uiSlice = append(uiSlice, v)
		}
	}

	for i, n := range uiSlice {
		// compare the UI to all other UIs
		for i2, n2 := range g.Nodes {
			if i != i2 {
				// if UI is same as other UI return false
				// i.e. graph is not totality-unique
				if reflect.DeepEqual(n, n2) {
					return false
				}
			}
		}
	}

	return true
}

// Covered returns true if all CDS nodes and edges are covered
func (g *Graph) Covered() bool {
	// grab all UI nodes
	uiSlice := make([]UI, 0)
	for _, v := range g.Nodes {
		if v.GetType() == UINode {
			uiSlice = append(uiSlice, v.(UI))
		}
	}

	// grab all CDS nodes and edges
	ds := *g.DS
	nodes := ds.ListNodes()
	edges := ds.ListEdges()

FIRST:
	// for every node in the CDS
	for _, v := range nodes {
		// check that at least one UI contains it
		for _, u := range uiSlice {
			s := u.GetSection()
			sp := *s
			uiCDSNodes := sp.ListNodes()
			// if UI contains node; check next CDS node
			if containsNode(uiCDSNodes, v) {
				continue FIRST
			}
		}
		// if CDS node is checked in every UI and does not show up
		return false
	}

SECOND:
	// for every edge in the CDS
	for _, v := range edges {
		// check that at least one UI contains it
		for _, u := range uiSlice {
			s := u.GetSection()
			sp := *s
			uiCDSEdges := sp.ListEdges()
			// if UI contains edge; check next CDS edge
			if containsEdge(uiCDSEdges, v) {
				continue SECOND
			}
		}

		// if CDS edge is checked in every UI and does not show up
		return false
	}

	return true
}

func (g *Graph) AddVUI(node UI) error {
	if !node.IsVirtual() {
		return fmt.Errorf("Not a virtual node.")
	} else {
		if !contains(g.Nodes, node) {
			g.Nodes = append(g.Nodes, node)
		} else {
			return fmt.Errorf("Node already exists in Dependency Graph.")
		}
	}
	return nil
}

func (g *Graph) RemoveVUI(node UI) error {
	newList := make([]DGNode, 0)
	if !node.IsVirtual() {
		return fmt.Errorf("Not a virtual node.")
	} else {
		for _, v := range g.Nodes {
			if node.ID() != v.ID() {
				newList = append(newList, v)
			}
		}

		g.Nodes = newList
	}

	return nil
}

func (g *Graph) Dependents(n DGNode) []DGNode {
	list := make([]DGNode, 0)
	for i, v := range g.Edges {
		if i.ID() != n.ID() {
			if contains(v, n) {
				list = append(list, i)
			}
		}
	}

	return list
}

func (g *Graph) Dependencies(n DGNode) []DGNode {
	emptyList := make([]DGNode, 0)
	if v, ok := g.Edges[n]; !ok {
		return emptyList
	} else {
		return v
	}
}

func (g *Graph) Type(n DGNode) NodeType {

	j, ok := n.(UI)
	if ok {
		if j.IsVirtual() {
			return VUINode
		}
		return UINode
	}
	k, ok := n.(Temporal)
	if ok {
		if k.IsVirtual() {
			return VirtualTemporalNode
		}
		return TemporalNode
	}
	_, ok = n.(Virtual)
	if ok {
		return VirtualNode
	}

	return Unknown
}
