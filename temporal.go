package fabric

// This is what is assigned to Threads beyond the first thread assigned to a particular UI
type Temporal interface {
	UI
	ListProcedures() []ProceduresList
}
