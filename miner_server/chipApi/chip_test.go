package chipApi

import (
	"fmt"
	"testing"
)

func TestChip(t *testing.T) {

	cardList := BMChipsInfos()
	fmt.Println(cardList)
	BurnChips("HQDZKC5BAAABJ0223", "000:d8:00:02", 6)
	SignMinerChips("HQDZKC5BBAJAH0146", "000:5e:00.2", "p2", "utility")

}
