cmake_minimum_required(VERSION 3.11)
set(TARGET_NAME "bmcpu")

project(${TARGET_NAME} C CXX)

set(CMAKE_CXX_STANDARD 14)

# add source file
#add_library(ut_miner SHARED chip.cpp main.cpp)

find_package(OpenSSL REQUIRED)

# add OpenSSL source file directory
include_directories(/usr/local/opt/openssl/include)
include_directories (${CMAKE_CURRENT_SOURCE_DIR}/../bmlib/include/)

# find libbmlib.so
LINK_LIBRARIES(${CMAKE_CURRENT_SOURCE_DIR}/libbmlib.so)

add_executable(ut_miner chip.cpp chipStart.cpp chipBurning.cpp chipSignature.cpp chipVerify.cpp main.cpp)

foreach(. ${TEST_SRC_PATH})
    get_filename_component(target ${src} NAME_WE)
    add_executable(${target} ${src})
    target_link_libraries(${target} ../bmlib OpenSSL::SSL OpenSSL::Crypto)
endforeach(.)

# link to OpenSSL
#target_link_libraries(uminer /usr/local/opt/openssl/lib/libssl.dylib /usr/local/opt/openssl/lib/libcrypto.dylib)
target_link_libraries(ut_miner OpenSSL::SSL OpenSSL::Crypto)

# link to pthread library
#target_link_libraries(ut_miner PRIVATE pthread)