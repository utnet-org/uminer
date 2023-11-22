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
	"fmt"

	"github.com/yanhuangpai/go-utility/common"
)

// CofferAPI provides an API to control the coffer.
type CofferAPI struct {
	e *Utility
}

// NewCofferAPI create a new CofferAPI instance.
func NewCofferAPI(e *Utility) *CofferAPI {
	return &CofferAPI{e}
}

func (api *CofferAPI) SuperAccount() common.Address {
	superAccount := common.HexToAddress("0x2116825a1f6De9C479f8BC36d4a0F32074182924")
	fmt.Printf("%s\n", superAccount)
	return superAccount
}
