package chronicler

import(
	"net/http"
)

type Route interface {
	Match(*http.Request) (bool)
	Story(http.ResponseWriter, *http.Request)
}
