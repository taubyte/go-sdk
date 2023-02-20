package event

import (
	"fmt"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
)

type httpRedirect struct {
	eventId uint32
	url     string
}

func (r httpRedirect) redirect(code uint32) error {
	err := httpEventSym.EventHttpRedirect(r.eventId, r.url, code)
	if err != 0 {
		return fmt.Errorf("redirecting HTTP request failed with: %s", err)
	}

	return nil
}

func (e Event) Redirect(url string) httpRedirect {
	return httpRedirect{uint32(e), url}
}

func (r httpRedirect) Temporary() error {
	return r.redirect(307)
}

func (r httpRedirect) Permanent() error {
	return r.redirect(308)
}

func (r httpRedirect) Code(code uint32) error {
	return r.redirect(code)
}
