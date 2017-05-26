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

func main() {
	log.Println("Fabric Example ...")
}
