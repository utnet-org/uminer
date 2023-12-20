#include <stdio.h>
#include <stdlib.h>
#define MAX_LEN                        10000
#define FLOAT_SIZE                     4
#define LOCAL_MEM_SIZE                 (1<<CONFIG_LOCAL_MEM_ADDRWIDTH)
#define NPU_NUM                        CONFIG_NPU_NUM

//#define CONFIG_GLOBAL_MEM_SIZE         0x10000000
//#define DIFF_DUMP_DATA_PATH            "/home/jinli/Git/bmdnn/build/test/"

#define GLOBAL_MEM_SIZE                CONFIG_GLOBAL_MEM_SIZE
#define DDR_DUMP
#ifndef DUMP_BIN
#define DUMP_BIN
#endif

int main(int argc, char**argv){
  argc = 4;

  char filename[MAX_LEN];
  char filename_index[MAX_LEN];
  char filename_tar[MAX_LEN];
  char filename_diff[MAX_LEN];
  int index = 0;
#ifdef DDR_DUMP
  unsigned int *mem_history =new unsigned int [GLOBAL_MEM_SIZE/FLOAT_SIZE];
  unsigned int *mem =new unsigned int [GLOBAL_MEM_SIZE/FLOAT_SIZE];
  int index_data = -1;
#else
  unsigned int mem_history[LOCAL_MEM_SIZE/FLOAT_SIZE*NPU_NUM];
  unsigned int mem[LOCAL_MEM_SIZE/FLOAT_SIZE];
#endif
  printf("%s %s %s %s\n",argv[1],argv[2],argv[3],argv[4]);
  if(argv[4] != NULL){
    index = 1;
#ifdef DDR_DUMP
    index_data = atoi(argv[4]);
#endif
  }
  int num_nodechip = atoi(argv[2]);
  int num_cycle = atoi(argv[3]);
  char path[MAX_LEN] = DIFF_DUMP_DATA_PATH;

  snprintf(filename_index, MAX_LEN, "%s%s_index_%d.dat",path,argv[1],num_nodechip);
  snprintf(filename_tar, MAX_LEN, "%s_%d_%d.dat",argv[1],num_nodechip,num_cycle);

  printf("filename_index = %s\n",filename_index);
  printf("filename_tar = %s\n",filename_tar);

  FILE *pt = fopen(filename_index,"r");
  FILE *cyclefile;
  FILE *pt_t  = NULL;
  if(!index) pt_t = fopen(filename_tar,"w");
  int diff_top = 0;
  int diff_bot = 0;

  snprintf(filename, MAX_LEN, "%s%s_%d_%d.dat",path,argv[1],num_nodechip,1); 
  printf("filename = %s\n",filename);
#ifndef DDR_DUMP
  cyclefile = fopen(filename,"r");
#ifdef DUMP_BIN
    fread(mem_history,1,LOCAL_MEM_SIZE*NPU_NUM,cyclefile);
#else
  for (int i = 0;i<LOCAL_MEM_SIZE/FLOAT_SIZE*NPU_NUM;i++){
    fscanf(cyclefile,"%x\n",&mem_history[i]);
  }
#endif  
  
  fclose(cyclefile);

  for(int i = 0;i<num_cycle;i++){
    snprintf(filename_diff, MAX_LEN, "%s%s_%d_%d.dat",path,argv[1],num_nodechip,i+1);
    for(int k = 0;k<NPU_NUM;k++){
//#ifdef DUMP_BIN
//      fread(&diff_top, sizeof(unsigned int), 1, pt);
//      fread(&diff_bot, sizeof(unsigned int), 1, pt);
//#else
      fscanf(pt,"%x\n",&diff_top);
      fscanf(pt,"%x\n",&diff_bot);
//#endif
      if((unsigned int)diff_top!=0xffffffff&&(unsigned int)diff_bot!=0xffffffff){

        FILE *p_diff = fopen(filename_diff,"r");
#ifdef DUMP_BIN
            fread(mem, 1,((diff_bot-diff_top)+1)*FLOAT_SIZE ,p_diff);
#else
          for(int l = 0;l<(int)(diff_bot-diff_top)+1;l++){
            fscanf(p_diff,"%x\n",&mem[l]);
          }
#endif
          for(int l = 0;l<(int)(diff_bot-diff_top)+1;l++){
            mem_history[diff_top+l] = mem[l];
          }
        fclose(p_diff);

      }

//#endif
    }
  }
#ifdef DUMP_BIN
    fwrite (mem_history, 1, LOCAL_MEM_SIZE*NPU_NUM,pt_t);
#else
  for(int i = 0;i<LOCAL_MEM_SIZE/FLOAT_SIZE*NPU_NUM;i++){
    fprintf(pt_t,"%x\n",mem_history[i]);
  }
#endif
  
#else//-------------------------------------------------------------------------
  cyclefile = fopen(filename,"r");
#ifdef DUMP_BIN
  fread(mem_history,1,GLOBAL_MEM_SIZE,cyclefile);
#else
  for (int i = 0;i<GLOBAL_MEM_SIZE/FLOAT_SIZE;i++){
    fscanf(cyclefile,"%x\n",&mem_history[i]);
  }
#endif
  if(index)printf("index %d data = %8.8x cycle %d \n",index_data,mem_history[index_data-1],1); 
  fclose(cyclefile);

 for(int i = 0;i<num_cycle;i++){
    snprintf(filename_diff, MAX_LEN, "%s%s_%d_%d.dat",path,argv[1],num_nodechip,i+1);
//#ifdef DUMP_BIN
//    fread(&diff_top, sizeof(unsigned int), 1, pt);
//    fread(&diff_bot, sizeof(unsigned int), 1, pt);
//#else
    fscanf(pt,"%x\n",&diff_top);
    fscanf(pt,"%x\n",&diff_bot);
//#endif
    if((unsigned int)diff_top!=0xffffffff&&(unsigned int)diff_bot!=0xffffffff){
      FILE *p_diff = fopen(filename_diff,"r");
#ifdef DUMP_BIN
          fread(mem, 1,((diff_bot-diff_top)+1)*FLOAT_SIZE ,p_diff);
#else
        for(int l = 0;l<(int)(diff_bot-diff_top)+1;l++){
          fscanf(p_diff,"%x\n",&mem[l]);
        }
#endif
        for(int l = 0;l<(int)(diff_bot-diff_top)+1;l++){
          mem_history[diff_top+l] = mem[l];
//          printf("l = %d %x\n",l,mem[l]);
        }
      fclose(p_diff);
      if(index)printf("index %d data = %8.8x cycle %d \n",index_data,mem_history[index_data-1],i+1);
    }
  }
#ifdef DUMP_BIN
  if(!index)fwrite (mem_history, 1, GLOBAL_MEM_SIZE,pt_t);
#else
  if(!index){
    for(int i = 0;i<GLOBAL_MEM_SIZE/FLOAT_SIZE;i++){
      fprintf(pt_t,"%8.8x\n",mem_history[i]);
    }
  }
#endif
  delete[] mem_history;
  delete[] mem;
#endif// DDR_DUMP-----------------------------------------------------------
  fclose(pt);
  if(!index)fclose(pt_t);
  return true;
}
