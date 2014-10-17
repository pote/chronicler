package chronicler

import(
	"net/http"
)

// Routes are what makes your application tick: while Stories group together a
// number of possible routes the routes themselves will take your request further
// into your stack and/or execute code, sometimes modifying the original request
// before passing it along.
type Route interface {

	// The Match method will determine if your route will respond to a specific
	// request, they can run arbitrary code in order to make this disctintion
	// ranging from evaluating the Request path and HTTP verb to checking the
	// time of the day, anything and everything you can imagine.
	Match(*http.Request) (bool)

	// The Perform method (we'll just call them Performances) is the code that
	// will be executed when the Route is matched, this is arbitrary code and
	// can fullfill one or more of several roles:
	//
	// ## Routing:
	//
	// Routing Performances take you to other places in your code, this is done by
	// registering a new Story with it's own sets of Routes, nesting Routes makes
	// makes it really easy to compose your application and routing tree by assigning
	// one responsibility to each Performance.
	//
	// ### Example:
	//
	// This Performance routes requests on a food delivery web app, we'll hit it with
	// a GET /orders HTTP request.
	//
	// ```
	// func (r *homeStory) Perform(w http.ResponseWriter, req *http.Request) {
	//   story := chronicler.NewStory()
	//   story.Register(&orders{})      // Our request will match this Route.
	//   story.Register(&restaurants{}) // This one won't even be evaluated
	//   story.Register(&sessions{})    // Nor this one.
	//
	//   story.Explore(w, req)
	// }
	//
	// type orders struct { }
	// func (r *orders) currentUser(req *http.Request) *User {
	//   token := req.Header.Get("Authorization")
	//   user := // get your user through your ORM of choice. :)
	//   return user
	// }
	// func (r *orders) Match(req *http.Request) (bool) {
	//   return strings.HasPrefix(req.URL.Path, "/orders")
	// }
	// func (r *orders) Perform(w http.ResponseWriter, req *http.Request) {
	//   story := chronicler.NewStory()
	//   story.Register(&newOrder{})
	//   story.Register(&orderIndex{})
	//
	//   story.Explore(w, req)
	// }
	// ```
	// The starting poing of this flow is a Routing Performance that registers several possible
	// routes, our GET /orders HTTP request will be evaluated against orders.Match successfully,
	// and so orders.Perform will be called, starting another Match cycle.
	//
	// ## Transformations
	//
	// Transformations are changes that any given Performance can apply to the request it
	// receives, different parts of your application will benefit from working under a given
	// set of circumstances which you can refine as the requests moves through the Routes in
	// your application.
	//
	// Transformations are similar to the concept of a [middleware stack](http://en.wikipedia.org/wiki/Middleware)
	// except they are naturally scoped to a given Route - and by extension to all routes nested within it.
	//
	// ### Example:
	//
	// ```
	// func (r *userOrders) Perform(w http.ResponseWriter, req *http.Request) {
	//   story := chronicler.NewStory()
	//   story.Register(&orders{})
	//
	//   // There's a helper for this, but we'll talk about that later. :)
	//   req.URL.Path = strings.TrimPrefix(req.URL.Path, "/user")
	//
	//   story.Explore(w, req)
	// }
	// ```
	// This is a common transformation of a request as you usually want the inner routes
	// to not have to deal with parts of the URL that have already been matched. This is
	// however only a tiny example of transformations: you are free to run arbitrary
	// code in your Performances so anything is game here. Transformations are commonly
	// used in Routing Performances.
	//
	// ## Conclusions
	//
	// Conclusions are Performances that don't propagate the request to any other routes,
	// they're the final destination for the request in the matched routing tree. They tipically
	// represent a specific flow in your application such as "creating a user and return it as json".
	//
	// ### Example:
	//
	// ```
	// func (r *ending) Perform(w http.ResponseWriter, req *http.Request) {
	//   io.WriteString(w, "And they lived happily ever after.")
	// }
	// ```
	//
	Perform(http.ResponseWriter, *http.Request)
}
