#ifndef CHIP_H
#define CHIP_H
#include <openssl/ec.h>
#include <openssl/ecdsa.h>
#include <openssl/sha.h>
#include <openssl/obj_mac.h>

// default value
#define ATTR_FAULT_VALUE         (int)0xFFFFFC00

// signature and public key
struct keyPairs{
    EC_KEY *Eckey;
    const EC_POINT *PubK;
};


// to cgo
#ifdef __cplusplus
extern "C" {
#endif

/* bm_smi_attr defines all the attributes fetched from kernel;
* it will be displayed by bm-smi and saved in files if needed.
* this struct is also defined in driver/bm_uapi.h.*/
struct bmsmi {
    int dev_id;
    int chip_id;
    int domain_bdf;  /* busid */
    int chip_mode;  /*0---pcie; 1---soc*/
    int status;      /* 0---Active 1---Fault */

    int mem_total;  // memory in total
    int mem_used;   // memory used
    int tpu_util;   // tpu util percentage

    int board_temp;  // board temperature
    int chip_temp;   // chip temperature
    int board_power;
    float tpu_power;
//    int fan_speed;

};

// api that feedback all chip information
struct bmsmi* queryChipDetails();

// chip signature
struct ChipSignature {
//    ECDSA_SIG *SignMsg;
//    const EC_POINT *PubK;
    unsigned char* SignMsg;
    unsigned char* PubK;
    int  status;
};
// burning on efuse: P2 + PubKey
struct ChipDeclaration {
    const char* EncryptedPrivK;
    const char* PubK;
};
// struct verification on block burst
struct ChipVerify {
    unsigned char* SignMsg;
    int ifVerifyPass;
};

/* method for chips */
// chip burning
int chipBurning(int dev_id);
// chip get P2 and pubkey
int chipGenPPubkey(int dev_id);
// read P2 and pubkey from file
struct ChipDeclaration readPPubkey(int dev_id);
// chip signature
struct ChipSignature* chipSignature(unsigned long chipId, const char* p2, const char* pubkey, const char* message, unsigned int size_p2, unsigned int  size_pubkey);
// chip verification
struct ChipVerify* chipVerify(unsigned long segment1, unsigned long segment2, const char* p2, const char* pubK, const char* message);



#ifdef __cplusplus
}

struct keyPairs getKeyPairs();

ChipSignature SignAtSPACC(int seq, unsigned char hash[SHA256_DIGEST_LENGTH]);

#endif

#endif // CHIP_H