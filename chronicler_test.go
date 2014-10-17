package chronicler

import(
	"io"
	"net/http"
)

type prologue struct { }

func (m *prologue) Match(req *http.Request) (bool) {
	return true
}

func (n *prologue) Story(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "It was night again. The Waystone Inn lay in silence, and it was a silence of three parts.\n")
}

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
