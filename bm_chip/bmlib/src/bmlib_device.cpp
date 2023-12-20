#ifdef USING_CMODEL
#include "bmlib_internal.h"
#include "bmlib_memory.h"
#include "cmodel_common.h"
#include "cmodel_runtime.h"
#include "l2_sram_table.h"
#include "message.h"
#ifdef BM_TV_GEN
    #include "bm_tv_gen_util.h"
#endif

#define CONFIG_HOST_STAGMEM_SIZE 0x100000

extern pthread_mutex_t msg_done_mutex[MAX_NODECHIP_NUM];
extern pthread_cond_t msg_done_cond[MAX_NODECHIP_NUM];
extern u32 msg_done[MAX_NODECHIP_NUM];

unsigned long long get_global_mem_size()
{
    auto env = std::getenv("CMODEL_GLOBAL_MEM_SIZE");
    if (env)
    {
        unsigned long long v;
        try {
            v = std::stoll(env);
        } catch (std::invalid_argument &) {
            printf("invalid CMODEL_GLOBAL_MEM_SIZE \"%s\"\n", env);
            throw;
        }
        printf("global mem size from env %lld\n", v);
        return v;
    } else {
        return CONFIG_GLOBAL_MEM_SIZE;
    }
}

bm_device::bm_device(int _dev_id)
  : dev_id(_dev_id),
    device_mem_pool(get_global_mem_size()),
    device_sync_last(0),
    device_sync_cpl(0) {
  thread_api_table.clear();

  printf("begin to cmodel init...\n");
  if (cmodel_init_without_l2_fill(dev_id, get_global_mem_size())
      != CMODEL_OK) {
    printf("BM: cmodel_init failed\n");
    exit(-1);
  }
  printf("l2_fill complete\n");

  set_cur_nodechip_idx(0);

  printf("npu num = %d\n", NPU_NUM);
  printf("eu num = %d\n", EU_NUM);
  printf("global mem size = %lld\n", GLOBAL_MEM_SIZE(dev_id));

  cmodel_nodechip_runtime_init(dev_id);

  // ctx->device_mem_size = cmodel_get_global_mem_size(devid);

  BM_CHECK_RET(bm_alloc_instr_reserved());
  BM_CHECK_RET(bm_alloc_arm_reserved());
  BM_CHECK_RET(bm_init_l2_sram());
  BM_CHECK_RET(bm_alloc_iommu_reserved());
  bm_wait_fwinit_done();

  pthread_mutex_init(&api_lock, nullptr);
  pthread_mutex_init(&arm_reserved_lock, nullptr);
  // init msg poll thread
  pthread_create(&msg_poll_thread, nullptr, bm_msg_done_poll, this);
}

bm_device::~bm_device() {
  // printf("destroy device %d\n", dev_id);
  pthread_cancel(msg_poll_thread);
  pthread_join(msg_poll_thread, nullptr);

  bm_free_arm_reserved();
  bm_free_instr_reserved();
  bm_free_iommu_reserved();

  bm_send_quit_message();

  cmodel_nodechip_runtime_exit(dev_id);
  cmodel_deinit(dev_id);
}

u64 bm_device::bm_device_alloc_mem(u32 size) {
  return device_mem_pool.bm_mem_pool_alloc(size);
}

void bm_device::bm_device_free_mem(u64 addr) {
  device_mem_pool.bm_mem_pool_free(addr);
}

void bm_device::_write_share_mem(u32 offset, u32 data) {
  u32 * write_addr = cmodel_get_share_memory_addr(offset, dev_id);
  *write_addr = data;
#ifdef BM_TV_GEN
  bm_wr_tv_dump_reg(SHARE_MEM_START_ADDR + offset * 4, data, HOST_REG);
#endif
}

void bm_device::_write_share_reg(u32 idx, u32 data) {
  cmodel_write_share_reg(idx, data, dev_id);
#ifdef BM_TV_GEN
  bm_wr_tv_dump_reg(SHARE_REG_BASE_ADDR + idx * 4, data, HOST_REG);
#endif
}

u32 bm_device::_read_share_reg(u32 idx) {
  return cmodel_read_share_reg(idx, dev_id);
}

