package dg

import (
	"github.com/JKhawaja/fabric"
)

// VDGPoset ...
type VDGPoset struct {
	Vdg *fabric.VDG
}

// NewVDGPoset ...
func NewVDGPoset(v *fabric.VDG) *VDGPoset {
	return &VDGPoset{
		Vdg: v,
	}
}

// VDG ...
func (v *VDGPoset) VDG() *fabric.VDG {
	return v.Vdg
}

// GenerateGraph ...
func (v *VDGPoset) GenerateGraph(nodes []fabric.Virtual) *fabric.VDG {
	return v.Vdg
}

// Order ...
func (v *VDGPoset) Order(node fabric.Virtual) {
	// TODO: order a new virtual node in the VDG according to Priority
}
