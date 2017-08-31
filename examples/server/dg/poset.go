package dg

import (
	"github.com/JKhawaja/fabric"
)

// VDGPoset ...
type VDGPoset struct {
	Vdg *fabric.VDG
}

// NewVDGPoset ...
func NewVDGPoset(v *fabric.VDG) fabric.VPoset {
	return &VDGPoset{
		Vdg: v,
	}
}

// VDG ...
func (v *VDGPoset) VDG() *fabric.VDG {
	return v.Vdg
}

// InitGraph ...
func (v *VDGPoset) InitGraph(nodes []fabric.Virtual) *fabric.VDG {
	return v.Vdg
}

// Order ...
func (v *VDGPoset) Order(node fabric.Virtual) error {
	// add node to VDG
	err := v.VDG().AddTopNode(node)
	if err != nil {
		return err
	}

	for vnode := range v.VDG().Top {
		if vnode.ID() != node.ID() {
			if vnode.GetPriority() <= node.GetPriority() && !vnode.Started() {
				// create an edge from all nodes with an equivalent or larger priority integer to this node
				err := v.VDG().AddVirtualEdge(vnode.ID(), node)
				if err != nil {
					continue
				}
			} else if vnode.GetPriority() > node.GetPriority() {
				// create an edge from the node to all nodes that have a smaller priority integer
				err := v.VDG().AddVirtualEdge(node.ID(), vnode)
				if err != nil {
					continue
				}
			}
		}

		// TODO: currently the Order() method orders ALL Update()s in order of their creation,
		// but the Update()s *only* need to be ordered if they are updating the *same* CDS node
	}

	return nil
}
