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

ChipSignature* chipSignature(unsigned long chipId, const char* P2Char, const char* PubKeyChar, const char* message, unsigned int size_p2, unsigned int  size_pubkey) {

    /*** sign by p2 with decryption of secret key ***/
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
    printf("p2 = %s\n", p2);
    unsigned char* pubkey = (unsigned char *)malloc(size_pubkey);
    strcpy(reinterpret_cast<char*>(pubkey), const_cast<char*>(PubKeyChar));

    /* produce signature */
    ret = bmcpu_gen_sign(handle, signature, &size_signature,
                         digest,sizeof(digest), p2, size_p2);
    if (ret != BM_SUCCESS) {
        printf("bmcpu_gen_sign error, ret = %d\n", ret);
        bm_dev_free(handle);
        ChipSignature* emptys = new ChipSignature[1];
        emptys->SignMsg = nullptr;
        emptys->PubK = nullptr;
        emptys->status = -1;
        return emptys;
    }
    // hex transform the signature binary bytes
    std::string signatureHex = byteArrayToHex(signature, size_signature);
    printf("signature size: %d\n",size_signature);
    printf("signature: %s\n",signature);
    std::cout << signatureHex << std::endl;

    /* get values */
    signatureList[0].SignMsg = new unsigned char[size_pubkey*2];  // multiply 2 here is to get rid of the wrong value output at cgo
    signatureList[0].PubK = new unsigned char[size_pubkey];
//    memcpy(signatureList[0].SignMsg, signature, size_pubkey);
    strcpy(reinterpret_cast<char*>(signatureList[0].SignMsg), const_cast<char*>(signatureHex.c_str()));
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