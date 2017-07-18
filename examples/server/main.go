package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/JKhawaja/fabric"
	"github.com/JKhawaja/fabric/examples/server/db"
	"github.com/JKhawaja/fabric/examples/server/dg"
)

// Global Session store
var sessions []Session

// Session is a user session object ...
type Session struct {
	ID     int
	VPoset *dg.VDGPoset
}

// NewSession ...
func NewSession(v *dg.VDGPoset) Session {
	return Session{
		ID:     GenSessionID(),
		VPoset: v,
	}
}

// GenSessionID ...
func GenSessionID() int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	for _, s := range sessions {
		if s.ID == id {
			id = GenSessionID()
		}
	}
	return id
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[0])
}

// middleware function that gets session id for user
func getSession(r *http.Request) (Session, error) {
	var sess Session

	// Get session id from url query param "id"
	v := r.URL.Query()
	v2, ok := v["id"]
	if !ok {
		return sess, fmt.Errorf("Please create a session first")
	}
	id, err := strconv.Atoi(v2[0])
	if err != nil {
		return sess, err
	}

	// Get Session object using id from global sessions store
	for _, s := range sessions {
		if s.ID == id {
			sess = s
			break
		}
	}

	// return Session object
	return sess, nil
}

// TODO: middleware function must verify that VDG dependencies have signaled complete
func signalCheck(g *fabric.Graph, s Session) bool {
	// add node to graph
	// blocking select statement that only returns a true value when all dependency complete signals have been recieved
	return true
}

func createSession(g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// create VUI
		sm1 := make(fabric.SignalingMap)
		s1 := make(fabric.SignalsMap)
		vu := dg.UI{
			Node: dg.Node{
				Id:        g.GenID(),
				Type:      fabric.UINode,
				Signalers: &sm1,
				Signals:   &s1,
			},
			Virtual: true,
		}

		// add vui to graph
		_, err := g.AddVUI(vu)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not create VUI!"))
		}

		// create VDG
		vdg, err := fabric.NewVDG(g)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not create a VDG!"))
		}

		// create a POSET for the VDG
		vposet := dg.NewVDGPoset(vdg)

		// create a Session
		session := NewSession(vposet)

		// return session id to user
		w.Write([]byte(string(session.ID)))
	}
}

func deleteSession(w http.ResponseWriter, r *http.Request) {
	// TODO: getSession()
	// remove that Session from global sessions store
}

func createNode(t *db.Tree, g *fabric.Graph) http.HandlerFunc {
	// TODO: middleware must verify that VDG dependencies
	return func(w http.ResponseWriter, r *http.Request) {
		// session and signal checks
		sess, err := getSession(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		cont := signalCheck(g, sess)
		if cont {
			v := r.URL.Query()
			value := v["value"]
			db.CreateNode(t, value[0])
			str := fmt.Sprintf("%v", t.Nodes)
			w.Write([]byte(str))
		}
	}
}

// TODO: create Tree CURDL handlers
// TODO: create session creation and deletion handlers
//		session creation will need to be passed graph or poset as argument
// 		sessions create a VUI that gets added to the tree poset
//		session also creates a VDG tree that request go-routines get added too
//		thus, there will be a VDG middleware that every handler gets wrapped in
// 		and this middleware will add the request to the tree and wait for its turn
// 		to call the handler.

func main() {
	log.Println("Server Example ...")
	// TODO: create a single Tree CDS
	// TODO: create a POSET for the CDS
	// TODO: have a single UI that covers the entire tree
	graph := fabric.NewGraph()
	log.Println(graph)
	tree := db.NewTree()

	// TODO: create handlers and http server ...
	// 	just have a basic server and demonstrate how a system can be designed to ensure that a user always avoids:
	// 	"lost updates" (two transactions are updating the same piece of data and one update gets overwritten),
	// 	"uncommited data" (two transactions access same data but first transaction rollsback while second transaction is executing),
	// 	"inconsistent retrievals" (when a transaction accesses data before and after another transaction finishes working with that data)
	// create session endpoint (login), delete session endpoint (logout), each session has a new poset, which generates a graph per user
	// make sure to test how one node will signal another node as complete and then and only then can a go routine handler proceed (blocking select statement)...

	http.HandleFunc("/", handler)
	http.HandleFunc("/v", createNode(tree, graph))
	// http.HandleFunc("/session")
	http.ListenAndServe(":8080", nil)
}
