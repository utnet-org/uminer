cmake_minimum_required(VERSION 3.13)
cmake_policy(SET CMP0046 NEW)

PROJECT(bm-smi)
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/)
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/src)
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/include)
include_directories(${PROJECT_ROOT}/config)
include_directories(${PROJECT_ROOT}/common/bm1684/include)
include_directories(${PROJECT_ROOT}/bmlib/include)
include_directories(${PROJECT_ROOT}/bmlib/src)
include_directories(${PROJECT_ROOT}/../bm_prebuilt_toolchains_win/gflags-windows/include)
include_directories(${PROJECT_ROOT}/../bm_prebuilt_toolchains_win/PDCurses-3.8/PDCurses-3.8)
include_directories(${PROJECT_ROOT}/../bm_prebuilt_toolchains_win/PDCurses-3.8/PDCurses-3.8/wincon)
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
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY_DEBUG ${CMAKE_INSTALL_PREFIX}/bm-smi)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY_DEBUG ${CMAKE_INSTALL_PREFIX}/bm-smi)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY_RELEASE ${CMAKE_INSTALL_PREFIX}/bm-smi)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY_RELEASE ${CMAKE_INSTALL_PREFIX}/bm-smi)
if(LINUX)
include_directories(${PROJECT_ROOT}/bmlib/src/linux)
elseif(WINDOWS)
include_directories(${PROJECT_ROOT}/bmlib/src/window)
endif()

link_directories(${PROJECT_ROOT}/../bm_prebuilt_toolchains_win/PDCurses-3.8/PDCurses-3.8/wincon)
if(TARGET_TYPE STREQUAL "release")
	link_directories(${PROJECT_ROOT}/../bm_prebuilt_toolchains_win/gflags-windows/lib/Release)
	file(COPY ${PROJECT_ROOT}/../bm_prebuilt_toolchains_win/gflags-windows/bin/Release/gflags.dll DESTINATION ${CMAKE_INSTALL_PREFIX}/bm-smi)
else()
	link_directories(${PROJECT_ROOT}/../bm_prebuilt_toolchains_win/gflags-windows/lib/Debug)
	file(COPY ${PROJECT_ROOT}/../bm_prebuilt_toolchains_win/gflags-windows/bin/Debug/gflags_debug.dll DESTINATION ${CMAKE_INSTALL_PREFIX}/bm-smi)
endif()
link_directories(${PROJECT_ROOT}/out/release/$ENV{TARGET_PROJECT}_CMAKE/lib)
SET(SRCS
  src/bm-smi.cpp
  src/bm_smi_cmdline.cpp
  src/bm_smi_test.cpp
  src/bm_smi_creator.cpp
  src/bm_smi_display.cpp
  src/bm_smi_ecc.cpp
  src/bm_smi_led.cpp
  src/bm_smi_recovery.cpp
  src/bm_smi_display_memory_detail.cpp
  src/bm_smi_display_util_detail.cpp
)

add_executable(bm-smi ${SRCS})
if(TARGET_TYPE STREQUAL "release")
	target_link_libraries(bm-smi pdcurses gflags)
else()
	target_link_libraries(bm-smi pdcurses gflags_debug)
endif()
add_dependencies(bm-smi libbmlib-static)
file(COPY ${PROJECT_ROOT}/../bm_prebuilt_toolchains_win/PDCurses-3.8/PDCurses-3.8/wincon/pdcurses.dll DESTINATION ${CMAKE_INSTALL_PREFIX}/bm-smi)
