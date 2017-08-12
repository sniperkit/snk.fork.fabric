package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
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
	VPoset fabric.VPoset
	VUI    fabric.UI
}

// NewSession ...
func NewSession(v fabric.VPoset) Session {
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

	if sess.ID == 0 {
		return sess, fmt.Errorf("Session not found. Please create a new session.")
	}

	// return Session object
	return sess, nil
}

// signalCheck is a middleware function that verifies that VDG dependencies have signaled complete
// NOTE: a node should mark itself started before calling SignalCheck;
// a node is considered blocked/spinning while SignalChecking, but it also has bounded itself
// from being added more dependencies.
func signalCheck(node fabric.Virtual) bool {
	cont := true

	depSignals := node.ListSignals()
	var wg sync.WaitGroup
	for _, channel := range depSignals {
		wg.Add(1)
		go signalHandler(channel, wg)
	}

	// Virtual Node blocks/spins
	wg.Wait()

	return cont
}

func signalHandler(c <-chan fabric.NodeSignal, wg sync.WaitGroup) {
	select {
	case sig := <-c:
		// NOTE: the switch cases could revolve around different access types, different signal values, and different UIs
		switch sig.Value {
		case fabric.Waiting:
			// do nothing
		case fabric.Started:
			// do nothing
		case fabric.Completed:
			wg.Done()
			return
		case fabric.Aborted:
			wg.Done()
			return
		case fabric.AbortRetry:
			wg.Done()
			break
		case fabric.PartialAbort:
			wg.Done()
			return
		}
	}
}

func createSession(c fabric.CDS, g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := c.(*db.Tree)
		// create VDG
		vdg, err := fabric.NewVDGWithRoot(g)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not create a VDG!"))
		}

		// wrap VDG in a VPOSET
		vposet := dg.NewVDGPoset(vdg)

		// create a Session
		session := NewSession(vposet)

		// create a section node in the tree using session id
		sn := t.NewSection(session.ID)

		// create a branch section using section node as root
		branch := fabric.NewBranch(sn, c)

		// create VUI
		vu := dg.NewVUI(g, branch)

		// add vui to graph
		_, err = g.AddVUI(vu)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not create VUI!"))
		}

		// set VUI for session
		session.VUI = vu

		// add session to global sessions store
		sessions = append(sessions, session)

		// return session id to user
		w.Write([]byte(strconv.Itoa(session.ID)))
	}
}

func deleteSession(g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get session
		sess, err := getSession(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// block till VDG has completed by signal checking the root node ...
		if signalCheck(sess.VPoset.VDG().Root) {
			// remove Session from global sessions store
			for i, s := range sessions {
				if s.ID == sess.ID {
					sessions = append(sessions[:i], sessions[i+1:]...)
					break
				}
			}

			// remove VDG
			g.RemoveVDG(sess.VPoset.VDG())

			// Remove VUI
			err = g.RemoveVUI(sess.VUI)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
		}
	}
}

func createNode(c fabric.CDS, g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := c.(*db.Tree)

		// grab session
		sess, err := getSession(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Create Virtual Node
		var list fabric.ProcedureList
		list = append(list, db.CreateNode)
		v := dg.NewVirtual(sess.VPoset.VDG(), sess.VUI, &list, db.CreateNode.Priority())

		// Order Virtual Node
		err = sess.VPoset.Order(v)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Start Virtual Node
		v.Start()

		// SignalCheck and then run logic
		if signalCheck(v) {
			val := r.URL.Query()
			value, ok := val["value"]
			if !ok {
				w.Write([]byte("Please provide a value for the node."))
				return
			}

			newNode, err := t.CreateNode(sess.VUI.GetSection(), value[0])
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			db.CreateNode.Commit(v.(fabric.DGNode))
			w.Write([]byte(strconv.Itoa(newNode.ID())))
		}

		// remove virtual node
		sess.VPoset.VDG().RemoveVirtualNode(v)
	}
}

func createEdge(c fabric.CDS, g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := c.(*db.Tree)
		// get session
		sess, err := getSession(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Create Virtual Node
		var list fabric.ProcedureList
		list = append(list, db.CreateEdge)
		v := dg.NewVirtual(sess.VPoset.VDG(), sess.VUI, &list, db.CreateEdge.Priority())

		// Order Virtual Node
		err = sess.VPoset.Order(v)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Start Virtual Node
		v.Start()

		// SignalCheck and then run logic
		if signalCheck(v) {
			val := r.URL.Query()
			node1 := val["n1"]
			node2 := val["n2"]
			node1id, _ := strconv.Atoi(node1[0])
			node2id, _ := strconv.Atoi(node2[0])
			var first fabric.Node
			var second fabric.Node
			for _, k := range c.ListNodes() {
				if k.ID() == node1id {
					first = k
				}
				if k.ID() == node2id {
					second = k
				}
			}

			newEdge, err := t.CreateEdge(sess.VUI.GetSection(), first, second)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			db.CreateEdge.Commit(v.(fabric.DGNode))
			w.Write([]byte(strconv.Itoa(newEdge.ID())))
		}

		// remove virtual node
		sess.VPoset.VDG().RemoveVirtualNode(v)
	}
}

