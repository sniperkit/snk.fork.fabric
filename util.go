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

func containsVirtual(s []Virtual, i Virtual) bool {
	for _, v := range s {
		if i.ID() == v.ID() {
			return true
		}
	}
	return false
}

// ContainsNode checks if a CDS node (reference) is in a NodeList
func ContainsNode(l NodeList, n Node) bool {
	for _, v := range l {
		if v.ID() == n.ID() {
			return true
		}
	}
	return false
}

// ContainsEdge checks if a CDS edge (reference) is in an Edglist
func ContainsEdge(l EdgeList, e Edge) bool {
	for _, v := range l {
		if v.ID() == e.ID() {
			return true
		}
	}
	return false
}
