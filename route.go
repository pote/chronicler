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
	// can fullfill one or more of several roles: routing, transformations and
	// conclusions.
	//
	Perform(http.ResponseWriter, *http.Request)
}
