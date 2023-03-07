package codec

import (
	"errors"

	"github.com/ipfs/go-cid"
)

const CidBufferSize = 64

type cidReader struct {
	data []byte
}

func CidReader() *cidReader {
	_cid := make([]byte, CidBufferSize)
	return &cidReader{
		_cid,
	}
}

func (r *cidReader) Ptr() *byte {
	return &r.data[0]
}

func (r *cidReader) Parse() (cid.Cid, error) {
	_, cidFromBytes, err := cid.CidFromBytes(r.data)
	if err != nil {
		return cid.Cid{}, err
	}

	return cidFromBytes, err
}

type cidWriter struct {
	data []byte
}

func CidWriter(cid cid.Cid) (*cidWriter, error) {
	bytes := cid.Bytes()
	if bytes == nil || len(bytes) == 0 {
		return nil, errors.New("invalid cid")
	}

	return &cidWriter{bytes}, nil
}

func (w *cidWriter) Ptr() *byte {
	return &w.data[0]
}
