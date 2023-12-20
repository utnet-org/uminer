#ifndef GLOBAL_LAYER_DATA_SPLIT_H_
#define GLOBAL_LAYER_DATA_SPLIT_H_

#include "common.h"

#if defined (__cplusplus)
extern "C" {
#endif


typedef struct pooling_secs_info {
    int nsecs;
    int hsecs;
    float Ratio;
}pooling_secs_info_t;

typedef struct depthwise_secs_info {
    int nsecs;
    int csecs;
    int hsecs;
    float Ratio;
}depthwise_secs_info;

typedef struct conv_secs_info {
  int ocsecs;
  int icsecs;
  int nsecs;
  int hsecs;
  int owsecs;
} conv_secs_info_t;

int conv16_ofmap_csize_local(int oh_slice);
// return 0: fail 1: success
int global_conv_data_split(
    int    input_n,
    int    input_c,
    int    input_h,
    int    input_w,
    int    output_c,
    int    output_h,
    int    output_w,
    int    groups,
    int    kh,
    int    kw,
    int    dh,
    int    dw,
    int    stride_h,
    int    stride_w,
    int    pad_h,
    int    pad_h_after,
    int    pad_w,
    int    pad_w_after,
    int    using_bias,
    conv_secs_info_t* p_secs_info);
// i return 0: fail 1: success
int global_conv_data_split_fix8b(
    int    input_n,
    int    input_c,
    int    input_h,
    int    input_w,
    int    output_c,
    int    output_h,
    int    output_w,
    int    groups,
    int    kh,
    int    kw,
    int    dh,
    int    dw,
    int    stride_h,
    int    stride_w,
    int    using_bias,
    int    ins_h,
    int    ins_w,
    conv_secs_info_t* p_secs_info);

int global_conv_data_split_fix16b(
    int    input_n,
    int    input_c,
    int    input_h,
    int    input_w,
    int    output_c,
    int    output_h,
    int    output_w,
    int    groups,
    int    kh,
    int    kw,
    int    dh,
    int    dw,
    int    stride_h,
    int    stride_w,
    int    using_bias,
    int    ins_h,
    int    ins_w,
    conv_secs_info_t* p_secs_info);


int global_pooling_data_split(
    int                 input_n,
    int                 input_c,
    int                 input_h,
    int                 input_w,
    int                 output_h,
    int                 output_w,
    int                 kh,
    int                 kw,
    int                 pad_h,
    int                 pad_w,
    int                 pad_h_after,
    int                 pad_w_after,
    int                 stride_h,
    int                 stride_w,
    int                 is_avg_pooling,
    pooling_secs_info_t *pooling_secs_info);

int global_pooling_train_data_split(
    int                 input_n,
    int                 input_c,
    int                 input_h,
    int                 input_w,
    int                 kh,
    int                 kw,
    int                 pad_h,
    int                 pad_w,
    int                 pad_h_after,
    int                 pad_w_after,
    int                 stride_h,
    int                 stride_w,
    int                 is_avg_pooling,
    pooling_secs_info_t *pooling_secs_info);

int global_pooling_bw_data_split(
    int                 input_n,
    int                 input_c,
    int                 input_h,
    int                 input_w,
    int                 kh,
    int                 kw,
    int                 pad_h,
    int                 pad_w,
    int                 pad_h_after,
    int                 pad_w_after,
    int                 stride_h,
    int                 stride_w,
    int                 is_avg_pooling,
    pooling_secs_info_t *pooling_secs_info);

int global_upsample_data_split(
    int                 input_n,
    int                 input_c,
    int                 input_h,
    int                 input_w,
    int                 kh,
    int                 kw,
    int                 pad_h,
    int                 pad_w,
    int                 stride_h,
    int                 stride_w,
    int                 is_avg_pooling,
    pooling_secs_info_t *pooling_secs_info);

int get_deconv_secs_info(
    bm_tensor_4d_t    input_shape,
    int bottom_c, int top_c, int kh, int kw,
    bm_tensor_4d_t    output_shape,
    int              using_bias,
    int stride_h,
    int dh, int pad_h,
    int groups,
    conv_secs_info_t *secs_info);

int get_depthwise_deconv_secs_info(
    bm_tensor_4d_t input_shape,
    int bottom_c,
    int kh,
    int kw,
    bm_tensor_4d_t output_shape,
    int using_bias,
    int stride_h,
    int dh,
    int pad_h,
    depthwise_secs_info* secs_info);

int lstm_data_split_select(
    int    batch_num,
    int    input_dim,
    int    output_dim,
    int    with_x_static,
    int*   batch_slice);

int lstm_output_dim_extend(int output_dim);

#if defined (__cplusplus)
}
#endif

#endif

