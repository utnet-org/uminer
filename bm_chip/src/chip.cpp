#include <stdlib.h>
#include <stdio.h>
#include <string>
#include <sstream>
#include <string.h>
#include <iostream>
#include <iomanip>

#include <openssl/ec.h>
#include <openssl/ecdsa.h>
#include <openssl/sha.h>
#include <openssl/obj_mac.h>

#include <openssl/rsa.h>
#include <openssl/pem.h>
#include <openssl/evp.h>
#include <openssl/pem.h>
#include <openssl/bn.h>
#include <openssl/bio.h>
#include <openssl/buffer.h>

#include "util.h"

keyPairs getKeyPairs(){
    keyPairs a;
    a.Eckey = NULL;
    a.PubK = NULL;
    // 选择 ECC 曲线（这里使用 secp256k1，比特币中使用的曲线）
    EC_KEY *eckey = EC_KEY_new_by_curve_name(NID_secp256k1);

    if (!eckey) {
        fprintf(stderr, "Error creating ECC key\n");
        return a;
    }
    // 生成 ECC 公钥和私钥
    if (EC_KEY_generate_key(eckey) != 1) {
        fprintf(stderr, "Error generating ECC key\n");
        return a;
    }

    // 获取 ECC 公钥
    const EC_POINT *pub_key = EC_KEY_get0_public_key(eckey);
    char *pub_key_hex = EC_POINT_point2hex(EC_KEY_get0_group(eckey), pub_key, POINT_CONVERSION_UNCOMPRESSED, NULL);
    a.Eckey = eckey;
    a.PubK = pub_key;
    return a;
}

/* convert byteArrays to hex */
std::string byteArrayToHex(const unsigned char* byteArray, size_t length) {
    std::stringstream ss;
    ss << std::hex << std::setfill('0');
    for (size_t i = 0; i < length; ++i)
        ss << std::setw(2) << static_cast<unsigned int>(byteArray[i]);
    return ss.str();
}
/* convert hex to byteArray */
unsigned char* hexToByteArray(const std::string& hexString) {
    size_t length = hexString.length() / 2;
    printf("hexString length: %zu\n", hexString.length());
    unsigned char* byteArray = new unsigned char[length];

    for (size_t i = 0; i < hexString.length(); i += 2) {
        std::string byteString = hexString.substr(i, 2);
        unsigned char byte = static_cast<unsigned char>(std::stoi(byteString, nullptr, 16));
        byteArray[i / 2] = byte; // 每两个字符转换为一个字节，存储到数组中
    }

    return byteArray;
}

/* base58 encode */
std::string Base58Encode(const unsigned char* input, size_t length) {
    // Convert the input binary data to OpenSSL BIGNUM
    BIGNUM* bn = BN_bin2bn(input, length, nullptr);
    if (!bn) {
        std::cout << "Error: Unable to create BIGNUM\n";
        return "";
    }

    // Create a BIO for Base58 encoding
    BIO* b58 = BIO_new(BIO_f_base64());
    BIO_set_flags(b58, BIO_FLAGS_BASE64_NO_NL);

    // Write BIGNUM to BIO for encoding
    BIO* mem = BIO_new(BIO_s_mem());
    BIO_push(b58, mem);
    if (!BN_print(b58, bn)) {
        std::cout << "Error: Unable to encode BIGNUM\n";
        BIO_free_all(b58);
        BN_free(bn);
        return "";
    }

    // Read Base58 encoded data from BIO
    BUF_MEM* bptr;
    BIO_get_mem_ptr(mem, &bptr);

    // Convert BIO data to string
    std::string encoded(reinterpret_cast<char*>(bptr->data), bptr->length);

    // Clean up
    BIO_free_all(b58);
    BN_free(bn);

    return encoded;
}
/* base58 decode to byteArray */
unsigned char* Base58Decode(const std::string& encoded){
    // Create a BIO for Base58 decoding
    BIO* b58 = BIO_new(BIO_f_base64());
    BIO_set_flags(b58, BIO_FLAGS_BASE64_NO_NL);

    // Write Base58 encoded data to BIO for decoding
    BIO* mem = BIO_new_mem_buf(encoded.c_str(), encoded.length());
    BIO_push(b58, mem);

    // Read decoded BIGNUM from BIO
    BIGNUM* bn = BN_new();
//    if (!BN_read(b58, bn)) {
//        std::cout << "Error: Unable to decode Base58\n";
//        BIO_free_all(b58);
//        BN_free(bn);
//        return nullptr;
//    }

    // Convert BIGNUM to binary data
    size_t length = BN_num_bytes(bn);
    unsigned char* data = (unsigned char*)malloc(length);
    if (!data) {
        std::cout << "Error: Memory allocation failed\n";
        BIO_free_all(b58);
        BN_free(bn);
        return nullptr;
    }
    BN_bn2bin(bn, data);

    // Clean up
    BIO_free_all(b58);
    BN_free(bn);

    return data;
}

/* Generate the RSA256 digest */
void generate_sha256_digest(const unsigned char *data, size_t data_size, unsigned char *digest) {
    SHA256_CTX sha256_ctx;
    SHA256_Init(&sha256_ctx);
    SHA256_Update(&sha256_ctx, data, data_size);
    SHA256_Final(digest, &sha256_ctx);
}

/* Verify digital signature */
int verify_with_public_key(const unsigned char *publicKeyData, unsigned int pub_key_length,
                           const unsigned char *digest, unsigned int digest_size,
                           const unsigned char *signature, unsigned int signature_length) {
    BIO *public_key_bio = BIO_new_mem_buf((void *)publicKeyData, pub_key_length);
    if (!public_key_bio) {
        perror("Error creating BIO for public key");
        return -1;
    }
    RSA *public_key = PEM_read_bio_RSAPublicKey(public_key_bio, NULL, NULL, NULL);
    if (!public_key) {
        perror("Error reading public key");
        BIO_free(public_key_bio);
        return -1;
    }

    BIO_free(public_key_bio);

    int result = RSA_verify(NID_sha256, digest, digest_size, signature, signature_length, public_key);

    RSA_free(public_key);

    return result;
}