func removeNode(c fabric.CDS, g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := c.(*db.Tree)
		// session and signal checks
		sess, err := getSession(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Create Virtual Node
		var list fabric.ProcedureList
		list = append(list, db.RemoveNode)
		v := dg.NewVirtual(sess.VPoset.VDG(), sess.VUI, &list, db.RemoveNode.Priority())

		// Order Virtual Node
		err = sess.VPoset.Order(v)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Start Virtual Node
		v.Start()

		if signalCheck(v) {
			val := r.URL.Query()
			node := val["node"]
			nodeID, _ := strconv.Atoi(node[0])
			err := t.RemoveNode(sess.VUI.GetSection(), nodeID)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			db.RemoveNode.Commit(v.(fabric.DGNode))
			w.Write([]byte("Node Removed successfully."))
		}

		// remove virtual node
		sess.VPoset.VDG().RemoveVirtualNode(v)
	}
}

func removeEdge(c fabric.CDS, g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := c.(*db.Tree)
		// session and signal checks
		sess, err := getSession(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Create Virtual Node
		var list fabric.ProcedureList
		list = append(list, db.RemoveEdge)
		v := dg.NewVirtual(sess.VPoset.VDG(), sess.VUI, &list, db.RemoveEdge.Priority())

		// Order Virtual Node
		err = sess.VPoset.Order(v)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Start Virtual Node
		v.Start()

		if signalCheck(v) {
			val := r.URL.Query()
			edge := val["edge"]
			edgeID, _ := strconv.Atoi(edge[0])
			err := t.RemoveEdge(sess.VUI.GetSection(), edgeID)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			db.RemoveEdge.Commit(v.(fabric.DGNode))
			w.Write([]byte("Edge removed successfully."))
		}

		// remove virtual node
		sess.VPoset.VDG().RemoveVirtualNode(v)
	}
}

func readNodeValue(c fabric.CDS, g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := c.(*db.Tree)
		// session and signal checks
		sess, err := getSession(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Create Virtual Node
		var list fabric.ProcedureList
		list = append(list, db.ReadNodeValue)
		v := dg.NewVirtual(sess.VPoset.VDG(), sess.VUI, &list, db.ReadNodeValue.Priority())

		// Order Virtual Node
		err = sess.VPoset.Order(v)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Start Virtual Node
		v.Start()

		// Signalcheck and then run logic
		if signalCheck(v) {
			val := r.URL.Query()
			node := val["node"]
			nodeID, _ := strconv.Atoi(node[0])
			value, err := t.ReadNodeValue(sess.VUI.GetSection(), nodeID)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			db.ReadNodeValue.Commit(v.(fabric.DGNode))
			w.Write([]byte(value.(string)))
		}

		// remove virtual node
		sess.VPoset.VDG().RemoveVirtualNode(v)
	}
}

func updateNodeValue(c fabric.CDS, g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := c.(*db.Tree)
		// session and signal checks
		sess, err := getSession(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Create Virtual Node
		var list fabric.ProcedureList
		list = append(list, db.UpdateNodeValue)
		v := dg.NewVirtual(sess.VPoset.VDG(), sess.VUI, &list, db.UpdateNodeValue.Priority())

		// Order Virtual Node
		err = sess.VPoset.Order(v)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Start Virtual Node
		v.Start()

		// SignalCheck and then run logic
		if signalCheck(v) {
			val := r.URL.Query()
			node := val["node"]
			value := val["value"]
			nodeID, _ := strconv.Atoi(node[0])
			err := t.UpdateNodeValue(sess.VUI.GetSection(), nodeID, value[0])
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			db.UpdateNodeValue.Commit(v.(fabric.DGNode))
			w.Write([]byte("Node updated successfully."))
		}

		// remove virtual node
		sess.VPoset.VDG().RemoveVirtualNode(v)
	}
}

func main() {
	log.Println("Fabric server example has started ...")
	log.Println("/createsession")
	log.Println("/deletesession?id=<session_id>")
	log.Println("/createnode?id=<session_id>&value=<my_value>")
	log.Println("/createedge?id=<session_id>&n1=<node_id>&n2=<node_id>")
	log.Println("/removenode?id=<session_id>&node=<node_id>")
	log.Println("/removeedge?id=<session_id>&edge=<edge_id>")
	log.Println("/readnodevalue?id=<session_id>&node=<node_id>")
	log.Println("/updatenodevalue?id=<session_id>&node=<node_id>&value=<my_value>")

	// create a Tree CDS
	tree := db.NewTree()

	// create a global DG
	graph := fabric.NewGraph()

	// TODO: have a single UI that covers the entire tree
	// create Section (entire CDS)
	// create UI
	// add UI to graph

	http.HandleFunc("/createsession", createSession(tree, graph))
	http.HandleFunc("/deletesession", deleteSession(graph))
	http.HandleFunc("/createnode", createNode(tree, graph))
	http.HandleFunc("/createedge", createEdge(tree, graph))
	http.HandleFunc("/removenode", removeNode(tree, graph))
	http.HandleFunc("/removeedge", removeEdge(tree, graph))
	http.HandleFunc("/readnodevalue", readNodeValue(tree, graph))
	http.HandleFunc("/updatenodevalue", updateNodeValue(tree, graph))
	http.ListenAndServe(":8080", nil)
}
