package dns

import (
	"fmt"

	dnsSym "github.com/taubyte/go-sdk-symbols/dns"
)

// Reroute reroutes the resolver to the provided address and network
// Return an error
func (r Resolver) Reroute(address, network string) error {
	err := dnsSym.DnsRerouteResolver(uint32(r), address, network)
	if err != 0 {
		return fmt.Errorf("Failed rerouting resolver with %v", err)
	}

	return nil
}

// Reset sets the current resolver to the default provided resolver
// Returns an error
func (r Resolver) Reset() error {
	err := dnsSym.DnsResetResolver(uint32(r))
	if err != 0 {
		return fmt.Errorf("Failed resetting resolver with %v", err)
	}

	return nil
}
