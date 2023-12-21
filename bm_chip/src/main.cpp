#include <stdlib.h>
#include <stdio.h>
#include <string>
#include <sstream>
#include <iomanip>
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

std::string byteArrayToHexTest(const unsigned char* byteArray, size_t length) {
    std::stringstream ss;
    ss << std::hex << std::setfill('0');
    for (size_t i = 0; i < length; ++i)
        ss << std::setw(2) << static_cast<unsigned int>(byteArray[i]);
    return ss.str();
}

void signDemo(int devId, unsigned int size_p2){
    FILE *file_p2 = fopen(("../key/p2_"+ std::to_string(devId)).c_str(), "r");
    FILE *file_pubkey = fopen(("../key/pubkey_"+ std::to_string(devId)).c_str(), "r");

    // p2
    unsigned int size_p2_padding;
    unsigned char* P2 = (unsigned char *)malloc(size_p2);
    if (file_p2) {
        fseek(file_p2, 0, SEEK_END);
        size_p2_padding = ftell(file_p2);
        fseek(file_p2, 0, SEEK_SET);

        fread(P2, 1, size_p2_padding, file_p2);
        fclose(file_p2);

    } else {
        printf("Error opening file.\n");
        return;
    }
    std::string str = byteArrayToHexTest(P2, size_p2_padding);
    const char* P2Byte = str.c_str();

    // pubK
    unsigned int size_pubkey;
    unsigned char* pubkey = (unsigned char*)malloc(size_pubkey);
    if (file_pubkey) {
        fseek(file_pubkey, 0, SEEK_END);
        size_pubkey = ftell(file_pubkey);
        fseek(file_pubkey, 0, SEEK_SET);

        fread(pubkey, 1, size_pubkey, file_pubkey);
        fclose(file_pubkey);

    } else {
        printf("Error opening file.\n");
        return;
    }
    std::string PubK(reinterpret_cast<char*>(pubkey));
    std::string msg = "utility";

    chipSignature(devId, P2Byte, PubK.c_str(), msg.c_str(), size_p2_padding, 426);

}

int main(int argc, char *argv[]) {
    if (argc < 3) {
        printf("insufficient argument for the main\n");
        return -1;
    }
    if (strcmp(argv[1], "start") == 0) {
        const char* fipBin = "/root/uminer/bm_chip/src/fip.bin";
        const char* rambootRootfs = "/root/uminer/bm_chip/src/ramboot_rootfs.itb";
        startCPU(atoi(argv[2]), fipBin, rambootRootfs);
    }
    if (strcmp(argv[1], "burn") == 0) {
        chipBurning(atoi(argv[2]));
    }
    if (strcmp(argv[1], "keygen") == 0){
        chipGenKeyPairs(atoi(argv[2]));
    }
    if (strcmp(argv[1], "sign") == 0) {
        signDemo(atoi(argv[2]), 1680);
    }

    return 0;

}

