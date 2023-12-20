#ifndef COMMON_H_
#define COMMON_H_

#include <math.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <unistd.h>
#include "bm_config.h"
#include "macro.h"
#include "memmap.h"
#include "op_code.h"
#ifdef __cplusplus
extern "C" {
#endif

#ifdef __x86_64__
void print_trace(void);
#else
#define print_trace()                                                                              \
  do {                                                                                             \
  } while (0)
#endif

#ifndef USING_CMODEL
#define hang(_ret) while (1)
#else
#define hang(_ret) exit(_ret)
#endif

#ifdef NO_PRINTF_IN_ASSERT
#define ASSERT_INFO(_cond, fmt, ...)                                                               \
  do {                                                                                             \
    if (!(_cond)) {                                                                                \
      hang(-1);                                                                                    \
    }                                                                                              \
  } while (0)
#else
#define ASSERT_INFO(_cond, fmt, ...)                                                               \
  do {                                                                                             \
    if (!(_cond)) {                                                                                \
      printf("ASSERT %s: %s: %d: %s\n", __FILE__, __func__, __LINE__, #_cond);                     \
      printf("ASSERT info: " fmt "\n", ##__VA_ARGS__);                                             \
      print_trace();                                                                               \
      hang(-1);                                                                                    \
    }                                                                                              \
  } while (0)
#endif

#define ASSERT(_cond) ASSERT_INFO(_cond, "none.")

#define ASSERT_RANGE(val, min, max)                                                                \
  ASSERT_INFO((val) >= (min) && (val) <= (max), #val "=%d must be in [%d, %d]", (val), (min), (max))

#define INLINE inline

#define UNUSED(x) (void)(x)

#define __ALIGN_MASK(x, mask) (((x) + (mask)) & ~(mask))
#define ALIGN(x, a) __ALIGN_MASK(x, (__typeof__(x))(a)-1)

#define ALIGN_SHIFT(x, s) (ALIGN(x, (1 << (s))))
#define NUM_ALIGN_SHFIT(x, s) (ALIGN_SHIFT((x), (s)) >> (s))

#define bm_min(x, y) (((x)) < ((y)) ? (x) : (y))
#define bm_max(x, y) (((x)) > ((y)) ? (x) : (y))

typedef unsigned char u8;
typedef unsigned short u16;
typedef unsigned int u32;
typedef unsigned long long u64;
typedef signed char i8;
typedef signed short i16;
typedef signed int i32;

#define __TRUE__     (1)
#define __FALSE__    (0)

typedef u32 stride_type;
typedef u32 size_type;

#define WORD_SIZE (32)
#define WORD_BITS (5)
#define WORD_MASK (0x1f)

#include "bm_api_struct.h"

typedef union {
  u32 ival;
  u16 sval;
  u8 cval;
} IC_VAL;

typedef union {
  i8 i8val;
  i16 i16val;
  i32 i32val;
  u8 u8val;
  u16 u16val;
  u32 u32val;
  float fval;
}VALUE;

typedef union {
  int ival;
  float fval;
} IF_VAL;

typedef struct {
  u32 where; /* low bit index. */
  u32 len;   /* bit length. */
} reg_id_t;

typedef enum {
  STORAGE_MODE_1N_FP32 = 0,
  STORAGE_MODE_1N_INT8 = 1,
  STORAGE_MODE_1N_INT16 = 2,
  STORAGE_MODE_2N_INT16 = 3,
  STORAGE_MODE_4N_INT8 = 4,
  STORAGE_MODE_2IC_FP32 = 5, // special for 2IC weight
  STORAGE_MODE_4N_4IC_4OC = 6,
  //  STORAGE_MODE_1N                   = 0,
  //  STORAGE_MODE_2N                   = 1,
  //  STORAGE_MODE_4N                   = 2,
  STORAGE_MODE_UNINITILIZED,
  STORAGE_MODE_END
} TENSOR_STORAGE_MODE;

typedef enum {
  OP_8BIT = 8,
  OP_16BIT = 16,
  OP_32BIT = 32,
} OP_BITWIDTH;

typedef enum {
  INT8 = 0,
  FP16 = 1,
  FP32 = 2,
  INT16 = 3,
  INT32 = 4,
  BFP16 = 5,
  INT64 = 6,
  PRECISION_MODE_END
} PRECISION_MODE;

typedef enum {
    ALIGNED = 0,
    COMPACT = 1,
    BIAS = 2,
    STRIDE = 3,
    STORE_FMT_END
} STORE_FMT;

#define INT8_PER_32BITS 4
#define INT16_PER_32BITS 2

#define SIZE_INT16_1N 2
#define SIZE_INT8_1N 1
#define SIZE_INT8_4N 4
#define SIZE_INT16_2N 4

typedef u32 tuple4_u32[4];

typedef struct tensor_info {
  u32 n, c, h, w;
  u32 w_stride, n_stride, c_stride, h_stride;
  u32 h_shift, w_shift;
  u32 address;
  u32 data_format;
  u32 neuron_matrix;     // 0: neuron, 1: matrix
  u32 matrix_col_margin; // the magin is not 0, when column_num%w_param!=0
  u32 lsize;             // local size of tensor
  TENSOR_STORAGE_MODE storage_mode;
  PRECISION_MODE prec;
} TENSOR_INFO;

typedef struct shape {
  u16 n, c, h, w;
} local_shape_t;

#define INT8_SIZE 1
#define FLOAT_SIZE 4
#define FLOAT_BITWIDTH 32
#define GET_U64(U32_H, U32_L) (((u64)(U32_H) << 32) | (u64)(U32_L))

typedef enum { NODECHIP_REG = 0, HOST_REG = 1 } REG_TYPE;

typedef enum {
  ENGINE_BD = 0,
  ENGINE_GDMA = 1,
  ENGINE_CDMA = 2,
  ENGINE_GDE = 3,
  ENGINE_NMS = 4,
  ENGINE_END
} ENGINE_ID;

typedef enum {
  ACTIVE_NULL = 0,
  ACTIVE_RELU = 1,
  ACTIVE_PRELU = 2,
} ACTIVE_TYPE_e;

typedef enum {
  CAFFE_SUPPORT = 0,
  TENSORFLOW_SUPPORT = 1,
  CAFFE_NEAREST = 2,
  TENSORFLOW_NEAREST = 3
} PLATFORM_SUPPORT;

typedef struct gdma_cmd_node_info_s {
  int n;
  int c;
  int h;
  int w;
  int direction;
  int src_format;
  int dest_format;
  bool setted;
} gdma_cmd_node_info_t;

typedef struct inst_profile {
  unsigned long long cycle;
  unsigned long long gdma_size;
  int gdma_direction;
  int src_format;
  int dst_format;
  double power;
  double compute_ability;
  bool b_gdma_use_l2;
} INST_PROFILE;

typedef struct cmd_id_node {
  unsigned int bd_cmd_id;
  unsigned int gdma_cmd_id;
  unsigned int nms_cmd_id;
  bool in_parallel_state;
#if defined(BM_STAS_GEN) || defined(BM_TV_GEN)
  unsigned long long cycle_count;
  unsigned long long cur_op_cycle;
#endif
#ifdef BM_STAS_GEN
  char cmd_name[16];
  char name_prefix[64];
  gdma_cmd_node_info_t gdma_cmd_info;
  INST_PROFILE inst_profile;
#endif
} CMD_ID_NODE;

#ifdef BM_STAS_GEN
static inline void set_gdma_cmd_info(CMD_ID_NODE *pid_node, int n, int c, int h, int w,
                                     int direction, int src_format, int dest_format) {
  gdma_cmd_node_info_t *the_info = &pid_node->gdma_cmd_info;
  the_info->n = n;
  the_info->c = c;
  the_info->h = h;
  the_info->w = w;
  the_info->direction = direction;
  the_info->src_format = src_format;
  the_info->dest_format = dest_format;
  the_info->setted = true;
}
#else
#define set_gdma_cmd_info(...)                                                                     \
  {}
#endif

void *create_cmd_id_node();
void destroy_cmd_id_node(void *pid_node);

#ifdef BM_STAS_GEN
void set_cmd_id_cycle(void *pid_node, int val);
int get_cmd_id_cycle(void *pid_node);
#endif

#ifdef USING_CMODEL
#define GLOBAL_MEM_SIZE(node_idx) cmodel_get_global_mem_size(node_idx)
#else
#define GLOBAL_MEM_SIZE (CONFIG_GLOBAL_MEM_SIZE)
#endif

typedef void *P_COMMAND;

int array_cmp_abs(float *p_exp, float *p_got, int len, const char *info_label, int exactly_matched,
                  float delta);

int array_cmp_rel(float *p_exp, float *p_got, int len, const char *info_label, float delta);

int array_cmp_int8(char *p_exp, char *p_got, int len, const char *info_label, float delta);

int array_cmp_fix8b(void *p_exp, void *p_got,
                    int sign, // 0: unsigned, 1: signed
                    int len, const char *info_label, int delta);

int array_cmp_fix16b(void *p_exp, void *p_got,
                     int sign, // 0: unsigned, 1: signed
                     int len, const char *info_label, int delta);

int array_cmp(float *p_exp, float *p_got, int len, const char *info_label, float delta);

int array_cmp_int(int *p_exp, int *p_got, int len, const char *info_label);

void dump_array_file(char *file, int row_num, int col_num, int transpose, float *parr);

INLINE static int get_num_shift(int num) {
  switch (num) {
    case 1:
      return 0;
    case 2:
      return 1;
    case 4:
      return 2;
    case 8:
      return 3;
    case 32:
      return 5;
    case 64:
      return 6;
    default:
      ASSERT(0);
  }
  return 0;
}

INLINE static int ceiling_func(int numerator, int denominator) {
  return (numerator + denominator - 1) / denominator;
}

INLINE static int ceiling_func_shift(int numerator, int shift) {
  return (numerator + (1 << shift) - 1) >> shift;
}

INLINE static int calc_offset(int *shape, int *offset) {
  return ((offset[0] * shape[1] + offset[1]) * shape[2] + offset[2]) * shape[3] + offset[3];
}

// All the size are in the units of bytes
INLINE static int get_index_csize_global(int h, int w, int index_bitwidth) {
  int size = h * w * index_bitwidth;
  // 32 bit align
  return (((size >> 5)) + ((size & 0x1f) != 0)) * FLOAT_SIZE;
}

INLINE static int get_index_cstride_global(int h, int w, int index_bitwidth) {
  int size = h * w * index_bitwidth;
  // 32 bit align
  return (((size >> 5)) + ((size & 0x1f) != 0)) * FLOAT_BITWIDTH / index_bitwidth;
}

INLINE static int get_addr_by_local_idx_offset(int localmem_idx, int localmem_offset) {
  return ((localmem_idx << LOCAL_MEM_ADDRWIDTH) + localmem_offset + LOCAL_MEM_START_ADDR);
}

INLINE static int get_local_idx_from_addr(int addr) {
  return ((addr | LOCAL_MEM_START_ADDR) - LOCAL_MEM_START_ADDR) >> LOCAL_MEM_ADDRWIDTH;
}

INLINE static int get_local_offset_from_addr(int addr) {
  return ((addr | LOCAL_MEM_START_ADDR) - LOCAL_MEM_START_ADDR) & ((1 << LOCAL_MEM_ADDRWIDTH) - 1);
}

INLINE static int get_index_width(int pool_size) {
  if (pool_size == 1) {
    return 1;
  }
  if (pool_size <= 16) {
    return 4;
  } else if (pool_size <= 256) {
    return 8;
  } else {
    return 16;
  }
}

INLINE static int greatest_common_divisor(int m, int n) {
  int tmp;
  while (m) {
    tmp = m;
    m = n % m;
    n = tmp;
  }
  return n;
}

INLINE static int least_common_multiple(int m, int n) {
  return m / greatest_common_divisor(m, n) * n;
}

#define GDMA_ENGINE_CMD_ALIGNED_BIT 7
#define BDC_ENGINE_CMD_ALIGNED_BIT 7

INLINE static int get_addr_by_featuremap_shift(int src_addr, int feature_map_shift, int c_stride) {
  int local_idx = (src_addr - LOCAL_MEM_START_ADDR) >> LOCAL_MEM_ADDRWIDTH;
  int local_offset = (src_addr - LOCAL_MEM_START_ADDR) & ((1 << LOCAL_MEM_ADDRWIDTH) - 1);
  if (feature_map_shift >= 0) {
    local_offset = local_offset + ((local_idx + feature_map_shift) >> NPU_SHIFT) * c_stride;
    local_idx = (local_idx + feature_map_shift) & NPU_MASK;
  } else {
    feature_map_shift = -feature_map_shift;
    int backward_step = feature_map_shift >> NPU_SHIFT;
    if ((feature_map_shift & NPU_MASK) > local_idx) {
      backward_step++;
      local_idx = (local_idx + NPU_NUM - (feature_map_shift & NPU_MASK));
    } else {
      local_idx = (local_idx - (feature_map_shift & NPU_MASK));
    }
    local_offset = local_offset - backward_step * c_stride;
  }
  return ((local_idx << LOCAL_MEM_ADDRWIDTH) + local_offset + LOCAL_MEM_START_ADDR);
}

INLINE static OP_BITWIDTH get_precision_bitwidth(PRECISION_MODE precision) {
  switch (precision) {
    case INT8:
      return OP_8BIT;
    case FP16:
    case INT16:
      return OP_16BIT;
    default:
      return OP_32BIT;
  }
}

INLINE static int get_precision_bytes(PRECISION_MODE precision) {
  return 32 / (int)get_precision_bitwidth(precision);
}

INLINE static u32 offset_to_laddr(u32 lmem_idx, u32 lmem_offset) {
  return LOCAL_MEM_START_ADDR + (lmem_idx << LOCAL_MEM_ADDRWIDTH) + lmem_offset;
}

INLINE static int get_hw_align_num(PRECISION_MODE precision) {
  return EU_NUM * (32 / get_precision_bitwidth(precision));
}

INLINE static int get_neuron_local_cstride(int h, int w, OP_BITWIDTH bw) {
  return ALIGN(h * w, EU_NUM * (32 / bw));
}

INLINE static int get_neuron_local_csize(int h, int w, OP_BITWIDTH bw) {
  return get_neuron_local_cstride(h, w, bw) * (bw / 8);
}

INLINE static int get_neuron_local_nstride(int c_stride, int c, int local_mem_idx) {
  return ceiling_func(local_mem_idx + c, NPU_NUM) * c_stride;
}

INLINE static int get_neuron_local_size(int n, int c, int h, int w, OP_BITWIDTH bw,
                                        int local_mem_idx) {
  int c_per_npu = ceiling_func_shift(local_mem_idx + c, NPU_SHIFT);
  return n * c_per_npu * get_neuron_local_csize(h, w, bw);
}

INLINE static int get_neuron_align_addr(int addr, OP_BITWIDTH bw) {
  return ALIGN(addr, EU_NUM * (32 / bw)) * bw / 8;
}

// Deprecated
INLINE static int get_nstride_local(int c_stride, int c, int local_mem_idx) {
  // 64 neurons align
  ASSERT(0);
  return ceiling_func_shift(local_mem_idx + c, NPU_SHIFT) * c_stride;
}

// Deprecated
INLINE static int get_neuron_csize_local(int h, int w) {
  ASSERT(0);
  int size = h * w;
  // EU_NUM neurons align
  return ALIGN(size, EU_NUM) * FLOAT_SIZE;
}

// Deprecated
INLINE static int get_neuron_csize_local_fix8b(int h, int w) {
  ASSERT(0);
  int size = h * w;
  // EU_NUM neurons align
  return ALIGN(size, EU_NUM) * sizeof(char);
}

// Deprecated
INLINE static int get_neuron_csize_local_fix16b(int h, int w) {
  ASSERT(0);
  int size = h * w;
  // EU_NUM neurons align
  return ALIGN(size, EU_NUM) * sizeof(short);
}

// Deprecated
INLINE static int get_csize_local_aligned_to_128bits(int h, int w) {
  ASSERT(0);
  int size = h * w;
  // EU_NUM neurons align
  return ALIGN(size, 4) * FLOAT_SIZE;
}

// Deprecated
INLINE static int get_cstride_local(int h, int w) {
  ASSERT(0);
  int size = h * w;
  // EU_NUM neurons align
  return ALIGN(size, EU_NUM);
}

// Deprecated
INLINE static int get_align_tensor_size(bm_tensor_4d_t shape) {
  ASSERT(0);
  int c_per_npu = ceiling_func_shift(shape.c, NPU_SHIFT);
  return shape.n * c_per_npu * get_neuron_csize_local(shape.h, shape.w);
}

// Deprecated
INLINE static int get_local_shape_size_align(local_shape_t shape) {
  ASSERT(0);
  int c_per_npu = ceiling_func_shift(shape.c, NPU_SHIFT);
  return (int)shape.n * c_per_npu * get_neuron_csize_local((int)shape.h, (int)shape.w);
}

// Deprecated
INLINE static int addr_EU_align(int addr) {
  ASSERT(0);
  addr = (addr + FLOAT_SIZE - 1) / FLOAT_SIZE; // for 16bits consideration
  return ALIGN(addr, EU_NUM) * FLOAT_SIZE;
}

INLINE static int nsecs_computing(int memsize_per_n, int input_n, int mem_capacity) {
  int mem_need = input_n * memsize_per_n;
  ASSERT(mem_need > mem_capacity);
  int nsecs = mem_need / mem_capacity;
  int nslice = (input_n + nsecs - 1) / nsecs;
  while (nslice * memsize_per_n > mem_capacity) {
    ASSERT(nsecs < input_n);
    if (nslice == 2) {
      nsecs = input_n;
    } else {
      nsecs++;
    }
    nslice = (input_n + nsecs - 1) / nsecs;
  }
  return nsecs;
}

#define LOCAL_MEM_BANKS 8
#define LOCAL_BANK_SIZE (LOCAL_MEM_SIZE / LOCAL_MEM_BANKS)
#define LAST_INI_REG_VAL 0x76125438
#define EU_CONFIG_LEN 1

INLINE static int pointer_wrap_around(u32 cur_pointer, int step, int len_bit_width) {
  u32 max_len = (1 << len_bit_width);
  u32 new_pointer = 0;

  new_pointer = cur_pointer + step;
  if (new_pointer >= max_len)
    new_pointer -= max_len;

  return (int)new_pointer;
}

typedef enum host_cdma_dir { HOST2CHIP, CHIP2HOST, CHIP2CHIP } HOST_CDMA_DIR;

INLINE void init_cmd_id_node(CMD_ID_NODE *p_id_node) {
  p_id_node->bd_cmd_id = 0;
  p_id_node->gdma_cmd_id = 0;
  p_id_node->nms_cmd_id = 0;
}

int conv_coeff_storage_convert(float *coeff_orig, float **coeff_reformat, unsigned int oc,
                               unsigned int ic, unsigned int kh, unsigned int kw,
                               unsigned int npu_num);

enum fw_downlod_stage { FW_START = 0, DDR_INIT_DONE = 1, DL_DDR_IMG_DONE = 2 };

INLINE static int gcd(int m, int n) {
  int tmp;
  while (m) {
    tmp = m;
    m = n % m;
    n = tmp;
  }
  return n;
}

INLINE static int lcm(int m, int n) { return m / gcd(m, n) * n; }

INLINE static void pipeline_move(int *array, int num) {
  for (int i = num - 1; i > 0; i--) {
    array[i] = array[i - 1];
  }
}

#ifdef USING_CMODEL
extern int get_cur_nodechip_idx();
#define GET_SHARE_MEM_ADDR(offset) cmodel_get_share_memory_addr(offset, get_cur_nodechip_idx())
#else
#define GET_SHARE_MEM_ADDR(offset) (volatile u32 *)(SHARE_MEM_START_ADDR + (offset)*4)
#endif

struct bm_api_append_data {
  /* ret_val = 0 SUCCESS; others FAILURE */
  u32 ret_val;
  /* how much us used by tpu to process this api */
  u32 tpu_process_time;
};

#ifdef __cplusplus
}
#endif
#endif /* COMMON_H_ */
