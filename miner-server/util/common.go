package util

import (
	"compress/gzip"
	rand2 "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ed25519"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

const alphanumerics = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 用 Gzip 压缩了，因此网关需判断头里是否有此 header 头，再用 compress/gzip 解压即可。
func GzipApi(res *http.Response) []byte {
	// 是否有 gzip
	gzipFlag := false
	for k, v := range res.Header {
		if strings.ToLower(k) == "content-encoding" && strings.ToLower(v[0]) == "gzip" {
			gzipFlag = true
		}
	}
	var content []byte
	if gzipFlag {
		// 创建 gzip.Reader
		gr, err := gzip.NewReader(res.Body)
		defer gr.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		content, _ = ioutil.ReadAll(gr)
	} else {
		content, _ = ioutil.ReadAll(res.Body)
	}
	return content
}

// RandomString generates a pseudo-random string of length n.
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphanumerics[rand.Int63()%int64(len(alphanumerics))]
	}
	return string(b)
}

// RSAKeysGeneration generates rsa2048 key pairs
func RSAKeysGeneration() (*rsa.PrivateKey, string) {
	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand2.Reader, 2048)
	if err != nil {
		fmt.Println("生成RSA密钥对失败:", err)
		return nil, ""
	}

	// 将私钥编码为PEM格式
	//privateKeyPEM := &pem.Block{
	//	Type:  "RSA PRIVATE KEY",
	//	Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	//}
	// 写入私钥到文件
	//privateFile, err := os.Create("private.pem")
	//if err != nil {
	//	fmt.Println("无法创建私钥文件:", err)
	//	return nil, ""
	//}
	//if err := pem.Encode(privateFile, privateKeyPEM); err != nil {
	//	fmt.Println("无法写入私钥文件:", err)
	//	return nil, ""
	//}

	// 生成公钥
	publicKey := privateKey.PublicKey
	// 将公钥编码为PEM格式
	publicKeyPEM, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		fmt.Println("公钥编码失败:", err)
		return nil, ""
	}

	// 写入公钥到文件
	//publicKeyBlock := &pem.Block{
	//	Type:  "PUBLIC KEY",
	//	Bytes: publicKeyPEM,
	//}
	//publicFile, err := os.Create("public.pem")
	//if err != nil {
	//	fmt.Println("无法创建公钥文件:", err)
	//	return nil, ""
	//}
	//if err := pem.Encode(publicFile, publicKeyBlock); err != nil {
	//	fmt.Println("无法写入公钥文件:", err)
	//	return nil, ""
	//}
	return privateKey, string(publicKeyPEM)
}

// ED25519KeysGeneration generates ed25519 key pairs
func ED25519KeysGeneration() (string, string) {
	// 生成ed25519密钥对
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		fmt.Println("生成密钥对失败:", err)
		return "", ""
	}

	// 将公钥转换为Base58编码
	publicKeyBase58 := base58.Encode(publicKey)
	// 将私钥转换为Base58编码
	privateKeyBase58 := base58.Encode(privateKey)

	// 将Base58编码的公钥解码为字节切片
	decodedPublicKey := base58.Decode(publicKeyBase58)

	// 将字节切片格式的公钥转换为十六进制格式
	publicKeyHex := hex.EncodeToString(decodedPublicKey)

	// 打印生成的密钥对及公钥的Base58编码和十六进制格式
	fmt.Println("私钥(Base58编码):", privateKeyBase58)
	fmt.Println("公钥(Base58编码):", publicKeyBase58)
	fmt.Println("公钥(十六进制):", publicKeyHex)

	// 假设要签名的消息是 "Hello, world!"
	message := []byte("Hello, world!")

	// 将私钥的 Base58 编码解码为字节切片
	decodedPrivateKey := base58.Decode(privateKeyBase58)
	// 使用私钥进行签名
	signature := ed25519.Sign(decodedPrivateKey, message)

	// 验证签名
	verified := ed25519.Verify(base58.Decode(publicKeyBase58), message, signature)
	fmt.Println("签名验证结果:", verified)

	return publicKeyBase58, privateKeyBase58
}

// ED25519AddressGeneration generates ed25519 key pairs with mnemonic,
func ED25519AddressGeneration(private string) (string, string, string) {

	var mnemonic string
	var privKey ed25519.PrivateKey
	var pubKey ed25519.PublicKey
	if private == "" {
		// generate keys
		entropy, err := bip39.NewEntropy(128)
		if err != nil {
			log.Fatal(err)
		}
		mnemonic, _ = bip39.NewMnemonic(entropy)
		seed := bip39.NewSeed("suit depth work bacon wine connect venue army blame better pause train", "")
		derivedPrivkey := ed25519.NewKeyFromSeed(seed[:32])
		privKey = derivedPrivkey
		pubKey = derivedPrivkey.Public().(ed25519.PublicKey)
	} else {
		// import the keys
		privKey = base58.Decode(private)
		pubKey = privKey.Public().(ed25519.PublicKey)
	}

	privateKey := base58.Encode(privKey)
	publicKey := base58.Encode(pubKey)
	address := hex.EncodeToString(pubKey)

	fmt.Println("address:", address)
	fmt.Println("publicKey:", publicKey)
	fmt.Println("privateKey:", privateKey)

	// 使用私钥进行签名
	signature := ed25519.Sign(privKey, []byte("message"))
	// 验证签名
	verified := ed25519.Verify(pubKey, []byte("message"), signature)
	fmt.Println("sign verify:", verified)

	return mnemonic, publicKey, privateKey

}
