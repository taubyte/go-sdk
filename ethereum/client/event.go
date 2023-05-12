package ethereum

import (
	"fmt"
	"time"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
)

func (c Contract) Subscribe(event, channel string, ttl time.Duration) error {
	ttlParsed := uint32(ttl.Seconds())
	if err0 := ethereumSym.EthSubscribeContractEvent(uint32(c.client), c.id, event, channel, ttlParsed); err0 != 0 {
		return fmt.Errorf("subscribing to contract event `%s` failed with: %s", event, err0)
	}

	return nil
}
