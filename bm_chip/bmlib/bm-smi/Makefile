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

BMSMI_SRCS_CXX    = $(BM_ROOT)/bmlib/bm-smi/src/*.cpp
BMSMI = $(OUT_DIR)/bmlib/bm-smi/bm-smi

ifneq ($(USING_CMODEL),1)
INSTALL_EXES += $(BMSMI)
RELEASE_EXES += $(BMSMI)
NNTC_RELEASE_EXES += $(BMSMI)
endif

GFLAGS = -l:libgflags_nothreads.a

INCLUDE_DIR += -I$(BM_ROOT)/bmlib/src -I$(BM_ROOT)/bmlib/include \
	       -I$(BM_ROOT)/bmlib/bm-smi/src -I$(BM_ROOT)/bmlib/bm-smi/include

ifeq ($(HOST_ARCH),x86)
CURSES      := -l:libbmncurses.a
INCLUDE_DIR += -I$(BM_ROOT)/../bm_prebuilt_toolchains/x86/ncurses/include \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/x86/ncurses/include/ncurses \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/x86/gflags/include

LDFLAGS     += -L$(OUT_DIR) -L$(BM_ROOT)/../bm_prebuilt_toolchains/x86/ncurses/lib \
 -L$(BM_ROOT)/../bm_prebuilt_toolchains/x86-64-core-i7--glibc--stable-2018.11-1/x86_64-buildroot-linux-gnu/sysroot/lib \
 -L$(BM_ROOT)/out/build -L$(BM_ROOT)/../bm_prebuilt_toolchains/x86/gflags/lib -l:bmlib.a -pthread
endif

ifeq ($(HOST_ARCH),aarch64)
CURSES      := -lncurses
INCLUDE_DIR += -I$(BM_ROOT)/../bm_prebuilt_toolchains/ncurses-5.9_arm64/build/include \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/ncurses-5.9_arm64/build/include/ncurses \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/aarch64/gflags/include

LDFLAGS     += -L$(OUT_DIR) -L$(BM_ROOT)/../bm_prebuilt_toolchains/ncurses-5.9_arm64/build/lib \
 -L$(BM_ROOT)/out/build -L$(BM_ROOT)/../bm_prebuilt_toolchains/aarch64/gflags/lib \
 -L$(OUT_DIR) -l:bmlib.a -pthread
endif

ifeq ($(HOST_ARCH),mips64)
CURSES      := -lncurses
INCLUDE_DIR += -I$(BM_ROOT)/../bm_prebuilt_toolchains/mips64/ncurses-6.2/include \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/mips64/ncurses-6.2/include/ncurses \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/mips64/gflags/build/include

LDFLAGS     += -L$(OUT_DIR) -L$(BM_ROOT)/../bm_prebuilt_toolchains/mips64/ncurses-6.2/lib \
 -L$(BM_ROOT)/out/build -L$(BM_ROOT)/../bm_prebuilt_toolchains/mips64/gflags/build/lib \
 -L$(OUT_DIR) -l:bmlib.a -pthread
endif

ifeq ($(HOST_ARCH),loongarch64)
CURSES      := -lncurses
INCLUDE_DIR += -I$(BM_ROOT)/../bm_prebuilt_toolchains/loongarch64/ncurses-6.1/include \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/loongarch64/ncurses-6.1/include/ncurses \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/loongarch64/gflags/include

LDFLAGS     += -L$(OUT_DIR) -L$(BM_ROOT)/../bm_prebuilt_toolchains/loongarch64/ncurses-6.1/lib \
 -L$(BM_ROOT)/out/build -L$(BM_ROOT)/../bm_prebuilt_toolchains/loongarch64/gflags/lib \
 -L$(OUT_DIR) -l:bmlib.a -pthread
endif

ifeq ($(HOST_ARCH),sw64)
CURSES      := -lncurses
INCLUDE_DIR += -I$(BM_ROOT)/../bm_prebuilt_toolchains/sw64/ncurses-6.2/include \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/sw64/ncurses-6.2/include/ncurses \
              -I$(BM_ROOT)/../bm_prebuilt_toolchains/sw64/gflags/build/include

LDFLAGS     += -L$(OUT_DIR) -L$(BM_ROOT)/../bm_prebuilt_toolchains/sw64/ncurses-6.2/lib \
 -L$(BM_ROOT)/out/build -L$(BM_ROOT)/../bm_prebuilt_toolchains/sw64/gflags/build/lib \
 -L$(OUT_DIR) -l:bmlib.a -pthread
endif

CXXFLAGS    += $(HOST_CXXFLAGS) $(INCLUDE_DIR)

ifneq ($(USING_CMODEL),1)
bm-smi: $(BMSMI_SRCS_CXX) $(BMLIB_STATIC)
	@$(MKDIR) $(OUT_DIR)/bmlib/bm-smi
	@$(QCOMPILE)$(HOST_CXX) $(CXXFLAGS) $(BMSMI_SRCS_CXX)  -o $(BMSMI) $(LDFLAGS) $(CURSES) $(GFLAGS) -static
else
bm-smi:
	@echo "bm-smi not support in cmodel"
endif

BM_SMI_DEP = $(patsubst %.cpp,%.d,$(BMSMI_SRCS_CXX))
-include $(BM_SMI_DEP)
