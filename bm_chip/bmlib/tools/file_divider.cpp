#include "stdio.h"
#include "stdlib.h"
#include "string.h"
#define  STR_MAX_LEN 1024
int main(int argc , char *argv[])
{
  FILE * fin;
  FILE * fout;
  int    status = 0;
  if ( argc < 4 )
  {
    printf("Usage : input_file_name total bytes, number sections\n");
    exit(-1);
  }
  
  fin = fopen(argv[1], "rb");
  if ( fin == NULL )
  {
       printf("	The input file does not exist\n");
       exit(-1);
  }
  
  
  int idx;
  int idx_len = strlen(argv[2]);
  unsigned long total_len = 0;
  for ( idx = 0 ; idx < idx_len; idx++ )
  {
       total_len *= 10;
       total_len += argv[2][idx] - '0';
  }

  int sec_num = atoi(argv[3]);
  int sec_len;
  printf("The total bytes %ld %d\n", total_len ,sec_num);
  if ( total_len%sec_num != 0 )
  {
      printf("The total_len should be divided by sec_num\n");
      exit(-1);
  }
  else
  {
      sec_len = total_len / sec_num;
  }

  int  sec_idx;
  char string_buf[STR_MAX_LEN];
  unsigned int *data_buf = new unsigned int[sec_len/4];
  for ( sec_idx = 0; sec_idx < sec_num; sec_idx++ )
  {
        snprintf(string_buf, STR_MAX_LEN, "%s%d", argv[1],sec_idx);
        fout = fopen(string_buf,"wb");
        status = fread(data_buf, sec_len, 1, fin);
        fwrite(data_buf, sec_len, 1, fout);
        printf("Section %d done\n", sec_idx);
        fclose(fout);
  }

  delete [] data_buf;

  return status;
}
