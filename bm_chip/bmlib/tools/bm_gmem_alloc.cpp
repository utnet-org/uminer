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
DEFINE_int32(heapid, 0xff, "which heap to alloc memroy from, 0xff is all");
DEFINE_int32(blocksize, 1024, "the size of one block memory, KBytes");
DEFINE_int32(blockcnt, 1, "how many blocks to alloc");
DECLARE_bool(help);
DECLARE_bool(helpshort);

int main(int argc, char *argv[])
{
	bm_device_mem_t *device_mem_arrar = NULL;
	int i = 0;
	int effective_blocknum = 0;
	static bm_handle_t handle;
	bm_status_t ret = BM_SUCCESS;
	int ch = 0;
	unsigned int gmem_heapid;

	gflags::SetUsageMessage("command line brew\n"
			"usage: bm_gmem_alloc [--devid=0] [--heapid=0] [--blocksize=1024] [--blockcnt=1]\n"
			"devid:\n"
			"  which sophon dev to alloc memroy from.\n"
			"heapid:\n"
			"  which heap to alloc memroy from, 0xff is all.\n"
			"blocksize:\n"
			"  the size of one block memory, KBytes.\n"
			"blockcnt:\n"
			"  how many blocks to alloc.");

	gflags::ParseCommandLineNonHelpFlags(&argc, &argv, true);
	if (FLAGS_help) {
		FLAGS_help = false;
		FLAGS_helpshort = true;
	}
	gflags::HandleCommandLineHelpFlags();

	device_mem_arrar = (bm_device_mem_t *)malloc(sizeof(bm_device_mem_t)*FLAGS_blockcnt);
	if(!device_mem_arrar) {
		printf("malloc host memroy failed\n");
		return -1;
	}

	ret = bm_dev_request(&handle, FLAGS_devid);
	if (BM_SUCCESS != ret) {
		printf("request dev%d failed\n", FLAGS_devid);
		free(device_mem_arrar);
		return BM_ERR_FAILURE;
	}
	for(i=0; i<FLAGS_blockcnt; i++) {
		if (FLAGS_heapid == 0xff) {
			ret = bm_malloc_device_byte(handle, &device_mem_arrar[i], FLAGS_blocksize*1024);
		} else {
			ret = bm_malloc_device_byte_heap(handle, &device_mem_arrar[i], FLAGS_heapid, FLAGS_blocksize*1024);
		}

		if (BM_SUCCESS != ret) {
			printf("request device memory block%d failed\n", i);
			effective_blocknum = i;
			break;
		}
		ret = bm_get_gmem_heap_id(handle, &device_mem_arrar[i], &gmem_heapid);
		if (BM_SUCCESS != ret) {
			printf("get gmem heap id failed\n");
			effective_blocknum = i;
			break;
		}
		if (FLAGS_heapid != 0xff && gmem_heapid != (unsigned int)FLAGS_heapid) {
			printf("get an error gmem heap id:%d\n", gmem_heapid);
			effective_blocknum = i;
			break;
		}
	}

	effective_blocknum = i;
	std::cin >> ch;

	for(i=0; i<effective_blocknum; i++) {
		bm_free_device(handle, device_mem_arrar[i]);
	}

	free(device_mem_arrar);
	bm_dev_free(handle);
}
