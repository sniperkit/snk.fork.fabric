package fabric

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

// Signal defines the possible signal values one dependency graph node can send to another
// RECOMMENDATION: every thread should have a proper reaction (which may be a non-reaction)
//  to each signal value for each access procedure in each dependency node.
// EXAMPLE: an example reaction to an abort signal could be "Abort Chain/Tree" where the dependents
// 	and their dependents, etc. all abort their operations if a signal value from a dependency node
// 	is an 'Abort' signal.
type Signal int

const (
	// Waiting can be used for an access procedure that has not begun but is in line too
	Waiting Signal = iota
	// Started can be used for an access procedure that is no longer waiting and has begun execution
	Started
	// Completed can be used for an access procedure that has finished execution successfully
	Completed
	// Aborted can be used for an access procedure that failed to finish execution
	Aborted
	// AbortRetry EXAMPLE: could use exponential backoff checks on retries for AbortRetry signals from dependencies ...
	AbortRetry
	// PartialAbort can be used to specify if an operation partially-completed before aborting)
	PartialAbort
)

// NodeSignal carries all the information a dependent node will need in order to know what
// action a dependent node has just taken.
type NodeSignal struct {
	AccessType int // should be equivalent to the ID() method return value for the Access Type
	Value      Signal
	Space      UI
}

// NodeType defines the possible values for types of dependency graph nodes
type NodeType int

const (
	// UINode are the spatial definitions usually assigned to a single thread
	UINode NodeType = iota
	// TemporalNode are assigned to threads which address the same UI as the temporals UI dependent
	TemporalNode
	// VirtualTemporalNode is a spawned temporary temporal node
	VirtualTemporalNode
	// VUINode is a temporary UI node
	VUINode
	// VDGNode is a node in a virtual dependency graph
	VDGNode
	// Unknown is a catch-all for an improperly constructed dependency graph node
	Unknown
)

// SignalingMap is a map of dependent node ids to a channel of signals
type SignalingMap map[int]chan NodeSignal

// SignalsMap is a map of dependency node ids to recieve-only channel of signals
type SignalsMap map[int]<-chan NodeSignal

// DGNode (Dependency Graph Node) ...
// every DGNode has an id, a Type, a state, and a set of Access Procedures
// NOTE: This will require assigning signals to their appropriate nodes
//		when setting up a dependency graph.
type DGNode interface {
	ID() int           // must be unique from all other DGNodes in our graph
	GetType() NodeType // specifies whether node is UI, VUI, etc.
	GetPriority() int  // not necessary, but can be useful
	ListProcedures() ProcedureList
	UpdateSignaling(SignalingMap, SignalsMap) // makes it possible to update the SignalingMap and SignalsMap for a DGNode
	ListSignalers() SignalingMap
	ListSignals() SignalsMap
	Signal(NodeSignal) // used to send the same signal to all dependents in signalers list
}

// Graph can be either UI DDAG, Temporal DAG or VDG
type Graph struct {
	DS  *CDS // reference to data structure that the dependency graph is for
	Top map[DGNode][]*DGNode
	VDG []*VDG
}

// NewGraph creates a new empty graph
func NewGraph() *Graph {
	return &Graph{
		Top: make(map[DGNode][]*DGNode),
		VDG: make([]*VDG, 0),
	}
}

// AddVDG ...
func (g *Graph) AddVDG(v *VDG) error {
	// check if VDG already exists in graph
	for _, vdg := range g.VDG {
		if vdg == v {
			return fmt.Errorf("VDG already exists in graph")
		}
	}

	g.VDG = append(g.VDG, v)

	return nil
}

// RemoveVDG ...
func (g *Graph) RemoveVDG(v *VDG) {
	for i, vdg := range g.VDG {
		if vdg == v {
			g.VDG = append(g.VDG[:i], g.VDG[i+1:]...)
		}
	}
}

// GenID ...
func (g *Graph) GenID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	for n := range g.Top {
		if n.ID() == id {
			id = g.GenID()
		}
	}
	return id
}

// IsLeafBoundary ...
func (g *Graph) IsLeafBoundary(n *DGNode) bool {
	if len(g.Dependencies(n)) == 0 {
		return true
	}

	return false
}

// IsRootBoundary ...
func (g *Graph) IsRootBoundary(n *DGNode) bool {
	if len(g.Dependents(n)) == 0 {
		return true
	}

	return false
}

