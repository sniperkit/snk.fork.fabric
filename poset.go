package fabric

// TODO: helpers for defining partial orderings on dependency graphs,
//		or at least on sub-graphs (e.g. Temporal DAGs, etc.)

// Permanent Graphs: (need Ordering Rules for each)
// 		Disjoint DAG of UIs
// 		Each UI is the root of a Temporal DAG

// TODO: Access Type Priority Ordering
//		if a DGNode has an Access type with priority lower than
//		all other Access Types in another DGNode, then it automatically
//		becomes a dependency of that node.

// Dynamic Nodes/Sub-Graphs: (need Ordering Rules for each)
// 		VUIs -- added to the disjoint DAG of UIS
// 		VDGs -- attached to some set of UIs (including VUIs)
