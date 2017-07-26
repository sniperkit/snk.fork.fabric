package fabric

import (
	"fmt"
	"math/rand"
	"time"
)

// Virtual is the interface definition that virtual nodes in a VDG graph
// need to satsify.
// NOTE: Virtual Dependency Graph Nodes can only have other
//		VDG nodes as dependents or dependencies
type Virtual interface {
	DGNode
	Start()        // tells the node to begin execution (usually by flipping a 'Started' boolean -- can also signal to dependents that it has started)
	Started() bool // specifies whether a node has started execution or not
	IsRoot() bool  // specifies whether VDG node is root node or not
	Subspace() UI  // the (V)UI that the Virtual node is associated to
}

// Root is the object type for pre-made root nodes
// Root is also is able to satisfy the UI interface in order to add itself as it's subspace
type Root struct {
	AccessProcedures *ProcedureList
	CDS              Section
	Isroot           bool
	IsLeaf           bool
	Signalers        *SignalingMap
	Signals          *SignalsMap
	Executing        bool
	Type             NodeType
}

// ID ...
func (r *Root) ID() int {
	return 0
}

// GetType ...
func (r *Root) GetType() NodeType {
	return VDGNode
}

// GetPriority ...
func (r *Root) GetPriority() int {
	return 0
}

// ListProcedures ...
func (r *Root) ListProcedures() ProcedureList {
	var p ProcedureList
	return p
}

// UpdateSignaling ...
func (r *Root) UpdateSignaling(sm SignalingMap, s SignalsMap) {
	*r.Signalers = sm
	*r.Signals = s
}

// ListSignalers ...
func (r *Root) ListSignalers() SignalingMap {
	return *r.Signalers
}

// ListSignals ...
func (r *Root) ListSignals() SignalsMap {
	return *r.Signals
}

// Signal ...
func (r *Root) Signal(p ProcedureSignals) {
	return
}

// Start ...
func (r *Root) Start() {
	r.Executing = true
}

// Started ...
func (r *Root) Started() bool {
	return r.Executing
}

// IsRoot ...
func (r *Root) IsRoot() bool {
	return true
}

// Subspace ...
func (r *Root) Subspace() UI {
	var i interface{} = r
	ir := i.(UI)
	return ir
}

// GetSection ...
func (r *Root) GetSection() Section {
	return r.CDS
}

// IsUnique ...
func (r *Root) IsUnique() bool {
	return false
}

// IsVirtual ...
func (r *Root) IsVirtual() bool {
	return true
}

// NOTE: Virtual Dependency Graphs are always trees with a root node

//		The set of (V)UIs that are associated with the VDG are the
//		set of (V)UIs that at least one node in the VDG accesses.

//		A VDG that only affects one (V)UI represents a virtual
//		*temporal* dependency graph.

//		A VDG that affects more than one (V)UI represents a virtual
//		*spatial* dependency graph.

// NOTE: all VDG nodes should be passed by reference, including the VDG itself

// IMPORTANT: **VDGs will run independent** of all other dependency graphs
//		in other words, a VDG does not have any dependents or dependencies
// 		on any (V)UIs! It simply has associations to (V)UIs. The purpose
// 		of a VDG is to order temporary threads (even if they are
//		associated with different UIs).

// VDG ...
type VDG struct {
	Global *Graph // a reference to the real global graph of the system
	Root   Virtual
	Top    map[Virtual][]Virtual
	Space  []int // the set of all (V)UI ids that at least one node in the VDG has access too
}

// NewVDG will return an empty VDG graph
func NewVDG(g *Graph) (*VDG, error) {
	// create VDG
	v := &VDG{
		Global: g,
		Top:    make(map[Virtual][]Virtual),
		Space:  make([]int, 0),
	}

	// add to graph
	err := g.AddVDG(v)
	if err != nil {
		return v, err
	}

	return v, nil
}

// NewVDGWithRoot ...
func NewVDGWithRoot(g *Graph) (*VDG, error) {
	// create Virtual root node
	sm1 := make(SignalingMap)
	s1 := make(SignalsMap)
	r := &Root{
		Signalers: &sm1,
		Signals:   &s1,
	}
	var i interface{} = r
	ir := i.(Virtual)

	// create VDG
	v := &VDG{
		Global: g,
		Root:   ir,
		Top:    make(map[Virtual][]Virtual),
		Space:  make([]int, 0),
	}

	// add to graph
	err := g.AddVDG(v)
	if err != nil {
		return v, err
	}

	return v, nil
}

// GenID can generate a unique integer id for a VDG node
func (g *VDG) GenID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	for n := range g.Top {
		if n.ID() == id {
			id = g.GenID()
		}
	}
	return id
}

// CreateSignalers ...
func (g *VDG) CreateSignalers(n Virtual) SignalingMap {
	sm := make(SignalingMap)

	deps := g.Dependents(n)
	for _, d := range deps {
		c := make(chan ProcedureSignals)
		sm[d.ID()] = c
	}

	return sm
}

