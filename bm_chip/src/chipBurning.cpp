#include <stdio.h>
#include <iostream>
#include <cstring>
#include <openssl/ec.h>
#include <openssl/ecdsa.h>
#include <openssl/sha.h>
#include <openssl/obj_mac.h>

#include "chip.h"


void BurnAtSPACC(int burnTimes, int seq, unsigned char*  p, unsigned char*  pubkey) {

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

ChipDeclaration cpuGenP_Pubkey(int burnTimes, int seq) {

    unsigned int size_p = 2048;
    unsigned int  size_pubkey = 2048;
    unsigned char* p = (unsigned char*)malloc(size_p);
    unsigned char* pubkey = (unsigned char*)malloc(size_pubkey);

    ChipDeclaration oneChip;

    BurnAtSPACC(burnTimes, seq, p, pubkey);
    // 如果烧录出问题直接返回nullptr
    if (NULL == p ||  NULL == pubkey ){
        fprintf(stderr, "Error burning the chip %d\n", seq);
        oneChip.EncryptedPrivK = nullptr;
        oneChip.PubK = nullptr;
        return oneChip;
    }

    std::string P = std::string(reinterpret_cast<char*>(p));
    std::string PubKey = std::string(reinterpret_cast<char*>(pubkey));

    oneChip.EncryptedPrivK = new unsigned char[P.length() + 1];
    oneChip.PubK = new unsigned char[PubKey.length() + 1];

    strcpy(reinterpret_cast<char*>(oneChip.EncryptedPrivK), P.c_str());
    strcpy(reinterpret_cast<char*>(oneChip.PubK), PubKey.c_str());

    return oneChip;

}

ChipDeclaration chipBurning(int dev_id) {

    printf("\nthe chip dev_id %d to be burned by producer ... \n", dev_id);

    // 从驱动烧录私钥再加密priK
    /*** 芯片烧录 ***/
    ChipDeclaration chip = cpuGenP_Pubkey(1, dev_id);
    if (chip.EncryptedPrivK == nullptr || chip.PubK == nullptr) {
        fprintf(stderr, "Error burning the chip dev_id %d\n", dev_id);
        ChipDeclaration empty;
        empty.EncryptedPrivK = nullptr;
        empty.PubK = nullptr;
        return empty;
    }

    return chip;

}