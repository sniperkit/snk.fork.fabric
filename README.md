# Fabric

A Concurrency control primitives package. Utilizing dependency graphs to avoid conflict in concurrent access to a single data structure.

This code package is less of a "here are some functions and objects. Use them." And more of a "here are some interfaces. Implement them." And this allows the package to behave more as a design guidance tool rather than a strict dependency.

One thing that this package does is enable a developer to turn any data structure implementation into a Concurrent Data Structure "Fabric-Friendly" package (or easily create a new CDS package out of the original data-structure code).

## Code Generator (WIP)

- Provide a small DSL that can be used to generate boilerplate for projects that will be using the `fabric` package.

## Extra Notes

**VDG:** temporary secondary dependency graph, which behaves independent of our main dependency graph, and only behaves on some given subset of UIs.

A VDG is basically equivalent to an internal partial-ordering of operations within a thread.

Except, that if we consider a VDG a thread, then we have to consider it as a thread node which has no dependents or dependencies. Thus, a VDG is like a virtualized thread node that allows us to manipulate multiple (V)UIs instead of only a single V(UI).

VDGs can present a potentiall concurrency-safety issue if you use them without taking into consideration the assumptions of what they are.

Make sure that all VDG operations are safe to operate independently from the main dependency graph.

In summation, using a VDG is making the following assumptions -- all access types that the VDG will be using are okay to be used without dependency management from our graph. As well as all VDG operations may occur at any given time independent of our global graph as well. They are simply a temporary secondary dependency graph that will operate on our CDS independent of our main graph.



