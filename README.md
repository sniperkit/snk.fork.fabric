# Fabric

Slides discussing the core concepts behind Fabric: http://jeremykhawaja.com/fabric/fabric.html

The concepts behind this package can be found here: https://github.com/JKhawaja/concurrency

A Concurrency control primitives package. Utilizing dependency graphs to avoid conflict in concurrent access to a single data structure. **Fine-Grained Blocking**.

This code package is less of a "here are some functions and objects. Use them." And more of a "here are some interfaces. Implement them." And this allows the package to behave more as a design guidance tool rather than a "strict" dependency.

In other words: **the purpose of this package is to use it in creating your own Concurrent-Data-Structure packages**. Another option is to take an existing data structure package and **fabric-ate** (hehe, get it?) a new package from it.

## Future (TODO)

- **Fabgen:** Provide a small code generation tool that can be used to generate boilerplate for projects that will be using the `fabric` package.
- **FabCheck:** methods for formal verification of system design using fabric package in order to avoid undesired behavior
- Make fabric work for multiple CDSs simultaneously
- Parallel distributed batch process scheduling

