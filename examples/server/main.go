package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JKhawaja/fabric"
	"github.com/JKhawaja/fabric/examples/server/db"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1])
}

func main() {
	log.Println("Server Example ...")
	graph := fabric.NewGraph()
	log.Println(graph)

	// TODO: create handlers and http server ...
	// 	just have a basic server and demonstrate how a system can be designed to ensure that a user always avoids:
	// 	"lost updates" (two transactions are updating the same piece of data and one update gets overwritten),
	// 	"uncommited data" (two transactions access same data but first transaction rollsback while second transaction is executing),
	// 	"inconsistent retrievals" (when a transaction accesses data before and after another transaction finishes working with that data)
	// create session endpoint (login), delete session endpoint (logout), each session has a new poset, which generates a graph per user
	// make sure to test how one node will signal another node as complete and then and only then can a go routine handler proceed (blocking select statement)...

	log.Println(db.MyReadFunc.ID())
	log.Println(db.MyOtherReadFunc.ID())

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
