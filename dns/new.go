package dns

import (
	"fmt"

	dnsSym "github.com/taubyte/go-sdk-symbols/dns"
)

// NewResolver creates and uses a default resolver
func NewResolver() (DefaultResolver, error) {
	var resolverId uint32
	err0 := dnsSym.DnsNewResolver(&resolverId)
	if err0 != 0 {
		return nil, fmt.Errorf("Creating new resolver failed with: %s", err0)
	}

	return Resolver(resolverId), nil
}
