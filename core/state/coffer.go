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

import "github.com/yanhuangpai/go-utility/common"

type Coffer struct {
	superAccount common.Address
	signers      []Signer                      // Array for ordered storage
	signerMap    map[common.Address]SignerInfo // Mapping for fast access
	totalPower   uint64                        // Total cumulative power for random selection
}

type Signer struct {
	address1 common.Address // Initial address set by the super account
	address2 common.Address // Address set by the owner of the chipset
	power    uint64         // Weight or power of the signer
}

type SignerInfo struct {
	index int    // Index in the Signers array
	power uint64 // Weight or power of the signer
}

func (s *StateDB) AddCofferSigner(addr1, addr2 common.Address, power uint64) {
	// Create a new Signer
	newSigner := Signer{
		address1: addr1,
		address2: addr2,
		power:    power,
	}

	// Append the new signer to the Signers slice
	s.coffer.signers = append(s.coffer.signers, newSigner)

	// Create a SignerInfo and add it to the SignerMap
	signerInfo := SignerInfo{
		index: len(s.coffer.signers) - 1, // Index of the new signer in the Signers slice
		power: power,
	}
	s.coffer.signerMap[addr1] = signerInfo

	// Update the total power in the Coffer
	s.coffer.totalPower += power
}
