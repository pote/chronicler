package chronicler

import(
	"log"
	"net/http"
)

// Stories are one of the two building blocks for Chronicler applications, they
// group together Routes and decide the path that a request being received
// will follow through your application
type Story struct {
	Routes  []Route
	Logging bool
}

func NewStory() *Story {
	story := &Story{
		Routes: []Route{},
		Logging: true,
	}

	return story
}

// Register takes a pointer to a Route interface and adds it to the stack,
// it's important to note that *the order in which we register routes is significant*
// as you'll read in the #Explore method.
func (n *Story) Register(route Route) {
	n.Routes = append(n.Routes, route)
}

// The Explore method will evaluate all it's routes and, if it finds one that matches
// the current request execute it's Performance.
func (n *Story) Explore(w http.ResponseWriter, req *http.Request) (bool) {
	for _, path := range n.Routes {
		if path.Match(req) {
			if n.Logging {
				log.Printf("Matched %v to %v\n", req.Method, req.URL.Path)
			}
			path.Perform(w, req)
			return true
		}
	}

	return false
}

func (n *Story) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	n.Explore(w, req)
}

// Serve is called on the parent Story of your application: it'll start
// a server bound to a given address and route all requests to it.
func (n *Story) Serve(addr string) {
	log.Println(`Chronicler shook his head and Bast gave a frustrated sigh. "How about plays? Have you seen The Ghost and the Goosegirl or The Ha'penny King?`)
	http.ListenAndServe(addr, n)
}
