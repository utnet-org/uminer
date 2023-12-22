#ifndef CHIP_H
#define CHIP_H
#include <openssl/ec.h>
#include <openssl/ecdsa.h>
#include <openssl/sha.h>
#include <openssl/obj_mac.h>

// default value
#define ATTR_FAULT_VALUE         (int)0xFFFFFC00


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
    const char* EncryptedPriK;
    const char* PubK;
    int EncryptedPriKSize;
    int PubKSize;
};

/* method for chips */
// start chip cpu
int startCPU(int dev_id, const char * boot_file, const char * core_file);
// chip burning
int chipBurning(int dev_id);
// chip get P2 and pubkey
int chipGenKeyPairs(int dev_id);
// read P2 and pubkey from file
struct ChipDeclaration readKeyPairs(int dev_id);
// chip signature
struct ChipSignature* chipSignature(unsigned long chipId, const char* p2, const char* pubkey, const char* message, unsigned int size_p2, unsigned int  size_pubkey);
// chip verification
int signatureVerify(const char* signature, const char* pubK, unsigned int size_signature, unsigned int size_pubkey, const char* message);



#ifdef __cplusplus
}

ChipSignature SignAtSPACC(int seq, unsigned char hash[SHA256_DIGEST_LENGTH]);

#endif

#endif // CHIP_H