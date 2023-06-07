package rpc

// TODO: Implement rest of Dial Options
type DialOptions struct {
	Headers map[string][]string `json:"Headers"`
}

type ClientOption func(d *DialOptions) error
