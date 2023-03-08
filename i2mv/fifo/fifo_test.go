package fifo

import (
	"io"
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/i2mv/fifo"
	"gotest.tools/v3/assert"
)

func TestNew(t *testing.T) {
	symbols.MockNew(0)

	if id, _ := New(true); id != 0 {
		t.Errorf("expected id `0` got `%d`", id)
	}

}

func TestOpen(t *testing.T) {
	symbols.MockOpen(0, false)

	ff, err := Open(0)
	assert.NilError(t, err)
	assert.Equal(t, ff.Id(), uint32(0))

	symbols.MockOpen(1, false)

	if _, err := Open(0); err == nil {
		t.Error("expected error")
	}
}

func TestWrite(t *testing.T) {
	symbols.MockAll(0, true, nil)
	fakeData := []byte("hello world")

	_, ff := New(true)
	n, err := ff.Write(fakeData)
	assert.NilError(t, err)
	assert.Equal(t, n, len(fakeData))

	symbols.MockPush(1)

	if _, err = ff.Write(fakeData); err == nil {
		t.Error("expected error")
		return
	}
}

func TestRead(t *testing.T) {
	fakeData := []byte("hello world")
	symbols.MockAll(0, false, fakeData)

	ff, err := Open(0)
	assert.NilError(t, err)

	data := make([]byte, 11)
	n, err := ff.Read(data)
	assert.NilError(t, err)
	assert.Equal(t, n, len(data))

	symbols.MockPop(1, fakeData)

	if _, err = ff.Read(data); err == nil {
		t.Error("expected error")
		return
	}

	symbols.MockPop(0, fakeData)

	data = make([]byte, 12)
	_, err = ff.Read(data)
	assert.Error(t, err, io.EOF.Error())
}

func TestClose(t *testing.T) {
	symbols.MockAll(0, true, nil)

	ff, err := Open(0)
	assert.NilError(t, err)

	err = ff.Close()
	assert.NilError(t, err)

	symbols.MockOpen(0, false)

	ff, err = Open(0)
	assert.NilError(t, err)

	err = ff.Close()
	assert.Error(t, err, "fifo `0` is not closable")
}
