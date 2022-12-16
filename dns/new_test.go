package dns

import (
	"testing"

	dnsSym "github.com/taubyte/go-sdk-symbols/dns"
	"github.com/taubyte/go-sdk/errno"
)

func TestNewError(t *testing.T) {
	dnsSym.DnsNewResolver = func(resolverIdPtr *uint32) (error errno.Error) {
		return 1
	}

	_, err := NewResolver()
	if err == nil {
		t.Error("Expected an error")
		return
	}
}