// SignalsAndSignalers will udpate the SignalingMaps and SignalsMaps for all DGNodes in the graph
func (g *Graph) SignalsAndSignalers() {

	// for all nodes in the graph
	for n, l := range g.Top {
		// create its SignalersMap
		sm := make(SignalingMap)
		deps := g.Dependents(&n)
		for _, d := range deps {
			c := make(chan NodeSignal)
			sm[d.ID()] = c
		}

		// create its SignalsMap
		s := make(SignalsMap)
		for _, np := range l {
			dep := *np
			channels := dep.ListSignalers()
			ch := channels[dep.ID()]
			s[dep.ID()] = ch
		}

		n.UpdateSignaling(sm, s)
	}
}

// BasicSignalHandler is the basic function type for handling signals from a dependency node
// Used in total-blocking, to call wg.Done() on certain Signal Values and return.
// will not allow for more complex signal handling e.g. handling an Abort or AbortRetry with more resilience (use with caution)
type BasicSignalHandler func(<-chan NodeSignal, sync.WaitGroup)

// TotalBlock is the most basic format for signal checking
// (used when a node wants to simply totally-block all further operations until its dependencies have signaled)
// as it only accepts a BasicSignalHandler it will not be a very powerful form of blocking (only use if lazy)
func (g *Graph) TotalBlock(nodeID int, handler BasicSignalHandler) bool {
	var wg sync.WaitGroup

	for n := range g.Top {
		if n.ID() == nodeID {
			depSignals := n.ListSignals()
			for _, channel := range depSignals {
				wg.Add(1)
				go handler(channel, wg)
			}
			break
		}
	}

	// Virtual Node blocks/spins
	wg.Wait()

	return true
}

// AddRealNode ...
// This should only be used for adding nodes to a graph
// to intialize the graph.
func (g *Graph) AddRealNode(node DGNode) (*DGNode, error) {
	var pointer *DGNode
	if !reflect.ValueOf(node).Type().Comparable() {
		return pointer, fmt.Errorf("Node type is not comparable and cannot be used in the graph topology. \n Try removing any slices, maps, and functions from struct definition.")
	}

	if _, ok := g.Top[node]; !ok {
		g.Top[node] = []*DGNode{}
	} else {
		return pointer, fmt.Errorf("Node already exists in Dependency Graph.")
	}

	for n := range g.Top {
		if n.ID() == node.ID() {
			pointer = &n
		}
	}
	return pointer, nil
}

// AddRealEdge will create an edge and an appropriate signaling channel between nodes
func (g *Graph) AddRealEdge(source int, dest *DGNode) {
	d := *dest

	for i, k := range g.Top {
		if i.ID() == source {
			if !containsDGNode(k, dest) {
				k = append(k, dest)
				g.Top[i] = k

				// update SignalingMap for destination
				depSig := d.ListSignalers()
				depS := d.ListSignals()
				depSig[i.ID()] = make(chan NodeSignal)
				d.UpdateSignaling(depSig, depS)

				// update SignalsMap for source
				signals := i.ListSignals()
				signalers := i.ListSignalers()
				for j, v := range d.ListSignalers() {
					if j == i.ID() {
						signals[d.ID()] = v
						break
					}
				}
				i.UpdateSignaling(signalers, signals)
			}
		}
	}
}

// CycleDetect will check whether a graph has cycles or not
func (g *Graph) CycleDetect() bool {
	var seen []DGNode
	var done []DGNode

	for i := range g.Top {
		if !contains(done, i) {
			result, d := g.cycleDfs(i, seen, done)
			done = d
			if result {
				return true
			}
		}
	}
	return false
}

// Allowed checks whether or not an access procedure is allowed to act on a node ...
func (g *Graph) Allowed(node DGNode, procedure AccessType) bool {
	allowed := false
	for _, p := range node.ListProcedures() {
		if p.ID() == procedure.ID() {
			allowed = true
		}
	}

	return allowed
}

