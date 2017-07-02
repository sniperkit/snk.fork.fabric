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
	IsRoot() bool // specifies whether VDG node is root node or not
	Subspace() UI // the (V)UI that the Virtual node is associated to
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
	Root   *Virtual
	Top    map[Virtual][]*Virtual
	Space  []int // the set of all (V)UI ids that at least one node in the VDG has access too
}

// NewVDG will return an empty VDG graph
func NewVDG(g *Graph) *VDG {
	return &VDG{
		Global: g,
		Top:    make(map[Virtual][]*Virtual),
	}

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
func (g *VDG) CreateSignalers(np *Virtual) SignalingMap {
	sm := make(SignalingMap)

	deps := g.Dependents(np)
	for _, d := range deps {
		c := make(chan ProcedureSignals)
		sm[d.ID()] = c
	}

	return sm
}

// Signals ...
func (g *VDG) Signals(np *Virtual) SignalsMap {
	n := *np
	sm := make(SignalsMap)

	deps := g.Dependencies(np)
	for _, d := range deps {
		channels := d.ListSignalers()
		ch := channels[n.ID()]
		sm[n.ID()] = ch
	}

	return sm
}

// Dependents ...
func (g *VDG) Dependents(np *Virtual) []Virtual {
	var list []Virtual
	n := *np

	for i, v := range g.Top {
		if i.ID() != n.ID() {
			if containsVirtual(v, np) {
				list = append(list, i)
			}
		}
	}

	return list
}

// Dependencies ...
func (g *VDG) Dependencies(np *Virtual) []Virtual {
	var list []Virtual
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

// AddVirtualNode adds a node to a VDG
func (g *VDG) AddVirtualNode(node Virtual) error {
	if _, ok := g.Top[node]; !ok {
		g.Top[node] = []*Virtual{}
	} else {
		return fmt.Errorf("Node already exists in Dependency Graph.")
	}
	// Add node's subspace to graph
	g.Space = append(g.Space, node.Subspace().ID())
	return nil
}

// RemoveVirtualNode is for removing nodes from a VDG
func (g *VDG) RemoveVirtualNode(np *Virtual) {
	n := *np
	delete(g.Top, n)

	// remove all references (edges) to node in other nodes edge slices
	for _, l := range g.Top {
		if containsVirtual(l, np) {
			for j, p := range l {
				k := *p
				if k.ID() == n.ID() {
					l = append(l[:j], l[j+1:]...)
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
}

// AddVirtualEdge adds an edge to a VDG
func (g *VDG) AddVirtualEdge(source int, dest *Virtual) {
	for i, k := range g.Top {
		if i.ID() == source {
			if !containsVirtual(k, dest) {
				k = append(k, dest)
			}
		}
	}
}

// RemoveVirtualEdge removes an edge from a VDG
func (g *VDG) RemoveVirtualEdge(source int, dest *Virtual) {
	for i, k := range g.Top {
		if i.ID() == source {
			dp := *dest
			for j, vp := range k {
				v := *vp
				if v.ID() == dp.ID() {
					k = append(k[:j], k[j+1:]...)
				}
			}
		}
	}
}
