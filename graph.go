package fabric

type Graph struct {
	Nodes []int
	Edges []int
}

func (g *Graph) CycleDetect() (bool, error) {
	var seen []int
	var unseen []int
	var done []int
}

func dfs(seen, unseen, done []int) {
}
