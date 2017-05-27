# Fabric

A Concurrency control primitives package. Utilizing dependency graphs to avoid conflict in concurrent access to a single data structure.

This code package is less of a "here are some functions and objects. Use them." And more of a "here are some interfaces. Implement them." And this allows the package to behave more as a design guidance tool rather than a strict dependency.

In other words: **the purpose of this package is to use it in creating your own CDS packages**. Another option is to take an existing data structure package and **fabric-ate** (hehe, get it?) a new package from it.

## Boilerplate Generator (WIP)

- Provide a small DSL that can be used to generate boilerplate for projects that will be using the `fabric` package.

## Extra Notes

**VDG:** temporary secondary dependency graph, which behaves independent of our main dependency graph, and only behaves on some given subset of UIs.

A VDG is basically equivalent to an internal partial-ordering of operations within a thread.

Except, that if we consider a VDG a thread, then we have to consider it as a thread node which has no dependents or dependencies. Thus, a VDG is like a virtualized thread node that allows us to manipulate multiple (V)UIs instead of only a single V(UI).

VDGs can present a potential concurrency-safety issue if you use them without taking into consideration the assumptions of what they are.

Make sure that all VDG operations are safe to operate independently from the main dependency graph.

In summation, using a VDG is making the following assumptions -- all access types that the VDG will be using are okay to be used without dependency management from our graph. As well as all VDG operations may occur at any given time independent of our global graph as well. They are simply a temporary secondary dependency graph that will operate on our CDS independent of our main graph.

### Handling CDS Manipulation

The answer to CDS manipulation is that the ordering of nodes as chosen by the developer should always be such that manipulation access procedures are ordered before or after a read or write e.g. the priority of a read and write procedure could be greater than the priority of a manipulation procedure.

So, while it is possible to use fabric to perform dirty reads and writes, etc. It is not the purpose of fabric to be full-proof, but rather to empower the developer to more easily implement full-proof (as defined by the spec of the system) concurrency control in their software.

Actually, the purpose of behavior avoidance is *not* to disallow a behavior from being designed into the system, but rather to be able to verify more easily that a certain behavior (or lack of behavior) exists in the system design.
