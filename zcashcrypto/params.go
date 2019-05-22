package zcashcrypto

type NetworkId [2]byte
type Prefix [1]byte

var SecretKeyPrefix = Prefix{140}

var MainnnetId NetworkId = NetworkId{0x1C}
var TestnetId NetworkId = NetworkId{0x1D, 0x25}
