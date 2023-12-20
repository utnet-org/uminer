//#include <stdio.h>
//#include <iostream>
//#include <cstring>
//#include <openssl/ec.h>
//#include <openssl/ecdsa.h>
//#include <openssl/sha.h>
//#include <openssl/obj_mac.h>

#include <stdlib.h>
#include <stdio.h>
#include <cstring>
#include <string>
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

#ifdef __linux__
#include <sys/syscall.h>
#else
#include <windows.h>
#endif

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

ChipSignature* chipSignature(unsigned long chipId, unsigned char* P2Char, const char* PubKeyChar, const char* message, unsigned int size_p2, unsigned int  size_pubkey) {

#if !defined(USING_CMODEL) && !defined(SOC_MODE)
    bm_handle_t handle;
    bm_status_t ret = BM_SUCCESS;
    unsigned int size_signature;
    unsigned int size_p2_padding;

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
    unsigned char* pubkey = (unsigned char *)malloc(size_pubkey);
//    strcpy(reinterpret_cast<char*>(p2), const_cast<char*>(P2Char));
    strcpy(reinterpret_cast<char*>(pubkey), const_cast<char*>(PubKeyChar));

//        FILE *file_p2 = fopen("../key/p2_10", "r");
//        unsigned char* p21 = (unsigned char *)malloc(size_p2);
//        if (file_p2) {
//            fseek(file_p2, 0, SEEK_END);
//            size_p2_padding = ftell(file_p2);
//            fseek(file_p2, 0, SEEK_SET);
//
//            fread(p21, 1, size_p2_padding, file_p2);
//
//            fclose(file_p2);
//
//        } else {
//            printf("Error opening file.\n");
//        }

    ret = bmcpu_gen_sign(handle, signature, &size_signature,
                         digest,sizeof(digest),P2Char,size_p2);
    if (ret != BM_SUCCESS) {
        printf("bmcpu_gen_sign error, ret = %d\n", ret);
        bm_dev_free(handle);
        ChipSignature* emptys = new ChipSignature[1];
        emptys->SignMsg = nullptr;
        emptys->PubK = nullptr;
        emptys->status = 0;
        return emptys;
    }
    printf("signature size :%d\n",size_signature);
    printf("signature :%s\n",signature);

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

    free(signature);
    bm_dev_free(handle);
#else
    printf("This test case is only valid in PCIe mode!\n");
#endif
    return signatureList;

}