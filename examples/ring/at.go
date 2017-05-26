package ring

import (
	"github.com/JKhawaja/fabric"
)

/* Access Types */

// TODO:

// Total-Invariance
//		Next(); READ
//		Previous(); READ
//		Front(); READ
//		Back(); READ

// No-Invariance
//		Remove(); MANIP

// Indempotent (but technically no invariance)
//		PushFront(); MANIP -- creates new node (w/ value) and puts in front
//		PushBack(); MANIP -- creates new node (w/ value) and puts in back
// 		InsertBefore(); MANIP -- creates new node (w/ value)
//		InsertAfter(); MANIP -- creates new node (w/ value)

// Value-Invariance
//		MoveToFront();MANIP
//		MoveToBack(); MANIP
//		MoveBefore(); MANIP
//		MoveAfter(); MANIP

// Adds multiple nodes and edges (inserts a ring)
//		PushFrontList(); MANIP
//		PushBackList(); MANIP
