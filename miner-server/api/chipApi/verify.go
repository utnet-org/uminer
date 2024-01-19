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
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
)

// verify at c++ by the openssl library
func VerifyMinerChips(signature string, pubK string, signatureSize int, pubKSize int, message string) bool {
	cSignature := C.CString(signature)
	cPubK := C.CString(pubK)
	cMessage := C.CString(message)
	res := C.signatureVerify(cSignature, cPubK, C.uint(signatureSize), C.uint(pubKSize), cMessage)

	if res == 1 {
		fmt.Println("signature is verified with success !")
		return true
	}
	fmt.Println("signature is verified with failure !")
	return false
}

// verify at go by crypto library
func hexStringToBytes(hexString string) ([]byte, error) {
	signatureBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, fmt.Errorf("error decoding hex string: %v", err)
	}
	return signatureBytes, nil
}
func VerifyChipsSignature(signature string, publicKey string, message string) bool {
	// parse publicKey
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		fmt.Println("failed to parse PEM block containing the public key")
		return false
	}

	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		fmt.Println("failed to parse DER encoded public key: ", err)
		return false
	}

	// Verify sign
	data := []byte(message)
	digest := sha256.Sum256(data)
	signatureBytes, err := hexStringToBytes(signature)
	if err != nil {
		fmt.Println("hex signature Error:", err)
		return false
	}

	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, digest[:], signatureBytes)
	if err != nil {
		fmt.Println("signature verification failed: ", err)
		return false
	}

	fmt.Println("Signature verified successfully")
	return true
}
