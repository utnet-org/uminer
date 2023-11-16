// Copyright 2015 The go-utility Authors
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
	"github.com/yanhuangpai/go-utility/common"
	"github.com/yanhuangpai/go-utility/common/hexutil"
)

// UnilityAPI provides an API to access Utility full node-related information.
type UnilityAPI struct {
	e *Utility
}

// NewUnilityAPI creates a new Utility protocol API for full nodes.
func NewUnilityAPI(e *Utility) *UnilityAPI {
	return &UnilityAPI{e}
}

// Unicrpytbase is the address that mining rewards will be sent to.
func (api *UnilityAPI) Unicrpytbase() (common.Address, error) {
	return api.e.Unicrpytbase()
}

// Coinbase is the address that mining rewards will be sent to (alias for Unicrpytbase).
func (api *UnilityAPI) Coinbase() (common.Address, error) {
	return api.Unicrpytbase()
}

// Hashrate returns the POW hashrate.
func (api *UnilityAPI) Hashrate() hexutil.Uint64 {
	return hexutil.Uint64(api.e.Miner().Hashrate())
}

// Mining returns an indication if this node is currently mining.
func (api *UnilityAPI) Mining() bool {
	return api.e.IsMining()
}
