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

func VerifyMinerChips(segmentStart uint64, segmentEnd uint64, p2 string, pubK string, message string) bool {
	cP2 := C.CString(p2)
	cPubK := C.CString(pubK)
	cMessage := C.CString(message)
	res := C.chipVerify(C.ulong(segmentStart), C.ulong(segmentEnd), cP2, cPubK, cMessage)

	numSignatures := segmentEnd - segmentStart
	chips := (*[1 << 30]C.struct_ChipVerify)(unsafe.Pointer(res))[:numSignatures:numSignatures]
	for i, chip := range chips {
		fmt.Printf("Signature %d: %s\n", i, C.GoString((*C.char)(unsafe.Pointer(chip.SignMsg))))
		if chip.ifVerifyPass == 0 {
			return false
		}

	}
	return true

}
