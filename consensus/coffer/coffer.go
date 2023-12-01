package coffer

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/yanhuangpai/go-utility/common"
	"github.com/yanhuangpai/go-utility/common/hexutil"
	"github.com/yanhuangpai/go-utility/core"
	"github.com/yanhuangpai/go-utility/core/state"
	"github.com/yanhuangpai/go-utility/crypto"
	"github.com/yanhuangpai/go-utility/uncdb"
)

// Coffer is the main structure for the Coffer consensus mechanism.
type Coffer struct {
	SuperAccount common.Address
	Signers      []Signer
	// Other necessary fields...
}

// CofferTransaction define a new transaction type in the blockchain.
// This involves modifying the data structure that represents a transaction.
type CofferTransaction struct {
	// Standard transaction fields
	Nonce    uint64
	GasPrice *big.Int
	GasLimit uint64
	To       *common.Address
	Value    *big.Int
	Data     []byte
	V, R, S  *big.Int

	// Fields specific to coffer
	Coffer Coffer
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

func NewCofferFromDB(db uncdb.Database) (*state.Coffer, error) {
	// Retrieve Coffer data from the database
	// Example: data := db.Get(/* coffer data key */)

	// Initialize and return a new Coffer instance using the retrieved data
	coffer := &state.Coffer{
		// Initialize fields using data
	}

	return coffer, nil
}

// ChangeSuperAccount updates the super account, need the old super account's private key to sign
func (c *Coffer) ChangeSuperAccount(signatureHex, newSuperAccount string) error {
	// Verify the super account
	result, _ := verifySig(c.SuperAccount.String(), newSuperAccount, signatureHex)
	if !result {

		return errors.New("only the current super account can update to a new super account")
	}
	// Update the super account
	c.SuperAccount = common.HexToAddress(newSuperAccount)
	return nil
}

func verifySig(from, msg, sigHex string) (bool, error) {
	sig, err := hexutil.Decode(sigHex)
	if err != nil {
		err = fmt.Errorf("invalid sig ('%s'), %w", sigHex, err)
		log.Fatal(err)
		return false, err
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
		return false, err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pk)
	fmt.Printf("recovered address is ('%s')", recoveredAddr)
	return strings.EqualFold(from, recoveredAddr.Hex()), nil
}
