#include <stdio.h>
#include <iostream>
#include <cstring>
#include <stdlib.h>
#include <sstream>
#include <openssl/ec.h>
#include <openssl/ecdsa.h>
#include <openssl/sha.h>
#include <openssl/obj_mac.h>

#include "chip.h"

std::string hexStringToByteArray(const std::string& hexString) {
    std::stringstream ss;
    for (size_t i = 0; i < hexString.length(); i += 2) {
        std::string byteString = hexString.substr(i, 2);
        unsigned char byte = static_cast<unsigned char>(std::stoi(byteString, nullptr, 16));
        ss << byte;
    }
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
    unsigned int  size_pubkey = 2048;
    unsigned char* p2 = (unsigned char*)malloc(size_p2);
    unsigned char* pubkey = (unsigned char*)malloc(size_pubkey);
    unsigned int size_p2_padding;

    ChipDeclaration oneChip;

    // read files to get results
//    size_p2 = 1675;
    FILE *file_pubkey = fopen(("../../bm_chip/src/pubkey_"+ std::to_string(seq)).c_str(), "r");
    FILE *file_p2 = fopen(("../../bm_chip/src/p2_"+ std::to_string(seq)).c_str(), "r");
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

    std::string byteArray = hexStringToByteArray("b8dbbe56e32ba53d9e5f17a45f07fb2551e05aeeeabe8cc7241c8779a9b5a796862fce3ea7eb9935fbf72ca32efee4ee68330b643df8efc10d66066cb801d72d9f1b75a8efa7df38327133b1a0947c0582762c349b3f2c8131de92374d2f0b3b84e3d1fc689c86861788feaeee641e2545d06d7353dbab871d78e506e0b2e246e9d9d8c02e14f414b30dec6f63f5");
    std::cout << "解码后的字节数据: " << byteArray << std::endl;

    std::string P2 = std::string(reinterpret_cast<char*>(p2));
    std::string PubKey = std::string(reinterpret_cast<char*>(pubkey));

    oneChip.EncryptedPrivK = new unsigned char[P2.length() + 1];
    oneChip.PubK = new unsigned char[PubKey.length() + 1];

    strcpy(reinterpret_cast<char*>(oneChip.EncryptedPrivK), P2.c_str());
    strcpy(reinterpret_cast<char*>(oneChip.PubK), PubKey.c_str());

    return oneChip;

}

int chipBurning(int dev_id) {

    printf("\nthe chip dev_id %d to be burned ... \n", dev_id);

    /*** 芯片烧录 ***/

    return 1;

}