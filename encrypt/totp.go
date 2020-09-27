package encrypt

import (
	"fmt"
	"hash"
	"crypto/sha512"
)

// OTP represents `OTP` struct
type OTP struct {
	secret	string 		// secret in base32 format
	digits	int 		// number of integers in the OTP. we expect 10 for the task
	harsher *Harsher	// digest function to use in the HMAC
}

// Harsher represents `Harsher` struct
type Harsher struct {
	HashName string
	Digest func() hash.Hash
}

// TOTP represents time-based counters
type TOTP struct {
	OTP
	interval int
}

// NewOTP represents function that will return new `OTP` type
func NewOTP (secret string, digits int, harsher *Harsher) OTP {
	if harsher == nil {
		harsher = &Harsher{
			HashName: "sha-512",
			Digest: sha512.New,
		}
	}
	return OTP{
		secret: secret,
		digits: digits,
		harsher: harsher,
	}
}

// NewTOTP represents function that will generate new TOTP type
// func NewTOTP(secret string, digits, interval int, harsher *Harsher) *TOTP {
// 	otp := NewTOTP()
// }

// Test is just a test
func Test() {
	fmt.Println("TEST")
}