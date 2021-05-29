package router

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/alecthomas/chroma/quick"
)

func SourceCodeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	f, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer f.Close()

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)
	if err := quick.Highlight(w, buf.String(), "go", "html", "github"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
