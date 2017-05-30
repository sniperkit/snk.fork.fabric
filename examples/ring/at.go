package ring

/* Access Types */

// NOTE: we are treating access types almost like classes of functions.
// 	These classes are defined by some function type. There are numerous other
// 	classes we could add to our collection here.

// 	For example: we could have a function type for removing multiple
// 	elements from the ring. Or, we could have function types for
// 	updating a value in an element e.g. if the value type is integer,
// 	the functions could be 'add' and 'subtract', etc.

// Total-Invariance
//	Next(); READ
//	Previous(); READ

type ElementRead func(*ElementNode) (*ElementNode, error)

func (r *ElementRead) Name() string {
	var n string
	return n
}

func (r *ElementRead) Priority() int {
	var p int
	return p
}

func (r *ElementRead) Commit() error {
	return nil
}

func (r *ElementRead) Rollback() error {
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

func (r *RingRead) Name() string {
	var n string
	return n
}

func (r *RingRead) Priority() int {
	var p int
	return p
}

func (r *RingRead) Commit() error {
	return nil
}

func (r *RingRead) Rollback() error {
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

func (e *ElementDelete) Name() string {
	var n string
	return n
}

func (e *ElementDelete) Priority() int {
	var p int
	return p
}

func (e *ElementDelete) Commit() error {
	return nil
}

func (e *ElementDelete) Rollback() error {
	return nil
}

func (e *ElementDelete) InvariantNode(node *ElementNode) bool {
	var b bool
	return b
}

func (e *ElementDelete) InvariantEdge(edge *ElementEdge) bool {
	var b bool
	return b
}

// Create (default position)
// 	PushFront(); MANIP -- creates new node (w/ value) and puts in front
// 	PushBack(); MANIP -- creates new node (w/ value) and puts in back

type CreateElement func(interface{}) (*ElementNode, error)

func (c *CreateElement) Name() string {
	var n string
	return n
}

func (c *CreateElement) Priority() int {
	var p int
	return p
}

func (c *CreateElement) Commit() error {
	return nil
}

func (c *CreateElement) Rollback() error {
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

func (c *CreateInsertElement) Name() string {
	var n string
	return n
}

func (c *CreateInsertElement) Priority() int {
	var p int
	return p
}

func (c *CreateInsertElement) Commit() error {
	return nil
}

func (c *CreateInsertElement) Rollback() error {
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

func (v *ValueInvariant) Name() string {
	var n string
	return n
}

func (v *ValueInvariant) Priority() int {
	var p int
	return p
}

func (v *ValueInvariant) Commit() error {
	return nil
}

func (v *ValueInvariant) Rollback() error {
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

func (m *MarkValueInvariant) Name() string {
	var n string
	return n
}

func (m *MarkValueInvariant) Priority() int {
	var p int
	return p
}

func (m *MarkValueInvariant) Commit() error {
	return nil
}

func (m *MarkValueInvariant) Rollback() error {
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

func (r *RingInsert) Name() string {
	var n string
	return n
}

func (r *RingInsert) Priority() int {
	var p int
	return p
}

func (r *RingInsert) Commit() error {
	return nil
}

func (r *RingInsert) Rollback() error {
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
