#include <stdio.h>
#include <iostream>
#include <cstring>
#include <openssl/ec.h>
#include <openssl/ecdsa.h>
#include <openssl/sha.h>
#include <openssl/obj_mac.h>

#include "chip.h"


struct ChipKeys {
    ECDSA_SIG *SignMsg;
    const EC_POINT *PubK;
};

ChipSignature SignAtSPACC(int seq, unsigned char hash[SHA256_DIGEST_LENGTH]) {

    /*** 芯片里得到key ***/
    keyPairs keys = getKeyPairs();

    // 签名消息
    ECDSA_SIG *keySignature = ECDSA_do_sign(hash, sizeof(hash), keys.Eckey);

//    ChipKeys signs;
//    signs.SignMsg = signature;
//    signs.PubK = keys.PubK;

    // 模拟签完
    ChipSignature signs;
    std::string signature = "chipSignature" + std::to_string(seq);
    std::string pubKey = "chipPubK" + std::to_string(seq);

    signs.SignMsg = new unsigned char[signature.length() + 1];
    signs.PubK = new unsigned char[pubKey.length() + 1];

    strcpy(reinterpret_cast<char*>(signs.SignMsg), signature.c_str());
    strcpy(reinterpret_cast<char*>(signs.PubK), pubKey.c_str());

    return signs;

}

ChipSignature* chipSignature(unsigned long chipId, const char* p2, const char* message) {

    // 生成 ChipSignature 结构体
    ChipSignature* signatureList = new ChipSignature[1];

    // 生成消息digest
    const char *messageBytes = message; //"Hello, ECC!";
    unsigned char digest[SHA256_DIGEST_LENGTH];
    SHA256(reinterpret_cast<const unsigned char*>(messageBytes), strlen(messageBytes), digest);  //用 message 的内容计算 SHA-256 哈希

    printf("\nchip to be signed ... \n");

    // 从驱动读取私钥再签名
    /*** 芯片签名digest ***/
    ChipSignature keySignature = SignAtSPACC(chipId, digest);
    if (keySignature.SignMsg == nullptr) {
        fprintf(stderr, "Error signing the message %ld\n", chipId);
        ChipSignature* emptys = new ChipSignature[1];
        emptys->SignMsg = nullptr;
        emptys->PubK = nullptr;
        return emptys;
    }
    // 将生成的 ChipSignature 结构体加入到数组中

//    signatureList->SignMsg = signature.SignMsg;
//    signatureList->PubK = signature.PubK;

    std::string signature = std::string(reinterpret_cast<char*>(keySignature.SignMsg));
    std::string pubKey = std::string(reinterpret_cast<char*>(keySignature.PubK));

    signatureList[0].SignMsg = new unsigned char[signature.length() + 1];
    signatureList[0].PubK = new unsigned char[pubKey.length() + 1];

    strcpy(reinterpret_cast<char*>(signatureList[0].SignMsg), reinterpret_cast<char*>(keySignature.SignMsg));
    strcpy(reinterpret_cast<char*>(signatureList[0].PubK), reinterpret_cast<char*>(keySignature.PubK));

    // 清理资源
//    ECDSA_SIG_free(signature.SignMsg);


    return signatureList;

}