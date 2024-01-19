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
	"strconv"
	"unsafe"
)

// ChipSign struct of chip signature
type ChipSign struct {
	Signature string
	Status    bool
}

func SignMinerChips(SerialNumber string, busId string, devId int, p2 string, pubKey string, p2Size int, pubKeySize int, message string) ChipSign {
	// locate the chipId
	list := BMChipsInfos("../../api/chipApi/bm_smi_1.txt")
	chipId := -1
	for _, item := range list {
		if item.SerialNum == SerialNumber {
			if item.Chips[0].BusId == busId {
				devId, _ := strconv.ParseInt(item.Chips[0].DevId, 10, 64)
				chipId = int(devId)
			} else if item.Chips[1].BusId == busId {
				devId, _ := strconv.ParseInt(item.Chips[1].DevId, 10, 64)
				chipId = int(devId)
			} else if item.Chips[2].BusId == busId {
				devId, _ := strconv.ParseInt(item.Chips[2].DevId, 10, 64)
				chipId = int(devId)
			}
		}
	}
	//if chipId == -1 {
	//	return ChipSign{
	//		Signature: "",
	//		Status:    false,
	//	}
	//}

	chipId = devId

	cP2 := C.CString(p2)
	cpubKey := C.CString(pubKey)
	cMessage := C.CString(message)
	res := C.chipSignature(C.ulong(chipId), cP2, cpubKey, cMessage, C.uint(p2Size), C.uint(pubKeySize))

	// Convert the C array to a Go slice for easier handling
	signatures := (*[1 << 30]C.struct_ChipSignature)(unsafe.Pointer(res))[:1:1]

	// get final ChipSignature result
	fmt.Printf("Signature of chip %d: %s\n", chipId, C.GoString((*C.char)(unsafe.Pointer(signatures[0].SignMsg))))
	fmt.Printf("PubKey of chip %d: %s\n", chipId, C.GoString((*C.char)(unsafe.Pointer(signatures[0].PubK))))
	return ChipSign{
		Signature: C.GoString((*C.char)(unsafe.Pointer(signatures[0].SignMsg))),
		Status:    true,
	}
}
