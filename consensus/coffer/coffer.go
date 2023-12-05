package coffer

import "github.com/yanhuangpai/go-utility/core"

// Coffer is the main structure for the Coffer consensus mechanism.
type CofferEngine struct {
	blockchain *core.BlockChain
	// Other fields as required
}

func New(blockchainInstance *core.BlockChain) *CofferEngine {
	return &CofferEngine{
		blockchain: blockchainInstance,
		// Set other necessary fields
	}
}
