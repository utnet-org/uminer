#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <string>
#include <sstream>

#ifndef UMINER_UTIL_H
#define UMINER_UTIL_H

// signature and public key
struct keyPairs{
    EC_KEY *Eckey;
    const EC_POINT *PubK;
};
struct keyPairs getKeyPairs();
unsigned char* hexToByteArray(const std::string& hexString);
std::string byteArrayToHex(const unsigned char* byteArray, size_t length);
void generate_sha256_digest(const unsigned char *data, size_t data_size, unsigned char *digest);
int verify_with_public_key(const unsigned char *publicKeyData, unsigned int pub_key_length,
                           const unsigned char *digest, unsigned int digest_size,
                           const unsigned char *signature, unsigned int signature_length);

#endif //UMINER_UTIL_H
