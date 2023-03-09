package node_test

import (
	"fmt"
	"time"

	"github.com/ipfs/go-cid"
	symbols "github.com/taubyte/go-sdk-symbols/p2p/node"
	"github.com/taubyte/go-sdk/p2p/node"
)

var (
	cid1 cid.Cid
	cid2 cid.Cid
	cid3 cid.Cid
)

// Initialize some CIDs test
func init() {
	var err error
	cid1, err = cid.Parse("bafzaajaiaejcatsa2r73dij2iewq47p2c6runxmvht2evx6agmojnk3pjflfcn52")
	if err != nil {
		panic(err)
	}

	cid2, err = cid.Parse("bafzaajaiaejcbiutf54aucf7ej47lrr27fuqarekshylvkbnwzd77p23wc2gzl3s")
	if err != nil {
		panic(err)
	}

	cid3, err = cid.Parse("bafzaajaiaejcbsmuszumjrisswnmvahkpycbjrlzhzrzdjryyqej4kww73kyxnnd")
	if err != nil {
		panic(err)
	}
}

func ExampleDiscover() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{
		DiscoverMax:     3,
		DiscoverTimeout: uint32(time.Second * 2),
		DiscoverId:      4,
		Peers:           []cid.Cid{cid1, cid2, cid3},
	}.Mock()

	peers, err := node.Discover(3, time.Second*2)
	if err != nil {
		return
	}

	for _, peer := range peers {
		fmt.Println(peer.String())
	}

	// Output:
	// bafzaajaiaejcatsa2r73dij2iewq47p2c6runxmvht2evx6agmojnk3pjflfcn52
	// bafzaajaiaejcbiutf54aucf7ej47lrr27fuqarekshylvkbnwzd77p23wc2gzl3s
	// bafzaajaiaejcbsmuszumjrisswnmvahkpycbjrlzhzrzdjryyqej4kww73kyxnnd
}
