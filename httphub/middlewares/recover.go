package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/ElMehdi19/httphub/httphub/helpers"
)

type stackLine struct {
	filePath   string
	lineNumber int
	memoryAddr string
}

func (l stackLine) toString() string {
	return fmt.Sprintf("<a href='/debug?path=%s&line=%d' target='blank'>%s</a>:%d %s", l.filePath, l.lineNumber, l.filePath, l.lineNumber, l.memoryAddr)
}

// parseStackLine takes a string and parses a stackLine out of it.
// It returns a non nil error if a valid stackLine couldn't be parsed.
func parseStackLine(line string) (stackLine, error) {
	// example line: /usr/local/go/src/runtime/debug/stack.go:24 +0x9f

	var l stackLine
	parts := strings.Split(line, ":")
	if len(parts) == 0 {
		return l, fmt.Errorf("no match")
	}

	// check if file actually exists.
	if _, err := os.Stat(strings.TrimSpace(parts[0])); os.IsNotExist(err) {
		fmt.Println(parts[0])
		return l, err
	}
	l.filePath = parts[0]

	// check if the filename is the only specified info.
	if len(parts) == 1 {
		return l, nil
	}

	parts = strings.Fields(parts[1])
	n, err := strconv.Atoi(parts[0])
	if err != nil {
		return l, nil
	}
	l.lineNumber = n

	// check if the memory address part is provided.
	if len(parts) > 1 {
		l.memoryAddr = parts[1]
	}

	return l, nil
}

// parseStack takes the stack trace string returns a new string with links to files
// highlighted as HTML anchor tags.
func parseStack(stack string) string {
	lines := strings.Split(stack, "\n")
	for i := 0; i < len(lines); i++ {
		// lines that include a file starts with a tab '\t'
		if !strings.HasPrefix(lines[i], "\t") {
			continue
		}
		l, err := parseStackLine(lines[i])
		if err != nil {
			continue
		}
		lines[i] = l.toString()
	}

	return strings.Join(lines, "\n")
}

// Recover middleware recovers from any errors that occured during the handling
// of incoming requests. If DEV_MODE is ON it will return the stack trace
// as the body of the response, otherwise a generic error message is returned.
func Recover(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if !helpers.IsDevMode() {
					http.Error(w, `{
						"error": true
					}`, http.StatusInternalServerError)
					return
				}

				stack := debug.Stack()
				parsedStack := parseStack(string(stack))
				w.Header().Set("content-type", "text/html")
				fmt.Fprintf(w, "<h1>Panic: %v</h1><pre>%s</pre>", err, parsedStack)
				return
			}
		}()
		h.ServeHTTP(w, r)
	})
}
