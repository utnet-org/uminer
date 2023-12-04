package types

import (
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/yanhuangpai/go-utility/common"
	"github.com/yanhuangpai/go-utility/crypto"
	"github.com/yanhuangpai/go-utility/rlp"
)

// SuperAccountUpdateTx is the data structure for super account update transactions
type SuperAccountUpdateTx struct {
	OldSuperAccount common.Address `json:"oldSuperAccount"`
	NewSuperAccount common.Address `json:"newSuperAccount"`
	Nonce           uint64         `json:"nonce"`
	GasPrice        *big.Int       `json:"gasPrice"`
	GasLimit        uint64         `json:"gasLimit"`
	Signature       []byte         `json:"signature"`
}

// Sign signs the transaction with the given signer and private key.
func (tx *SuperAccountUpdateTx) Sign(signer Signer, prv *ecdsa.PrivateKey) error {
	if len(tx.Signature) == 0 {
		return errors.New("no signature present in transaction")
	}

	// Serialize the transaction excluding the signature
	encodedTx, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}

	// Hash the serialized transaction
	h := crypto.Keccak256Hash(encodedTx)
	signature, err := crypto.Sign(h[:], prv)
	if err != nil {
		return err
	}
	tx.Signature = signature
	return nil
}

// VerifySignature verifies the signature of the transaction.
func (tx *SuperAccountUpdateTx) VerifySignature(signer Signer) (common.Address, error) {
	if len(tx.Signature) == 0 {
		return common.Address{}, errors.New("no signature")
	}
	if len(tx.Signature) == 0 {
		return common.Address{}, errors.New("no signature present in transaction")
	}

	// Serialize the transaction excluding the signature
	encodedTx, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return common.Address{}, err
	}

	// Hash the serialized transaction
	h := crypto.Keccak256Hash(encodedTx)
	pubkey, err := crypto.Ecrecover(h[:], tx.Signature)
	if err != nil {
		return common.Address{}, err
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(pubkey[1:])[12:])
	return addr, nil
}

// EncodeRLP serializes the transaction into the Ethereum RLP format.
func (tx *SuperAccountUpdateTx) EncodeRLP() ([]byte, error) {
	return rlp.EncodeToBytes(tx)
}

// DecodeRLP deserializes the Ethereum RLP format into the transaction.
func (tx *SuperAccountUpdateTx) DecodeRLP(encodedTx []byte) error {
	return rlp.DecodeBytes(encodedTx, tx)
}

// Sender returns the address derived from the signature (i.e., the sender of the transaction).
func (tx *SuperAccountUpdateTx) Sender() (common.Address, error) {
	if len(tx.Signature) == 0 {
		return common.Address{}, errors.New("no signature present in transaction")
	}

	// Serialize the transaction excluding the signature
	encodedTx, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return common.Address{}, err
	}

	// Hash the serialized transaction
	h := crypto.Keccak256Hash(encodedTx)

	// Recover the public key from the signature
	pubkey, err := crypto.SigToPub(h.Bytes(), tx.Signature)
	if err != nil {
		return common.Address{}, err
	}

	return crypto.PubkeyToAddress(*pubkey), nil
}
