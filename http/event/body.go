package event

import (
	"fmt"
	"io"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
	"github.com/taubyte/go-sdk/errno"
)

func (e Event) Body() EventBody {
	return EventBody(e)
}

func (b EventBody) Read(p []byte) (int, error) {
	var counter uint32
	err := httpEventSym.ReadHttpEventBody(uint32(b), &p[0], uint32(len(p)), &counter)
	if err != 0 {
		if err == errno.ErrorEOF {
			return int(counter), io.EOF
		} else {
			return 0, fmt.Errorf("Reading HTTP body Failed with: %s", err)
		}
	}

	return int(counter), nil
}

func (b EventBody) Close() error {
	err := httpEventSym.CloseHttpEventBody(uint32(b))
	if err != (0) {
		return fmt.Errorf("Failed closing http body with: %s", err)
	}
	return nil
}
