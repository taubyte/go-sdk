package dns

import (
	"testing"
	"unsafe"

	dnsSym "github.com/taubyte/go-sdk-symbols/dns"
	"github.com/taubyte/go-sdk/errno"
)

func TestTxTError(t *testing.T) {
	testId := uint32(18)
	testKey := "taubyte.com"
	testData := map[string][]byte{}
	dnsSym.MockTxT(testId, testData)

	// Error from _size
	_, err := Resolver(testId).LookupTXT(testKey)
	if err == nil {
		t.Error("Expected Error")
		return
	}

	// Size 0
	testData[testKey] = []byte{}
	_, err = Resolver(testId).LookupTXT(testKey)
	if err != nil {
		t.Error(err)
		return
	}

	// Id doesn't match
	_, err = Resolver(7).LookupTXT(testKey)
	if err == nil {
		t.Error("Expected Error")
		return
	}

	// Error from _data
	testData[testKey] = []byte("Hello, world")
	dnsSym.DnsLookupTxT = func(resolverId uint32, name string, txtPtr *byte) (error errno.Error) {
		return 1
	}
	_, err = Resolver(testId).LookupTXT(testKey)
	if err == nil {
		t.Error("Expected Error")
		return
	}

	// conversion error
	dnsSym.DnsLookupTxT = func(resolverId uint32, name string, txtPtr *byte) (error errno.Error) {
		d := unsafe.Slice(txtPtr, 22)
		copy(d, []byte("Hello, world"))
		return 0
	}
	dnsSym.DnsLookupTxTSize = func(resolverId uint32, name string, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 12
		return 0
	}
	_, err = Resolver(testId).LookupTXT(testKey)
	if err == nil {
		t.Error("Expected Error")
		return
	}
}

func TestAddressError(t *testing.T) {
	testId := uint32(18)
	testKey := "taubyte.com"
	testData := map[string][]byte{}
	dnsSym.MockAddress(testId, testData)

	// Error from _size
	_, err := Resolver(testId).LookupAddress(testKey)
	if err == nil {
		t.Error("Expected Error")
		return
	}

	// Size 0
	testData[testKey] = []byte{}
	_, err = Resolver(testId).LookupAddress(testKey)
	if err != nil {
		t.Error(err)
		return
	}

	// Id doesn't match
	_, err = Resolver(7).LookupAddress(testKey)
	if err == nil {
		t.Error("Expected Error")
		return
	}

	// Error from _data
	testData[testKey] = []byte("Hello, world")
	dnsSym.DnsLookupAddress = func(resolverId uint32, name string, txtPtr *byte) (error errno.Error) {
		return 1
	}
	_, err = Resolver(testId).LookupAddress(testKey)
	if err == nil {
		t.Error("Expected Error")
		return
	}

	// Conversion error
	dnsSym.DnsLookupAddress = func(resolverId uint32, name string, txtPtr *byte) (error errno.Error) {
		d := unsafe.Slice(txtPtr, 22)
		copy(d, []byte("Hello, world"))
		return 0
	}
	dnsSym.DnsLookupAddressSize = func(resolverId uint32, name string, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 12
		return 0
	}
	_, err = Resolver(testId).LookupAddress(testKey)
	if err == nil {
		t.Error("Expected Error")
		return
	}
}

func TestCnameError(t *testing.T) {
	testId := uint32(18)
	testKey := "taubyte.com"
	testData := map[string][]byte{}
	dnsSym.MockCname(testId, testData)

	_, err := Resolver(testId).LookupCNAME(testKey)
	if err == nil {
		t.Error("Expected Error")
		return
	}

	// Size 0
	testData[testKey] = []byte{}
	_, err = Resolver(testId).LookupCNAME(testKey)
	if err != nil {
		t.Error(err)
		return
	}

	// Id doesn't match
	_, err = Resolver(7).LookupCNAME(testKey)
	if err == nil {
		t.Error("Expected Error")
		return
	}

	// Error from _data
	testData[testKey] = []byte("Hello, world")
	dnsSym.DnsLookupCNAME = func(resolverId uint32, name string, cnamePtr *byte) (error errno.Error) {
		return 1
	}
	_, err = Resolver(testId).LookupCNAME(testKey)
	if err == nil {
		t.Error("Expected Error")
		return
	}
}

func TestMxError(t *testing.T) {
	testId := uint32(18)
	testKey := "taubyte.com"
	testData := map[string][]byte{}
	dnsSym.MockMx(testId, testData)

	// Error from _size
	_, err := Resolver(testId).LookupMX(testKey)
	if err == nil {
		t.Error("Expected Error")
		return
	}

	// Size 0
	testData[testKey] = []byte{}
	_, err = Resolver(testId).LookupMX(testKey)
	if err != nil {
		t.Error(err)
		return
	}

	// Id doesn't match
	_, err = Resolver(7).LookupMX(testKey)
	if err == nil {
		t.Error("Expected Error")
		return
	}

	// Error from _data
	testData[testKey] = []byte("Hello, world")
	dnsSym.DnsLookupMX = func(resolverId uint32, name string, txtPtr *byte) (error errno.Error) {
		return 1
	}
	_, err = Resolver(testId).LookupMX(testKey)
	if err == nil {
		t.Error("Expected Error")
		return
	}

	// Conversion error
	dnsSym.DnsLookupMX = func(resolverId uint32, name string, txtPtr *byte) (error errno.Error) {
		d := unsafe.Slice(txtPtr, 300)
		copy(d, []byte("32144124/24231315326245315332415321253"))
		return 0
	}
	dnsSym.DnsLookupMXSize = func(resolverId uint32, name string, sizePtr *uint32) (error errno.Error) {
		*sizePtr = 300
		return 0
	}
	_, err = Resolver(testId).LookupMX(testKey)
	if err == nil {
		t.Error("Expected Error")
		return
	}

}
