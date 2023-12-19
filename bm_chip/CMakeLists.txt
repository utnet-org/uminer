cmake_minimum_required(VERSION 3.21)
project(ut_miner)

set(CMAKE_CXX_STANDARD 14)

# 添加源文件
#add_library(ut_miner SHARED chip.cpp main.cpp)

find_package(OpenSSL REQUIRED)

# 添加 OpenSSL 头文件路径
include_directories(/usr/local/opt/openssl/include)

#add_executable(ut_miner src/util.cpp src/mechanism.cpp src/block.cpp src/chip.cpp src/crypto/aes.cpp src/main.cpp)
add_executable(ut_miner src/chip.cpp src/chipBurning.cpp src/chipSignature.cpp src/chipVerify.cpp src/main.cpp)

# 链接 OpenSSL 库
target_link_libraries(ut_miner /usr/local/opt/openssl/lib/libssl.dylib /usr/local/opt/openssl/lib/libcrypto.dylib)

# 链接 pthread 库
#target_link_libraries(ut_miner PRIVATE pthread)


