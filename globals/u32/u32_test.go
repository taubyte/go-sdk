package u32

import (
	"fmt"
	"testing"

	globalSym "github.com/taubyte/go-sdk-symbols/globals"
)

func TestNew(t *testing.T) {
	globalSym.MockData{}.Mock()

	_, err := Get("hello")
	if err == nil {
		t.Error("Expected error")
		return
	}

	u, err := GetOrCreate("hello")
	if err != nil {
		t.Error(err)
		return
	}

	u.Set(u.Value() + 6)

	fmt.Println(u.Value())

	u2, err := Get("hello")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(u2.Value())
}
