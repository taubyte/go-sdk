package logs

import (
	"github.com/mailru/easyjson"
	"github.com/taubyte/go-sdk/ethereum/client/bytes"
)

type EventLog struct {
	Log   *Log   `json:"log"`
	Error string `json:"error"`
}

type Log struct {
	Address     *bytes.Address `json:"address"`
	Topics      []*bytes.Hash  `json:"topics"`
	Data        []byte         `json:"data"`
	BlockNumber uint64         `json:"blockNumber"`
	TxHash      *bytes.Hash    `json:"transactionHash"`
	TxIndex     uint           `json:"transactionIndex"`
	BlockHash   *bytes.Hash    `json:"blockHash"`
	Index       uint           `json:"logIndex"`
	Removed     bool           `json:"removed"`
}

func (l *Log) Unmarshal(data []byte) error {
	return easyjson.Unmarshal(data, l)
}
