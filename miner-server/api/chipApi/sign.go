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

// ChipSign the structure of chip signature
type ChipSign struct {
	Signature string
	Status    bool
}

// SignMinerChips the api through cgo to sign a message at the chip with p2 key
func SignMinerChips(devId int, p2 string, pubKey string, p2Size int, pubKeySize int, message string) ChipSign {

	cP2 := C.CString(p2)
	cpubKey := C.CString(pubKey)
	cMessage := C.CString(message)
	// cgo chipSignature func handler at c++ driver: sign the chip along with message and p2 key
	res := C.chipSignature(C.ulong(devId), cP2, cpubKey, cMessage, C.uint(p2Size), C.uint(pubKeySize))

	// convert the array from c++ to the array by golang slice for easier handling
	signatures := (*[1 << 30]C.struct_ChipSignature)(unsafe.Pointer(res))[:1:1]

	// get final chip signature result
	fmt.Printf("Signature of chip %d: %s\n", devId, C.GoString((*C.char)(unsafe.Pointer(signatures[0].SignMsg))))
	fmt.Printf("PubKey of chip %d: %s\n", devId, C.GoString((*C.char)(unsafe.Pointer(signatures[0].PubK))))
	if signatures[0].status == "-1" {
		fmt.Printf("signature of chip %d: fails due to wrong p2 key or internal chip error", devId)
		return ChipSign{
			Signature: "",
			Status:    false,
		}
	}
	if signatures[0].status == "0" {
		fmt.Printf("signature of chip %d: fails due to failure of verification of public key", devId)
		return ChipSign{
			Signature: C.GoString((*C.char)(unsafe.Pointer(signatures[0].SignMsg))),
			Status:    false,
		}
	}
	return ChipSign{
		Signature: C.GoString((*C.char)(unsafe.Pointer(signatures[0].SignMsg))),
		Status:    true,
	}
}

// the fake function for operating at miner server when lack of chips and driver

//func SignMinerChips(devId int, p2 string, pubKey string, p2Size int, pubKeySize int, message string) ChipSign {
//	return ChipSign{
//		Signature: "signature",
//		Status:    true,
//	}
//}
