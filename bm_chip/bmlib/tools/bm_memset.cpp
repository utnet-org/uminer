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
DEFINE_int32(mode, 1, "which set mode to use");
DEFINE_int32(size, 1, "how many bytes to set");
DECLARE_bool(help);
DECLARE_bool(helpshort);

bm_status_t	test_memset_byte(bm_handle_t handle, int size)
{
	bm_status_t ret = BM_SUCCESS;
	bm_device_mem_t device_mem;
	unsigned char* cmp_data = NULL;
	unsigned char src_byte = 0;
	int i = 0;

	cmp_data = (unsigned char *)malloc(size);
	if(!cmp_data) {
		printf("malloc host memroy failed\n");
		return BM_ERR_NOMEM;
	}

	ret = bm_malloc_device_byte(handle, &device_mem, size);
	if (BM_SUCCESS != ret) {
		free(cmp_data);
		printf("request device memory failed, size = 0x%x\n", size);
		return BM_ERR_NOMEM;
	}

	src_byte = rand() % 256;
	ret = bm_memset_device_ext(handle, &src_byte, 1, device_mem);
	if (BM_SUCCESS != ret) {
		free(cmp_data);
		bm_free_device(handle, device_mem);
		printf("device memory byte set failed\n");
		return BM_ERR_FAILURE;
	}

	ret = bm_memcpy_d2s(handle, cmp_data, device_mem);
	if (BM_SUCCESS != ret) {
		free(cmp_data);
		bm_free_device(handle, device_mem);
		printf("d2s failed in %s\n", __func__);
		return BM_ERR_FAILURE;
	}

	for(i=0; i<size; i++) {
		if(cmp_data[i] != src_byte) {
			printf("cmp failed in byte mode, src_byte = %d, got = %d on index %d\n",
					src_byte, cmp_data[i], i);
			free(cmp_data);
			bm_free_device(handle, device_mem);
			return BM_ERR_FAILURE;
		}
	}

	free(cmp_data);
	bm_free_device(handle, device_mem);

	return ret;
}

bm_status_t	test_memset_2bytes(bm_handle_t handle, int size)
{
	bm_status_t ret = BM_SUCCESS;
	bm_device_mem_t device_mem;
	unsigned char* cmp_data = NULL;
	unsigned short src_short = 0;
	unsigned short *psrc = NULL;
	int i = 0;

	cmp_data = (unsigned char *)malloc(size);
	if(!cmp_data) {
		printf("malloc host memroy failed\n");
		return BM_ERR_NOMEM;
	}

	ret = bm_malloc_device_byte(handle, &device_mem, size);
	if (BM_SUCCESS != ret) {
		free(cmp_data);
		printf("request device memory failed, size = 0x%x\n", size);
		return BM_ERR_NOMEM;
	}

	src_short = rand() % 65536;
	ret = bm_memset_device_ext(handle, &src_short, 2, device_mem);
	if (BM_SUCCESS != ret) {
		free(cmp_data);
		bm_free_device(handle, device_mem);
		printf("device memory set 2bytes failed\n");
		return BM_ERR_FAILURE;
	}

	ret = bm_memcpy_d2s(handle, cmp_data, device_mem);
	if (BM_SUCCESS != ret) {
		free(cmp_data);
		bm_free_device(handle, device_mem);
		printf("d2s failed in %s\n", __func__);
		return BM_ERR_FAILURE;
	}

	psrc = (unsigned short *)cmp_data;
	for(i=0; i<(size/2); i++) {
		if(psrc[i] != src_short) {
			printf("cmp failed in 2bytes mode, src_short = %d, got = %d on index %d\n",
					src_short, psrc[i], i);
			free(cmp_data);
			bm_free_device(handle, device_mem);
			return BM_ERR_FAILURE;
		}
	}

	free(cmp_data);
	bm_free_device(handle, device_mem);

	return ret;
}

bm_status_t	test_memset_3bytes(bm_handle_t handle, int size)
{
	bm_status_t ret = BM_SUCCESS;
	bm_device_mem_t device_mem;
	unsigned char* cmp_data = NULL;
	unsigned char rgb[3];
	int i = 0;

	cmp_data = (unsigned char *)malloc(size);
	if(!cmp_data) {
		printf("malloc host memroy failed\n");
		return BM_ERR_NOMEM;
	}

	ret = bm_malloc_device_byte(handle, &device_mem, size);
	if (BM_SUCCESS != ret) {
		free(cmp_data);
		printf("request device memory failed, size = 0x%x\n", size);
		return BM_ERR_NOMEM;
	}

	rgb[0] = rand() % 256;
	rgb[1] = rand() % 256;
	rgb[2] = rand() % 256;
	ret = bm_memset_device_ext(handle, rgb, 3, device_mem);
	if (BM_SUCCESS != ret) {
		free(cmp_data);
		bm_free_device(handle, device_mem);
		printf("device memory set 3bytes failed\n");
		return BM_ERR_NOMEM;
	}

	ret = bm_memcpy_d2s(handle, cmp_data, device_mem);
	if (BM_SUCCESS != ret) {
		free(cmp_data);
		bm_free_device(handle, device_mem);
		printf("d2s failed in %s\n", __func__);
		return BM_ERR_FAILURE;
	}

	for(i=0; i<size; i++) {
		if(cmp_data[i] != rgb[i%3]) {
			printf("cmp failed in 3bytes mode, rgb[] = %d, got = %d on index %d\n",
					rgb[i%3], cmp_data[i], i);
			free(cmp_data);
			bm_free_device(handle, device_mem);
			return BM_ERR_FAILURE;
		}
	}

	free(cmp_data);
	bm_free_device(handle, device_mem);

	return ret;
}

