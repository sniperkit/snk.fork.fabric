package ring

import (
	"github.com/JKhawaja/fabric"
)

/* Access Types */

// NOTE: we are treating access types almost like classes of functions.
//		These classes are defined by some function type. There are numerous other
//		classes we could add to our collection here.

//		For example: we could have a function type for removing multiple
//		elements from the ring. Or, we could have function types for
//		updating a value in an element e.g. if the value type is integer,
//		the functions could be 'add' and 'subtract', etc.

// Total-Invariance
//	Next(); READ
//	Previous(); READ

type ElementRead func(*Element) *Element

func (r *ElementRead) Name() string {
	return ""
}

func (r *ElementRead) Priority() int {
	return 0
}

func (r *ElementRead) Commit() error {
	return nil
}

func (r *ElementRead) Rollback() error {
	return nil
}

func (r *ElementRead) InvariantNodes(s fabric.Section) fabric.NodeList {
	var nl fabric.NodeList

	return nl
}

func (r *ElementRead) InvariantEdges(s fabric.Section) fabric.EdgeList {
	var el fabric.EdgeList

	return el
}

// Total-Invariance
//	Front(); READ
//	Back(); READ

type RingRead func(*Ring) *Element

func (r *RingRead) Name() string {
	return ""

}

func (r *RingRead) Priority() int {
	return 0

}

func (r *RingRead) Commit() error {
	return nil
}

func (r *RingRead) Rollback() error {
	return nil
}

func (r *RingRead) InvariantNodes(s fabric.Section) fabric.NodeList {
	var nl fabric.NodeList
	return nl
}

func (r *RingRead) InvariantEdges(s fabric.Section) fabric.EdgeList {
	var el fabric.EdgeList
	return el
}

// Delete an element
// Remove(); MANIP

type ElementDelete func(*Element) interface{}

func (e *ElementDelete) Name() string {
	return ""

}

func (e *ElementDelete) Priority() int {
	return 0

}

func (e *ElementDelete) Commit() error {
	return nil
}

func (e *ElementDelete) Rollback() error {
	return nil
}

func (e *ElementDelete) InvariantNodes(s fabric.Section) fabric.NodeList {
	var nl fabric.NodeList
	return nl
}

func (e *ElementDelete) InvariantEdges(s fabric.Section) fabric.EdgeList {
	var el fabric.EdgeList
	return el
}

// Create (default position)
//	PushFront(); MANIP -- creates new node (w/ value) and puts in front
//	PushBack(); MANIP -- creates new node (w/ value) and puts in back

type CreateElement func(interface{}) *Element

func (c *CreateElement) Name() string {
	return ""

}

func (c *CreateElement) Priority() int {
	return 0

}

func (c *CreateElement) Commit() error {
	return nil
}

func (c *CreateElement) Rollback() error {
	return nil
}

func (c *CreateElement) InvariantNodes(s fabric.Section) fabric.NodeList {
	var nl fabric.NodeList
	return nl
}

func (c *CreateElement) InvariantEdges(s fabric.Section) fabric.EdgeList {
	var el fabric.EdgeList
	return el
}

// Create (w/ chosen position)
// 	InsertBefore(); MANIP -- creates new node (w/ value)
//	InsertAfter(); MANIP -- creates new node (w/ value)

type CreateInsertElement func(interface{}, *Element) *Element

func (c *CreateInsertElement) Name() string {
	return ""

}

func (c *CreateInsertElement) Priority() int {
	return 0

}

func (c *CreateInsertElement) Commit() error {
	return nil
}

func (c *CreateInsertElement) Rollback() error {
	return nil
}

func (c *CreateInsertElement) InvariantNodes(s fabric.Section) fabric.NodeList {
	var nl fabric.NodeList
	return nl
}

func (c *CreateInsertElement) InvariantEdges(s fabric.Section) fabric.EdgeList {
	var el fabric.EdgeList
	return el
}

// Value-Invariance
//		MoveToFront();MANIP
//		MoveToBack(); MANIP

type ValueInvariant func(*Element)

func (v *ValueInvariant) Name() string {
	return ""

}

func (v *ValueInvariant) Priority() int {
	return 0

}

func (v *ValueInvariant) Commit() error {
	return nil
}

func (v *ValueInvariant) Rollback() error {
	return nil
}

func (v *ValueInvariant) InvariantNodes(s fabric.Section) fabric.NodeList {
	var nl fabric.NodeList
	return nl
}

func (v *ValueInvariant) InvariantEdges(s fabric.Section) fabric.EdgeList {
	var el fabric.EdgeList
	return el
}

// MoveBefore(); MANIP
// MoveAfter(); MANIP

type MarkValueInvariant func(*Element, *Element)

func (m *MarkValueInvariant) Name() string {
	return ""

}

func (m *MarkValueInvariant) Priority() int {
	return 0

}

func (m *MarkValueInvariant) Commit() error {
	return nil
}

func (m *MarkValueInvariant) Rollback() error {
	return nil
}

func (m *MarkValueInvariant) InvariantNodes(s fabric.Section) fabric.NodeList {
	var nl fabric.NodeList
	return nl
}

func (m *MarkValueInvariant) InvariantEdges(s fabric.Section) fabric.EdgeList {
	var el fabric.EdgeList
	return el
}

// Adds multiple nodes and edges (inserts a ring)
//		PushFrontList(); MANIP
//		PushBackList(); MANIP

type RingInsert func(*Ring)

func (r *RingInsert) Name() string {
	return ""

}

func (r *RingInsert) Priority() int {
	return 0

}

func (r *RingInsert) Commit() error {
	return nil
}

func (r *RingInsert) Rollback() error {
	return nil
}

func (r *RingInsert) InvariantNodes(s fabric.Section) fabric.NodeList {
	var nl fabric.NodeList
	return nl
}

func (r *RingInsert) InvariantEdges(s fabric.Section) fabric.EdgeList {
	var el fabric.EdgeList
	return el
}
