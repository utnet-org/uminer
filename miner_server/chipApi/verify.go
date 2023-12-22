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
	"fmt"
)

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

//func VerifyChipsSignature(signature string, publicKey string, message string) bool {
//	// 解析公钥
//	block, _ := pem.Decode([]byte(publicKey))
//	if block == nil {
//		fmt.Println("failed to parse PEM block containing the public key")
//		return false
//	}
//
//	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
//	if err != nil {
//		fmt.Println("failed to parse DER encoded public key: ", err)
//		return false
//	}
//
//	// 验证签名
//	data := []byte(message)
//	digest := sha256.Sum256(data)
//	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, digest[:], []byte(signature))
//	if err != nil {
//		fmt.Println("signature verification failed: ", err)
//		return false
//	}
//
//	fmt.Println("Signature verified successfully")
//	return true
//}
