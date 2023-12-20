cmake_minimum_required(VERSION 3.13)
cmake_policy(SET CMP0046 NEW)

include_directories(${CMAKE_CURRENT_SOURCE_DIR}/)
include_directories(${PROJECT_ROOT}/config)
include_directories(${PROJECT_ROOT}/common/bm1684/include)
include_directories(${PROJECT_ROOT}/c_model/bm1684/include)
include_directories(${PROJECT_ROOT}/c_model/bm1684/src)
include_directories(${PROJECT_ROOT}/bmlib/include)
include_directories(${PROJECT_ROOT}/bmdnn/bm1684/include)
include_directories(${PROJECT_ROOT}/firmware_core/bm1684/include)
include_directories(${PROJECT_ROOT}/bmlib/src)
if(LINUX)
include_directories(${PROJECT_ROOT}/bmlib/src/linux)
elseif(WINDOWS)
include_directories(${PROJECT_ROOT}/bmlib/src/window)
endif()

link_directories(${PROJECT_ROOT}/out/release/$ENV{TARGET_PROJECT}_CMAKE/lib)
if(TARGET_TYPE STREQUAL "release")
	if (RUNTIME_LIB STREQUAL "MD")
		set(CMAKE_CXX_FLAGS_RELEASE "/MD")
		set(CMAKE_C_FLAGS_RELEASE "/MD")
	else()
		set(CMAKE_CXX_FLAGS_RELEASE "/MT")
		set(CMAKE_C_FLAGS_RELEASE "/MT")
	endif()
else()
	if (RUNTIME_LIB STREQUAL "MD")
		set(CMAKE_CXX_FLAGS_DEBUG "/MDd")
		set(CMAKE_C_FLAGS_DEBUG "/MDd")
	else()
		set(CMAKE_CXX_FLAGS_DEBUG "/MTd")
		set(CMAKE_C_FLAGS_DEBUG "/MTd")
	endif()
endif()
if(LINUX)
set(SRCS
  src/bmkernel_runtime.cpp
  src/bmlib_mmpool.cpp
  src/bmcpu_runtime.cpp
  src/bmlib_runtime.cpp
  src/bmlib_util.cpp
  src/bm1682_runtime.cpp
  src/bmlib_log.cpp
  src/bmlib_device.cpp
  src/bmlib_memory.cpp

  src/bmlib_profile.cpp
  src/linux/bmlib_ioctl.cpp
)
elseif(WINDOWS)

set(CMAKE_RUNTIME_OUTPUT_DIRECTORY_DEBUG ${CMAKE_INSTALL_PREFIX}/bin)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY_DEBUG ${CMAKE_INSTALL_PREFIX}/lib)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY_RELEASE ${CMAKE_INSTALL_PREFIX}/bin)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY_RELEASE ${CMAKE_INSTALL_PREFIX}/lib)

set(SRCS
  # src/bmkernel_runtime.cpp
  # src/bmlib_mmpool.cpp
  src/bmcpu_runtime.cpp
  src/bmlib_runtime.cpp
  # src/bmlib_util.cpp
  # src/bm1682_runtime.cpp
  src/bmlib_log.cpp
  # src/bmlib_device.cpp
  src/bmlib_memory.cpp

  src/bmlib_profile.cpp
  src/window/bmlib_ioctl.cpp
)
endif()

add_library(libbmlib SHARED ${SRCS})
if(WINDOWS)
    set(CMAKE_MODULE_PATH ${PROJECT_ROOT}/.)
    message(STATUS "cmake_module_path: ${CMAKE_MODULE_PATH}")
    find_package(SETUPAPILIB_LIBRARY)

    if(SETUPAPILIB_FOUND)
        message(STATUS "found SetupAPI ${SETUPAPILIB_INCLUDE_DIR} ${SETUPAPILIB_LIBRARY}")
        include_directories(${SETUPAPILIB_INCLUDE_DIR})
        target_link_libraries(libbmlib ${SETUPAPILIB_LIBRARY})
        #target_link_libraries(bmlib ${CMAKE_CURRENT_SOURCE_DIR}/SetupAPI.Lib)
    else()
        message(STATUS "err: not found SETUPAPILIB_FOUND")
    endif(SETUPAPILIB_FOUND)

elseif(LINUX)
    if(USING_CMODEL)
        target_link_libraries(libbmlib libcommon cmodel firmware_core dl)
    else()
        target_link_libraries(libbmlib libcommon)
    endif()
endif()

add_library(libbmlib-static STATIC ${SRCS})
if(WINDOWS)
    include_directories(${SETUPAPILIB_INCLUDE_DIR})
    target_link_libraries(libbmlib-static libcommon-static ${SETUPAPILIB_LIBRARY})
elseif(LINUX)
    if(USING_CMODEL)
        target_link_libraries(libbmlib-static libcommon-static cmodel-static firmware_core-static dl)
    else()
        target_link_libraries(libbmlib-static libcommon-static)
    endif()
endif()
install(TARGETS libbmlib libbmlib-static DESTINATION lib)
