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

BMLIB_DIR = $(BM_ROOT)/bmlib

BMLIB_SRCS_C      = $(wildcard $(BMLIB_DIR)/src/*.c)
BMLIB_SRCS_C      += $(COMMON_SRCS_C)

BMLIB_SRCS_CXX    = $(wildcard $(BMLIB_DIR)/src/*.cpp)
BMLIB_SRCS_CXX    += $(wildcard $(BMLIB_DIR)/src/linux/*.cpp)

INCLUDE_DIR += -I$(BMLIB_DIR)/include
INCLUDE_DIR += -I$(BMLIB_DIR)/src/linux
ifeq ($(CHIP_NAME),bm1684)
INCLUDE_DIR += -I$(BM_ROOT)/firmware_top/bm1684/boot_test/include
else ifeq ($(CHIP_NAME),bm1686)
INCLUDE_DIR += -I$(BM_ROOT)/firmware_top/bm1686/boot_test/include
endif

ifeq ($(USING_CMODEL),1)
BMLIB_SRCS_C      += $(FIRMWARE_CORE_SRCS_C)
INCLUDE_DIR += -I$(FIRMWARE_CORE)/include

BMLIB_SRCS_C      += $(C_MODEL_SRCS_C)
BMLIB_SRCS_CXX      += $(C_MODEL_SRCS_CXX)
INCLUDE_DIR += -I$(C_MODEL)/include
endif

BMLIB_C_OBJS1 = $(call SOURCES_TO_OBJS,$(BMLIB_SRCS_C),c)
BMLIB_CXX_OBJS1 = $(call SOURCES_TO_OBJS,$(BMLIB_SRCS_CXX),cpp)

BMLIB_C_OBJS2 = $(patsubst $(BM_ROOT)/%.o,%.o, $(BMLIB_C_OBJS1))
BMLIB_CXX_OBJS2 = $(patsubst $(BM_ROOT)/%.o,%.o, $(BMLIB_CXX_OBJS1))

BMLIB_C_OBJS = $(addprefix $(OUT_DIR)/,$(BMLIB_C_OBJS2))
BMLIB_CXX_OBJS = $(addprefix $(OUT_DIR)/,$(BMLIB_CXX_OBJS2))

BMLIB_OBJ_DIR = $(sort $(dir $(BMLIB_C_OBJS) $(BMLIB_CXX_OBJS)))

#in backend USEING_FULLNET=0, USING_MULTI_THREAD_ENGINE=0 STAS_GEN=1
BMLIB_STATIC = $(OUT_DIR)/bmlib.a
BMLIB_BACKEND_STATIC = $(OUT_DIR)/bmlib_backend.a
BMLIB_DYN = $(OUT_DIR)/libbmlib.so
BMLIB_BACKEND_X86_CFLGAS = -DBM_STAS_GEN $(X86_CFLAGS) -Wextra
BMLIB_BACKEND_X86_CXXFLGAS = -DBM_STAS_GEN $(X86_CXXFLAGS) -Wextra

INSTALL_LIBS += $(BMLIB_DYN)
RELEASE_LIBS += $(BMLIB_DYN)
RELEASE_HEADERS += $(BMLIB_DIR)/include/*

NNTC_INSTALL_LIBS += $(BMLIB_DYN)
NNTC_RELEASE_LIBS += $(BMLIB_DYN)
NNTC_RELEASE_HEADERS += $(BMLIB_DIR)/include/*

$(BMLIB_OBJ_DIR):
	@$(MKDIR) $@

ifeq ($(USING_CMODEL),1)
BMLIB_RUNTIME_C_OBJS = $(patsubst %.o,%_runtime.o, $(BMLIB_C_OBJS))
BMLIB_RUNTIME_CXX_OBJS = $(patsubst %.o,%_runtime.o, $(BMLIB_CXX_OBJS))
BMLIB_BACKEND_C_OBJS = $(patsubst %.o,%_backend.o, $(BMLIB_C_OBJS))
BMLIB_BACKEND_CXX_OBJS = $(patsubst %.o,%_backend.o, $(BMLIB_CXX_OBJS))
BMLIB_RUNTIME_X86_CFLGAS = -DUSING_FULLNET -DUSING_MULTI_THREAD_ENGINE $(X86_CFLAGS) -Wextra
BMLIB_RUNTIME_X86_CXXFLGAS = -DUSING_FULLNET -DUSING_MULTI_THREAD_ENGINE $(X86_CXXFLAGS) -Wextra

# only in cmodel runtime, USEING_FULLNET=1, USING_MULTI_THREAD_ENGINE=1
BMLIB_RUNTIME_STATIC = $(OUT_DIR)/bmlib_runtime.a
BMLIB_RUNTIME_DYN = $(OUT_DIR)/libbmlib_runtime.so

$(BMLIB_C_OBJS): $(OUT_DIR)/%.o:$(BM_ROOT)/%.c | $(BMLIB_OBJ_DIR)
	@echo cc -c $< -o $@
	@$(X86_CC) $(INCLUDE_DIR) $(X86_CFLAGS) -Wextra $(X86_LDFLAGS) -c $< -o $@

$(BMLIB_CXX_OBJS): $(OUT_DIR)/%.o:$(BM_ROOT)/%.cpp | $(BMLIB_OBJ_DIR)
	@echo cc -c $< -o $@
	@$(X86_CXX) $(INCLUDE_DIR) $(X86_CXXFLAGS) -Wextra $(X86_LDFLAGS) -c $< -o $@

$(BMLIB_RUNTIME_C_OBJS): $(OUT_DIR)/%_runtime.o:$(BM_ROOT)/%.c | $(BMLIB_OBJ_DIR)
	@echo cc -c $< -o $@
	@$(X86_CC) $(INCLUDE_DIR) $(BMLIB_RUNTIME_X86_CFLGAS) $(X86_LDFLAGS) -c $< -o $@

$(BMLIB_RUNTIME_CXX_OBJS): $(OUT_DIR)/%_runtime.o:$(BM_ROOT)/%.cpp | $(BMLIB_OBJ_DIR)
	@echo cc -c $< -o $@
	@$(X86_CXX) $(INCLUDE_DIR) $(BMLIB_RUNTIME_X86_CXXFLGAS) $(X86_LDFLAGS) -c $< -o $@

$(BMLIB_BACKEND_C_OBJS): $(OUT_DIR)/%_backend.o:$(BM_ROOT)/%.c | $(BMLIB_OBJ_DIR)
	@echo cc -c $< -o $@
	@$(X86_CC) $(INCLUDE_DIR) $(BMLIB_BACKEND_X86_CFLGAS) $(X86_LDFLAGS) -c $< -o $@

$(BMLIB_BACKEND_CXX_OBJS): $(OUT_DIR)/%_backend.o:$(BM_ROOT)/%.cpp | $(BMLIB_OBJ_DIR)
	@echo cc -c $< -o $@
	@$(X86_CXX) $(INCLUDE_DIR) $(BMLIB_BACKEND_X86_CXXFLGAS) $(X86_LDFLAGS) -c $< -o $@

$(BMLIB_STATIC): $(BMLIB_RUNTIME_C_OBJS) $(BMLIB_RUNTIME_CXX_OBJS) $(BMLIB_C_OBJS) $(BMLIB_CXX_OBJS) $(BMLIB_BACKEND_C_OBJS) $(BMLIB_BACKEND_CXX_OBJS)
	@echo AR $(BMLIB_RUNTIME_STATIC)
	@$(X86_AR) rcs $(BMLIB_RUNTIME_STATIC) $(BMLIB_RUNTIME_C_OBJS) $(BMLIB_RUNTIME_CXX_OBJS)
	@echo AR $(BMLIB_STATIC)
	@$(X86_AR) rcs $(BMLIB_STATIC) $(BMLIB_C_OBJS) $(BMLIB_CXX_OBJS)

$(BMLIB_RUNTIME_STATIC): $(BMLIB_RUNTIME_C_OBJS) $(BMLIB_RUNTIME_CXX_OBJS)
	@echo AR $(BMLIB_RUNTIME_STATIC)
	@$(X86_AR) rcs $(BMLIB_RUNTIME_STATIC) $(BMLIB_RUNTIME_C_OBJS) $(BMLIB_RUNTIME_CXX_OBJS)

$(BMLIB_BACKEND_STATIC): $(BMLIB_BACKEND_C_OBJS) $(BMLIB_BACKEND_CXX_OBJS)
	@echo AR $(BMLIB_BACKEND_STATIC)
	@$(X86_AR) rcs $(BMLIB_BACKEND_STATIC) $(BMLIB_BACKEND_C_OBJS) $(BMLIB_BACKEND_CXX_OBJS)

$(BMLIB_DYN): $(BMLIB_STATIC)
	@echo cc $(BMLIB_DYN)
	@$(X86_CXX) -shared -o $(BMLIB_DYN) -Wl,--whole-archive $(BMLIB_RUNTIME_STATIC) -Wl,--no-whole-archive

$(BMLIB_RUNTIME_DYN): $(BMLIB_RUNTIME_STATIC)
	@echo cc $(BMLIB_RUNTIME_DYN)
	@$(HOST_CXX) -shared -o $(BMLIB_RUNTIME_DYN) -Wl,--whole-archive $(BMLIB_RUNTIME_STATIC) -Wl,--no-whole-archive

BMLIB_DEP = $(patsubst %.o,%.d,$(BMLIB_CXX_OBJS) $(BMLIB_C_OBJS) $(BMLIB_RUNTIME_C_OBJS) $(BMLIB_RUNTIME_CXX_OBJS) $(BMLIB_BACKEND_C_OBJS) $(BMLIB_BACKEND_CXX_OBJS))
-include $(BMLIB_DEP)
else

#in device runtime, USEING_FULLNET=0, USING_MULTI_THREAD_ENGINE=0 STAS_GEN=0
$(BMLIB_C_OBJS): $(OUT_DIR)/%.o:$(BM_ROOT)/%.c | $(BMLIB_OBJ_DIR)
	@echo cc -c $< -o $@
	@$(HOST_CC) $(INCLUDE_DIR) $(HOST_CFLAGS) $(HOST_LDFLAGS) -c $< -o $@

$(BMLIB_CXX_OBJS): $(OUT_DIR)/%.o:$(BM_ROOT)/%.cpp | $(BMLIB_OBJ_DIR)
	@echo cc -c $< -o $@
	@$(HOST_CXX) $(INCLUDE_DIR) $(HOST_CXXFLAGS) $(HOST_LDFLAGS) -c $< -o $@

$(BMLIB_STATIC): $(BMLIB_C_OBJS) $(BMLIB_CXX_OBJS)
	@echo AR $(BMLIB_STATIC)
	@$(HOST_AR) rcs $(BMLIB_STATIC) $(BMLIB_C_OBJS) $(BMLIB_CXX_OBJS)

$(BMLIB_DYN): $(BMLIB_STATIC)
	@echo cc $(BMLIB_DYN)
	@$(HOST_CXX) $(HOST_CXXFLAGS) -shared -o $(BMLIB_DYN) -Wl,--whole-archive $(BMLIB_STATIC) -Wl,--no-whole-archive $(HOST_LDFLAGS)

BMLIB_DEP = $(patsubst %.o,%.d,$(BMLIB_CXX_OBJS) $(BMLIB_C_OBJS))
-include $(BMLIB_DEP)
endif

bmlib: $(BMLIB_DYN) $(BMLIB_RUNTIME_DYN)

# bmlib_backend and bmlib_runtime are for nntoolchain
bmlib_backend: $(BMLIB_BACKEND_STATIC)
bmlib_runtime: $(BMLIB_RUNTIME_DYN)
