#include <stdio.h>
#include <bmlib_runtime.h>

#include "chip.h"

int startCPU(int seq, const char * boot_file, const char * core_file) {
#if !defined(USING_CMODEL) && !defined(SOC_MODE)
    bm_handle_t handle;
    bm_status_t ret;

    ret = bm_dev_request(&handle, seq);
    if ((ret != BM_SUCCESS) || (handle == NULL)) {
        printf("bm_dev_request error, ret = %d\n", ret);
        return -1;
    }

    char* boot_file_path = const_cast<char*>(boot_file);
    char* core_file_path = const_cast<char*>(core_file);
    ret = bmcpu_start_cpu(handle, boot_file_path,  core_file_path);
    if (ret != BM_SUCCESS) {
        printf("ERROR!!! start cpu error!\n");
        bm_dev_free(handle);
        return 0;
    } else {
        printf("Start cpu success!\n");
    }

    bm_dev_free(handle);
#else

#endif
    return 1;
}
