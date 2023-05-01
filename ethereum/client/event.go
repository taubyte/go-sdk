package ethereum

import (
	"fmt"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
	qf "github.com/taubyte/go-sdk/ethereum/client/queryFilter"
)

func (c Client) Subscribe(channelName string, filter *qf.QueryFilter) error {
	if err := filter.Parse(); err != nil {
		return err
	}

	if err0 := ethereumSym.EthSubscribeEvent(filter.SymReturn(uint32(c), channelName)); err0 != 0 {
		return fmt.Errorf("subscribing to event filter failed with: %s", err0)
	}

	return nil
}
