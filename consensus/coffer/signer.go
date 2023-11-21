package coffer

import "errors"

var superAccountAddress string // Set this during initialization
// Signer represents a participant in the Coffer consensus mechanism.
type Signer struct {
	PublicKey1 string // Original public key for verification
	PublicKey2 string // Secondary public key for signing blocks
	Power      int    // Power level assigned by the super account
	// Other fields...
}

// NewSigner creates a new signer. Can only be called by the super account.
func NewSigner(callerAddress, publicKey string, power int) *Signer {
	if callerAddress != superAccountAddress {
		return nil // Or handle the error as per your design
	}
	// Proceed to create a new Signer
	return &Signer{
		PublicKey1: publicKey,
		Power:      power,
	}
}

func (s *Signer) SetPublicKey2(callerKey, publicKey2 string) error {
	if callerKey != s.PublicKey1 {
		return errors.New("unauthorized caller")
	}
	s.PublicKey2 = publicKey2
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
