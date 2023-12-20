package main

//#cgo CXXFLAGS: -std=c++11
//#cgo CFLAGS: -I../../bm_chip/src
//#cgo LDFLAGS: -L../../bm_chip/src -lchip -lstdc++
//#cgo LDFLAGS: -L/usr/local/opt/openssl/lib -lcrypto
//#include <chip.h>
//#include <openssl/ec.h>
//#include <openssl/ecdsa.h>
//#include <openssl/obj_mac.h>
//
import "C"
import (
	"encoding/hex"
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
func BurnChips(SerialNumber string, busId string, chipId int) bool {
	res := C.chipBurning(C.int(chipId))

	fmt.Println("result  ", res)

	return true

}

// ReadChipKeyPairs is
func ReadChipKeyPairs(SerialNumber string, busId string, chipId int) ChipKeyPairs {
	res := C.readPPubkey(C.int(chipId))

	// Convert the C array to a Go slice for easier handling
	chipArray := (*[1 << 30]C.struct_ChipDeclaration)(unsafe.Pointer(&res))[:1:1]
	chip := chipArray[0]

	keyPairs := ChipKeyPairs{
		SerialNumber: SerialNumber,
		BusId:        busId,
		P2:           C.GoString((*C.char)(unsafe.Pointer(chip.EncryptedPrivK))),
		PubKey:       C.GoString((*C.char)(unsafe.Pointer(chip.PubK))),
	}
	p2 := hex.EncodeToString([]byte(keyPairs.P2))
	decodeP2, _ := hex.DecodeString(p2)
	fmt.Printf("P2 encode: %s\n", keyPairs.P2)
	fmt.Printf("P2 %d: %s\n", chipId, p2)
	fmt.Printf("PubKey %d: %s\n", chipId, keyPairs.PubKey)
	fmt.Printf("P2 decode: %s\n", decodeP2)

	return keyPairs

}