bm_status_t	test_memset_4bytes(bm_handle_t handle, int size)
{
	bm_status_t ret = BM_SUCCESS;
	bm_device_mem_t device_mem;
	unsigned char* cmp_data = NULL;
	float fp32 = 0.0;
	float *psrc = NULL;
	int i = 0;

	cmp_data = (unsigned char *)malloc(size);
	if(!cmp_data) {
		printf("malloc host memroy failed\n");
		return BM_ERR_NOMEM;
	}

	ret = bm_malloc_device_byte(handle, &device_mem, size);
	if (BM_SUCCESS != ret) {
		free(cmp_data);
		printf("request device memory failed, size = 0x%x\n", size);
		return BM_ERR_NOMEM;
	}

	fp32 = rand() + 3.9;
	ret = bm_memset_device_ext(handle, &fp32, 4, device_mem);
	if (BM_SUCCESS != ret) {
		free(cmp_data);
		bm_free_device(handle, device_mem);
		printf("device memory set 4bytes failed\n");
		return BM_ERR_FAILURE;
	}

	ret = bm_memcpy_d2s(handle, cmp_data, device_mem);
	if (BM_SUCCESS != ret) {
		free(cmp_data);
		bm_free_device(handle, device_mem);
		printf("d2s failed in %s\n", __func__);
		return BM_ERR_FAILURE;
	}

	psrc = (float*)cmp_data;
	for(i=0; i<size/4; i++) {
		if(psrc[i] != fp32) {
			printf("cmp failed in 4bytes mode, fp32 = %f, got = %f on index %d\n",
					fp32, psrc[i], i);
			free(cmp_data);
			bm_free_device(handle, device_mem);
			return BM_ERR_FAILURE;
		}
	}

	fp32 = rand() + 5.8;
	int int32 = 0;
	memcpy(&int32, &fp32, 4);
	ret = bm_memset_device(handle, int32, device_mem);
	if (BM_SUCCESS != ret) {
		free(cmp_data);
		bm_free_device(handle, device_mem);
		printf("device memory set 4bytes failed\n");
		return BM_ERR_FAILURE;
	}

	ret = bm_memcpy_d2s(handle, cmp_data, device_mem);
	if (BM_SUCCESS != ret) {
		free(cmp_data);
		bm_free_device(handle, device_mem);
		printf("d2s failed in %s\n", __func__);
		return BM_ERR_FAILURE;
	}

	psrc = (float*)cmp_data;
	for(i=0; i<size/4; i++) {
		if(psrc[i] != fp32) {
			printf("cmp failed in 4bytes mode, fp32 = %f, got = %f on index %d\n",
					fp32, psrc[i], i);
			free(cmp_data);
			bm_free_device(handle, device_mem);
			return BM_ERR_FAILURE;
		}
	}

	free(cmp_data);
	bm_free_device(handle, device_mem);

	return ret;
}

int main(int argc, char *argv[])
{
	bm_handle_t handle;
	bm_status_t ret = BM_SUCCESS;
	struct timespec tp;

	gflags::SetUsageMessage("command line brew\n"
			"usage: bm_memset [--devid=0] [--mode=1] [size=0x1]\n"
			"devid:\n"
			"  which sophon dev to alloc memroy from.\n"
			"mode:\n"
			"  mode1:byte set, mode2:fp16 set, mode3:rgb set, mode4:fp32 set.\n"
			"size:\n"
			"  how many bytes to set.\n"
			);

	gflags::ParseCommandLineNonHelpFlags(&argc, &argv, true);
	if (FLAGS_help) {
		FLAGS_help = false;
		FLAGS_helpshort = true;
	}
	gflags::HandleCommandLineHelpFlags();

	clock_gettime(CLOCK_THREAD_CPUTIME_ID, &tp);
	srand(tp.tv_nsec);
	printf("random seed %lu\n", tp.tv_nsec);

	ret = bm_dev_request(&handle, FLAGS_devid);
	if (BM_SUCCESS != ret) {
		printf("request dev%d failed\n", FLAGS_devid);
		return BM_ERR_FAILURE;
	}

	if (1==FLAGS_mode) {
		ret = test_memset_byte(handle, FLAGS_size);
	} else if(2==FLAGS_mode) {
		ret = test_memset_2bytes(handle, FLAGS_size);
	} else if(3==FLAGS_mode) {
		ret = test_memset_3bytes(handle, FLAGS_size);
	} else if(4==FLAGS_mode) {
		ret = test_memset_4bytes(handle, FLAGS_size);
	} else {
		ret = BM_ERR_PARAM;
		printf("error mode %d\n", FLAGS_mode);
	}

	bm_dev_free(handle);
	printf("memset mode %d test %s\n", FLAGS_mode, ret? "failed" : "pass");
	return (int)ret;
}
