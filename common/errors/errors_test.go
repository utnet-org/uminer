package errors

import (
	"fmt"
	"testing"
)

func TestFromError(t *testing.T) {
	err := fmt.Errorf("%w", Errorf(nil, ErrorUnknown))
	fmt.Println(FromError(err))
}
