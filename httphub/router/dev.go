package router

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

// SourceCodeHandler displays the content of the file specified in the path
// query argument with syntax highlighting and highlights the line passed
// in the line query argument.
func SourceCodeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")

	f, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer f.Close()
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)

	line := r.FormValue("line")
	lineNum, err := strconv.Atoi(line)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	lines := [][2]int{
		{lineNum, lineNum},
	}

	formatter := html.New(html.TabWidth(4), html.WithLineNumbers(true), html.HighlightLines(lines))
	style := styles.Get("github")
	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, buf.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "text/html")
	formatter.Format(w, style, iterator)
}
