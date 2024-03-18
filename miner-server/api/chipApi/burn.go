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

// ChipKeyPairs struct of p2-pubKey pairs for a bmchip after burning and encrypting
type ChipKeyPairs struct {
	SerialNumber string
	BusId        string
	P2           string
	P2Size       int
	PubKey       string
	PubKeySize   int
}

// StartChips api through cgo to drive bmchip to activate cpu
func StartChips(chipId int, fipBin string, rambootRootfs string) bool {
	cfipBin := C.CString(fipBin)
	cRambootRootfs := C.CString(rambootRootfs)
	res := C.startCPU(C.int(chipId), cfipBin, cRambootRootfs)

	if res == 1 {
		fmt.Println("chip ", chipId, " activate success !")
		return true
	}
	fmt.Println("chip ", chipId, " activate failed !")
	return false

}

// BurnChips api through cgo to drive bmchip to burn and get p2-pubKey pairs
func BurnChips(SerialNumber string, busId string, chipId int) bool {
	res := C.chipBurning(C.int(chipId))

	if res == 1 {
		fmt.Println("chip ", chipId, " burned at efuse success !")
		return true
	}
	fmt.Println("chip ", chipId, " burned at efuse failed !")
	return false

}

// GenChipsKeyPairs is generating the key pairs from PKA after burning and restarting the machine, keys are stored in files
func GenChipsKeyPairs(SerialNumber string, busId string, chipId int) bool {
	res := C.chipGenKeyPairs(C.int(chipId))
	if res == 1 {
		fmt.Println("chip ", chipId, "generate p2 and pubKey success !")
		return true
	} else if res == 0 {
		fmt.Println("chip ", chipId, ": error opening file !")
		return false
	} else if res == -1 {
		fmt.Println("chip ", chipId, ": generate p2 and pubKey error !")
		return false
	} else if res == -3 {
		fmt.Println("chip ", chipId, ": bm_dev_request error !")
		return false
	}

	return false

}

// ReadChipKeyPairs is to read the keys from the stored files
func ReadChipKeyPairs(SerialNumber string, busId string, chipId int) ChipKeyPairs {
	res := C.readKeyPairs(C.int(chipId))

	// Convert the C array to a Go slice for easier handling
	chipArray := (*[1 << 30]C.struct_ChipDeclaration)(unsafe.Pointer(&res))[:1:1]
	chip := chipArray[0]

	keyPairs := ChipKeyPairs{
		SerialNumber: SerialNumber,
		BusId:        busId,
		P2:           C.GoString((*C.char)(unsafe.Pointer(chip.EncryptedPriK))),
		P2Size:       int(chip.EncryptedPriKSize),
		PubKey:       C.GoString((*C.char)(unsafe.Pointer(chip.PubK))),
		PubKeySize:   int(chip.PubKSize),
	}

	fmt.Printf("P2 %d: %s, size is %d\n", chipId, keyPairs.P2, keyPairs.P2Size)
	fmt.Printf("PubKey %d: %s, size is %d\n", chipId, keyPairs.PubKey, keyPairs.PubKeySize)

	return keyPairs

}

// the fake function for operating at miner server

//func StartChips(chipId int, fipBin string, rambootRootfs string) bool {
//	return false
//
//}
//func BurnChips(SerialNumber string, busId string, chipId int) bool {
//	return false
//
//}
//func GenChipsKeyPairs(SerialNumber string, busId string, chipId int) bool {
//	return false
//
//}
//func ReadChipKeyPairs(SerialNumber string, busId string, chipId int) ChipKeyPairs {
//	keyPairs := ChipKeyPairs{
//		SerialNumber: SerialNumber,
//		BusId:        busId,
//		P2:           "P2",
//		P2Size:       int(1),
//		PubKey:       "PubKey",
//		PubKeySize:   int(1),
//	}
//	return keyPairs
//}
