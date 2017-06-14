package ring

import (
	"github.com/JKhawaja/fabric"
)

/* Access Types */

// NOTE: we are treating access types almost like classes of functions.
// 	These classes are defined by some function type. There are numerous other
// 	classes we could add to our collection here.

// 	For example: we could have a function type for removing multiple
// 	elements from the ring. Or, we could have function types for
// 	updating a value in an element e.g. if the value type is 'integer',
// 	the functions could be 'add' and 'subtract', etc.

// Total-Invariance
//	Next(); READ
//	Previous(); READ

type ElementRead func(*ElementNode) (*ElementNode, error)

func (r *ElementRead) Class() string {
	return "ElementRead"
}

func (r *ElementRead) Priority() int {
	var p int
	return p
}

func (r *ElementRead) Commit(np *fabric.DGNode) error {
	return nil
}

func (r *ElementRead) Rollback(np fabric.RestoreNodes, el fabric.RestoreEdges) error {
	return nil
}

func (e *ElementRead) InvariantNode(node *ElementNode) bool {
	var b bool
	return b
}

func (e *ElementRead) InvariantEdge(edge *ElementEdge) bool {
	var b bool
	return b
}

// Total-Invariance
//	Front(); READ
//	Back(); READ

type RingRead func(*Ring) (*ElementNode, error)

func (r *RingRead) Class() string {
	return "RingRead"
}

func (r *RingRead) Priority() int {
	var p int
	return p
}

func (r *RingRead) Commit(np *fabric.DGNode) error {
	return nil
}

func (r *RingRead) Rollback(np *fabric.DGNode) error {
	return nil
}

func (r *RingRead) InvariantNode(n *ElementNode) bool {
	var b bool
	return b
}

func (r *RingRead) InvariantEdge(e *ElementEdge) bool {
	var b bool
	return b
}

// Delete an element
// Remove(); MANIP

type ElementDelete func(*ElementNode) (interface{}, error)

func (e *ElementDelete) Class() string {
	return "ElementDelete"
}

func (e *ElementDelete) Priority() int {
	var p int
	return p
}

func (e *ElementDelete) Commit(n *fabric.DGNode) error {
	// TODO: Commit can be used to Signal when an operation is complete
	// to the dependent nodes for the current DG Node (which is passed
	// as a reference argument to this method).
	return nil
}

func (e *ElementDelete) Rollback(nl fabric.RestoreNodes, el fabric.RestoreEdges) error {

	// TODO: Rollbacks should work based on a snapshot i.e.
	//	a rollback should do nothing more than a restoration overwrite
	//  the beginning of a procedure takes the value of a node
	return nil
}

func (e *ElementDelete) InvariantNode(node *ElementNode) bool {
	var b bool
	// TODO: pass a CDS node into this method to check if it is invariant
	// 	to procedures of the function type that this method is defined for.
	// EXAMPLE: In this case, the ElementDelete function type may not be
	//	effective on "root" or "leaf" nodes, etc.
	// NOTE: that all invariant node and edge methods should check for if
	//	the entity is marked as Immutable or not.
	if node.Immutable() {
		b = true
	}

	return b
}

func (e *ElementDelete) InvariantEdge(edge *ElementEdge) bool {
	var b bool
	// TODO: pass a CDS edge into this method to check if it is invariant
	// 	to procedures of the function type that this method is defined for.
	if edge.Immutable() {
		b = true
	}
	return b
}

// Create (default position)
// 	PushFront(); MANIP -- creates new node (w/ value) and puts in front
// 	PushBack(); MANIP -- creates new node (w/ value) and puts in back

type CreateElement func(interface{}) (*ElementNode, error)

func (c *CreateElement) Class() string {
	return "CreateElement"
}

func (c *CreateElement) Priority() int {
	var p int
	return p
}

func (c *CreateElement) Commit(np *fabric.DGNode) error {
	return nil
}

func (c *CreateElement) Rollback(np *fabric.DGNode) error {
	return nil
}

func (c *CreateElement) InvariantNode(node *ElementNode) bool {
	var b bool
	return b
}

func (c *CreateElement) InvariantEdge(edge *ElementEdge) bool {
	var b bool
	return b
}

// Create (w/ chosen position)
// 	InsertBefore(); MANIP -- creates new node (w/ value)
//	InsertAfter(); MANIP -- creates new node (w/ value)

type CreateInsertElement func(interface{}, *ElementNode) (*ElementNode, error)

func (c *CreateInsertElement) Class() string {
	return "CreateInsertElement"
}

func (c *CreateInsertElement) Priority() int {
	var p int
	return p
}

func (c *CreateInsertElement) Commit(np *fabric.DGNode) error {
	return nil
}

func (c *CreateInsertElement) Rollback(np *fabric.DGNode) error {
	return nil
}

func (c *CreateInsertElement) InvariantNode(node *ElementNode) bool {
	var b bool
	return b
}

func (c *CreateInsertElement) InvariantEdge(edge *ElementEdge) bool {
	var b bool
	return b
}

// Value-Invariance
// 	MoveToFront();MANIP
// 	MoveToBack(); MANIP

type ValueInvariant func(*ElementNode) error

func (v *ValueInvariant) Class() string {
	return "ValueInvariant"
}

func (v *ValueInvariant) Priority() int {
	var p int
	return p
}

func (v *ValueInvariant) Commit(np *fabric.DGNode) error {
	return nil
}

func (v *ValueInvariant) Rollback(np *fabric.DGNode) error {
	return nil
}

func (v *ValueInvariant) InvariantNode(node *ElementNode) bool {
	var b bool
	return b
}

func (v *ValueInvariant) InvariantEdge(edge *ElementEdge) bool {
	var b bool
	return b
}

// MoveBefore(); MANIP
// MoveAfter(); MANIP

type MarkValueInvariant func(*ElementNode, *ElementNode) error

func (m *MarkValueInvariant) Class() string {
	return "MarkValueInvariant"
}

func (m *MarkValueInvariant) Priority() int {
	var p int
	return p
}

func (m *MarkValueInvariant) Commit(np *fabric.DGNode) error {
	return nil
}

func (m *MarkValueInvariant) Rollback(np *fabric.DGNode) error {
	return nil
}

func (m *MarkValueInvariant) InvariantNode(node *ElementNode) bool {
	var b bool
	return b
}

func (m *MarkValueInvariant) InvariantEdge(edge *ElementEdge) bool {
	var b bool
	return b
}

// Adds multiple nodes and edges (inserts a ring)
// 	PushFrontList(); MANIP
// 	PushBackList(); MANIP

type RingInsert func(*Ring) error

func (r *RingInsert) Class() string {
	return "RingInsert"
}

func (r *RingInsert) Priority() int {
	var p int
	return p
}

func (r *RingInsert) Commit(np *fabric.DGNode) error {
	return nil
}

func (r *RingInsert) Rollback(np *fabric.DGNode) error {
	return nil
}

func (r *RingInsert) InvariantNode(node *ElementNode) bool {
	var b bool
	return b
}

func (r *RingInsert) InvariantEdge(edge *ElementEdge) bool {
	var b bool
	return b
}
