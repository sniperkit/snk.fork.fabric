package db

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/JKhawaja/fabric"
)

// Tree ...
type Tree struct {
	Root     fabric.Node
	Sections fabric.NodeList
	Nodes    fabric.NodeList
	Edges    fabric.EdgeList
}

// NewSection takes a session id and creates a root node for a branch section
// that will be dedicated to that session (here session ids are behaving like user ids)
func (t *Tree) NewSection(id int) fabric.Node {
	// create section node (the value will be the session id)
	n := NewTreeNode(t, id)
	t.Sections = append(t.Sections, n)

	// create edge from root to section node
	e := NewTreeEdge(t, t.Root, n)
	t.Edges = append(t.Edges, e)

	return n
}

func containsNode(l fabric.NodeList, id int) bool {
	for _, v := range l {
		if v.ID() == id {
			return true
		}
	}
	return false
}

func containsEdge(l fabric.EdgeList, id int) bool {
	for _, v := range l {
		if v.ID() == id {
			return true
		}
	}
	return false
}

// CreateNode will create a node and add it to the section and tree data store
func (t *Tree) CreateNode(s fabric.Section, value interface{}) (fabric.Node, error) {

	n := CreateNode(t, value)

	// update section with new node
	nodes := *s.ListNodes()
	nodes = append(nodes, n)
	s.UpdateNodeList(&nodes)

	return n, nil
}

// TODO: this function will iterate through the sections node list twice,
// rewrite function so it does everything it needs to on a single iteration of
// the sections nodelist.
func (t *Tree) RemoveNode(s fabric.Section, id int) error {
	// verify that node is in section before being removed
	nodes := *s.ListNodes()
	if containsNode(nodes, id) {
		RemoveNode(t, id)
	} else {
		return fmt.Errorf("Node is not in section. Cannot remove.")
	}

	// remove node from section list and update section with new list
	for i, n := range nodes {
		if n.ID() == id {
			nodes = append(nodes[:i], nodes[i+1:]...)
			break
		}
	}

	s.UpdateNodeList(&nodes)

	return nil
}

// CreateEdge ...
func (t *Tree) CreateEdge(s fabric.Section, n1, n2 fabric.Node) (fabric.Edge, error) {
	// TODO: verify that both nodes are in the section ...
	e := CreateEdge(t, n1, n2)

	// update section with new edge
	elp := s.ListEdges()
	edges := *elp
	edges = append(edges, e)
	s.UpdateEdgeList(&edges)

	return e, nil
}

// TODO: this function will iterate through the sections edge list twice,
// rewrite function so it does everything it needs to on a single iteration of
// the sections edgelist.
func (t *Tree) RemoveEdge(s fabric.Section, id int) error {
	// verify that edge is in section before being removed
	edges := *s.ListEdges()
	if containsEdge(edges, id) {
		RemoveEdge(t, id)
	} else {
		return fmt.Errorf("Edge is not in section. Cannot remove.")
	}

	// remove edge from section list and update section with new list
	for i, e := range edges {
		if e.ID() == id {
			edges = append(edges[:i], edges[i+1:]...)
			break
		}
	}

	s.UpdateEdgeList(&edges)

	return nil
}

// ReadNodeValue ...
func (t *Tree) ReadNodeValue(s fabric.Section, id int) (interface{}, error) {
	// verify that node is in section before being read
	var value interface{}
	nodes := *s.ListNodes()
	if containsNode(nodes, id) {
		value = ReadNodeValue(t, id)
	} else {
		return value, fmt.Errorf("Node is not in section. Cannot read value.")
	}

	return value, nil
}

// UpdateNodeValue ...
func (t *Tree) UpdateNodeValue(s fabric.Section, id int, value interface{}) error {
	// verify that node is in section before being updated
	nodes := *s.ListNodes()
	if containsNode(nodes, id) {
		UpdateNodeValue(t, id, value)
	} else {
		return fmt.Errorf("Node is not in section. Cannot update value.")
	}

	return nil
}

// NewTree ...
func NewTree() fabric.CDS {
	t := &Tree{}
	var i interface{}
	n := NewTreeNode(t, i)
	t.Root = n

	var nl fabric.NodeList
	nl = append(nl, n)
	t.Nodes = nl

	el := make(fabric.EdgeList, 0)
	t.Edges = el

	return t
}

// GenNodeID ...
// Generate an ID for a CDS Node
func (t Tree) GenNodeID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	for _, n := range t.Nodes {
		if n.ID() == id {
			id = t.GenNodeID()
		}
	}

	return id
}

// GenEdgeID ...
// Generate an ID for a CDS Edge
func (t Tree) GenEdgeID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	for _, e := range t.Edges {
		if e.ID() == id {
			id = t.GenEdgeID()
		}
	}

	return id

}

// ListNodes ...
func (t Tree) ListNodes() fabric.NodeList {
	return t.Nodes
}

// ListEdges ...
func (t Tree) ListEdges() fabric.EdgeList {
	return t.Edges
}

// TreeNode satisfies fabric.Node interface
type TreeNode struct {
	Id    int
	Value interface{}
	Imm   bool
}

// NewTreeNode ...
func NewTreeNode(c fabric.CDS, value interface{}) fabric.Node {
	return &TreeNode{
		Id:    c.GenNodeID(),
		Value: value,
	}
}

// ID ...
func (t *TreeNode) ID() int {
	return t.Id
}

// Immutable ...
func (t *TreeNode) Immutable() bool {
	return t.Imm
}

// TreeEdge satisfies the fabric.Edge interface
type TreeEdge struct {
	Id          int
	Source      fabric.Node
	Destination fabric.Node
	Imm         bool
}

// NewTreeEdge ...
func NewTreeEdge(c fabric.CDS, s, d fabric.Node) fabric.Edge {
	return &TreeEdge{
		Id:          c.GenEdgeID(),
		Source:      s,
		Destination: d,
	}
}

// ID ...
func (t *TreeEdge) ID() int {
	return t.Id
}

// GetSource ...
func (t *TreeEdge) GetSource() fabric.Node {
	return t.Source
}

// GetDestination ...
func (t *TreeEdge) GetDestination() fabric.Node {
	return t.Destination
}

// Immutable ...
func (t *TreeEdge) Immutable() bool {
	return t.Imm
}
