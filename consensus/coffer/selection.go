package coffer

import "math/rand"

// SelectSigner selects a signer from a slice of signers based on their power.
func SelectSigner(signers []Signer) *Signer {
	totalPower := 0
	for _, signer := range signers {
		totalPower += signer.Power
	}

	if totalPower == 0 {
		return nil // No signers or all signers have zero power
	}

	randPoint := rand.Intn(totalPower)
	for _, signer := range signers {
		randPoint -= signer.Power
		if randPoint <= 0 {
			return &signer
		}
	}
	return nil // Fallback, should not happen
}
