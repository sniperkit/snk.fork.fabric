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
