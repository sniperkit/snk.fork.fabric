package db

import (
	"fmt"

	"github.com/JKhawaja/fabric"
)

// AddTreeNode ...
type AddTreeNode func(*Tree, interface{}) fabric.Node

// ID ...
func (a AddTreeNode) ID() int {
	return 0
}

// Priority ...
func (a AddTreeNode) Priority() int {
	return 3
}

// Commit ...
func (a AddTreeNode) Commit(n fabric.DGNode) error {

	// Get the UI being affected
	var ui fabric.UI
	switch n.GetType() {
	case fabric.UINode, fabric.VUINode:
		u, ok := n.(fabric.UI)
		if !ok {
			return fmt.Errorf("Could not convert node to (V)UI node.")
		}
		ui = u
	case fabric.TemporalNode, fabric.VirtualTemporalNode:
		t, ok := n.(fabric.Temporal)
		if !ok {
			return fmt.Errorf("Could not convert node to (V)Temporal node.")
		}
		uis := t.GetRoots()
		ui = uis[0]
	case fabric.VDGNode:
		v, ok := n.(fabric.Virtual)
		if !ok {
			return fmt.Errorf("Could not convert node to Virtual (VDG) node.")
		}
		ui = v.Subspace()
	}

	// Create the signal
	s := fabric.NodeSignal{
		AccessType: a.ID(),
		Value:      fabric.Completed,
		Space:      ui,
	}

	// Send the Signal
	n.Signal(s)
	return nil
}

// Rollback ...
func (a AddTreeNode) Rollback(np fabric.RestoreNodes, el fabric.RestoreEdges) error {
	return nil
}

// AddTreeEdge ...
type AddTreeEdge func(*Tree, fabric.Node, fabric.Node) fabric.Edge

// ID ...
func (a AddTreeEdge) ID() int {
	return 1
}

// Priority ...
func (a AddTreeEdge) Priority() int {
	return 2
}

// Commit ...
func (a AddTreeEdge) Commit(n fabric.DGNode) error {
	// Get the UI being affected
	var ui fabric.UI
	switch n.GetType() {
	case fabric.UINode, fabric.VUINode:
		u, ok := n.(fabric.UI)
		if !ok {
			return fmt.Errorf("Could not convert node to (V)UI node.")
		}
		ui = u
	case fabric.TemporalNode, fabric.VirtualTemporalNode:
		t, ok := n.(fabric.Temporal)
		if !ok {
			return fmt.Errorf("Could not convert node to (V)Temporal node.")
		}
		uis := t.GetRoots()
		ui = uis[0]
	case fabric.VDGNode:
		v, ok := n.(fabric.Virtual)
		if !ok {
			return fmt.Errorf("Could not convert node to Virtual (VDG) node.")
		}
		ui = v.Subspace()
	}

	// Create the signal
	s := fabric.NodeSignal{
		AccessType: a.ID(),
		Value:      fabric.Completed,
		Space:      ui,
	}

	// Send the Signal
	n.Signal(s)
	return nil
}

// Rollback ...
func (a AddTreeEdge) Rollback(np fabric.RestoreNodes, el fabric.RestoreEdges) error {
	return nil
}

// DeleteTreeEntity ...
type DeleteTreeEntity func(*Tree, int)

// ID ...
func (d DeleteTreeEntity) ID() int {
	return 2
}

// Priority ...
func (d DeleteTreeEntity) Priority() int {
	return 1
}

// Commit ...
func (d DeleteTreeEntity) Commit(n fabric.DGNode) error {
	// Get the UI being affected
	var ui fabric.UI
	switch n.GetType() {
	case fabric.UINode, fabric.VUINode:
		u, ok := n.(fabric.UI)
		if !ok {
			return fmt.Errorf("Could not convert node to (V)UI node.")
		}
		ui = u
	case fabric.TemporalNode, fabric.VirtualTemporalNode:
		t, ok := n.(fabric.Temporal)
		if !ok {
			return fmt.Errorf("Could not convert node to (V)Temporal node.")
		}
		uis := t.GetRoots()
		ui = uis[0]
	case fabric.VDGNode:
		v, ok := n.(fabric.Virtual)
		if !ok {
			return fmt.Errorf("Could not convert node to Virtual (VDG) node.")
		}
		ui = v.Subspace()
	}

	// Create the signal
	s := fabric.NodeSignal{
		AccessType: d.ID(),
		Value:      fabric.Completed,
		Space:      ui,
	}

	// Send the Signal
	n.Signal(s)
	return nil
}

// Rollback ...
func (d DeleteTreeEntity) Rollback(np fabric.RestoreNodes, el fabric.RestoreEdges) error {
	return nil
}

// ReadTreeNode ...
type ReadTreeNode func(*Tree, int) interface{}

// ID ...
func (r ReadTreeNode) ID() int {
	return 3
}

// Priority ...
func (r ReadTreeNode) Priority() int {
	return 4
}

// Commit ...
func (r ReadTreeNode) Commit(n fabric.DGNode) error {
	// Get the UI being affected
	var ui fabric.UI
	switch n.GetType() {
	case fabric.UINode, fabric.VUINode:
		u, ok := n.(fabric.UI)
		if !ok {
			return fmt.Errorf("Could not convert node to (V)UI node.")
		}
		ui = u
	case fabric.TemporalNode, fabric.VirtualTemporalNode:
		t, ok := n.(fabric.Temporal)
		if !ok {
			return fmt.Errorf("Could not convert node to (V)Temporal node.")
		}
		uis := t.GetRoots()
		ui = uis[0]
	case fabric.VDGNode:
		v, ok := n.(fabric.Virtual)
		if !ok {
			return fmt.Errorf("Could not convert node to Virtual (VDG) node.")
		}
		ui = v.Subspace()
	}

	// Create the signal
	s := fabric.NodeSignal{
		AccessType: r.ID(),
		Value:      fabric.Completed,
		Space:      ui,
	}

	// Send the Signal
	n.Signal(s)
	return nil
}

// Rollback ...
func (r ReadTreeNode) Rollback(np fabric.RestoreNodes, el fabric.RestoreEdges) error {
	return nil
}

// UpdateTreeNode ...
type UpdateTreeNode func(*Tree, int, interface{})

// ID ...
func (u UpdateTreeNode) ID() int {
	return 4
}

// Priority ...
func (u UpdateTreeNode) Priority() int {
	return 5
}

// Commit ...
func (u UpdateTreeNode) Commit(n fabric.DGNode) error {
	// Get the UI being affected
	var ui fabric.UI
	switch n.GetType() {
	case fabric.UINode, fabric.VUINode:
		u, ok := n.(fabric.UI)
		if !ok {
			return fmt.Errorf("Could not convert node to (V)UI node.")
		}
		ui = u
	case fabric.TemporalNode, fabric.VirtualTemporalNode:
		t, ok := n.(fabric.Temporal)
		if !ok {
			return fmt.Errorf("Could not convert node to (V)Temporal node.")
		}
		uis := t.GetRoots()
		ui = uis[0]
	case fabric.VDGNode:
		v, ok := n.(fabric.Virtual)
		if !ok {
			return fmt.Errorf("Could not convert node to Virtual (VDG) node.")
		}
		ui = v.Subspace()
	}

	// Create the signal
	s := fabric.NodeSignal{
		AccessType: u.ID(),
		Value:      fabric.Completed,
		Space:      ui,
	}

	// Send the Signal
	n.Signal(s)
	return nil
}

// Rollback ...
func (u UpdateTreeNode) Rollback(np fabric.RestoreNodes, el fabric.RestoreEdges) error {
	return nil
}
