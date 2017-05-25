# Fabric

A Concurrency control primitives package. Utilizing dependency graphs to avoid conflict in concurrent access to a single data structure.

This code package is less of a "here are some functions and objects. Use them." And more of a "here are some interfaces. Implement them." And this allows the package to behave more as a design guidance tool rather than a strict dependency.

One thing that this package does is enable a developer to turn any data structure implementation into a Concurrent Data Structure "Fabric-Friendly" package (or easily create a new CDS package out of the original data-structure code).

## Code Generator (WIP)

- Provide a small DSL that can be used to generate boilerplate for projects that will be using the `fabric` package.