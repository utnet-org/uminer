#ifndef BMCPU_INTERNAL_H
#define BMCPU_INTERNAL_H

#if defined(__cplusplus)
extern "C" {
#endif
#ifdef WIN32
#include "..\..\common\bm1684\include_win\common_win.h"
#else
#include "../../common/bm1684/include/common.h"
#endif
#include "bmlib_runtime.h"
#include "message.h"

#define BMCPU_RUNTIME_LOG_TAG "bmcpu_runtime"

#define LOGLIB_ITEM_MAX_SIZE 128
#define LOGLIB_ITEM_MAX_SIZE_SHIFT 7 /* (2 << 7) = 128 */


#define BM_SVC_PROTECT_START_ADDR       (0x314100000ULL)
#define BM_SVC_PROTECT_END_ADDR         (0x314200000ULL)
#define RSA_PUBLIC_KEY_ADDR             BM_SVC_PROTECT_START_ADDR
#define RSA_PRIVATE_KEY_ADDR_RAW        (RSA_PUBLIC_KEY_ADDR + 0x1000)
#define RSA_PRIVATE_KEY_ADDR            (RSA_PRIVATE_KEY_ADDR_RAW + 0x1000)
#define AES_KEY_ADDR                    (RSA_PRIVATE_KEY_ADDR + 0x1000)
#define RSA_CIPHER_ADDR                 (AES_KEY_ADDR + 0x1000)
#define RSA_DIGEST_ADDR                 (RSA_CIPHER_ADDR + 0x1000)
#define BM_GEN_TEST_PUBLIC_ADDR         (BM_SVC_PROTECT_START_ADDR+0X10000)
#define BM_GEN_TEST_PRIVITE_ADDR        (BM_SVC_PROTECT_START_ADDR+0X11000)


typedef enum arm9_fw_mode {
    FW_PCIE_MODE,
    FW_SOC_MODE,
    FW_MIX_MODE
} bm_arm9_fw_mode;

/* bm api message struct
 * represent the api information to be sent to driver
 */
typedef struct bm_api_data {
    bm_api_id_t api_id;
    u64         api_handle;
    u64         data;
    int         timeout;
} bm_api_data_t;
/* bm api message struct
 * represent the api information to be sent to driver
 */
/* bm api message struct
 * represent the api information to be sent to driver
 */
typedef struct bm_api_ext {
    bm_api_id_t api_id;
    const u8 *  api_addr;
    u32         api_size;
    u64         api_handle;
} bm_api_ext_t;

typedef struct {
    u64 u64LogIndex;
} log_runtime_info_t;

#pragma pack(1)
struct sgcpu_comm_data {
    char *data;
    int len;
};

typedef struct bm_api_start_cpu {
    u32 flag;
} bm_api_start_cpu_t;

typedef struct bm_api_open_process {
    u32 flags;
} bm_api_open_process_t;

typedef struct bm_api_set_time {
    u32 tv_sec;
    u32 tv_usec;
    u32 tz_minuteswest;
    u32 tz_dsttime;
} bm_api_set_time_t;

typedef struct bm_api_cpu_load_library_internal {
    u64   process_handle;
    u8 *  library_path;
    void *library_addr;
    u32   size;
    u8    library_name[28];
} bm_api_cpu_load_library_internal_t;

typedef struct bm_api_cpu_exec_func_internal {
    u64 process_handle;
    u8 *function_name;
    u8 *function_param;
    u32 param_size;
    u32 opt;
    u8  local_function_name[28];
    u8  local_function_param[80];
} bm_api_cpu_exec_func_internal_t;

typedef struct bm_api_cpu_load_library {
    u64 process_handle;
    u8 *library_path;
} bm_api_cpu_load_library_t;

typedef struct bm_api_set_log {
    u64 log_addr;
    u32 log_size;
    u32 log_opt;
} bm_api_set_log_t;

typedef struct bm_api_cpu_exec_func {
    u64 process_handle;
    u8 *function_name;
    u8 *function_param;
    u32 param_size;
    u32 opt;
} bm_api_cpu_exec_func_t;

typedef struct bm_api_cpu_map_addr {
    u64 process_handle;
    u8 *phyaddr;
    u32 size;
} bm_api_cpu_map_addr_t;

typedef struct bm_api_close_process {
    u64 process_handle;
} bm_api_close_process_t;

typedef struct bm_api_get_log {
    u64 process_handle;
    u64 reserve;
} bm_api_get_log_t;

typedef struct bm_aeskey {
    u64   aeskey_addr;
    u32   size_aeskey;
} bm_aeskey_t;

typedef struct bm_p2_pubkey {
    u64   p2_addr;
    u32   size_p2;
    u64   pubkey_addr;
    u32   size_pubkey;
} bm_p2_pubkey_t;

typedef struct bm_p2_digest {
    u32   size_p2;
    u32   size_digest;
    u64   p2_addr;
    u64   pubkey_addr;
} bm_p2_digest_t;


#pragma pack()
#ifdef __linux__
bm_status_t bmcpu_start_mix_cpu(bm_handle_t handle,
                            char *      boot_file,
                            char *      core_file);
bm_status_t bm_setup_veth(bm_handle_t handle);
bm_status_t bm_remove_veth(bm_handle_t handle);
bm_status_t bm_force_reset_bmcpu(bm_handle_t handle);
bm_status_t bmcpu_set_arm9_fw_mode(bm_handle_t handle, bm_arm9_fw_mode mode);
bm_status_t bmcpu_load_boot(bm_handle_t handle, char *boot_file);
bm_status_t bmcpu_load_kernel(bm_handle_t handle, char *kernel_file);
bm_status_t bm_trigger_a53(bm_handle_t handle, int id);
bm_status_t bm_connect_a53(bm_handle_t handle, int timeout);
int bm_write_data(bm_handle_t handle, char *buf, int len);
int bm_read_data(bm_handle_t handle, char *buf, int len);
int bm_write_msg(bm_handle_t handle, char *buf, int len);
int bm_read_msg(bm_handle_t handle, char *buf, int len);
#endif
bm_status_t bm_query_api_data(bm_handle_t handle,
                              bm_api_id_t api_id,
                              u64         api_handle,
                              u64 *       data,
                              int         timeout);
#if defined(__cplusplus)
}
#endif

#endif /* BMCPU_INTERNAL_H */
