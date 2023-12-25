package chipApi

import (
	"fmt"
	"testing"
)

func TestChip(t *testing.T) {

	cardList := BMChipsInfos()
	fmt.Println(cardList)
	BurnChips("HQDZKC5BAAABJ0223", "000:d8:00:02", 6)
	GenChipsKeyPairs("HQDZKC5BAAABJ0223", "000:d8:00:02", 10)
	ReadChipKeyPairs("HQDZKC5BAAABJ0223", "000:d8:00:02", 10)
	SignMinerChips("HQDZKC5BBAJAH0146", "000:5e:00.2", 10, "p2", "pubkey", 1680, 426, "utility")

}
