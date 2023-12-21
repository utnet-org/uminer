#include <stdio.h>
#include <stdlib.h>
#include <iostream>
#include <cstring>
#include <sstream>
#include <iomanip>
#include <openssl/ec.h>
#include <openssl/ecdsa.h>
#include <openssl/sha.h>
#include <openssl/obj_mac.h>

#include "chip.h"

std::string byteArrayToHex(const unsigned char* byteArray, size_t length) {
    std::stringstream ss;
    ss << std::hex << std::setfill('0');
    for (size_t i = 0; i < length; ++i)
        ss << std::setw(2) << static_cast<unsigned int>(byteArray[i]);
    return ss.str();
}

void BurnAtSPACCDemo(int burnTimes, int seq, unsigned char*  p, unsigned char*  pubkey) {

    std::string str1 = "";
    std::string str2 = "";
    if (burnTimes == 1) {
        // 第一次烧录
        printf("aes key programed success the 1st time!\n");
        // 模拟两个随意的string
        str1 = "chipP1_" + std::to_string(seq);
        str2 = "chipPubkey_" + std::to_string(seq);
    }
    if (burnTimes == 2) {
        // 第二次烧录
        printf("aes key programed success the 2nd time!\n");
        // 模拟两个随意的string
        str1 = "chipP2_" + std::to_string(seq);
        str2 = "chipPubkey_" + std::to_string(seq);
    }

    // 得到加密私钥和公钥并赋值
    strcpy(reinterpret_cast<char*>(p), str1.c_str());
    strcpy(reinterpret_cast<char*>(pubkey), str2.c_str());

}

int chipGenPPubkey(int seq) {

    // 从驱动烧录结束后 读取私钥再加密priK

    unsigned int size_p2 = 2048;
    unsigned int  size_pubkey = 2048;
    unsigned char* p2 = (unsigned char*)malloc(size_p2);
    unsigned char* pubkey = (unsigned char*)malloc(size_pubkey);

    ChipDeclaration oneChip;

    // get p2 and pubkey and store in files

    // 如果烧录出问题直接返回nullptr
    if (NULL == p2 ||  NULL == pubkey ){
        fprintf(stderr, "Error burning the chip %d\n", seq);
        return 0;
    }

    return 1;

}

ChipDeclaration readPPubkey(int seq) {
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
        oneChip.EncryptedPrivK = nullptr;
        oneChip.PubK = nullptr;
        return oneChip;
    }

    // hex
    std::string str = byteArrayToHex(p2, size_p2_padding);
    const char* P2Byte = str.c_str();

    // get value
    oneChip.EncryptedPrivK = P2Byte;
    oneChip.PubK = reinterpret_cast<const char*>(pubkey);

    free(p2);
    free(pubkey);

    return oneChip;

}

int chipBurning(int dev_id) {

    printf("\nthe chip dev_id %d to be burned ... \n", dev_id);

    /*** 芯片烧录 ***/

    return 1;

}