#include <stdio.h>
#include <stdlib.h>
#include <iostream>
#include <cstring>
#include <sstream>
#include <iomanip>

#include <openssl/rsa.h>
#include <openssl/pem.h>
#include <openssl/evp.h>
#include <openssl/pem.h>
#include <openssl/bio.h>
#include <openssl/buffer.h>
#include <bmlib_runtime.h>

#include "chip.h"
#include "util.h"

#ifdef __linux__
#include <sys/syscall.h>
#else
#include <windows.h>
#endif


int chipGenKeyPairs(int seq) {

    /*** read p2 and pubKey by SPACC and PKA ***/
#if !defined(USING_CMODEL) && !defined(SOC_MODE)
    bm_handle_t handle;
    bm_status_t ret = BM_SUCCESS;
    unsigned int size_p2 = 2048;
    unsigned int size_pubkey = 2048;
    unsigned int size_signature;

    ret = bm_dev_request(&handle, seq);
    if ((ret != BM_SUCCESS) || (handle == NULL)) {
        printf("bm_dev_request error, ret = %d\n", ret);
        return -3;
    }

    unsigned char* pubkey = (unsigned char*)malloc(size_pubkey);
    unsigned char* p2     = (unsigned char*)malloc(size_p2);

    /* get key pairs */
    ret = bmcpu_gen_p2_pubkey(handle, p2, &size_p2, pubkey, &size_pubkey);
    if (ret != BM_SUCCESS) {
        printf("bmcpu_gen_p2_pubkey error, ret = %d\n", ret);
        bm_dev_free(handle);
        return -1;
    }
//    strcpy((char*)pubkey, "hello");
//    strcpy((char*)p2, "world");


    /* Must be 16btye aligned storage, otherwise spacc cannot decrypt correctly*/
    unsigned int size_p2_padding = (size_p2 + 15) & ~15;
    printf("size_pubkey: %d size_p2: %d\n", size_pubkey, size_p2);
    printf("pubkey: %s  \n", pubkey);

    FILE *file_pubkey = fopen(("../../bm_chip/src/key/pubkey_"+ std::to_string(seq)).c_str(), "w");
    FILE *file_p2 = fopen(("../../bm_chip/src/key/p2_"+ std::to_string(seq)).c_str(), "w");

    // this directory is for c++ demo
//    FILE *file_pubkey = fopen(("../key/pubkey_"+ std::to_string(seq)).c_str(), "w");
//    FILE *file_p2 = fopen(("../key/p2_"+ std::to_string(seq)).c_str(), "w");

    if (file_pubkey) {
        size_t bytes_written = fwrite(pubkey, 1, size_pubkey, file_pubkey);
        if (bytes_written == size_pubkey) {
            printf("Data written to pubkey successfully.\n");
        } else {
            printf("Error writing data to pubkey\n");
        }
        fclose(file_pubkey);
    } else {
        printf("Error opening file.\n");
        return 0;
    }

    if (file_p2) {
        size_t bytes_written = fwrite(p2, 1, size_p2_padding, file_p2);
        if (bytes_written == size_p2_padding) {
            printf("Data written to p2 successfully.\n");
        } else {
            printf("Error writing data to p2.\n");
        }
        fclose(file_p2);
    } else {
        printf("Error opening file.\n");
        return 0;
    }

    free(pubkey);
    free(p2);
    bm_dev_free(handle);
#else

#endif
    return 1;

}

ChipDeclaration readKeyPairs(int seq) {

    unsigned int size_p2 = 2048;
    unsigned int size_pubkey = 2048;
    unsigned char* p2 = (unsigned char*)malloc(size_p2);
    unsigned char* pubkey = (unsigned char*)malloc(size_pubkey);
    unsigned int size_p2_padding;

    ChipDeclaration oneChip;

    // read files to get results
    FILE *file_pubkey = fopen(("../../bm_chip/src/key/pubkey_"+ std::to_string(seq)).c_str(), "r");
    FILE *file_p2 = fopen(("../../bm_chip/src/key/p2_"+ std::to_string(seq)).c_str(), "r");
    if (file_pubkey) {
        fseek(file_pubkey, 0, SEEK_END);
        size_pubkey = ftell(file_pubkey);
        fseek(file_pubkey, 0, SEEK_SET);

        fseek(file_p2, 0, SEEK_END);
        size_p2_padding = ftell(file_p2);
        fseek(file_p2, 0, SEEK_SET);

        fread(pubkey, 1, size_pubkey, file_pubkey);
        fread(p2, 1, size_p2_padding, file_p2);

        fclose(file_pubkey);
        fclose(file_p2);

//        printf("size_pubkey: %d size_p2 :%d size_p2_padding: %d \n", size_pubkey, size_p2, size_p2_padding);
//        printf("Pubkey:\n%s\n", pubkey);
    } else {
        printf("Error opening file.\n");
        oneChip.EncryptedPriK = nullptr;
        oneChip.PubK = nullptr;
        return oneChip;
    }

    // hex
    std::string str = byteArrayToHex(p2, size_p2_padding);
    const char* P2Byte = str.c_str();

    // get value
    oneChip.EncryptedPriK = P2Byte;
    oneChip.PubK = reinterpret_cast<const char*>(pubkey);
    oneChip.EncryptedPriKSize = size_p2_padding;
    oneChip.PubKSize = size_pubkey;

    free(p2);
    free(pubkey);

    return oneChip;

}

int chipBurning(int dev_id) {

    /*** burning at efuse ***/
#if !defined(USING_CMODEL) && !defined(SOC_MODE)
    printf("\nthe chip dev_id %d to be burned ... \n", dev_id);
    bm_handle_t handle;
    bm_status_t ret = BM_SUCCESS;

    ret = bm_dev_request(&handle, dev_id);
    if ((ret != BM_SUCCESS) || (handle == NULL)) {
        printf("bm_dev_request error, ret = %d\n", ret);
        return -3;
    }

    ret = bmcpu_gen_aes_key(handle);
    if (ret != BM_SUCCESS) {
        printf("burn_aes_key error, error = %d\n", ret);
        bm_dev_free(handle);
        return ret;
    }

    bm_dev_free(handle);
    printf("aes key programed success!\n");

#else

#endif
    return 1;

}