u32 bm_device::_poll_message_fifo_cnt() {
  u32 wp, rp;
  wp = _read_share_reg(SHARE_REG_MESSAGE_WP);
  rp = _read_share_reg(SHARE_REG_MESSAGE_RP);

  u32 wp_tog = wp >> SHAREMEM_SIZE_BIT;
  u32 rp_tog = rp >> SHAREMEM_SIZE_BIT;
  u32 wp_offset = wp - (wp_tog << SHAREMEM_SIZE_BIT);
  u32 rp_offset = rp - (rp_tog << SHAREMEM_SIZE_BIT);

  if (wp_tog == rp_tog)
    return (1 << SHAREMEM_SIZE_BIT) - wp_offset + rp_offset;
  else
    return rp_offset - wp_offset;
}

void bm_device::copy_message_to_sharemem(const u32 *src_msg_buf,
    u32 *wp, u32 size, u32 api_id) {
  u32 cur_wp = _read_share_reg(SHARE_REG_MESSAGE_WP);
  *wp = cur_wp;

  _write_share_mem(cur_wp & SHAREMEM_MASK, api_id);
  u32 next_wp = pointer_wrap_around(cur_wp, 1, SHAREMEM_SIZE_BIT) &
        SHAREMEM_MASK;
  _write_share_mem(next_wp & SHAREMEM_MASK, size);

  if (api_id == BM_API_QUIT)
    return;
  for (u32 idx = 0; idx < size; idx++) {
    next_wp = pointer_wrap_around(*wp, 2 + idx, SHAREMEM_SIZE_BIT);
    _write_share_mem(next_wp & SHAREMEM_MASK, src_msg_buf[idx]);
  }
}

bm_status_t bm_device::bm_device_send_api(
      bm_api_id_t api_id, const u8 *api, u32 size) {
  pthread_mutex_lock(&api_lock);
  pthread_t thd_id = pthread_self();
  // get thread api info
  thread_api_info *thd_api_info = bm_get_thread_api_info(thd_id);
  if (!thd_api_info) {
    bm_add_thread_api_info(thd_id);
    thd_api_info = bm_get_thread_api_info(thd_id);
  }

  // update thread api last seq
  thd_api_info->last_seq++;

  // add api queue entry into fifo
  pending_api_queue.push({thd_id, thd_api_info->last_seq, 0});
  /*
  printf("SEND API: thread %lu --- seq_id %d\n",
          thd_id, thd_api_info->last_seq);
  */

  u32 fifo_empty_number = API_MESSAGE_EMPTY_SLOT_NUM * (size/sizeof(u32) + 2);
  while (_poll_message_fifo_cnt() <= fifo_empty_number) {
  }

  u32 wp;
  copy_message_to_sharemem(reinterpret_cast < const u32 *>(api), &wp,
        size/sizeof(u32), api_id);
  u32 next_wp = pointer_wrap_around(wp, size/sizeof(u32) + 2,
        SHAREMEM_SIZE_BIT);
  _write_share_reg(SHARE_REG_MESSAGE_WP, next_wp);

  pthread_mutex_unlock(&api_lock);
  return BM_SUCCESS;
}

bm_status_t bm_device::bm_device_sync() {
  pthread_mutex_lock(&api_lock);
  u32 dev_sync_last = ++device_sync_last;
  pending_api_queue.push({DEVICE_SYNC_MARKER, 0, dev_sync_last});
  pthread_mutex_unlock(&api_lock);

  printf("SYNC DEVICE API: device last seq %d\n", dev_sync_last);
  while (dev_sync_last != device_sync_cpl) {
  }

  #ifdef BM_TV_GEN
  bm_read32_wait_eq_tv(SHARE_REG_BASE_ADDR + SHARE_REG_MESSAGE_RP * 4,
      cmodel_read_share_reg(SHARE_REG_MESSAGE_WP, dev_id), 32, 0, HOST_REG);
  #endif
  // while (_poll_message_fifo_cnt() != (1 << SHAREMEM_SIZE_BIT));
  return BM_SUCCESS;
}

