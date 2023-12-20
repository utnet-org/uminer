##############################################################################
#
#    Copyright (c) 2016-2022 by Bitmain Technologies Inc. All rights reserved.
#
#    The material in this file is confidential and contains trade secrets
#    of Bitmain Technologies Inc. This is proprietary information owned by
#    Bitmain Technologies Inc. No part of this work may be disclosed,
#    reproduced, copied, transmitted, or used in any way for any purpose,
#    without the express written permission of Bitmain Technologies Inc.
#
##############################################################################
BMLIB_TOOLS_DIR = $(BM_ROOT)/bmlib/tools

BMLIB_TOOLS_SRCS_CXX = $(wildcard $(BMLIB_TOOLS_DIR)/*.cpp)

BMLIB_TOOLS_CXX_EXE1 = $(patsubst %.cpp,%,$(BMLIB_TOOLS_SRCS_CXX))
BMLIB_TOOLS_CXX_EXE2 = $(patsubst $(BMLIB_TOOLS_DIR)/%,%,$(BMLIB_TOOLS_CXX_EXE1))
BMLIB_TOOLS_CXX_EXE = $(addprefix $(OUT_DIR)/bmlib/tools/,$(BMLIB_TOOLS_CXX_EXE2))

BMLIB_TOOLS_OBJ_DIR = $(sort $(dir $(BMLIB_TOOLS_CXX_EXE)))
DAILY_TEST_EXES += $(BMLIB_TOOLS_CXX_EXE)
DAILY_TEST_EXES += $(BM_ROOT)/../bm_prebuilt_toolchains/sc5-bins/bl1.bin
DAILY_TEST_EXES += $(BM_ROOT)/../bm_prebuilt_toolchains/sc5-bins/spi_flash.bin

GFLAGS = -l:libgflags_nothreads.a

ifeq ($(CHIP_NAME),bm1684)
INCLUDE_DIR += -I$(BM_ROOT)/config \
              -I$(BM_ROOT)/common/bm1684/include \
              -I$(BM_ROOT)/bmlib/include \
              -I$(BM_ROOT)/bmlib/src \
              -I$(BM_ROOT)/firmware_core/bm1684/include \
              -I$(BM_ROOT)/firmware_core/bm1684/src \
              -I$(BM_ROOT)/driver/include \
              -I$(BM_ROOT)/c_model/bm1684/include \
              -I$(BM_ROOT)/c_model/bm1684/src
else ifeq ($(CHIP_NAME),bm1686)
INCLUDE_DIR += -I$(BM_ROOT)/config \
              -I$(BM_ROOT)/common/bm1686/include \
              -I$(BM_ROOT)/bmlib/include \
              -I$(BM_ROOT)/bmlib/src \
              -I$(BM_ROOT)/firmware_core/bm1686/include \
              -I$(BM_ROOT)/firmware_core/bm1686/src \
              -I$(BM_ROOT)/driver/include \
              -I$(BM_ROOT)/c_model/bm1686/include \
              -I$(BM_ROOT)/c_model/bm1686/src
endif

ifeq ($(HOST_ARCH),x86)
CURSES      := -l:libbmncurses.a
INCLUDE_DIR += -I$(BM_ROOT)/../bm_prebuilt_toolchains/x86/ncurses/include \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/x86/ncurses/include/ncurses \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/x86/gflags/include

BMLIB_TOOLS_LDFLAGS     += $(HOST_LDFLAGS) -L$(OUT_DIR) -L$(BM_ROOT)/out/build -L$(BM_ROOT)/../bm_prebuilt_toolchains/x86/ncurses/lib \
 -L$(BM_ROOT)/../bm_prebuilt_toolchains/x86-64-core-i7--glibc--stable-2018.11-1/x86_64-buildroot-linux-gnu/sysroot/lib \
 -L$(BM_ROOT)/out/build -L$(BM_ROOT)/../bm_prebuilt_toolchains/x86/gflags/lib -l:bmlib.a
endif

ifeq ($(HOST_ARCH),aarch64)
CURSES      := -lncurses
INCLUDE_DIR += -I$(BM_ROOT)/../bm_prebuilt_toolchains/ncurses-5.9_arm64/build/include \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/ncurses-5.9_arm64/build/include/ncurses \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/aarch64/gflags/include

BMLIB_TOOLS_LDFLAGS     += $(HOST_LDFLAGS) -L$(OUT_DIR) -L$(BM_ROOT)/out/build -L$(BM_ROOT)/../bm_prebuilt_toolchains/ncurses-5.9_arm64/build/lib \
 -L$(BM_ROOT)/out/build -L$(BM_ROOT)/../bm_prebuilt_toolchains/aarch64/gflags/lib -l:bmlib.a
endif

ifeq ($(HOST_ARCH),mips64)
CURSES      := -lncurses
INCLUDE_DIR += -I$(BM_ROOT)/../bm_prebuilt_toolchains/mips64/ncurses-6.2/include \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/mips64/ncurses-6.2/include/ncurses \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/mips64/gflags/build/include

BMLIB_TOOLS_LDFLAGS     += -L$(OUT_DIR) -L$(BM_ROOT)/../bm_prebuilt_toolchains/mips64/ncurses-6.2/lib \
 -L$(BM_ROOT)/out/build -L$(BM_ROOT)/../bm_prebuilt_toolchains/mips64/gflags/build/lib \
 -L$(OUT_DIR) -l:bmlib.a -pthread
endif

ifeq ($(HOST_ARCH),loongarch64)
CURSES      := -lncurses
INCLUDE_DIR += -I$(BM_ROOT)/../bm_prebuilt_toolchains/loongarch64/ncurses-6.1/include \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/loongarch64/ncurses-6.1/include/ncurses \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/loongarch64/gflags/include

BMLIB_TOOLS_LDFLAGS     += -L$(OUT_DIR) -L$(BM_ROOT)/../bm_prebuilt_toolchains/loongarch64/ncurses-6.1/lib \
 -L$(BM_ROOT)/out/build -L$(BM_ROOT)/../bm_prebuilt_toolchains/loongarch64/gflags/lib \
 -L$(OUT_DIR) -l:bmlib.a -pthread
endif

ifeq ($(HOST_ARCH),sw64)
CURSES      := -lncurses
INCLUDE_DIR += -I$(BM_ROOT)/../bm_prebuilt_toolchains/sw64/ncurses-6.2/include \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/sws64/ncurses-6.2/include/ncurses \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/sw64/gflags/build/include

BMLIB_TOOLS_LDFLAGS     += -L$(OUT_DIR) -L$(BM_ROOT)/../bm_prebuilt_toolchains/sw64/ncurses-6.2/lib \
 -L$(BM_ROOT)/out/build -L$(BM_ROOT)/../bm_prebuilt_toolchains/sw64/gflags/build/lib \
 -L$(OUT_DIR) -l:bmlib.a -pthread
endif

ifeq ($(USING_CMODEL),1)
bmlib_tools:
	@echo "bmlib_tools not support in cmodel"
else
INSTALL_EXES += $(BMLIB_TOOLS_CXX_EXE)
ifneq ($(SOC_MODE), 1)
RELEASE_EXES += $(OUT_DIR)/bmlib/tools/bm_firmware_update
RELEASE_EXES += $(OUT_DIR)/bmlib/tools/update_boot_info
endif

$(BMLIB_TOOLS_OBJ_DIR):
	$(MKDIR) $@
$(BMLIB_TOOLS_CXX_EXE): $(OUT_DIR)/%:$(BM_ROOT)/%.cpp $(BMLIB_STATIC) | $(BMLIB_TOOLS_OBJ_DIR)
	@echo -o $@ $<
	@$(HOST_CXX) $(CXXFLAGS)  -o $@ $< $(CURSES) $(GFLAGS) $(BMLIB_TOOLS_LDFLAGS) -static

bmlib_tools: $(BMLIB_TOOLS_CXX_EXE)

BMLIB_TOOLS_DEP = $(patsubst %.cpp,%.d,$(BMLIB_TOOLS_SRCS_CXX))
-include $(BMLIB_TOOLS_DEP)

endif
