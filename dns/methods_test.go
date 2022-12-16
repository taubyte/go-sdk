package dns

import (
	"testing"

	dnsSym "github.com/taubyte/go-sdk-symbols/dns"
)

func TestRerouteError(t *testing.T) {
	testId := uint32(8)
	testAddress := "8.8.8.8:53"
	testNet := "UDP"
	dnsSym.MockReroute(testId, testAddress, testNet)

	if Resolver(0).Reroute(testAddress, testNet) == nil {
		t.Error("Expected error")
		return
	}

	if Resolver(testId).Reroute(testAddress+"other", testNet) == nil {
		t.Error("Expected error")
		return
	}

	if Resolver(testId).Reroute(testAddress, testNet+"netty") == nil {
		t.Error("Expected error")
		return
	}
}

func TestResetError(t *testing.T) {
	testId := uint32(9)
	dnsSym.MockReset(testId)

	if Resolver(0).Reset() == nil {
		t.Error("Expected error")
		return
	}
}