bm_status_t bm_device::bm_send_quit_message() {
  printf("BMLIB Send Quit Message\n");
  bm_device_send_api((bm_api_id_t)BM_API_QUIT, nullptr, 0);
  sleep(1);
  return BM_SUCCESS;
}

void bm_device::bm_wait_fwinit_done() {
#ifdef BM_TV_GEN
  bm_read32_wait_eq_tv(SHARE_REG_BASE_ADDR + SHARE_REG_FW_STATUS * 4,
                  LAST_INI_REG_VAL, 32, 0, HOST_REG);
#endif
  while (_read_share_reg(SHARE_REG_FW_STATUS) != LAST_INI_REG_VAL) {
  }
}

bm_status_t bm_device::bm_malloc_device_dword(
    bm_device_mem_t *pmem, int cnt) {
  u32 size = cnt * FLOAT_SIZE;
  u64 addr = 0;

  addr = device_mem_pool.bm_mem_pool_alloc(size);

  pmem->u.device.device_addr = addr;
  pmem->flags.u.mem_type = BM_MEM_TYPE_DEVICE;
  pmem->size = size;
  return BM_SUCCESS;
}

void bm_device::bm_free_device(bm_device_mem_t mem) {
  u64 addr = (u64)bm_mem_get_device_addr(mem);
  device_mem_pool.bm_mem_pool_free(addr);
}

bm_status_t bm_device::bm_alloc_arm_reserved() {
  BM_CHECK_RET(bm_malloc_device_dword(&arm_reserved_dev_mem,
        CONFIG_COUNT_RESERVED_DDR_ARM));
  return BM_SUCCESS;
}

void bm_device::bm_free_arm_reserved() {
  bm_free_device(arm_reserved_dev_mem);
}

bm_status_t bm_device::bm_alloc_instr_reserved() {
  BM_CHECK_RET(bm_malloc_device_dword(&instr_reserved_mem,
        COUNT_RESERVED_DDR_INSTR / 4));
  return BM_SUCCESS;
}

void bm_device::bm_free_instr_reserved() {
  bm_free_device(instr_reserved_mem);
}

#define SMMU_RESERVED_SIZE (0x40000 * 4)
bm_status_t bm_device::bm_alloc_iommu_reserved() {
  BM_CHECK_RET(bm_malloc_device_dword(&iommu_reserved_dev_mem,
        SMMU_RESERVED_SIZE/sizeof(float)));
  return BM_SUCCESS;
}

void bm_device::bm_free_iommu_reserved() {
  bm_free_device(iommu_reserved_dev_mem);
}

bm_status_t bm_device::bm_init_l2_sram() {
  bm_device_mem_t  pmem;

  bm_set_device_mem(&pmem, sizeof(l2_sram_table), L2_SRAM_START_ADDR
                  + EX_INT_TABLE_OFFSET);
  BM_CHECK_RET(bm_device_memcpy_s2d(pmem, l2_sram_table));

  return BM_SUCCESS;
}

bm_status_t bm_device::bm_device_memcpy_s2d(bm_device_mem_t dst, void *src) {
  u32 size_total = bm_mem_get_size(dst);
  u64 dst_addr = bm_mem_get_device_addr(dst);
  u64 size_step;
  u64 realmem_size = CONFIG_HOST_STAGMEM_SIZE;
  for (u32 pass_idx = 0, cur_addr_inc = 0;
      pass_idx < (size_total + realmem_size - 1) / realmem_size;
      pass_idx++) {
    if ((pass_idx + 1) * realmem_size < size_total) {
      size_step = realmem_size;
    } else {
      size_step = size_total - pass_idx * realmem_size;
    }
    u8 *src_copy = reinterpret_cast < u8 *>(src) + cur_addr_inc;
    int mem_type;
    u8 *dst_copy;
    u64 cur_dst_addr = dst_addr + cur_addr_inc;
    if (cur_dst_addr >= GLOBAL_MEM_START_ADDR) {
        dst_copy = reinterpret_cast < u8 *>(get_global_memaddr(dev_id))
                       + cur_dst_addr - GLOBAL_MEM_START_ADDR;
        mem_type = MEM_TYPE_GLOBAL;
    } else if (cur_dst_addr >=  L2_SRAM_START_ADDR &&
                cur_dst_addr < (L2_SRAM_START_ADDR + L2_SRAM_SIZE)) {
        dst_copy = reinterpret_cast < u8 *>(get_l2_sram(dev_id))
                       + cur_dst_addr - L2_SRAM_START_ADDR;
        mem_type = MEM_TYPE_L2;
    } else {
        printf("The dst address of cdma is invalid\n");
        ASSERT(0);
    }
    host_dma_copy_cmodel(dst_copy, src_copy, size_step,
                    HOST2CHIP, dev_id, mem_type);
    cur_addr_inc += size_step;
  }
  return BM_SUCCESS;
}

