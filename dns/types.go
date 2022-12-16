package dns

type Resolver uint32
type DefaultResolver interface {
	LookupTXT(name string) ([]string, error)
	LookupAddress(name string) ([]string, error)
	LookupCNAME(name string) (string, error)
	LookupMX(name string) ([]*MxResp, error)
	Reroute(address, network string) error
	Reset() error
}

type MxResp struct {
	Host string
	Pref uint16
}
