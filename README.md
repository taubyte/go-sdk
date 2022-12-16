# taubyte/go-sdk

## Usage 

```golang

import (
	"fmt"
	"io"

	"github.com/taubyte/go-sdk"
)

func main() {}

//export ping
func ping(e sdk.Event) uint32 {
	fmt.Println("Event Type", e.Type())
	h, _ := e.HTTP()

// Gets the HTTP Method
	method, _ := h.Method()
	fmt.Println("HTTP Method", method)

// Sets Header of key 'Test' to value 'Test Header'
	h.Headers().Set("X-Test", "Test Header")
	body, _ := io.ReadAll(h.Body())
	host, _ := h.Host()
	path, _ := h.Path()
	userAgent, _ := h.UserAgent()
	acceptEncoding, _ := h.Headers().Get("Accept-Encoding")
	headers, _ := h.Headers().List()
	queries, _ := h.Query().List()
	query, _ := h.Query().Get("name")

// Writes response body
	h.Write([]byte(fmt.Sprintf(`{"ping": "pong","body": "%v","host": "%s","path": "%s","useragent": "%s","acceptencoding": "%s","query": "%s"}`, body, host, path, userAgent, acceptEncoding, query)))

// Closes the Body
	err := h.Body().Close()
	if err != nil {
		panic(err)
	}

// Returns status of 404 
	h.Return(404)
	return 0
}

```