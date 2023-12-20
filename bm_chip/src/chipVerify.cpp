#include <stdio.h>
#include <iostream>
#include <cstring>
#include <openssl/ec.h>
#include <openssl/ecdsa.h>
#include <openssl/sha.h>
#include <openssl/obj_mac.h>

#include "chip.h"

ChipVerify* chipVerify(unsigned long segment1, unsigned long segment2, const char* p2, const char* pubK, const char* message) {

    unsigned long segmentCount = segment2 - segment1;
    // 生成 ChipSignature 结构体
    ChipVerify* Chips = new ChipVerify[segmentCount];

    // 生成消息
    const char *messageBytes = message;
    unsigned char hash[SHA256_DIGEST_LENGTH];
    SHA256(reinterpret_cast<const unsigned char*>(messageBytes), strlen(messageBytes), hash);  //用 message 的内容计算 SHA-256 哈希

    printf("\n%lu chips to be verified ... \n", segmentCount);
    for (int i = 0; i < segmentCount; ++i) {

        keyPairs keys = getKeyPairs();
        // 从驱动读取私钥再签名(真正环境下把p2传入得到signature)
        /*** 芯片签名digest ***/
        ECDSA_SIG *signature = ECDSA_do_sign(hash, sizeof(hash), keys.Eckey);
        if (!signature) {
            fprintf(stderr, "Error signing the message\n");
            ChipVerify* emptys = new ChipVerify[1];
            emptys->ifVerifyPass = 0;
            emptys->SignMsg = nullptr;
            return emptys;
        }

        // 从驱动验签(真正环境下把pubK传入验证signature)
        /*** 验证签名digest ***/
        EC_KEY *eckey0 = EC_KEY_new_by_curve_name(NID_secp256k1);
        EC_KEY_set_public_key(eckey0, keys.PubK);
        if (ECDSA_do_verify(hash, sizeof(hash), signature, eckey0) != 1) {
            fprintf(stderr, "Signature verification failed\n");
            ChipVerify* emptys = new ChipVerify[1];
            emptys->ifVerifyPass = 0;
            emptys->SignMsg = nullptr;
            return emptys;
            return 0;
        } else {
            printf("chip %d verified successfully\n", i);
            std::string signature = "chipSignature" + std::to_string(i);
            Chips[i].ifVerifyPass = 1;
            Chips[i].SignMsg = new unsigned char[signature.length() + 1];

            strcpy(reinterpret_cast<char*>(Chips[i].SignMsg), signature.c_str());
        }

        // 清理资源
//        ECDSA_SIG_free(signature);
//        OPENSSL_free(keys.PubK);
//        EC_KEY_free(keys.Eckey);
    }

    return Chips;
}