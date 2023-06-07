package rpc

func Header(key, value string) ClientOption {
	return func(d *DialOptions) error {
		if headers, exists := d.Headers[key]; !exists {
			d.Headers[key] = []string{value}
		} else {
			d.Headers[key] = append(headers, value)
		}

		return nil
	}
}
