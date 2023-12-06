package chipApi

import (
	"testing"
)

func TestChip(t *testing.T) {

	BMChipsInfos()

	SignChips(1, 2, "p2", "utility")

}
