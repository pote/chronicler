package chronicler

import(
	"net/http"
	"testing"
)

func TestMatches(t *testing.T) {
	tale := NewTale()
	tale.Register(&prologue{})

	writer := NewFakeWriter()
	req := http.Request{
		Method: "GET",
	}

	tale.Travel(writer, &req)

	if writer.Text() != "It was night again. The Waystone Inn lay in silence, and it was a silence of three parts.\n" {

		t.Error("Could not match the the parent tell")
	}
}
