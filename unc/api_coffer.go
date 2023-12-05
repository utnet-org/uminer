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
	"fmt"
	"strconv"
	"strings"

	"github.com/yanhuangpai/go-utility/common"
	"github.com/yanhuangpai/go-utility/common/hexutil"
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
func (api *CofferAPI) UpdateSuperAccount(oldSuperAccount, newSuperAccount common.Address, nonceStr, gasLimitStr, gasPriceStr, signatureStr string) (string, error) {
	currentSuper, err := api.SuperAccount()
	if err != nil {
		return "", err
	}
	if strings.ToLower(oldSuperAccount.Hex()) != strings.ToLower(currentSuper.Hex()) {
		return "", errors.New("unauthorized: oldSuperAccount the current super account")
	}
	// Parse the nonce from hex string to uint64
	nonce, err := strconv.ParseUint(strings.TrimPrefix(nonceStr, "0x"), 16, 64)
	if err != nil {
		return "", fmt.Errorf("invalid nonce: %v", err)
	}

	// Parse the gasLimit from hex string to uint64
	gasLimit, err := strconv.ParseUint(strings.TrimPrefix(gasLimitStr, "0x"), 16, 64)
	if err != nil {
		return "", fmt.Errorf("invalid gas limit: %v", err)
	}

	// Parse the gasPrice from hex string to big.Int
	gasPrice, err := hexutil.DecodeBig(gasPriceStr)
	if err != nil {
		return "", fmt.Errorf("invalid gas price: %v", err)
	}

	// Decode the signature from hex string
	signature, err := hexutil.Decode(signatureStr)
	if err != nil {
		return "", fmt.Errorf("invalid signature: %v", err)
	}

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
