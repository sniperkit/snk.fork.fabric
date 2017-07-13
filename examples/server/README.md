# Server Example

- REST API => Resources
- *Each* resource is it's *own* CDS
- Thus, each resource could have a POSET which wraps its main global graph
- Each resource will have a set of UIs which covers (addresses) the entirety of all nodes and edges (relationships between nodes) in the resource CDS (often a DB table, etc.)
    + A single UI is fine as well (and will be a common case)
- Every **user session** creates a single **VUI** *per resource*.
    + Each VUI contains all the user's entities of that resource type
- VUIs can be ordered in the resource's global graph by the resource POSET if necessary
- Each VUI is the space for a VDG tree. This VDG tree is the **User Request Tree** (User Request POSET) for that resource.

**Important:** when designing the system it is critical that structural updates (any updates, really) made to the CDS are reflected in each DG node.

## Idea

A single tree resource. A tree with a root node, and each branch of the tree belongs to a different user. A user can add, remove, update, and read nodes in its branch.