// Recursive Depth-First-Search; used for Cycle Detection
func (g *Graph) cycleDfs(start DGNode, seen, done []DGNode) (bool, []DGNode) {
	seen = append(seen, start)
	adj := g.Top[start]
	for _, vp := range adj {
		v := *vp
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

// GetAdjacents will return the list of nodes that a node is connected too
func (g *Graph) GetAdjacents(node DGNode) []DGNode {
	var list []DGNode

	for n, l := range g.Top {
		if n.ID() == node.ID() {
			// Add all dependents to list
			for n2, l2 := range g.Top {
				if containsDGNode(l2, &n) {
					list = append(list, n2)
				}
			}
			// Add all dependencies to list
			for _, np := range l {
				list = append(list, *np)
			}
		}
	}

	return list
}

// TotalityUnique is a Totality-Uniqueness check for the UI nodes of a graph...
// should only be called once when creating the UI dependency graph;
// can be called with the creation of each UI if needed for
// more "real-time" verification.
func (g *Graph) TotalityUnique() bool {
	// grab all UI nodes
	var uiSlice []DGNode
	for i := range g.Top {
		if i.GetType() == UINode {
			uiSlice = append(uiSlice, i)
		}
	}

	var done []DGNode

	// for every UI Node
	for i, n := range uiSlice {
		// compare it against every other UI node
		for j, n2 := range uiSlice {
			if !contains(done, n2) {
				if j != i {
					if reflect.DeepEqual(n, n2) {
						return false
					}
				}
			}
		}
		done = append(done, n)
	}

	return true
}

// Covered returns true if all CDS nodes and edges are covered
func (g *Graph) Covered() bool {
	// grab all UI nodes
	var uiSlice []UI
	for v := range g.Top {
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
			uiCDSNodesPointer := s.ListNodes()
			uiCDSNodes := *uiCDSNodesPointer
			// if UI contains node; check next CDS node
			if ContainsNode(uiCDSNodes, v) {
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
			uiCDSEdgesPointer := s.ListEdges()
			uiCDSEdges := *uiCDSEdgesPointer
			// if UI contains edge; check next CDS edge
			if ContainsEdge(uiCDSEdges, v) {
				continue SECOND
			}
		}

		// if CDS edge is checked in every UI and does not show up
		return false
	}

	return true
}

// AddVUI requires that the node return a true value for its IsVirtual method
func (g *Graph) AddVUI(node UI) (*DGNode, error) {
	var pointer *DGNode

	if !node.IsVirtual() {
		return pointer, fmt.Errorf("Not a virtual node.")
	}

	var nodeSlice []DGNode
	for n := range g.Top {
		nodeSlice = append(nodeSlice, n)
	}

	if !contains(nodeSlice, node) {
		g.Top[node.(DGNode)] = []*DGNode{}
	} else {
		return pointer, fmt.Errorf("Node already exists in Dependency Graph")
	}

	for n := range g.Top {
		if n.ID() == node.ID() {
			pointer = &n
		}
	}

	return pointer, nil
}

// RemoveVUI ...
func (g *Graph) RemoveVUI(n DGNode) error {
	// FIXME: sometimes this check fails randomly ...
	node, ok := n.(UI)
	if !ok {
		return fmt.Errorf("Not a UI node")
	}

	// FIXME: sometimes this check fails randomly ...
	if !node.IsVirtual() {
		return fmt.Errorf("Not a virtual node")
	}

	for n1 := range g.Top {
		if n1.ID() == n.ID() {
			if len(g.Dependencies(&n1)) != 0 {
				return fmt.Errorf("VUI node still has dependencies")
			}
		}
	}

	// Remove VUI from Signals maps in depedent nodes
	for n1 := range g.Top {
		if n1.ID() == n.ID() {
			for n2, l := range g.Top {
				if containsDGNode(l, &n1) {
					signals := n2.ListSignals()
					delete(signals, n1.ID())
					n2.UpdateSignaling(n2.ListSignalers(), signals)
				}
			}
		}
	}

	// remove node from graph
	delete(g.Top, n)

	return nil
}

// Dependents ...
func (g *Graph) Dependents(np *DGNode) []DGNode {
	var list []DGNode
	n := *np

	for i, v := range g.Top {
		if i.ID() != n.ID() {
			if containsDGNode(v, np) {
				list = append(list, i)
			}
		}
	}

	return list
}

// Dependencies ...
func (g *Graph) Dependencies(np *DGNode) []DGNode {
	var list []DGNode

	n := *np
	v, ok := g.Top[n]
	if !ok {
		return list
	}

	for _, p := range v {
		pp := *p
		list = append(list, pp)
	}
	return list
}

// Type will return the proper NodeType value for a given DGNode argument
func (g *Graph) Type(n DGNode) NodeType {
	if j, ok := n.(UI); ok {
		if j.IsVirtual() {
			return VUINode
		}
		return UINode
	}
	if k, ok := n.(Temporal); ok {
		if k.IsVirtual() {
			return VirtualTemporalNode
		}
		return TemporalNode
	}
	if _, ok := n.(Virtual); ok {
		return VDGNode
	}

	return Unknown
}
