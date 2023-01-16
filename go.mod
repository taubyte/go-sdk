module github.com/taubyte/go-sdk

go 1.18

// Direct Taubyte imports
require github.com/taubyte/go-sdk-symbols v0.1.4

// Direct imports
require (
	github.com/ipfs/go-cid v0.0.7
	golang.org/x/crypto v0.1.0
	golang.org/x/exp v0.0.0-20221026153819-32f3d567a233
)

replace github.com/taubyte/go-sdk-symbols => ../go-sdk-symbols

// Indirect imports
require (
	github.com/klauspost/cpuid/v2 v2.0.12 // indirect
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/multiformats/go-base32 v0.0.3 // indirect
	github.com/multiformats/go-base36 v0.1.0 // indirect
	github.com/multiformats/go-multibase v0.0.3 // indirect
	github.com/multiformats/go-multihash v0.0.15 // indirect
	github.com/multiformats/go-varint v0.0.6 // indirect
	golang.org/x/sys v0.1.0 // indirect
)
