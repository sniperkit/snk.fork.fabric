package db

// ElementNode satisfies fabric.Node interface
type ElementNode struct {
	Id    int
	Value interface{}
	L     *List
	Imm   bool
}

func (e ElementNode) ID() int {
	return e.Id
}

func (e ElementNode) Immutable() bool {
	return e.Imm
}
