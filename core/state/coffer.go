// Copyright 2023 The go-utility Authors
// This file is part of the go-utility library.
//
// The go-utility library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-utility library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-utility library. If not, see <http://www.gnu.org/licenses/>.

package state

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/yanhuangpai/go-utility/common"
	"github.com/yanhuangpai/go-utility/core/types"
)

type Coffer struct {
	SuperAccount common.Address
	Signers      []Signer                      // Array for ordered storage
	SignerMap    map[common.Address]SignerInfo // Mapping for fast access
	TotalPower   uint64                        // Total cumulative power for random selection
}

type Signer struct {
	Address1 common.Address // Initial address set by the super account
	Address2 common.Address // Address set by the owner of the chipset
	power    uint64         // Weight or power of the signer
}

type SignerInfo struct {
	index int    // Index in the Signers array
	power uint64 // Weight or power of the signer
}

// Validate checks the validity of the SuperAccountUpdateTx transaction.
func (s *StateDB) Validate(tx *types.SuperAccountUpdateTx) error {
	// Verify the nonce
	currentNonce := s.GetNonce(tx.OldSuperAccount)
	if tx.Nonce != currentNonce {
		return fmt.Errorf("invalid nonce: got %d, want %d", tx.Nonce, currentNonce)
	}

	// Check the signature (simplified example)
	sender, err := tx.Sender()
	if err != nil {
		return fmt.Errorf("failed to verify signature: %v", err)
	}

	// Ensure the sender is the current super account
	if strings.ToLower(sender.Hex()) != strings.ToLower(tx.OldSuperAccount.Hex()) {
		return fmt.Errorf("unauthorized: sender is not the current super account")
	}

	// Check the account's balance for gas fees (simplified)
	balance := s.GetBalance(tx.OldSuperAccount)
	gasCost := new(big.Int).Mul(tx.GasPrice, new(big.Int).SetUint64(tx.GasLimit))
	if balance.Cmp(gasCost) < 0 {
		return fmt.Errorf("insufficient balance to pay for gas")
	}

	// Additional checks can be added here

	return nil
}

// ProcessTransaction updates the Coffer structure based on the transaction.
func (s *StateDB) ProcessTransaction(tx *types.SuperAccountUpdateTx) error {

	// Validate the transaction
	if err := s.Validate(tx); err != nil {
		return fmt.Errorf("transaction validation failed: %v", err)
	}
	// Update the super account
	s.Coffer.SuperAccount = tx.NewSuperAccount

	// Additional logic to process the transaction can be added here

	return nil
}

// func (s *StateDB) AddCofferSigner(addr1, addr2 common.Address, power uint64) {
// 	// Create a new Signer
// 	newSigner := Signer{
// 		Address1: addr1,
// 		Address2: addr2,
// 		power:    power,
// 	}

// 	// Append the new signer to the Signers slice
// 	s.coffer.Signers = append(s.coffer.Signers, newSigner)

// 	// Create a SignerInfo and add it to the SignerMap
// 	signerInfo := SignerInfo{
// 		index: len(s.coffer.Signers) - 1, // Index of the new signer in the Signers slice
// 		power: power,
// 	}
// 	s.coffer.SignerMap[addr1] = signerInfo

// 	// Update the total power in the Coffer
// 	s.coffer.TotalPower += power
// }
