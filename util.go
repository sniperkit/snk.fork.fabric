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

func edgeContains(l EdgeList, e Edge) bool {
	for _, v := range l {
		if e.ID() == v.ID() {
			return true
		}
	}
	return false
}

// TODO: change to containsNode
func containsNode(l NodeList, n Node) bool {
	for _, v := range l {
		if v.ID() == n.ID() {
			return true
		}
	}
	return false
}

// TODO: change to containsEdge
func containsEdge(l EdgeList, e Edge) bool {
	for _, v := range l {
		if v.ID() == e.ID() {
			return true
		}
	}
	return false
}
