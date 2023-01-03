package event

import (
	"fmt"
	"net/http"

	eventSym "github.com/taubyte/go-sdk-symbols/event"
)

type httpRedirect struct {
	eventId uint32
	url     string
}

func (r httpRedirect) redirect(code uint32) error {
	err := eventSym.EventHttpRedirect(r.eventId, r.url, code)
	if err != 0 {
		return fmt.Errorf("redirecting HTTP request failed with: %s", err)
	}

	return nil
}

func (e HttpEvent) Redirect(url string) httpRedirect {
	return httpRedirect{uint32(e), url}
}

func (r httpRedirect) Temporary() error {
	return r.redirect(http.StatusTemporaryRedirect)
}

func (r httpRedirect) Permanent() error {
	return r.redirect(http.StatusPermanentRedirect)
}
