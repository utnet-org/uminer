#include <openssl/ec.h>
#include <openssl/ecdsa.h>
#include <openssl/sha.h>
#include <openssl/obj_mac.h>

#include "chip.h"

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
