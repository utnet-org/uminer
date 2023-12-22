
#include "chip.h"
#include <iostream>

// 芯片的驱动读取芯片数量
//bool bm_dev_getCount(int *devCount) {
//    *devCount = 10;
//    return true;
//};
//
//bool DeviceIoControl(const char* bmctl_device, int* id, bmsmi* device){
//    // 芯片驱动读取device的参数
//
//    device->chip_id = *id;
//
//    device->domain_bdf = *id % 3;
//
//    device->chip_mode = 0;
//
//    device->status = 0;
//
//    device->tpu_util = 0;
//
//    device->board_temp = 48;
//
//    device->chip_temp = 51;
//
//    device->board_power = 75;
//
//    device->tpu_power = 2.0;
//
//    return true;
//
//}
//
//void bm_total_gmem(const char* handle, int *mem){
//    *mem = 100 * 1024 * 1024;
//}
//void bm_avail_gmem(const char* handle, int *mem){
//    *mem = 1 * 1024 * 1024;
//}
//
//
//
//bmsmi* queryChipDetails() {
//    int dev_cnt = 0;
//
//    bmsmi* fail_g_attr = new bmsmi[1];
//
//    fail_g_attr[0].chip_id = ATTR_FAULT_VALUE;
//    fail_g_attr[0].chip_mode = ATTR_FAULT_VALUE;  // 0---pcie = ATTR_FAULT_VALUE; 1---soc
//    fail_g_attr[0].domain_bdf = ATTR_FAULT_VALUE;
//    fail_g_attr[0].status = ATTR_FAULT_VALUE;
//
//    fail_g_attr[0].mem_used = 0;
//    fail_g_attr[0].mem_total = 0;
//    fail_g_attr[0].tpu_util = ATTR_FAULT_VALUE;
//
//    fail_g_attr[0].board_temp = ATTR_FAULT_VALUE;
//    fail_g_attr[0].chip_temp = ATTR_FAULT_VALUE;
//    fail_g_attr[0].board_power = ATTR_FAULT_VALUE;
//    fail_g_attr[0].tpu_power = ATTR_FAULT_VALUE;
//
//    if (bm_dev_getcount(&dev_cnt) != true) {
//        printf("get devcount failed!\n");
//        return fail_g_attr;
//    }
//
//    // initialized
//    printf("we have %d dev\n", dev_cnt);
//    bmsmi* g_attr = new bmsmi[dev_cnt];
//
//    for (int dev_id = 0; dev_id < 0 + dev_cnt; dev_id++) {
//        g_attr[dev_id].dev_id = dev_id;
//        // memory
//        const char* handle;
//        int mem_total, mem_avail;
//        bm_total_gmem(handle, &mem_total);
//        bm_avail_gmem(handle, &mem_avail);
//        g_attr[dev_id].mem_total = mem_total / 1024 / 1024;
//        g_attr[dev_id].mem_used  = (mem_total - mem_avail) / 1024 / 1024;
//        // other params
//        const char* bmctl_device;
//        bool status = DeviceIoControl(bmctl_device, &dev_id, &g_attr[dev_id]);
//        if (status == false) {
//            printf("DeviceIoControl BMCTL_GET_PROC_GMEM failed \n");
//            return fail_g_attr;
//        }
//
//    }
//    return g_attr;
//
//}