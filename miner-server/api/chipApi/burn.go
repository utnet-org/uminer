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
	"encoding/hex"
	"fmt"
	"os"
)

// ChipKeyPairs the structure of p2Key-pubKey pairs for a bm-chip after burning and encryption
type ChipKeyPairs struct {
	SerialNumber string
	BusId        string
	P2           string
	P2Size       int
	PubKey       string
	PubKeySize   int
}

// StartChips the api through cgo to call the driver to activate bm-chip
func StartChips(chipId int, fipBin string, rambootRootfs string) bool {
	// files for activate A53 process of the chip
	cfipBin := C.CString(fipBin)
	cRambootRootfs := C.CString(rambootRootfs)
	// cgo startCPU func handler at c++ driver: activate the chip
	res := C.startCPU(C.int(chipId), cfipBin, cRambootRootfs)

	if res == 1 {
		fmt.Println("chip ", chipId, " activate success !")
		return true
	}
	fmt.Println("chip ", chipId, " activate failed !")
	return false

}

// BurnChips the api through cgo to drive bm-chip to burn secret key at EFUSE
func BurnChips(SerialNumber string, busId string, chipId int) bool {
	// cgo chipBurning func handler at c++ driver: burn secret key at EFUSE
	res := C.chipBurning(C.int(chipId))

	if res == 1 {
		fmt.Println("chip ", chipId, " burned at efuse success !")
		return true
	}
	fmt.Println("chip ", chipId, " burned at efuse failed !")
	return false

}

// GenChipsKeyPairs the api through cgo to generate the key pairs from PKA after burning and restarting the machine, keys are stored in files
func GenChipsKeyPairs(SerialNumber string, busId string, chipId int) bool {
	// cgo chipGenKeyPairs func handler at c++ driver: generate the p2 key and public key simultaneously
	res := C.chipGenKeyPairs(C.CString(SerialNumber), C.CString(busId), C.int(chipId))

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

// ReadChipKeyPairs the api through cgo to read the keys from the stored files
//func ReadChipKeyPairs(SerialNumber string, busId string, chipId int) ChipKeyPairs {
//	// cgo chipGenKeyPairs func handler at c++ driver: generate the p2 key and public key simultaneously
//	res := C.readKeyPairs(C.CString(SerialNumber), C.CString(busId), C.int(chipId))
//
//	// convert the array from c++ to the array by golang slice for easier handling
//	chipArray := (*[1 << 30]C.struct_ChipDeclaration)(unsafe.Pointer(&res))[:1:1]
//	chip := chipArray[0]
//
//	keyPairs := ChipKeyPairs{
//		SerialNumber: SerialNumber,
//		BusId:        busId,
//		P2:           C.GoString((*C.char)(unsafe.Pointer(chip.EncryptedPriK))),
//		P2Size:       int(chip.EncryptedPriKSize),
//		PubKey:       C.GoString((*C.char)(unsafe.Pointer(chip.PubK))),
//		PubKeySize:   int(chip.PubKSize),
//	}
//
//	fmt.Printf("P2 %d: %s, size is %d\n", chipId, keyPairs.P2, keyPairs.P2Size)
//	fmt.Printf("PubKey %d: %s, size is %d\n", chipId, keyPairs.PubKey, keyPairs.PubKeySize)
//
//	return keyPairs
//
//}

// ReadChipKeyPairs the api to read the keys from the stored files
func ReadChipKeyPairs(SerialNumber string, busId string, chipId int) ChipKeyPairs {
	sizeP2 := 1680
	sizePubKey := 2048
	p2 := make([]byte, sizeP2)
	pubKey := make([]byte, sizePubKey)
	var sizeP2Padding int

	keyPairs := ChipKeyPairs{}

	// read files to get keys results
	filename := fmt.Sprintf("%s_%s", SerialNumber, busId)
	filePubKey, err := os.Open(fmt.Sprintf("../../../bm_chip/src/key/pubkey_%s", filename))
	if err != nil {
		fmt.Println("Error opening pubkey file:", err)
		return keyPairs
	}
	fileP2, err := os.Open(fmt.Sprintf("../../../bm_chip/src/key/p2_%s", filename))
	if err != nil {
		fmt.Println("Error opening p2 file:", err)
		return keyPairs
	}

	// read
	sizeP2Padding, err = fileP2.Read(p2)
	if err != nil {
		fmt.Println("Error reading pubkey file:", err)
		return keyPairs
	}
	p2Hex := hex.EncodeToString(p2)
	sizePubKey, err = filePubKey.Read(pubKey)
	if err != nil {
		fmt.Println("Error reading pubkey file:", err)
		return keyPairs
	}

	keyPairs.P2 = p2Hex
	keyPairs.P2Size = sizeP2Padding
	keyPairs.PubKey = string(pubKey)
	keyPairs.PubKeySize = sizePubKey

	fmt.Printf("P2 %d: %s, size is %d\n", chipId, keyPairs.P2, keyPairs.P2Size)
	fmt.Printf("PubKey %d: %s, size is %d\n", chipId, keyPairs.PubKey, keyPairs.PubKeySize)

	return keyPairs

}

// the fake function for operating at miner server when lack of chips and driver

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
