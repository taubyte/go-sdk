package event

import (
	"fmt"

	eventSym "github.com/taubyte/go-sdk-symbols/event"
)

func (e HttpEvent) Redirect(url string, code uint32) error {
	err := eventSym.EventHttpRedirect(uint32(e), url, code)
	if err != 0 {
		return fmt.Errorf("Redirecting HTTP request failed with: %s", err)
	}

	return nil
}
