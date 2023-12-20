#include "stdlib.h"
#include "bmlib_utils.h"
#ifdef __linux__
#include "common.h"
#else
#include "..\..\common\bm1684\include_win\common_win.h"
#endif


void random_conv_param(
    int &n, int &ic, int &ih, int &iw, int &oc,
    int &kh, int &kw, int &dh, int &dw,
    int &ph, int &pw, int &sh, int &sw) {
  n = rand() % 256 + 1;
  ic = rand() % (NPU_NUM * 16) + 1;
  oc = rand() % (NPU_NUM * 16) + 1;
  int max_kernel = 0;
  int max_dilation = 0;
  if (rand() % 10 < 8) {
    max_kernel = 5;
    max_dilation = 10;
  }else {
    max_kernel = 15;
    max_dilation = 3;
  }
  if (rand() % 2 == 0) {
    kh = kw = rand() % max_kernel + 1;
    dh = dw = rand() % max_dilation + 1;
  }else {
    kh = rand() % max_kernel + 1;
    kw = rand() % max_kernel + 1;
    dh = rand() % max_dilation + 1;
    dw = rand() % max_dilation + 1;
  }
  ih = rand() % (EU_NUM * 16) + dh * (kh - 1) + 1;
  iw = rand() % (EU_NUM * 4) + dw * (kw - 1) + 1;
  ph = rand() % kh;
  pw = rand() % kw;
  if (rand() % 2 == 0) {
    sh = 1;
    sw = 1;
  }else {
    sh = rand() % kh + 1;
    sw = rand() % kw + 1;
  }
  while(((u64)(n * ic) ) * ((u64)(ih * iw)) > (1 << 25) ||
          ((u64)((u64)(n * oc) * ((u64)ih * iw))) > ( 1 << 25)) {
    n = n / 2;
    ic = ic / 2 + 1;
    oc = oc / 2 + 1;
  }
}

void random_param(
    int &n, int &c, int &h, int &w,
    int &kh, int &kw, int &ph, int &pw, int &sh, int &sw,
    int &oc) {
  int oh, ow;
  static bool big_n = true;
  int n_max, c_max, h_max, w_max;
  if (big_n) {
    n_max = 64; c_max = 64;
    h_max = 256; w_max = 256;
  } else {
    n_max = 16; c_max = 16;
    h_max = 256; w_max = 256;
  }
  big_n = !big_n;

again:

  n = rand() % n_max;
  if ( n == 0 ) {
    n++;
  }
  c = rand() % c_max;
  if ( c == 0 ) {
    c++;
  }
  if ((rand() % 2) == 0 ) {
    kh = rand() % 4;
    kw = rand() % 4;
    if (kh == 0)
      kh++;
    if (kw == 0)
      kw++;
  } else if ((rand() % 3) == 0) {
    kh = rand() % 8;
    kw = rand() % 8;
    if (kh == 0)
      kh++;
    if (kw == 0)
      kw++;
  } else {
    kh = rand() % 16;
    kw = rand() % 16;
    if (kh == 0)
      kh++;
    if (kw == 0)
      kw++;
  }
  kh = kh + 1;
  kw = kw + 1;

  ph = rand() % (kh / 2);
  pw = rand() % (kw / 2);

  if ( pw == 0 ) {
    pw++;
  }
  if ( ph == 0 ) {
    ph++;
  }

  if (kh == 2 * ph)
    kh = kh + 1;
  if (kw == 2 * pw)
    kw = kw + 1;
  h = rand() % h_max;
  if ( h == 0 ) {
    h++;
  }
  w = rand() % w_max;
  if ( w == 0) {
    w++;
  }
  if ( h <= 2 * kh ) {
    h = 2 * kh;
  }
  if ( w <= 2 * kw ) {
    w = 2 * kw;
  }

  sh = (rand() % 2 * kh) % 15;
  sw = (rand() % 2 * kw) % 15;
  if ( sh == 0 ) {
    sh++;
  }
  if ( sw == 0 ) {
    sw++;
  }

  if ((h + 2 * ph - kh) % sh != 0)
    h = h - (h + 2 * ph - kh) % sh;
  if ((w + 2 * pw - kw) % sw != 0)
    w = w - (w + 2 * pw - kw) % sw;

  if (n * c * h * w > 1024 * 1024 * 4)
    goto again;

  oc = rand() % c_max;
  if ( oc == 0 ) {
    oc++;
  }
  oh = (h + 2 * ph - kh) / sh + 1;
  ow = (w + 2 * pw - kw) / sw + 1;

  if (n * oc * oh * ow > 1024 * 1024 * 4)
    goto again;

  /* for 1x1 conv */
  if (n * oc * h * w > 1024 * 1024 * 4)
    goto again;
}

