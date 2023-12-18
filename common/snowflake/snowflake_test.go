package snowflake

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNextUID(t *testing.T) {
	UID := NextUID()
	fmt.Println(UID)
	fmt.Println(strconv.FormatUint(UID, 10))
	fmt.Printf("%s", fmt.Sprintf("%d", UID))
}
