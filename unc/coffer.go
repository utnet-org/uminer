package unc

import "github.com/yanhuangpai/go-utility/consensus/coffer"

type CofferAPI struct {
	coffer *coffer.Coffer
}

func NewCofferAPI(coffer *coffer.Coffer) *CofferAPI {
	return &CofferAPI{coffer: coffer}
}

func (api *CofferAPI) GetCurrentSigner() (*coffer.Signer, error) {
	return api.coffer.GetCurrentSigner()
}

func (api *CofferAPI) AddNewSigner(newSigner *coffer.Signer) error {
	return api.coffer.AddNewSigner(newSigner)
}
