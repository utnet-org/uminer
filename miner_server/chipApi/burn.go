package chipApi

//#cgo CXXFLAGS: -std=c++11
//#cgo CFLAGS: -I/Users/mac/Desktop/UtilityChain/ut_miner/src
//#cgo LDFLAGS: -L/Users/mac/Desktop/UtilityChain/ut_miner/src -lchip -lstdc++
//#cgo LDFLAGS: -L/usr/local/opt/openssl/lib -lcrypto
//#include <chip.h>
//#include <openssl/ec.h>
//#include <openssl/ecdsa.h>
//#include <openssl/obj_mac.h>
//
import "C"
import (
	"fmt"
	"unsafe"
)

// ChipKeyPairs struct of p2-pubKey pairs for a bmchip after burning and encrypting
type ChipKeyPairs struct {
	SerialNumber string
	BusId        string
	P2           string
	PubKey       string
}

// BurnChips api through cgo to drive bmchip to burn and get p2-pubKey pairs
func BurnChips(SerialNumber string, busId string, chipId int) {
	res := C.chipBurningStepOne(C.int(chipId))

	// Convert the C array to a Go slice for easier handling
	chipArray := (*[1 << 30]C.struct_ChipDeclarationOne)(unsafe.Pointer(&res))[:1:1]
	chip := chipArray[0]

	keyPairs := ChipKeyPairs{
		SerialNumber: SerialNumber,
		BusId:        busId,
		P2:           C.GoString((*C.char)(unsafe.Pointer(chip.EncryptedPrivK))),
		PubKey:       C.GoString((*C.char)(unsafe.Pointer(chip.PubK))),
	}
	fmt.Printf("P2 %d: %s\n", chipId, keyPairs.P2)
	fmt.Printf("PubKey %d: %s\n", chipId, keyPairs.PubKey)

	// Pack tx for chain upload

}
