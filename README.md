# Fabric

## Concepts

Slides discussing the core concepts behind Fabric: http://jeremykhawaja.com/fabric/fabric.html

The initial paper outlining the core concepts (needs revision): https://github.com/JKhawaja/concurrency

## Project State

**Note to Developers:** Fabric is still a work-in-progress. The current implementation is usable but will likely undergo several breaking changes in the near future. Please vendor if used, and/or use with caution. 

Please don't hesitate to complain or make a recommendation with an issue on GitHub. :)

## Summary

A Concurrency control primitives package. Utilizing dependency graphs to avoid conflict in concurrent access to a single data structure. **Fine-Grained Blocking**.

This code package is less of a "here are some functions and objects. Use them." And more of a "here are some interfaces. Implement them." And this allows the package to behave more as a design guidance tool rather than a "strict" dependency.

In other words: **the purpose of this package is to use it in creating your own Concurrent-Data-Structure packages**. Another option is to take an existing data structure package and **fabric-ate** (hehe, get it?) a new package from it.

## In Progress

- **Fabgen:** A small code generation tool that can be used to generate boilerplate for projects that will be using the `fabric` package.

## Future

- **FabCheck:** methods for formal verification of system design using fabric package in order to avoid undesired behavior
- Make fabric work for multiple CDSs simultaneously
- Parallel distributed batch process scheduling (taking a batch of e.g. transactions, and scheduling them across a set of machines in parallel, dependent on their relationship with each other within a fine-grained blocking dependency graph).

