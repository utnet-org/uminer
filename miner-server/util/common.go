package util

import (
	"compress/gzip"
	rand2 "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ed25519"
	"io/ioutil"
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

	return publicKeyBase58, privateKeyBase58
}

func MinerSignTx(privateKey string, msg string) string {

	return "sign"
}
