#include <stdio.h>
#include <unistd.h>

#include "common.h"
#include "bmlib_runtime.h"
#include "bmlib_internal.h"
#include "bmlib_memory.h"
#include "warp_affine_const.h"
#ifdef USE_NPU_BILINEAR_IMAGE_SCALE
#include "bilinear_coeff_tbl.h"
#endif
#ifndef USING_CMODEL

static bm_status_t bm_load_warp_affine_constant(bm_handle_t ctx) {
  BM_CHECK_RET(bm_malloc_device_byte(ctx, &ctx->warp_affine_constant_mem,
      CONSTANT_ENTRY_WARP_TBL * sizeof(int)));
  BM_CHECK_RET(bm_memcpy_s2d(ctx, ctx->warp_affine_constant_mem,
                             warp_affine_tbl));
  return BM_SUCCESS;
}

static void bm_free_warp_affine_constant(bm_handle_t ctx) {
  bm_free_device(ctx, ctx->warp_affine_constant_mem);
}

#ifdef USE_NPU_BILINEAR_IMAGE_SCALE
static bm_status_t bm_load_bilinear_image_scale_constant(bm_handle_t ctx) {
  BM_CHECK_RET(bm_malloc_device_dword(ctx, &ctx->bilinear_coeff_table_dev_mem,
      CONSTANT_ENTRY_BILINEAR_TBL));
  BM_CHECK_RET(bm_memcpy_s2d(ctx, ctx->bilinear_coeff_table_dev_mem,
                             bilinear_coeff_tbl));
  return BM_SUCCESS;
}

static void bm_free_bilinear_image_scale_constant(bm_handle_t ctx) {
  bm_free_device(ctx, ctx->bilinear_coeff_table_dev_mem);
}

static bm_status_t bm_alloc_bilinear_image_scale_reserved(bm_handle_t ctx) {
  BM_CHECK_RET(bm_malloc_device_dword(ctx, &ctx->image_scale_reserved_mem,
      COUNT_RESERVED_DDR_IMAGE_SCALE/4));
  return BM_SUCCESS;
}

static void bm_free_binlinear_image_scale_reserved(bm_handle_t ctx) {
  bm_free_device(ctx, ctx->image_scale_reserved_mem);
}
#endif

bm_status_t bm1682_dev_request(bm_context_t *ctx) {
  BM_CHECK_RET(bm_load_warp_affine_constant(ctx));
#ifdef USE_NPU_BILINEAR_IMAGE_SCALE
  BM_CHECK_RET(bm_alloc_bilinear_image_scale_reserved(ctx));
  BM_CHECK_RET(bm_load_bilinear_image_scale_constant(ctx));
#endif
  return BM_SUCCESS;
}

void bm1682_dev_free(bm_context_t *ctx) {
#ifdef USE_NPU_BILINEAR_IMAGE_SCALE
  bm_free_bilinear_image_scale_reserved(ctx);
  bm_free_bilinear_image_scale_constant(ctx);
#endif
  bm_free_warp_affine_constant(ctx);
}
#endif
