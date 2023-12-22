#include <stdio.h>
#include <iostream>
#include <cstring>
#include <openssl/ec.h>
#include <openssl/ecdsa.h>
#include <openssl/sha.h>
#include <openssl/obj_mac.h>

#include "chip.h"
#include "util.h"

int signatureVerify(const char* signatureHex, const char* pubK, unsigned int size_signature, unsigned int  size_pubkey, const char* message) {

    /* Generate digest*/
    const char* data = message;
    size_t data_size = sizeof(data) - 1;
    unsigned char digest[SHA256_DIGEST_LENGTH];
    generate_sha256_digest(reinterpret_cast<const unsigned char*>(data), data_size, digest);

    /* digital signature and pubkey */
    unsigned char* signature = (unsigned char*)malloc(size_pubkey);
    std:: string str(signatureHex);
    signature = hexToByteArray(str);
    unsigned char* pubkey = (unsigned char *)malloc(size_pubkey);
    strcpy(reinterpret_cast<char*>(pubkey), const_cast<char*>(pubK));

    /* Verify digital signature */
    int verify_result = verify_with_public_key (pubkey, size_pubkey,
                                                digest, SHA256_DIGEST_LENGTH,
                                                signature, size_signature);

    free(pubkey);
    free(signature);

    if (verify_result == 1) {
        printf("sign verify success\n");
        return 1;
    } else {
        printf("sign verify fail\n");
        return 0;
    }

}