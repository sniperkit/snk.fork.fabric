package fabric

// TODO: figure out the types of virtuals and what each virtual
//		interface will require in terms of UI or Temporal Node
//		assignments, etc. virtual nodes will need a
//		lifecycle / lifespan.

type Life int

const (
	Idle Life = iota
	Running
	Finished
)

// NOTE: Virtual Nodes can have UI and temporal dependents
//		and dependencies.
type Virtual interface {
	UI
	Dependents() []Node
	Dependencies() []Node
	ListProcedures() ProceduresList
	Lifecycle() Life
}
