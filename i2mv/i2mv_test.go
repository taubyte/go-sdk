package i2mv

import (
	"io"
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/i2mv"
)

func TestNew(t *testing.T) {
	symbols.MockNew(0)

	// Successes
	if _, _, err := New([]byte{0}, false); err != nil {
		t.Errorf("new failed with: %s", err)
		return
	}

	if _, _, err := New([]byte{0}, true); err != nil {
		t.Errorf("new failed with: %s", err)
		return
	}

	// Failures
	if _, _, err := New(nil, false); err == nil {
		t.Error("expected error")
		return
	}

	symbols.MockNew(-1)
	if _, _, err := New([]byte{0}, false); err == nil {
		t.Error("expected error")
		return
	}
}

func TestOpen(t *testing.T) {
	// Success
	var testId uint32 = 1
	var testSize uint32 = 4
	symbols.MockSize(testSize)

	mv, err := Open(testId)
	if err != nil {
		t.Errorf("open failed with: %s", err)
		return
	}

	if mv.Id() != testId {
		t.Errorf("expected memory view id `%d` got `%d`", testId, mv.Id())
		return
	}

	if mv.Size() != testSize {
		t.Errorf("expected memory view size `%d` got `%d`", testSize, mv.Size())
		return
	}

	if mv.OffSet() != 0 {
		t.Errorf("expected memory view offset `0` got `%d`", mv.OffSet())
	}

	// Failure
	symbols.MockSize(0)

	if _, err = Open(0); err == nil {
		t.Errorf("expected error")
		return
	}
}

func TestRead(t *testing.T) {
	var testId uint32 = 1
	var testSize uint32 = 4
	symbols.MockSize(testSize)
	symbols.MockRead(testId)

	// Successes
	mv, err := Open(testId)
	if err != nil {
		t.Errorf("opening memory view failed with: %s", err)
		return
	}

	data := make([]byte, 1)

	if _, err = mv.Read(data); err != nil {
		t.Errorf("reading memory view failed with: %s", err)
		return
	}

	if mv.OffSet() != int64(len(data)) {
		t.Errorf("expected offset to be shifted from 0 to `%d` got `%d`", len(data), mv.OffSet())
		return
	}

	// Failures
	if _, err = mv.Read(nil); err == nil {
		t.Error("expected error")
		return
	}

	symbols.MockRead(testId + 1)

	if _, err = mv.Read(data); err == nil {
		t.Errorf("expected error")
		return
	}
}

func TestSeek(t *testing.T) {
	var testId uint32 = 1
	var testSize uint32 = 4
	symbols.MockSize(testSize)

	mv, err := Open(testId)
	if err != nil {
		t.Errorf("opening memory view failed with: %s", err)
		return
	}

	// Successes
	if _, err = mv.Seek(int64(testSize), io.SeekCurrent); err != nil {
		t.Errorf("seek from current failed with: %s", err)
		return
	}

	if _, err = mv.Seek(-4, io.SeekEnd); err != nil {
		t.Errorf("seek from end failed with: %s", err)
		return
	}

	if _, err = mv.Seek(5, io.SeekStart); err != io.EOF {
		t.Errorf("seek from start past max index failed with: %s", err)
		return
	}

	// Failures
	if _, err = mv.Seek(0, io.SeekEnd+2); err == nil {
		t.Errorf("expected error")
		return
	}

	if _, err = mv.Seek(-1, io.SeekStart); err == nil {
		t.Errorf("expected error")
		return
	}

}

func TestClose(t *testing.T) {
	var testId uint32 = 1
	var testSize uint32 = 4
	symbols.MockSize(testSize)

	mv, err := Open(testId)
	if err != nil {
		t.Errorf("opening memory view failed with: %s", err)
		return
	}

	symbols.MockClose(true)

	if err := mv.Close(); err != nil {
		t.Errorf("closing memory view failed with: %s", err)
		return
	}

	symbols.MockClose(false)
	if err := mv.Close(); err == nil {
		t.Error("expected error")
	}

}
