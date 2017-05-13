package fabric

// wrap data structure elements to become generic CDS Nodes
type Node interface {
	ID() int // returns node id
	Immutable() bool
}

type NodeList []Node

// REFACTOR: Determine how to add an immutable boolean to an edge
type EdgesMap map[Node][]Node

// add these methods to data structure objects to use as CDS
type CDS interface {
	ListNodes() (NodeList, error)               // traverse CDS and convert CDS elements into being "Nodes" (return new list of Nodes)
	ListEdges(nodes NodeList) (EdgesMap, error) // use Nodes List generated by ListNodes() to add elements edges to map
}

/*
	DS = Data Structure; used when a UI will have
	access to entire CDS.
*/
type DS struct {
	Nodes []int
	Edges map[int][]int
}

func NewDS(nodes []int, edges map[int][]int) *DS {
	return &DS{
		Nodes: nodes,
		Edges: edges,
	}
}

func (s *DS) NodeCount() int {
	return len(s.Nodes)
}

func (s *DS) EdgeCount() int {
	var total int
	for i, v := range s.Edges {
		total += len(v)
	}
	return total
}
