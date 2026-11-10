package sized

import "github.com/taubyte/go-sdk/errno"

// attempts caps how often Read asks again for a value that keeps growing.
const attempts = 5

// Read fetches a value the host hands over in two steps: its size, then its
// data. The host refuses the data step with ErrorSizeMismatch when the value
// outgrew the size it reported, so the size is asked for again.
// A zero size yields an empty value without a data step.
func Read(size func(*uint32) errno.Error, data func(*byte) errno.Error) (value []byte, sizeErr, dataErr errno.Error) {
	for i := 1; ; i++ {
		var n uint32
		if sizeErr = size(&n); sizeErr != 0 {
			return nil, sizeErr, 0
		}
		if n == 0 {
			return []byte{}, 0, 0
		}

		value = make([]byte, n)
		dataErr = data(&value[0])
		if dataErr != errno.ErrorSizeMismatch || i == attempts {
			return value, 0, dataErr
		}
	}
}
