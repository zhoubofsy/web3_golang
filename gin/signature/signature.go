package signature

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

type Signer struct {
	PriateKey *ecdsa.PrivateKey
}

func NewSigner(privKey string) *Signer {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		fmt.Printf("Failed to create private key: %v\n%v", privKey, err)
		return nil
	}
	return &Signer{PriateKey: privateKey}
}

func (s *Signer) MakeSignature(data []byte) ([]byte, error) {
	dataHash := crypto.Keccak256(data)
	return crypto.Sign(dataHash, s.PriateKey)
}

func (s *Signer) VerfiySignature(data, signature []byte) bool {
	dataHash := crypto.Keccak256(data)
	publicKey := s.PriateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("error casting public key to ECDSA")
		return false
	}
	bytePub := crypto.FromECDSAPub(publicKeyECDSA)
	return crypto.VerifySignature(bytePub, dataHash, signature[:len(signature)-1])
}
