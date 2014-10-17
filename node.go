package chronicler

import(
	"log"
	"net/http"
)

type Node struct {
	Routes  []Route
	Logging bool
}

func NewNode() *Node {
	node := &Node{
		Routes: []Route{},
		Logging: true,
	}

	return node
}

// Register takes a pointer to a Route interface and adds it to the stack,
// it's important to note that *the order in which we register routes is significant*
// as you'll read in the #Travel method description.
func (n *Node) Register(route Route) {
	n.Routes = append(n.Routes, route)
}

// The Dispatch method will evaluate all it's routes and, if it finds one that matches
// the current request execute it's Performance.
func (n *Node) Dispatch(w http.ResponseWriter, req *http.Request) (bool) {
	for _, path := range n.Routes {
		if path.Match(req) {
			if n.Logging {
				log.Printf("Matched request: %q\n", req)
			}
			path.Perform(w, req)
			return true
		}
	}

	return false
}

func (n *Node) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	n.Dispatch(w, req)
}

// Serve is called on the parent Node of your application: it'll start
// a server bound to a given address and route all requests to it.
func (n *Node) Serve(addr string) {
	http.ListenAndServe(addr, n)
}
