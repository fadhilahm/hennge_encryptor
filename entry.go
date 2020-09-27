package main

import (
	// "./encrypt"
	// "fmt"
	"github.com/xlzd/gotp"
	"crypto/sha512"
	"hash"
)

func main() {
	// encrypt.Test()
	
	// try to use other people's package

	// initialize your own hasher
	// hasher := sha512.New()
	myHasher := gotp.Hasher{
		HashName: "sha-512",
		Digest: sha512.New,
	}

	totp := gotp.NewTOTP("fadhilahmetra@gmail.comHENNGECHALLENGE003", 10, 30, &myHasher)
	totp.At(1)
}

// Hasher represents `Hasher` struct
type Hasher struct {
	HashName string
	Digest func() hash.Hash
}