// Signals ...
func (g *VDG) Signals(n Virtual) SignalsMap {
	sm := make(SignalsMap)

	deps := g.Dependencies(n)
	for _, d := range deps {
		channels := d.ListSignalers()
		ch := channels[n.ID()]
		sm[n.ID()] = ch
	}

	return sm
}

// Dependents ...
func (g *VDG) Dependents(n Virtual) []Virtual {
	var list []Virtual

	for i, v := range g.Top {
		if i.ID() != n.ID() {
			if containsVirtual(v, n) {
				list = append(list, i)
			}
		}
	}

	return list
}

// Dependencies ...
func (g *VDG) Dependencies(n Virtual) []Virtual {
	var list []Virtual

	v, ok := g.Top[n]
	if !ok {
		return list
	}

	for _, p := range v {
		list = append(list, p)
	}

	return list
}

// AddVirtualNode adds a node to a VDG
func (g *VDG) AddVirtualNode(node Virtual) (Virtual, error) {
	var ret Virtual
	if _, ok := g.Top[node]; !ok {
		g.Top[node] = []Virtual{}
	} else {
		return ret, fmt.Errorf("Node already exists in Dependency Graph.")
	}
	// Add node's subspace to graph
	g.Space = append(g.Space, node.Subspace().ID())
	for n := range g.Top {
		if n.ID() == node.ID() {
			ret = n
		}
	}
	return ret, nil
}

// AddTopNode will add a node to the VDG and create an edge pointing from the root node to it
func (g *VDG) AddTopNode(node Virtual) (Virtual, error) {
	var ret Virtual
	if _, ok := g.Top[node]; !ok {
		g.Top[node] = []Virtual{}
	} else {
		return ret, fmt.Errorf("Node already exists in Dependency Graph.")
	}

	// Add node's subspace to graph
	g.Space = append(g.Space, node.Subspace().ID())
	for n := range g.Top {
		if n.ID() == node.ID() {
			ret = n
			// Add edge from root node to our new node
			root := g.Root
			g.AddVirtualEdge(root.ID(), n)
		}
	}
	return ret, nil
}

// RemoveVirtualNode is for removing a single node from a VDG
// It will also remove all edges that have the node as
// the destination node of the edge. And it will remove the (V)UI
// subspace if not required by any other node in the VDG.
func (g *VDG) RemoveVirtualNode(n Virtual) error {

	for node, list := range g.Top {
		if node.ID() == n.ID() {
			if len(list) > 0 {
				return fmt.Errorf("Virtual node still has dependencies. Cannot be deleted.")
			}
		}
	}

	delete(g.Top, n)

	// remove all references (edges) to node in other nodes edge slices
	for n1, l := range g.Top {
		if containsVirtual(l, n) {
			for j, k := range l {
				if k.ID() == n.ID() {
					l = append(l[:j], l[j+1:]...)
					g.Top[n1] = l
				}
			}
		}
	}

	// Remove VUI subspace from VDG (if not subspace for another Virtual node)
	id := n.Subspace().ID()
	remove := true
	for i := range g.Top {
		if i.Subspace().ID() == id {
			remove = false
			break
		}
	}
	if remove == true {
		for i, k := range g.Space {
			if k == id {
				g.Space = append(g.Space[:i], g.Space[i+1:]...)
				break
			}
		}
	}

	return nil
}

// AddVirtualEdge adds an edge to a VDG
func (g *VDG) AddVirtualEdge(source int, d Virtual) error {
	for i, k := range g.Top {
		if i.ID() == source {
			if i.Started() {
				return fmt.Errorf("Node has already started. Cannot add dependencies")
			}
			if !containsVirtual(k, d) {
				k = append(k, d)
				g.Top[i] = k

				// update SignalingMap for destination
				depSig := d.ListSignalers()
				depS := d.ListSignals()
				depSig[i.ID()] = make(chan ProcedureSignals)
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

	return nil
}

// RemoveVirtualEdge removes a single edge from the VDG
// Useful for when a dependency node is not being removed but
// the dependent node no longer requires it as a dependency.
func (g *VDG) RemoveVirtualEdge(source int, d Virtual) {
	for i, k := range g.Top {
		if i.ID() == source {
			for j, v := range k {
				if v.ID() == d.ID() {
					k = append(k[:j], k[j+1:]...)
					g.Top[i] = k
				}
			}
		}
	}
}

// CycleDetect will check whether a graph has cycles or not
func (g *VDG) CycleDetect() bool {
	var seen []Virtual
	var done []Virtual

	for i := range g.Top {
		if !containsV(done, i) {
			result, d := g.cycleDfs(i, seen, done)
			done = d
			if result {
				return true
			}
		}
	}
	return false
}

// Recursive Depth-First-Search; used for Cycle Detection
func (g *VDG) cycleDfs(start Virtual, seen, done []Virtual) (bool, []Virtual) {
	seen = append(seen, start)
	adj := g.Top[start]
	for _, v := range adj {
		if containsV(done, v) {
			continue
		}

		if containsV(seen, v) {
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
