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

package unc

import (
	"errors"
	"math/big"

	"github.com/yanhuangpai/go-utility/common"
	"github.com/yanhuangpai/go-utility/core/state"
	"github.com/yanhuangpai/go-utility/core/types"
)

// CofferAPI provides an API to control the coffer.
type CofferAPI struct {
	e *Utility
}

// NewCofferAPI create a new CofferAPI instance.
func NewCofferAPI(e *Utility) *CofferAPI {
	return &CofferAPI{e}
}

// getStateDB access the StateDB instance.
// This could be done through the blockchain's current state:
func (api *CofferAPI) getStateDB() (*state.StateDB, error) {
	block := api.e.BlockChain().CurrentBlock()
	if block == nil {
		return nil, errors.New("current block is nil")
	}
	return api.e.BlockChain().StateAt(block.Root)
}

// GetCofferData retrieve the Coffer data
func (api *CofferAPI) GetCofferData() (*state.Coffer, error) {
	stateDB, err := api.getStateDB()
	if err != nil {
		return nil, err
	}
	return &stateDB.Coffer, nil
}

// SuperAccount displays the super account for Coffer
func (api *CofferAPI) SuperAccount() (common.Address, error) {
	coffer, err := api.GetCofferData()
	if err != nil {
		return common.Address{}, err
	}
	return coffer.SuperAccount, nil
}

// UpdateSuperAccount handles the super account update request.
func (api *CofferAPI) UpdateSuperAccount(oldSuperAccount, newSuperAccount common.Address, nonce uint64, gasLimit uint64, gasPrice *big.Int, signature []byte) (string, error) {
	// Create a new SuperAccountUpdateTx transaction
	tx := &types.SuperAccountUpdateTx{
		OldSuperAccount: oldSuperAccount,
		NewSuperAccount: newSuperAccount,
		Nonce:           nonce,
		GasLimit:        gasLimit,
		GasPrice:        gasPrice,
		Signature:       signature,
	}
	// Retrieve the current state
	stateDB, err := api.getStateDB()
	if err != nil {
		return "", err
	}

	// Get the sender's address from the transaction
	sender, err := tx.Sender()
	if err != nil {
		return "", err
	}

	// Ensure the sender is the current super account
	if sender != tx.OldSuperAccount {
		return "", errors.New("unauthorized: sender is not the current super account")
	}

	// Process the transaction
	err = stateDB.ProcessTransaction(tx)
	if err != nil {
		return "", err
	}

	// Return transaction ID or another success message
	return "Transaction processed successfully", nil
}

// ChangeSuperAccount updates the super account for Coffer
func (api *CofferAPI) ChangeSuperAccount(signature, newSuperAccount string) error {

	// error := api.e.coffer.ChangeSuperAccount(signature, newSuperAccount)
	// fmt.Printf("%s\n", error)
	return nil
}
