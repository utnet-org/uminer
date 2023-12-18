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

func SignChips(segmentStart uint64, segmentEnd uint64, p2 string, message string) {
	cP2 := C.CString(p2)
	cMessage := C.CString(message)
	res := C.chipSignature(C.ulong(segmentStart), C.ulong(segmentEnd), cP2, cMessage)
	//fmt.Println(cMessage, *res) // Custom print logic for ECDSA_SIG

	// Accessing the returned array
	numSignatures := segmentEnd - segmentStart // Replace with the actual number of signatures

	// Convert the C array to a Go slice for easier handling
	signatures := (*[1 << 30]C.struct_ChipSignature)(unsafe.Pointer(res))[:numSignatures:numSignatures]

	// Print the content of each ChipSignature
	for i, signature := range signatures {
		fmt.Printf("Signature %d: %s\n", i, C.GoString((*C.char)(unsafe.Pointer(signature.SignMsg))))
		fmt.Printf("PubKey %d: %s\n", i, C.GoString((*C.char)(unsafe.Pointer(signature.PubK))))
	}

}
