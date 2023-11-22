package coffer

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/yanhuangpai/go-utility/common"
	"github.com/yanhuangpai/go-utility/core"
	"github.com/yanhuangpai/go-utility/crypto"
)

// Coffer is the main structure for the Coffer consensus mechanism.
type Coffer struct {
	SuperAccount common.Address
	Signers      []Signer
	// Other necessary fields...
}

// NewCoffer creates a new Coffer consensus mechanism.
func NewCoffer(genesis *core.Genesis) (*Coffer, error) {
	superAccount, err := ExtractSuperAccountFromGenesis(genesis)
	if err != nil {
		return nil, err
	}

	return &Coffer{
		SuperAccount: superAccount,
		// ... initialize other fields ...
	}, nil
}

// UpdateSuperAccount updates the super account, need the old super account's private key to sign
func (c *Coffer) UpdateSuperAccount(hexSignature string, newSuperAccount common.Address) error {

	trimmedSignature := hexSignature
	if len(hexSignature) >= 2 && hexSignature[:2] == "0x" {
		trimmedSignature = hexSignature[2:]
	}

	// Decode the hex string to a byte slice
	signature, err := hex.DecodeString(trimmedSignature)
	if err != nil {
		log.Fatal(err)
	}
	if len(signature) != 65 {
		return errors.New("signature must be 65 bytes long")
	}

	// Construct the message that was supposedly signed
	message := fmt.Sprintf("UpdateSuperAccount:%s", newSuperAccount.Hex())
	messageHash := crypto.Keccak256Hash([]byte(message))

	// Recover the public key from the signature
	publicKey, err := crypto.SigToPub(messageHash.Bytes(), signature)
	if err != nil {
		return errors.New("failed to verify signature")
	}

	// Convert the public key to an Ethereum address
	signer := crypto.PubkeyToAddress(*publicKey)

	// Check if the signer is the current super account
	if signer != c.SuperAccount {
		return errors.New("only the current super account can update to a new super account")
	}

	// Update the super account
	c.SuperAccount = newSuperAccount
	return nil
}
