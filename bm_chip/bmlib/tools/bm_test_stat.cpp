#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <sys/ioctl.h>
#include <ncurses.h>
#include <gflags/gflags.h>
#include <errno.h>
#include <ctime>
#include <cstdio>
#include <cstdlib>
#include <iostream>
#include <sstream>
#include <fstream>
#include <string>
#include "bmlib_runtime.h"
#include "bmlib_internal.h"
#include "message.h"
#include "gflags/gflags.h"

DEFINE_int32(devid, 0, "which sophon dev to alloc memroy from");
DECLARE_bool(help);
DECLARE_bool(helpshort);

int main(int argc, char *argv[])
{
	bm_handle_t handle;
	bm_status_t ret = BM_SUCCESS;
	bm_dev_stat_t stat;
	int i = 0;

	gflags::SetUsageMessage("command line brew\n"
			"usage: bm_gmem_alloc [--devid=0] \n"
			"devid:\n"
			"  which sophon dev to test on.\n"
			);

	gflags::ParseCommandLineNonHelpFlags(&argc, &argv, true);
	if (FLAGS_help) {
		FLAGS_help = false;
		FLAGS_helpshort = true;
	}
	gflags::HandleCommandLineHelpFlags();

	ret = bm_dev_request(&handle, FLAGS_devid);
	if (BM_SUCCESS != ret) {
		printf("request dev%d failed\n", FLAGS_devid);
		return BM_ERR_FAILURE;
	}
	ret = bm_get_stat(handle, &stat);
	if (BM_SUCCESS != ret) {
		printf("get stat data failed\n");
	} else {
		printf("there are total %d MB gmem and used %d MB, tpu util is %d\n", stat.mem_total, stat.mem_used, stat.tpu_util);
		for (i = 0; i < stat.heap_num; i++) {
			printf("heap:%d mem_total: %dMB, mem_avail: %dMB, mem_used: %dMB\n",
					i, stat.heap_stat[i].mem_total, stat.heap_stat[i].mem_avail, stat.heap_stat[i].mem_used);
		}
	}

	bm_dev_free(handle);
	return (int)ret;
}
