package chipApi

//#cgo CXXFLAGS: -std=c++11
//#cgo CFLAGS: -I../../../bm_chip/src
//#cgo LDFLAGS: -L../../../bm_chip/src -lchip -lstdc++
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

// ChipSign struct of chip signature
type ChipSign struct {
	Signature string
	Status    bool
}

// SignMinerChips is to sign message at the chip with p2
func SignMinerChips(devId int, p2 string, pubKey string, p2Size int, pubKeySize int, message string) ChipSign {

	cP2 := C.CString(p2)
	cpubKey := C.CString(pubKey)
	cMessage := C.CString(message)
	res := C.chipSignature(C.ulong(devId), cP2, cpubKey, cMessage, C.uint(p2Size), C.uint(pubKeySize))

	// Convert the C array to a Go slice for easier handling
	signatures := (*[1 << 30]C.struct_ChipSignature)(unsafe.Pointer(res))[:1:1]

	// get final ChipSignature result
	fmt.Printf("Signature of chip %d: %s\n", devId, C.GoString((*C.char)(unsafe.Pointer(signatures[0].SignMsg))))
	fmt.Printf("PubKey of chip %d: %s\n", devId, C.GoString((*C.char)(unsafe.Pointer(signatures[0].PubK))))
	return ChipSign{
		Signature: C.GoString((*C.char)(unsafe.Pointer(signatures[0].SignMsg))),
		Status:    true,
	}
}

// the fake function for operating at miner server

//func SignMinerChips(devId int, p2 string, pubKey string, p2Size int, pubKeySize int, message string) ChipSign {
//	return ChipSign{
//		Signature: "signature",
//		Status:    true,
//	}
//}
