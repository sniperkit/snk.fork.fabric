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
	VPoset *dg.VDGPoset
	VUI    *fabric.DGNode
}

// NewSession ...
func NewSession(v *dg.VDGPoset, vui *fabric.DGNode) Session {
	return Session{
		ID:     GenSessionID(),
		VPoset: v,
		VUI:    vui,
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
func signalCheck(np *fabric.Virtual) bool {
	cont := true
	node := *np

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

func signalHandler(c <-chan fabric.ProcedureSignals, wg sync.WaitGroup) {
	for {
		select {
		case sigMap := <-c:
			for _, value := range sigMap {
				// NOTE: the switch cases could revolve around different access types and different signal values
				switch value {
				case fabric.Waiting:
					continue
				case fabric.Started:
					continue
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
		default:
			continue
		}
	}
}

func createSession(g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// create VUI
		// NOTE: in a real application, this VUI could be constructed to only allow
		// access procedures to access a CDS section.
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
		vui, err := g.AddVUI(vu)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not create VUI!"))
		}

		// create VDG
		vdg, err := fabric.NewVDGWithRoot(g)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not create a VDG!"))
		}

		// wrap VDG in a VPOSET
		vposet := dg.NewVDGPoset(vdg)

		// create a Session
		session := NewSession(vposet, vui)

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
			g.RemoveVDG(sess.VPoset.Vdg)

			// Remove VUI
			var iv interface{} = *sess.VUI
			vui := iv.(fabric.DGNode)
			err = g.RemoveVUI(vui)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
		}
	}
}

func createNode(t *db.Tree, g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// grab session
		sess, err := getSession(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Create Virtual Node
		sm1 := make(fabric.SignalingMap)
		s1 := make(fabric.SignalsMap)
		var list fabric.ProcedureList
		list = append(list, db.CreateNode)
		v := dg.Virtual{
			Node: dg.Node{
				Id:               sess.VPoset.VDG().GenID(),
				Type:             fabric.VDGNode,
				Priority:         db.CreateNode.Priority(),
				AccessProcedures: &list,
				Signalers:        &sm1,
				Signals:          &s1,
			},
			Start: false,
			Root:  false,
			Space: sess.VUI,
		}

		// Order Virtual Node
		vn := sess.VPoset.Order(v)

		// Start Virtual Node
		v.Start = true

		// SignalCheck and then run logic
		if signalCheck(vn) {
			val := r.URL.Query()
			value := val["value"]
			newNode := db.CreateNode(t, value[0])
			w.Write([]byte(strconv.Itoa(newNode.ID())))
		}

		// remove virtual node
		sess.VPoset.VDG().RemoveVirtualNode(vn)
	}
}

func createEdge(t *db.Tree, g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get session
		sess, err := getSession(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Create Virtual Node
		sm1 := make(fabric.SignalingMap)
		s1 := make(fabric.SignalsMap)
		var list fabric.ProcedureList
		list = append(list, db.CreateEdge)
		v := dg.Virtual{
			Node: dg.Node{
				Id:               sess.VPoset.VDG().GenID(),
				Type:             fabric.VDGNode,
				Priority:         db.CreateEdge.Priority(),
				AccessProcedures: &list,
				Signalers:        &sm1,
				Signals:          &s1,
			},
			Start: false,
			Root:  false,
			Space: sess.VUI,
		}

		// Order Virtual Node
		vn := sess.VPoset.Order(v)

		// Start Virtual Node
		v.Start = true

		// SignalCheck and then run logic
		if signalCheck(vn) {
			val := r.URL.Query()
			node1 := val["n1"]
			node2 := val["n2"]
			node1id, _ := strconv.Atoi(node1[0])
			node2id, _ := strconv.Atoi(node2[0])
			var first *db.TreeNode
			var second *db.TreeNode
			for _, kp := range t.Nodes {
				k := *kp
				if k.ID() == node1id {
					kn := k.(db.TreeNode)
					first = &kn
				}
				if k.ID() == node2id {
					kn := k.(db.TreeNode)
					second = &kn
				}
			}
			newEdge := db.CreateEdge(t, first, second)
			w.Write([]byte(strconv.Itoa(newEdge.ID())))
		}

		// remove virtual node
		sess.VPoset.VDG().RemoveVirtualNode(vn)
	}
}

func removeNode(t *db.Tree, g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// session and signal checks
		sess, err := getSession(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Create Virtual Node
		sm1 := make(fabric.SignalingMap)
		s1 := make(fabric.SignalsMap)
		var list fabric.ProcedureList
		list = append(list, db.RemoveNode)
		v := dg.Virtual{
			Node: dg.Node{
				Id:               sess.VPoset.VDG().GenID(),
				Type:             fabric.VDGNode,
				Priority:         db.RemoveNode.Priority(),
				AccessProcedures: &list,
				Signalers:        &sm1,
				Signals:          &s1,
			},
			Start: false,
			Root:  false,
			Space: sess.VUI,
		}

		// Order Virtual Node
		vn := sess.VPoset.Order(v)

		// Start Virtual Node
		v.Start = true

		if signalCheck(vn) {
			val := r.URL.Query()
			node := val["node"]
			nodeID, _ := strconv.Atoi(node[0])
			db.RemoveNode(t, nodeID)
			w.Write([]byte("Node Removed successfully."))
		}

		// remove virtual node
		sess.VPoset.VDG().RemoveVirtualNode(vn)
	}
}

func removeEdge(t *db.Tree, g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// session and signal checks
		sess, err := getSession(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Create Virtual Node
		sm1 := make(fabric.SignalingMap)
		s1 := make(fabric.SignalsMap)
		var list fabric.ProcedureList
		list = append(list, db.RemoveEdge)
		v := dg.Virtual{
			Node: dg.Node{
				Id:               sess.VPoset.VDG().GenID(),
				Type:             fabric.VDGNode,
				Priority:         db.RemoveEdge.Priority(),
				AccessProcedures: &list,
				Signalers:        &sm1,
				Signals:          &s1,
			},
			Start: false,
			Root:  false,
			Space: sess.VUI,
		}

		// Order Virtual Node
		vn := sess.VPoset.Order(v)

		// Start Virtual Node
		v.Start = true

		if signalCheck(vn) {
			val := r.URL.Query()
			edge := val["edge"]
			edgeID, _ := strconv.Atoi(edge[0])
			db.RemoveEdge(t, edgeID)
			w.Write([]byte("Edge removed successfully."))
		}

		// remove virtual node
		sess.VPoset.VDG().RemoveVirtualNode(vn)
	}
}

func readNodeValue(t *db.Tree, g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// session and signal checks
		sess, err := getSession(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Create Virtual Node
		sm1 := make(fabric.SignalingMap)
		s1 := make(fabric.SignalsMap)
		var list fabric.ProcedureList
		list = append(list, db.ReadNodeValue)
		v := dg.Virtual{
			Node: dg.Node{
				Id:               sess.VPoset.VDG().GenID(),
				Type:             fabric.VDGNode,
				Priority:         db.ReadNodeValue.Priority(),
				AccessProcedures: &list,
				Signalers:        &sm1,
				Signals:          &s1,
			},
			Start: false,
			Root:  false,
			Space: sess.VUI,
		}

		// Order Virtual Node
		vn := sess.VPoset.Order(v)

		// Start Virtual Node
		v.Start = true

		// Signalcheck and then run logic
		if signalCheck(vn) {
			val := r.URL.Query()
			node := val["node"]
			nodeID, _ := strconv.Atoi(node[0])
			value := db.ReadNodeValue(t, nodeID)
			w.Write([]byte(value.(string)))
		}

		// remove virtual node
		sess.VPoset.VDG().RemoveVirtualNode(vn)
	}
}

func updateNodeValue(t *db.Tree, g *fabric.Graph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// session and signal checks
		sess, err := getSession(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// Create Virtual Node
		sm1 := make(fabric.SignalingMap)
		s1 := make(fabric.SignalsMap)
		var list fabric.ProcedureList
		list = append(list, db.UpdateNodeValue)
		v := dg.Virtual{
			Node: dg.Node{
				Id:               sess.VPoset.VDG().GenID(),
				Type:             fabric.VDGNode,
				Priority:         db.UpdateNodeValue.Priority(),
				AccessProcedures: &list,
				Signalers:        &sm1,
				Signals:          &s1,
			},
			Start: false,
			Root:  false,
			Space: sess.VUI,
		}

		// Order Virtual Node
		vn := sess.VPoset.Order(v)

		// Start Virtual Node
		v.Start = true

		// SignalCheck and then run logic
		if signalCheck(vn) {
			val := r.URL.Query()
			node := val["node"]
			value := val["value"]
			nodeID, _ := strconv.Atoi(node[0])
			db.UpdateNodeValue(t, nodeID, value[0])
			w.Write([]byte("Node updated successfully."))
		}

		// remove virtual node
		sess.VPoset.VDG().RemoveVirtualNode(vn)
	}
}

func main() {
	log.Println("Server Example has started ...")

	// create a Tree CDS
	tree := db.NewTree()

	// create a global DG
	graph := fabric.NewGraph()

	// TODO: have a single UI that covers the entire tree

	http.HandleFunc("/createsession", createSession(graph))
	http.HandleFunc("/deletesession", deleteSession(graph))
	http.HandleFunc("/createnode", createNode(tree, graph))
	http.HandleFunc("/createedge", createEdge(tree, graph))
	http.HandleFunc("/removenode", removeNode(tree, graph))
	http.HandleFunc("/removeedge", removeEdge(tree, graph))
	http.HandleFunc("/readnodevalue", readNodeValue(tree, graph))
	http.HandleFunc("/updatenodevalue", updateNodeValue(tree, graph))
	http.ListenAndServe(":8080", nil)
}
