#include <stdlib.h>
#include <stdio.h>
#include <string>

#ifndef UMINER_UTIL_H
#define UMINER_UTIL_H

// signature and public key
struct keyPairs{
    EC_KEY *Eckey;
    const EC_POINT *PubK;
};
struct keyPairs getKeyPairs();
std::string byteArrayToHex(const unsigned char* byteArray, size_t length);

#endif //UMINER_UTIL_H
