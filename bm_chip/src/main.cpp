#include <stdlib.h>
#include <stdio.h>
#include <string>
#include <sstream>
#include <iostream>
#include <iomanip>
#include <string.h>
#include <pthread.h>
#include <bmlib_runtime.h>

#include <openssl/rsa.h>
#include <openssl/pem.h>
#include <openssl/evp.h>
#include <openssl/pem.h>
#include <openssl/bn.h>
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

std::string Base58EncodeTest(const unsigned char* input, size_t length) {
    // Convert the input binary data to OpenSSL BIGNUM
    BIGNUM* bn = BN_bin2bn(input, length, nullptr);
    if (!bn) {
        std::cout << "Error: Unable to create BIGNUM\n";
        return "";
    }

    // Create a BIO for Base58 encoding
    BIO* b58 = BIO_new(BIO_f_base64());
    BIO_set_flags(b58, BIO_FLAGS_BASE64_NO_NL);

    // Write BIGNUM to BIO for encoding
    BIO* mem = BIO_new(BIO_s_mem());
    BIO_push(b58, mem);
    if (!BN_print(b58, bn)) {
        std::cout << "Error: Unable to encode BIGNUM\n";
        BIO_free_all(b58);
        BN_free(bn);
        return "";
    }

    // Read Base58 encoded data from BIO
    BUF_MEM* bptr;
    BIO_get_mem_ptr(mem, &bptr);

    // Convert BIO data to string
    std::string encoded(reinterpret_cast<char*>(bptr->data), bptr->length);

    // Clean up
    BIO_free_all(b58);
    BN_free(bn);

    return encoded;
}

void signDemo(int devId, unsigned int size_pubkey, unsigned int size_p2){
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
    printf("p2 = %s\n", P2);
    std::string str = byteArrayToHexTest(P2, size_p2_padding);
    const char* P2Byte = str.c_str();

    // pubK
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
        chipGenKeyPairs(argv[2], argv[3], atoi(argv[4]));
    }
    if (strcmp(argv[1], "sign") == 0) {
        signDemo(atoi(argv[2]), 426, 1680);
    }

    return 0;

}

