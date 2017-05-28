package ring

import (
	"github.com/JKhawaja/fabric"
)

// NOTE: The channels in Signalers are the same channels that all the dependents of
// of a node need to be reading from.

type DGNode struct {
	Id               int
	Type             fabric.NodeType
	Signalers        fabric.SignalingMap
	AccessProcedures fabric.ProcedureList
	Dependents       []fabric.DGNode
	Dependencies     []fabric.DGNode
	Signals          fabric.SignalsMap
	IsRoot           bool
	IsLeaf           bool
}
