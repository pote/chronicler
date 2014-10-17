package chronicler

import(
	"net/http"
)

type Tale struct {
	Routes  []Route
}

func NewTale() *Tale {
	tale := &Tale{
		Routes: []Route{},
	}

	return tale
}

func (n *Tale) Register(route Route) {
	n.Routes = append(n.Routes, route)
}

func (n *Tale) Travel(w http.ResponseWriter, req *http.Request) (bool) {
	for _, plan := range n.Routes {
		if plan.Match(req) {
			plan.Story(w, req)
			return true
		}
	}

	return false
}

func (n *Tale) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	go n.Travel(w, req)
}

func (n *Tale) Perform(addr string) {
	http.ListenAndServe(addr, n)
	select {}
}
