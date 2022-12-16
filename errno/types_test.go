package errno

import (
	"testing"
)

func TestCap(t *testing.T) {
	if ErrorCap.String() != "ErrorCap" {
		t.Errorf("Expected ErrorCap to represent final error, but got `%s`", ErrorCap.String())
	}
}