bm_status_t bm_device::bm_device_memcpy_d2s(void *dst, bm_device_mem_t src) {
  u32 size_total = bm_mem_get_size(src);
  u64 src_addr = bm_mem_get_device_addr(src);

  u64 size_step;
  u64 realmem_size = CONFIG_HOST_STAGMEM_SIZE;
  for (u32 pass_idx = 0, cur_addr_inc = 0;
      pass_idx < (size_total + realmem_size - 1) / realmem_size;
      pass_idx++) {
    if ((pass_idx + 1) * realmem_size < size_total)
      size_step = realmem_size;
    else
      size_step = size_total - pass_idx * realmem_size;
    u8 *dst_copy = reinterpret_cast < u8 *>(dst) + cur_addr_inc;
    int mem_type;
    u64 cur_src_addr = src_addr + cur_addr_inc;
    u8 *src_copy;
    if (cur_src_addr >= GLOBAL_MEM_START_ADDR) {
        src_copy = reinterpret_cast < u8 *>(get_global_memaddr(dev_id))
                       + cur_src_addr - GLOBAL_MEM_START_ADDR;
        mem_type = MEM_TYPE_GLOBAL;
    } else if (cur_src_addr >= L2_SRAM_START_ADDR &&
                    cur_src_addr < (L2_SRAM_START_ADDR + L2_SRAM_SIZE)) {
        src_copy = reinterpret_cast < u8 *>(get_global_memaddr(dev_id))
                       + cur_src_addr - L2_SRAM_START_ADDR;
        mem_type = MEM_TYPE_L2;
    } else {
        printf("The src address of cdma is invalid\n");
        ASSERT(0);
    }
    host_dma_copy_cmodel(dst_copy, src_copy, size_step,
                    CHIP2HOST, dev_id, mem_type);
    cur_addr_inc += size_step;
  }
  return BM_SUCCESS;
}

u64 bm_device::bm_device_arm_reserved_req() {
  pthread_mutex_lock(&arm_reserved_lock);
  return arm_reserved_dev_mem.u.device.device_addr;
}
void bm_device::bm_device_arm_reserved_rel() {
  pthread_mutex_unlock(&arm_reserved_lock);
}

bm_status_t bm_device::bm_device_thread_sync() {
  // should add volatile, if not, thread_api_info will be
  // optimized by c++ in nodebug mode
  volatile thread_api_info *thd_api_info;
  thd_api_info = bm_get_thread_api_info(pthread_self());
  if (!thd_api_info) {
    printf("Error: thread api info %lu is not found!\n", pthread_self());
    ASSERT(0);
    return BM_ERR_FAILURE;
  }

  while (thd_api_info->last_seq != thd_api_info->cpl_seq) {
  }

  #ifdef BM_TV_GEN
  bm_read32_wait_eq_tv(SHARE_REG_BASE_ADDR + SHARE_REG_MESSAGE_RP * 4,
      cmodel_read_share_reg(SHARE_REG_MESSAGE_WP, dev_id), 32, 0, HOST_REG);
  #endif
  return BM_SUCCESS;
}

thread_api_info *bm_device::bm_get_thread_api_info(pthread_t thd_id) {
  std::map < pthread_t, thread_api_info>::iterator it;
  it = thread_api_table.find(thd_id);
  if (it != thread_api_table.end())
    return &it->second;
  else
    return nullptr;
}

