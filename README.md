# Fabric

A Concurrency control primitives package. Utilizing dependency graphs to avoid conflict in concurrent access to a single data structure.

This code package is less of a "here are some functions and objects. Use them." And more of a "here are some interfaces. Implement them." And this allows the package to behave more as a design guidance tool rather than a strict dependency.

One thing that this package does is enable a developer to turn any data structure implementation into a Concurrent Data Structure "Fabric-Friendly" package (or easily create a new CDS package out of the original data-structure code).

## Code Generator (WIP)

- Provide a small DSL that can be used to generate boilerplate for projects that will be using the `fabric` package.

## Extra Notes

Any real node thread can spawn VDGs attached to the same UI that they are attached to as long as they do not allow the VDG nodes to be one of it's own dependencies.

This might be useful for when we want too execute one atomic procedure, signal completed, let the dependent node execute its procedure, but then execute the next atomic procedure in our thread AFTER the dependent has signaled completed. However, we cannot have cyclic dependencies, so instead we spawn a virtual thread containing the next operation that will have the current dependent node as one of it's *dependency* nodes and will execute once our current dependent node signals complete.

The problem with the above is that VDG virtual nodes are *not allowed* to have real dependencies or dependents!! Thus, the idea of virtual spawning goes right out the window!! ...