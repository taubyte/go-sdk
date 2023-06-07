module github.com/taubyte/go-sdk

go 1.18

// Direct Taubyte imports
require github.com/taubyte/go-sdk-symbols v0.2.6

// Direct imports
require (
	github.com/ipfs/go-cid v0.0.7
	golang.org/x/crypto v0.1.0
	golang.org/x/exp v0.0.0-20221026153819-32f3d567a233
	gotest.tools v2.2.0+incompatible
	gotest.tools/v3 v3.4.0

)

// Indirect imports
require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/ethereum/go-ethereum v1.10.26 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/cpuid/v2 v2.0.12 // indirect
	github.com/mailru/easyjson v0.7.7
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/multiformats/go-base32 v0.0.3 // indirect
	github.com/multiformats/go-base36 v0.1.0 // indirect
	github.com/multiformats/go-multibase v0.0.3 // indirect
	github.com/multiformats/go-multihash v0.0.15 // indirect
	github.com/multiformats/go-varint v0.0.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.1.0 // indirect
)
