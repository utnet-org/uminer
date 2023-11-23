package coffer

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/yanhuangpai/go-utility/common"
	"github.com/yanhuangpai/go-utility/common/hexutil"
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
func (c *Coffer) UpdateSuperAccount(signatureHex, newSuperAccount string) error {

	// signatureHex = strings.TrimPrefix(signatureHex, "0x")
	// signatureBytes, err := hex.DecodeString(signatureHex)
	// if err != nil {
	// 	log.Fatalf("Failed to decode hex string: %v", err)
	// }
	// if len(signatureBytes) != 65 {
	// 	log.Fatalf("Invalid signature length: got %d bytes, want 65 bytes", len(signatureBytes))
	// }

	// if signatureBytes[crypto.RecoveryIDOffset] != 27 && signatureBytes[crypto.RecoveryIDOffset] != 28 {
	// 	log.Fatalf("invalid Ethereum signature (V is not 27 or 28)")
	// }

	// signatureBytes[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	// // Construct the message that was supposedly signed
	// message := fmt.Sprintf(newSuperAccount.Hex())
	// messageHash := crypto.Keccak256Hash([]byte(message))

	// // Recover the public key from the signature
	// publicKey, err := crypto.SigToPub(messageHash.Bytes(), signatureBytes)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return errors.New("failed to verify signature")
	// }

	// // Convert the public key to an Ethereum address
	// signer := crypto.PubkeyToAddress(*publicKey)

	// // Check if the signer is the current super account
	// if signer != c.SuperAccount {
	// 	return errors.New("only the current super account can update to a new super account : " + signer.String())
	// }

	// Update the super account
	result, recoverAddr := verifySig(c.SuperAccount.String(), newSuperAccount, signatureHex)
	if !result {
		return errors.New("only the current super account can update to a new super account : " + recoverAddr.String())
	}
	c.SuperAccount = common.HexToAddress(newSuperAccount)
	return nil
}

func verifySig(from, msg, sigHex string) (bool, *common.Address) {
	sig, err := hexutil.Decode(sigHex)
	if err != nil {
		err = fmt.Errorf("invalid sig ('%s'), %w", sigHex, err)
		log.Fatal(err)
		return false, nil
	}

	messageToHash := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	msgHash := crypto.Keccak256Hash([]byte(messageToHash))
	// ethereum "black magic" :(
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27
	}

	pk, err := crypto.SigToPub(msgHash.Bytes(), sig)
	if err != nil {
		err = fmt.Errorf("failed to recover public key from sig ('%s'), %w", sigHex, err)
		log.Fatal(err)
		return false, nil
	}

	recoveredAddr := crypto.PubkeyToAddress(*pk)
	fmt.Printf("recovered address is ('%s')", recoveredAddr)
	return strings.EqualFold(from, recoveredAddr.Hex()), &recoveredAddr
}
