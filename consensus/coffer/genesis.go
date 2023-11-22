package coffer

import (
	"errors"

	"github.com/yanhuangpai/go-utility/common"
	"github.com/yanhuangpai/go-utility/core"
)

func ExtractSuperAccountFromGenesis(genesis *core.Genesis) (common.Address, error) {
	if genesis == nil {
		return common.Address{}, errors.New("genesis block is nil")
	}

	// Define the offset where the address starts in ExtraData (32 bytes in this case)
	addressOffset := 32

	if len(genesis.ExtraData) < addressOffset+common.AddressLength {
		return common.Address{}, errors.New("extradata is too short to contain a super account address")
	}

	var addr common.Address
	copy(addr[:], genesis.ExtraData[addressOffset:addressOffset+common.AddressLength])
	return addr, nil
}
