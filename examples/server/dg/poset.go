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

// GenerateGraph ...
func (v *VDGPoset) GenerateGraph(nodes []fabric.Virtual) *fabric.VDG {
	return v.Vdg
}

// Order ...
func (v *VDGPoset) Order(node fabric.Virtual) fabric.Virtual {
	// add node to VDG
	n, _ := v.VDG().AddTopNode(node)

	for vnode := range v.VDG().Top {
		if vnode.GetPriority() < node.GetPriority() && !vnode.Started() {
			// create an edge from all nodes with a larger priority integer to this node
			err := v.VDG().AddVirtualEdge(vnode.ID(), n)
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

		// TODO: order Update nodes behind other already existing update nodes;
		// technically, the Updates only need to be ordered if they are updating the same node
		// this is to avoid: "lost updates" (two transactions are updating the same piece of data and one update gets overwritten)
	}

	// return pointer to nodes location in VDG
	return n
}
