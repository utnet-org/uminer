#include "stdio.h"
#include "stdlib.h"

int main(int argc, char * argv[])
{
    int rd_status = 1;
    if ( argc < 5 )
    { 
        printf("The arguments should be ifilename, ofilename, word_num, array_name\n");
    }

    FILE * inf  = fopen(argv[1], "r");
    if ( inf == NULL )
    {
        printf("The input file does not exist\n");
        return -1;  
    }
    FILE * outf = fopen(argv[2], "w");
    if ( outf == NULL )
    {
        printf("The output file does not exist\n");
        return -1;
    }

    int word_num = atoi(argv[3]);
    fprintf(outf, "unsigned int  %s[%d] = {\n", argv[4], word_num);
    for ( int word_idx = 0; word_idx < word_num; word_idx++ )
    {
         unsigned int rd_data;

         rd_status = fscanf(inf,"%x\n", &rd_data);
         if ( word_idx < word_num - 1 )
         {
             fprintf(outf, "0x%8.8x,\n", rd_data);
         }
         else
         {
             fprintf(outf, "0x%8.8x\n", rd_data);
         }
    }
    fprintf(outf, "};\n");
    fprintf(outf, "#define CONSTANT_ENTRY_SFU_TBL %d\n",word_num);
    printf("done\n");
	return rd_status;
}


