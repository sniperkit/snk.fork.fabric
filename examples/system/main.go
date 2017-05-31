package main

import (
	"log"

	"github.com/JKhawaja/fabric"
	"github.com/JKhawaja/fabric/examples/ring"
)

// TODO: How will we address the dynamics of the ring structure, since it is very likely
//		that nodes in the structure will be added and removed very often.

//		One possible solution: could be that we assign the entire ring to a UI, then we can
//		assign each structural node to its own VUI. This will note the fact that any given
//		structural node can be temporary relative to the overall data structure.

//		The other idea: is that we assign certain nodes to their own UIs, and only allow
//		certain access types within certain UIs. For example, we could assign the root of
//		the ring to its own UI...

/*
	The Assignment function needs to also be aware of which unique independents have a thread
	or a virtual DDAG (Dependency-DAG) already assigned to them. It needs to be able to (re-)assign
	threads to available unique independents. This will be useful for when a thread is e.g. a processing
	request on a server, and will have a finite lifespan. (In other words, since virtual nodes are
	basically just threads, the assignment function will need to keep track of thread orderings per
	real node.)

    Can we run into an issue with the Assignment Function being blocked from updating its assignment
    list for a unique independent (??)

    Assignments can also be classified: e.g. a thread can be assigned to a unique independent but it
    is only allowed to read from the data structure, whereas another thread can be assigned that is
    allowed to modify the data, etc.
*/

func main() {
	log.Println("Fabric Example ...")
}
