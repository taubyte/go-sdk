package queryfilter

import (
	"errors"

	"github.com/taubyte/go-sdk/ethereum/client/bytes"
	"github.com/taubyte/go-sdk/utils/booleans"
	"github.com/taubyte/go-sdk/utils/codec"
)

func (q *QueryFilter) SymReturn(clientId uint32, channel string) (_clientId uint32, blockIdentifierPtr *byte, blockIdentifierSize uint32, isHash uint32, fromBlockPtr *byte, fromBlockSize uint32, toBlockPtr *byte, toBlockSize uint32, addressesPtr *byte, addressesSize uint32, topicsPtr *byte, topicsSize uint32, _channel string, ttl uint32) {
	return clientId, q.raw.blockIdentifierPtr, q.raw.blockIdentifierSize, q.raw.isHash, q.raw.fromBlockPtr, q.raw.fromBlockSize, q.raw.toBlockPtr, q.raw.toBlockSize, q.raw.addressesPtr, q.raw.addressesSize, q.raw.topicsPtr, q.raw.topicsSize, channel, q.raw.ttl
}

func (q *QueryFilter) Parse() error {
	var nilByte = byte(0)

	raw := queryFilterRaw{}

	if q.BlockHash != nil {
		if q.FromBlock != nil || q.ToBlock != nil {
			return errors.New("cannot filter specific block, and range of blocks")
		}

		raw.blockIdentifierPtr = &q.BlockHash[0]
		raw.blockIdentifierSize = bytes.HashByteLength
		raw.isHash = booleans.FromBool(true)
	} else {
		raw.blockIdentifierPtr = &nilByte
	}

	if q.FromBlock != nil {
		raw.fromBlockPtr = &q.FromBlock.Bytes()[0]
		raw.fromBlockSize = uint32(len(q.FromBlock.Bytes()))
	} else {
		raw.fromBlockPtr = &nilByte
	}

	if q.ToBlock != nil {
		raw.toBlockPtr = &q.ToBlock.Bytes()[0]
		raw.toBlockSize = uint32(len(q.ToBlock.Bytes()))
	} else {
		raw.toBlockPtr = &nilByte
	}

	if q.Addresses != nil {
		var addresses []byte
		if err := codec.Convert(q.Addresses).To(&addresses); err != nil {
			return err
		}

		raw.addressesPtr = &addresses[0]
		raw.addressesSize = uint32(len(addresses))
	} else {
		raw.addressesPtr = &nilByte
	}

	if q.Topics != nil {
		var topics []byte
		if err := codec.Convert(q.Topics).To(&topics); err != nil {
			return err
		}

		raw.topicsPtr = &topics[0]
		raw.topicsSize = uint32(len(topics))
	} else {
		raw.topicsPtr = &nilByte
	}

	return nil
}
