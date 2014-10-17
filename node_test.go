package chronicler

import(
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSingleRouteNode(t *testing.T) {
	node := NewNode()
	node.Register(&episode{})

	writer := NewFakeWriter()
	req := http.Request{}

	node.Dispatch(writer, &req)

	if writer.Text() != "It was night again. The Waystone Inn lay in silence, and it was a silence of three parts.\n" {
		t.Error("Could not match the main node")
	}
}

func TestNesting(t *testing.T) {
	node := NewNode()
	node.Register(&arch{})

	writer := NewFakeWriter()
	req := http.Request{}

	node.Dispatch(writer, &req)

	if writer.Text() != "I am a comedy, and so we will laugh."{
		t.Error("Did not match the intended route")
	}
}

func TestServe(t *testing.T) {
	node := NewNode()
	node.Register(&episode{})
	go node.Serve(":2814")

	resp, err :=  http.Get("http://localhost:2814")
	if err != nil {
		t.Error(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if string(body) != "It was night again. The Waystone Inn lay in silence, and it was a silence of three parts.\n" {
		t.Errorf("Incorrect response body: %v\n", string(body))
	}
}
