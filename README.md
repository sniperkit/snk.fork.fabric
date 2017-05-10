# Fabric

A Concurrency primitives package.

Any real node thread can spawn virtual DAG node threads (not VUIs) attached to the same UI that they are attached to as long as they do not allow the virtual node thread to be one of it's dependencies.

This might be useful for when we want too execute one atomic procedure, signal completed, let the dependent node execute its procedure, but then exectue the next atomic procedure in our thread AFTER the dependent has signaled completed. However, we cannot have cyclic dependencies, so instead we spawn a virtual thread containing the next operation that will have the current dependent node as one of it's *dependency* nodes and will execute once our current dependent node signals complete.

Virtual Threads are often spawned by routers ...