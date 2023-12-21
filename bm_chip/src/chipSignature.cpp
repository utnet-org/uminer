//#include <openssl/ec.h>
//#include <openssl/ecdsa.h>
//#include <openssl/sha.h>
//#include <openssl/obj_mac.h>

#include <stdlib.h>
#include <stdio.h>
#include <cstring>
#include <string>
#include <iostream>
#include <sstream>
#include <string.h>
#include <pthread.h>
#include <bmlib_runtime.h>

#include <openssl/rsa.h>
#include <openssl/pem.h>
#include <openssl/evp.h>
#include <openssl/pem.h>
#include <openssl/bio.h>
#include <openssl/buffer.h>

#include "chip.h"
#include "util.h"

#ifdef __linux__
#include <sys/syscall.h>
#else
#include <windows.h>
#endif

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

ChipSignature* chipSignature(unsigned long chipId, const char* P2Char, const char* PubKeyChar, const char* message, unsigned int size_p2, unsigned int  size_pubkey) {

#if !defined(USING_CMODEL) && !defined(SOC_MODE)
    bm_handle_t handle;
    bm_status_t ret = BM_SUCCESS;
    unsigned int size_signature;

    ret = bm_dev_request(&handle, chipId);
    if ((ret != BM_SUCCESS) || (handle == NULL)) {
        printf("bm_dev_request error, ret = %d\n", ret);
        ChipSignature* emptys = new ChipSignature[1];
        emptys->SignMsg = nullptr;
        emptys->PubK = nullptr;
        emptys->status = 0;
        return emptys;
    }

    unsigned char* signature = (unsigned char*)malloc(size_pubkey);
    ChipSignature* signatureList = new ChipSignature[1];
    /* Generate digest*/
    const char* data = message;
    size_t data_size = sizeof(data) - 1;
    unsigned char digest[SHA256_DIGEST_LENGTH];
    generate_sha256_digest(reinterpret_cast<const unsigned char*>(data), data_size, digest);

    /* digital signature */
    unsigned char* p2 = (unsigned char *)malloc(size_p2);
    std:: string str2(P2Char);
    p2 = hexToByteArray(str2);
    unsigned char* pubkey = (unsigned char *)malloc(size_pubkey);
    strcpy(reinterpret_cast<char*>(pubkey), const_cast<char*>(PubKeyChar));

    ret = bmcpu_gen_sign(handle, signature, &size_signature,
                         digest,sizeof(digest), p2, size_p2);
    if (ret != BM_SUCCESS) {
        printf("bmcpu_gen_sign error, ret = %d\n", ret);
        bm_dev_free(handle);
        ChipSignature* emptys = new ChipSignature[1];
        emptys->SignMsg = nullptr;
        emptys->PubK = nullptr;
        emptys->status = 0;
        return emptys;
    }
    // hex
    std::string signatureHex = byteArrayToHex(signature, size_signature);
    printf("signature size: %d\n",size_signature);
    printf("signature: %s\n",signature);
    std::cout << signatureHex << std::endl;

    /* get values */
    signatureList[0].SignMsg = new unsigned char[size_pubkey];
    signatureList[0].PubK = new unsigned char[size_pubkey];
    memcpy(signatureList[0].SignMsg, signature, size_pubkey);
    strcpy(reinterpret_cast<char*>(signatureList[0].PubK), const_cast<char*>(PubKeyChar));

    /* Verify digital signature */
    int verify_result = verify_with_public_key (pubkey, size_pubkey,
                                                digest, SHA256_DIGEST_LENGTH,
                                                signature, size_signature);
    if (verify_result == 1) {
        printf("sign verify success\n");
        signatureList[0].status = 1;
    } else {
        printf("sign verify fail\n");
        signatureList[0].status = 0;
    }

    free(p2);
    free(pubkey);
    free(signature);
    bm_dev_free(handle);
#else
    printf("This test case is only valid in PCIe mode!\n");
#endif
    return signatureList;

}