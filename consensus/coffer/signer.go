package coffer

import (
	"errors"

	"github.com/yanhuangpai/go-utility/common" // Path to the package where Block is defined
)

// Signer represents a participant in the Coffer consensus mechanism.
type Signer struct {
	Address1 common.Address // Original public key for verification
	Address2 common.Address // Secondary public key for signing blocks
	Power    int            // Power level assigned by the super account
	// Other fields...
}

// NewSigner creates a new signer. Can only be called by the super account.
func (c *Coffer) NewSigner(signatureHex, address1 string, power int) (Signer, error) {
	// Verify the super account
	result, _ := verifySig(c.SuperAccount.String(), address1, signatureHex)
	if !result {

		emptySigner := &Signer{
			Address1: common.HexToAddress(address1),
			Power:    0,
		}
		return *emptySigner, errors.New("only the current super account can update to a new super account")
	}
	// Proceed to create a new Signer
	signer := &Signer{
		Address1: common.HexToAddress(address1),
		Power:    power,
	}

	return *signer, nil
}

func (s *Signer) ActivateSigner(callerAddress, address2 common.Address) error {
	if callerAddress != s.Address1 {
		return errors.New("unauthorized caller")
	}
	s.Address2 = address2
	return nil
}

// UpdatePower updates the power of a signer.
func (s *Signer) UpdatePower(newPower int) {
	s.Power = newPower
}

// Example function to get the current signer
func (c *Coffer) GetCurrentSigner() (*Signer, error) {
	// Logic to return the current signer
	return nil, nil
}

// Example function to add a new signer (only callable by the super account)
func (c *Coffer) AddNewSigner(newSigner *Signer) error {
	// Logic to add a new signer
	return nil
}

// Other signer-related functions...
