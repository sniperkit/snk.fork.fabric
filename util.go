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

func containsDGNode(s []*DGNode, ip *DGNode) bool {
	i := *ip
	for _, vp := range s {
		v := *vp
		if i.ID() == v.ID() {
			return true
		}
	}

	return false
}

func containsVirtual(s []*Virtual, ip *Virtual) bool {
	for _, vp := range s {
		v := *vp
		i := *ip
		if i.ID() == v.ID() {
			return true
		}
	}
	return false
}

// containsNode checks if a CDS node (reference) is in a NodeList
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

// containsEdge checks if a CDS edge (reference) is in an Edglist
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
