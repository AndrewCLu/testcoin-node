package crypto

import (
	"testing"
)

// Tests signing and verification using newly generated digital signature keys
func TestSignVerifyWithDigitalSignatureKeys(t *testing.T) {
	encodedPublicKey, encodedPrivateKey, err := NewDigitalSignatureKeys()
	if err != nil {
		t.Fatalf(`Failed to create new digital signature keys`)
	}

	bytesA := []byte("Go BTC")
	bytesB := []byte("Go ETH")

	s1, _ := SignByteArray(bytesA, encodedPrivateKey)
	verify1, _ := VerifyByteArray(bytesA, encodedPublicKey, s1)
	if !verify1 {
		t.Fatalf(`Verification false for byte array: %v`, bytesA)
	}

	s2, _ := SignByteArray(bytesA, encodedPrivateKey)
	verify2, _ := VerifyByteArray(bytesB, encodedPublicKey, s2)
	if verify2 {
		t.Fatalf(`Verification succeeded for two unequal byte arrays: %v and %v`, bytesA, bytesB)
	}
}
