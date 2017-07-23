package db

import (
	"github.com/JKhawaja/fabric"
)

/*
	NOTE: The process of creating Access Types:
		- first create your basic access methods e.g. read, write, etc.
		- Then create function types matching the function signatures of the access methods
		- Then create fabric methods for those function types, and convert the access methods to those function types
		- be sure to provide a return error value for any procedures that you want to be able to commit/rollback
	Now you have all your original access methods but in a fabricated format.
*/

/*
	NOTE: we can also wrap all these access procedures in methods:
	func (t *Tree) AddNode(value interface{}) *TreeNode {
		return CreateNode(t, value)
	}
*/

func addNode(t *Tree, value interface{}) *TreeNode {
	n := TreeNode{
		Id:    t.GenNodeID(),
		Value: value,
	}

	var i interface{} = n
	in := i.(fabric.Node)

	t.Nodes = append(t.Nodes, &in)

	return &n
}

// CreateNode ...
// EXAMPLE: CreateNode(myTreepointer, myValue)
var CreateNode = AddTreeNode(addNode)

func addEdge(t *Tree, s, d *TreeNode) *TreeEdge {
	e := TreeEdge{
		Id:          t.GenEdgeID(),
		Source:      s,
		Destination: d,
	}

	var i interface{} = e
	ie := i.(fabric.Edge)

	t.Edges = append(t.Edges, &ie)

	return &e
}

// CreateEdge ...
// EXAMPLE: CreateTreeEdge(myTreepointer, myFirstNode, mySecondNode)
var CreateEdge = AddTreeEdge(addEdge)

func deleteNode(t *Tree, id int) {
	// remove node from node list
	for i, np := range t.Nodes {
		node := *np
		if node.ID() == id {
			t.Nodes = append(t.Nodes[:i], t.Nodes[i+1:]...)
			break
		}
	}

	// remove all edges that have node as destination
	for i, ep := range t.Edges {
		edge := *ep
		dest := *edge.GetDestination()
		if dest.ID() == id {
			t.Edges = append(t.Edges[:i], t.Edges[i+1:]...)
		}
	}
}

// RemoveNode ...
// EXAMPLE: RemoveNode(myTreepointer, myNodeID)
var RemoveNode = DeleteTreeEntity(deleteNode)

func deleteEdge(t *Tree, id int) {
	for i, ep := range t.Edges {
		edge := *ep
		if edge.ID() == id {
			t.Edges = append(t.Edges[:i], t.Edges[i+1:]...)
			break
		}
	}
}

// RemoveEdge ...
// EXAMPLE: RemoveEdge(myTreepointer, myEdgeID)
var RemoveEdge = DeleteTreeEntity(deleteEdge)

func readNode(t *Tree, id int) interface{} {
	for _, node := range t.Nodes {
		n := *node
		if n.ID() == id {
			tn := n.(TreeNode)
			return tn.Value
		}
	}

	return nil
}

// ReadValue ...
// EXAMPLE: ReadNodeValue(myTreepointer, myNodeID)
var ReadNodeValue = ReadTreeNode(readNode)

func updateNode(t *Tree, id int, v interface{}) {
	for i, np := range t.Nodes {
		node := *np
		if node.ID() == id {
			tn := node.(TreeNode)
			tn.Value = v
			var in interface{} = tn
			newNode := in.(fabric.Node)
			t.Nodes[i] = &newNode
		}
	}
}

// UpdateValue ...
// EXAMPLE: UpdateNodeValue(myTreepointer, myNodeID, myValue)
var UpdateNodeValue = UpdateTreeNode(updateNode)
