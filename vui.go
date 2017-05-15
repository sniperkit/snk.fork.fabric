package fabric

type VUI interface {
	UI
	Dependents() []Node
	Dependencies() []Node
	ListProcedures() ProceduresList
}