bm_status_t bm_device::bm_add_thread_api_info(pthread_t thd_id) {
  thread_api_table.insert(std::pair < pthread_t, thread_api_info > (thd_id,
              {thd_id, 0, 0}));
  return BM_SUCCESS;
}

bm_status_t bm_device::bm_remove_thread_api_info(pthread_t thd_id) {
  std::map < pthread_t, thread_api_info>::iterator it;
  it = thread_api_table.find(thd_id);
  if (it != thread_api_table.end())
    thread_api_table.erase(it);
  return BM_SUCCESS;
}

void *bm_device::bm_msg_done_poll(void *arg) {
  bm_device *bm_dev = reinterpret_cast < bm_device *>(arg);
  while (1) {
    while (!bm_dev->pending_api_queue.empty()) {
      api_queue_entry api_front = bm_dev->pending_api_queue.front();
      if (api_front.thd_id == DEVICE_SYNC_MARKER) {
// device sync
        bm_dev->device_sync_cpl = api_front.dev_seq;
        pthread_mutex_lock(&bm_dev->api_lock);
        bm_dev->pending_api_queue.pop();
        pthread_mutex_unlock(&bm_dev->api_lock);
        pthread_yield();
      } else {
// msg api pending
        pthread_mutex_lock(&msg_done_mutex[bm_dev->dev_id]);
        while (0 == msg_done[bm_dev->dev_id])
          pthread_cond_wait(&msg_done_cond[bm_dev->dev_id],
                            &msg_done_mutex[bm_dev->dev_id]);
        msg_done[bm_dev->dev_id]--;
        pthread_mutex_unlock(&msg_done_mutex[bm_dev->dev_id]);

        if (api_front.thd_id != 0) {
          bm_dev->thread_api_table[api_front.thd_id].cpl_seq =
                  api_front.thd_seq;
          pthread_mutex_lock(&bm_dev->api_lock);
          bm_dev->pending_api_queue.pop();
          pthread_mutex_unlock(&bm_dev->api_lock);
        } else {
          ASSERT(0);
        }
      }
    }
// busy waiting sleep 200ms, reduce cpu usage
#if defined(USING_CMODEL) && !defined(USING_MULTI_THREAD_ENGINE)
    usleep(200000);
#endif
    pthread_testcancel();
  }
  return nullptr;
}

bm_device_manager::bm_device_manager(int _max_dev_cnt)
  : dev_cnt(0),
    max_dev_cnt(_max_dev_cnt),
    bm_dev_list(nullptr) {
  bm_dev_list = new bm_device *[max_dev_cnt];
  if (!bm_dev_list)
      return;
  for (int i = 0; i < max_dev_cnt; i++)
      bm_dev_list[i] = nullptr;
}

bm_device_manager::~bm_device_manager() {
  if (bm_dev_list) {
    for (int i = 0; i < max_dev_cnt; i++) {
      if (bm_dev_list[i]) {
        delete bm_dev_list[i];
        bm_dev_list[i] = nullptr;
      }
    }
    delete []bm_dev_list;
  }
}

bm_device_manager *bm_device_manager::get_dev_mgr() {
  pthread_mutex_lock(&init_lock);
  if (!bm_dev_mgr)
    bm_dev_mgr = new bm_device_manager(MAX_NODECHIP_NUM);
  pthread_mutex_unlock(&init_lock);
  return bm_dev_mgr;
}

bm_device *bm_device_manager::get_bm_device(int dev_id) {
  ASSERT(bm_dev_list);
  ASSERT(dev_id < max_dev_cnt);
  pthread_mutex_lock(&init_lock);
  if (!bm_dev_list[dev_id]) {
    bm_dev_list[dev_id] = new bm_device(dev_id);
    dev_cnt++;
  }
  pthread_mutex_unlock(&init_lock);
  return bm_dev_list[dev_id];
}

void bm_device_manager::destroy_dev_mgr() {
  // std::cout << "bm_dev_mgr "<<bm_dev_mgr <<std::endl;
  if (bm_dev_mgr)
    delete bm_dev_mgr;
}

bm_device_manager *bm_device_manager::bm_dev_mgr = nullptr;
pthread_mutex_t bm_device_manager::init_lock = PTHREAD_MUTEX_INITIALIZER;
#endif


