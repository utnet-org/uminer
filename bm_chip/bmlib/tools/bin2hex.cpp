#include "stdio.h"
#include "stdlib.h"

int main(int argc , char *argv[])
{
  FILE * fin;
  FILE * fout;
  if (argc < 4)
  {
    printf("Usage : input_file_name output_file_name, number dword(or lines in hex)\n");
    exit(-1);
  }
  fin = fopen(argv[1], "rb");
  if ( fin == NULL )
  {
       printf("	The input file does not exist\n");
       exit(-1);
  }
  fout = fopen(argv[2], "w");
  if ( fin == NULL )
  {
       printf("	The output file does not exist\n");
       exit(-1);
  }

  unsigned int word;
  int          status = 0;
  for ( int word_idx = 0 ; word_idx < atoi(argv[3]); word_idx++ )
  {
       status = fread(&word, 4, 1, fin);
       fprintf(fout, "%8.8x\n", word);  
  }
  fclose(fin);
  fclose(fout);
  return status;
}
