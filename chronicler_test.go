package chronicler

import(
	"io"
	"net/url"
	"net/http"
)

// A writer we can test easily.
type fakeWriter struct {
	data   []byte
}

func (f *fakeWriter) Write(newData []byte) (int, error) {
	f.data = append(f.data, newData...)
	return 0, nil
}

func (f *fakeWriter) Header() http.Header {
	return http.Header{}
}


func (f *fakeWriter) WriteHeader(int) {
	// No Op
}

func (f *fakeWriter) Text() string {
	return string(f.data)
}

func NewFakeWriter() *fakeWriter {
	return &fakeWriter{
		data: []byte{},
	}
}

func NewFakeRequest() *http.Request {
	return &http.Request{
		Method: "POST",
		URL: &url.URL{
			Path: "/messages",
		},
	}
}



// A short route that will always be matched
type episode struct { }

func (r *episode) Match(req *http.Request) (bool) {
	return true
}

func (r *episode) Perform(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "It was night again. The Waystone Inn lay in silence, and it was a silence of three parts.\n")
}

// A Performance with three possible paths.
type arch struct { }

func (r *arch) Match(req *http.Request) (bool) {
	return true
}

func (r *arch) Perform(w http.ResponseWriter, req *http.Request) {
	development := NewStory()
	development.Register(&tragedy{})
	development.Register(&comedy{})
	development.Register(&legend{})
	development.Explore(w, req)
}

type tragedy struct { }

func (r *tragedy) Match(req *http.Request) (bool) {
	return false
}

func (r *tragedy) Perform(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am a tragedy, and so we must endure.")
}

type comedy struct { }

func (r *comedy) Match(req *http.Request) (bool) {
	return true
}

func (r *comedy) Perform(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am a comedy, and so we will laugh.")
}

type legend struct { }

func (r *legend) Match(req *http.Request) (bool) {
	return true
}

func (r *legend) Perform(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am a legend, so I will never end.")
}
