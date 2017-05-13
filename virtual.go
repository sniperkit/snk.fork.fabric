package fabric

// TODO: figure out the types of virtuals and what each virtual interface will
// 		require in terms of UI or Temporal Node assignments, etc.
//		virtual nodes will need a lifecycle / lifespan

type Life int

const (
	Idle Life = iota
	Running
	Finished
)

type Virtual interface {
	UI
	Lifecycle() Life
}
