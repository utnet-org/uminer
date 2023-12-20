#ifndef BM_DEVICE_H_
#define BM_DEVICE_H_

#ifdef USING_CMODEL

#include "bmlib_runtime.h"
#include "bmlib_mmpool.h"
#include "common.h"
#include <map>
#include <queue>
#include <atomic>

#define DEVICE_SYNC_MARKER 0xFFFF

struct thread_api_info {
    pthread_t thd_id;
    u32 last_seq;
    u32 cpl_seq;
};

struct api_queue_entry {
    pthread_t thd_id;
    u32 thd_seq;
    u32 dev_seq;
};

class bm_device {
  public:
    bm_device(int _dev_id);
    ~bm_device();
    int bm_device_id() {
        return dev_id;
    }
    bm_status_t bm_device_send_api(
        bm_api_id_t api_id, const u8 *api, u32 size);
    bm_status_t bm_device_sync();
    bm_status_t bm_device_thread_sync();
    u64 bm_device_alloc_mem(u32 size);
    void bm_device_free_mem(u64 addr);
    bm_status_t bm_device_memcpy_s2d(bm_device_mem_t dst, void *src);
    bm_status_t bm_device_memcpy_d2s(void *dst, bm_device_mem_t src);
    u64 bm_device_arm_reserved_req();
    void bm_device_arm_reserved_rel();

  private:
    /* add lock */
    void _write_share_mem(u32 offset, u32 data);
    void _write_share_reg(u32 idx, u32 data);
    u32 _read_share_reg(u32 idx);

    u32 _poll_message_fifo_cnt();
    void copy_message_to_sharemem(const u32 *src_msg_buf,
                                  u32 *wp, u32 size, u32 api_id);

    bm_status_t bm_send_quit_message();
    void bm_wait_fwinit_done();

    bm_status_t bm_alloc_arm_reserved();
    void bm_free_arm_reserved();
    bm_status_t bm_alloc_instr_reserved();
    void bm_free_instr_reserved();
    bm_status_t bm_alloc_iommu_reserved();
    void bm_free_iommu_reserved();
    bm_status_t bm_malloc_device_dword(bm_device_mem_t *pmem, int cnt);
    void bm_free_device(bm_device_mem_t mem);
    bm_status_t bm_init_l2_sram();

    thread_api_info *bm_get_thread_api_info(pthread_t thd_id);
    bm_status_t bm_add_thread_api_info(pthread_t thd_id);
    bm_status_t bm_remove_thread_api_info(pthread_t thd_id);

    static void *bm_msg_done_poll(void *arg);

    int                dev_id;
    // u64                device_mem_size;
    bm_mem_pool        device_mem_pool;
    bm_device_mem_t    instr_reserved_mem;
    bm_device_mem_t    arm_reserved_dev_mem;
    bm_device_mem_t    iommu_reserved_dev_mem;
    /* arm reserved memory lock */
    pthread_mutex_t    arm_reserved_lock;

    pthread_mutex_t    api_lock;
    pthread_t msg_poll_thread;

    std::map < pthread_t, thread_api_info> thread_api_table;
    std::queue<api_queue_entry> pending_api_queue;
    u32 device_sync_last;
    std::atomic<u32> device_sync_cpl;
};

class bm_device_manager {
  public:
    bm_device *get_bm_device(int dev_id);
    static bm_device_manager *get_dev_mgr();
    static void destroy_dev_mgr();

  private:
    bm_device_manager(int max_dev_cnt);
    ~bm_device_manager();

    int dev_cnt;
    int max_dev_cnt;
    bm_device **bm_dev_list;
    static bm_device_manager *bm_dev_mgr;
    static pthread_mutex_t init_lock;
};

class bm_device_manager_control {
  public:
    bm_device_manager_control() {
    }
    ~ bm_device_manager_control() {
        bm_device_manager::destroy_dev_mgr();
    }
};

#endif
#endif
