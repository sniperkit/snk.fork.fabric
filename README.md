# Fabric

The concepts behind this package can be found here: https://github.com/JKhawaja/concurrency

A Concurrency control primitives package. Utilizing dependency graphs to avoid conflict in concurrent access to a single data structure.

This code package is less of a "here are some functions and objects. Use them." And more of a "here are some interfaces. Implement them." And this allows the package to behave more as a design guidance tool rather than a "strict" dependency.

In other words: **the purpose of this package is to use it in creating your own Concurrent-Data-Structure packages**. Another option is to take an existing data structure package and **fabric-ate** (hehe, get it?) a new package from it.

## Boilerplate Generator (TODO)

- Provide a small DSL that can be used to generate boilerplate for projects that will be using the `fabric` package.

