package dns

import (
	"fmt"
	"strconv"
	"strings"

	dnsSym "github.com/taubyte/go-sdk-symbols/dns"
	"github.com/taubyte/go-sdk/utils/codec"
)

// LookupTXT returns the DNS TXT records for the given domain name.
func (r Resolver) LookupTXT(fqdn string) ([]string, error) {
	var size uint32
	err := dnsSym.DnsLookupTxTSize(uint32(r), fqdn, &size)
	if err != 0 {
		return nil, fmt.Errorf("Failed lookupTxTSize with %v", err)
	}
	if size == 0 {
		return []string{}, nil
	}

	recordBytes := make([]byte, size)
	err = dnsSym.DnsLookupTxT(uint32(r), fqdn, &recordBytes[0])
	if err != 0 {
		return nil, fmt.Errorf("Failed lookup txt with %v", err)
	}

	var records []string
	err0 := codec.Convert(recordBytes).To(&records)
	if err0 != nil || len(records) == 0 {
		return nil, fmt.Errorf("Converting TxT records to []strings with %s", err0)
	}

	return records, nil
}

// LookupAddr performs a reverse lookup for the given address, returning a list
// of names mapping to that address and an error.
func (r Resolver) LookupAddress(ip string) ([]string, error) {
	var size uint32
	err := dnsSym.DnsLookupAddressSize(uint32(r), ip, &size)
	if err != 0 {
		return nil, fmt.Errorf("Failed lookupAddressSize with %v", err)
	}
	if size == 0 {
		return []string{}, nil
	}

	addressBytes := make([]byte, size)
	err = dnsSym.DnsLookupAddress(uint32(r), ip, &addressBytes[0])
	if err != 0 {
		return nil, fmt.Errorf("Failed lookup address with %v", err)
	}

	var address []string
	err0 := codec.Convert(addressBytes).To(&address)
	if err0 != nil || len(address) == 0 {
		return nil, fmt.Errorf("Converting address to []strings with %s", err0)
	}

	return address, nil
}

// LookupCNAME returns the canonical name for the given host.
//
// LookupCNAME does not return an error if the given name does not
// contain DNS "CNAME" records, as long as the name resolves to
// address records.
func (r Resolver) LookupCNAME(fqdn string) (string, error) {
	var size uint32
	err := dnsSym.DnsLookupCNAMESize(uint32(r), fqdn, &size)
	if err != 0 {
		return "", fmt.Errorf("Failed lookupCNAMESize with %v", err)
	}
	if size == 0 {
		return "", nil
	}

	cname := make([]byte, size)
	err = dnsSym.DnsLookupCNAME(uint32(r), fqdn, &cname[0])
	if err != 0 {
		return "", fmt.Errorf("Failed CNAME lookup with %v", err)
	}

	host := strings.Trim(string(cname), "\x00")

	return host, nil
}

// LookupMX returns the DNS MX records for the given name and an error.
func (r Resolver) LookupMX(fqdn string) ([]*MxResp, error) {
	mxRecords := make([]*MxResp, 0)
	var size uint32
	err := dnsSym.DnsLookupMXSize(uint32(r), fqdn, &size)
	if err != 0 {
		return nil, fmt.Errorf("Failed lookupMXSize with %v", err)
	}
	if size == 0 {
		return []*MxResp{}, nil
	}

	_mxRecords := make([]byte, size)
	err = dnsSym.DnsLookupMX(uint32(r), fqdn, &_mxRecords[0])
	if err != 0 {
		return nil, fmt.Errorf("Failed MX lookup with %v", err)
	}

	// TODO: use string slice conversion
	values := strings.Split(string(_mxRecords), "/")
	for idx, value := range values {
		// Skip every other except first index and last index->so we get host/pref
		if idx%2 == 1 && idx != 0 || idx == len(values)-1 {
			continue
		}

		_value, err := strconv.ParseInt(strings.Trim(values[idx+1], "\x00"), 10, 32)
		if err != nil {
			return nil, fmt.Errorf("Failed parsing to uint with %v", err)
		}

		host := strings.Trim(value, "\x00")
		pref := uint16(_value)

		newMx := &MxResp{
			Host: host,
			Pref: pref,
		}

		mxRecords = append(mxRecords, newMx)

	}

	return mxRecords, nil
}
