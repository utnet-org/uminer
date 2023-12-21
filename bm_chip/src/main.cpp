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
void signDemo(unsigned int size_p2){
    FILE *file_p2 = fopen("../key/p2_10", "r");
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
    }
    std::string str = byteArrayToHexTest(P2, size_p2_padding);
    const char* P2Byte = str.c_str();

    std::string PubK = "-----BEGIN RSA PUBLIC KEY-----\n"
                       "MIIBCgKCAQEA6fn2R5LBtnJ+P7mINn6rv+xUzsZ4ojfft7ISMyYFTNqgfgk7E8H+\n"
                       "lWrm5xDqY0axE9zWyBSeCunWmX/KLMlvleDWyTvRk4ZJn8tY5bTxBLmRXI6DC8pr\n"
                       "mjVegpojico4PYz8fCKwpzM8kUpl3qPkreRk+qwu8mV/l4FdfK+DKGXrqkhAsAma\n"
                       "Iz3lSpcybJrNzIeRvGX7Y7Z20hY8Bm8QIIlr+vLwlhKwCghbYcjhrPU77de5bvAU\n"
                       "QYxLoE+MN2Ux65d46+VAVKpmKLCEvdJ5ezCksTkPFaOYtVdOpaAjwLv6eEdGV9IQ\n"
                       "UedEqPGLRBclMElR3r9WI6GNIsPAa/w/uQIDAQAB\n"
                       "-----END RSA PUBLIC KEY-----";
    std::string msg = "utility";
    chipSignature(10, P2Byte, PubK.c_str(), msg.c_str(), size_p2_padding, 426);
}

int main() {
    signDemo(1680);
    return 0;

}

