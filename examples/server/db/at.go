package db

import (
	"github.com/JKhawaja/fabric"
)

// ElementRead ...
type ElementRead func(*ElementNode) (*ElementNode, error)

// ID ...
func (r *ElementRead) ID() int {
	return 0
}

// Priority ...
func (r *ElementRead) Priority() int {
	return 0
}

// Commit ...
func (r *ElementRead) Commit(np *fabric.DGNode) error {
	return nil
}

// Rollback ...
func (r *ElementRead) Rollback(np fabric.RestoreNodes, el fabric.RestoreEdges) error {
	return nil
}

// AddNode will create a new Tree node
func (t *Tree) AddNode() *TreeNode {
	n := TreeNode{
		Id: t.GenNodeID(),
	}

	var i interface{} = n
	in := i.(fabric.Node)

	t.Nodes = append(l.Nodes, &in)

	return &n
}

// AddEdge ...
func (t *Tree) AddEdge(s, d *TreeNode) {
	e := TreeEdge{
		Id:          t.GenEdgeID(),
		Source:      s,
		Destination: d,
	}

	var i interface{} = e
	ie := i.(fabric.Edge)

	t.Edges = append(t.Edges, &ie)
}

// DeleteNode ...
func (t *Tree) DeleteNode(n *TreeNode) {

	// TODO: remove node from node list
	// remove all edges that have node as destination

}

// DeleteEdge ...
func (t *Tree) DeleteEdge(e *TreeEdge) {
	// TODO: remove edge from edge list
}

// ReadNode ...
func (t *Tree) ReadNode(n *TreeNode) interface{} {
	// TODO: return value stored at node
}

// UpdateNode ...
func (t *Tree) UpdateNode(n *TreeNode, v interface{}) {
	// TODO: update node with supplied value
}
