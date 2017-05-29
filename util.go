package fabric

// contains checks if DGNode is already in DGNode slice or not
func contains(s []DGNode, i DGNode) bool {
	for _, v := range s {
		if i.ID() == v.ID() {
			return true
		}
	}
	return false
}

func containsNode(l NodeList, np *Node) bool {
	n := *np
	for _, vp := range l {
		v := *vp
		if v.ID() == n.ID() {
			return true
		}
	}
	return false
}

func containsEdge(l EdgeList, ep *Edge) bool {
	e := *ep
	for _, vp := range l {
		v := *vp
		if v.ID() == e.ID() {
			return true
		}
	}
	return false
}
