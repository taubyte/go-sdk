package ethereum

import (
	"fmt"
	"time"

	ethereumSym "github.com/taubyte/go-sdk-symbols/ethereum/client"
)

func (e Event) Subscribe(channel string, ttl time.Duration) error {
	ttlParsed := uint32(ttl.Seconds())
	if err0 := ethereumSym.EthSubscribeContractEvent(uint32(e.client), e.contractID, e.name, channel, ttlParsed); err0 != 0 {
		return fmt.Errorf("subscribing to contract event `%s` failed with: %s", e.name, err0)
	}

	return nil
}
