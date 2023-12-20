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

BM_ROOT     ?= $(PWD)
export      BM_ROOT
SHELL       := /bin/bash
##############################################################################
# Include common definitions.
##############################################################################
include $(BM_ROOT)/../../Makefile.def
OUT_DIR     ?= $(BM_ROOT)
#$(BM_ROOT)/build/tools

SRCS_C      = $(wildcard *.c)
SRCS_CXX    = $(wildcard *.cpp)

TGTS_C      = $(addprefix $(OUT_DIR)/, $(notdir $(SRCS_C:.c=)))
TGTS_CXX    = $(addprefix $(OUT_DIR)/, $(notdir $(SRCS_CXX:.cpp=)))

ifeq ($(CHIP_NAME),bm1684)
INCLUDE_DIR = -I$(BM_ROOT)/common/bm1684/include \
              -I$(BM_ROOT)/firmware_core/bm1684/include \
              -I$(BM_ROOT)/firmware_core/bm1684/src \
              -I$(BM_ROOT)/driver/include \
              -I$(BM_ROOT)/c_model/bm1684/include \
              -I$(BM_ROOT)/c_model/bm1684/src
else ifeq ($(CHIP_NAME),bm1686)
INCLUDE_DIR = -I$(BM_ROOT)/common/bm1686/include \
              -I$(BM_ROOT)/firmware_core/bm1686/include \
              -I$(BM_ROOT)/firmware_core/bm1686/src \
              -I$(BM_ROOT)/driver/include \
              -I$(BM_ROOT)/c_model/bm1686/include \
              -I$(BM_ROOT)/c_model/bm1686/src
endif

CFLAGS      += $(INCLUDE_DIR) -DBUILD_LIBBMDNN -Wno-unused-function 
CXXFLAGS    += $(INCLUDE_DIR) -DBUILD_LIBBMDNN -Wno-unused-function -DDIFF_DUMP_DATA_PATH=\"/home/jinli/Git/bmdnn/build/test/\"
LDFLAGS     += -L$(BM_ROOT)/build -lstdc++
              
all: $(OUT_DIR) $(TGTS_C) $(TGTS_CXX) 

$(OUT_DIR):
	$(MKDIR) $@

$(OUT_DIR)/%: %.c
	$(CC) $(CFLAGS) $< -o $@ $< $(LDFLAGS)

$(OUT_DIR)/%: %.cpp
	$(CXX) $(CXXFLAGS) -o $@ $< $(LDFLAGS)

clean:
	$(RM) $(TGTS_C) $(TGTS_CXX)
#	$(RM) $(OUT_DIR)

