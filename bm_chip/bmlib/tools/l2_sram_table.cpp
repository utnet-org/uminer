#include <cassert>
#include <cmath>
#include <cstdio>
#include <cstdlib>
#include <cstring>
#include "memmap.h"

enum table_type { EX_INT, EX, LNX, EX_TAILOR, LNX_TAILOR, ARCSINX_TAILOR, UINT2FLOAT, INT2FLOAT, SINX, WARP };

int construct_table(FILE *pFile, float *table, int size, enum table_type type) {
    int level, idx;
    float coeff;
    if (type != WARP) {
        memset(table, 0, sizeof(float) * size);
    }
    switch (type) {
        case EX_INT:
            for (int i = -103; i <= 88; i++) {
                unsigned char j = (unsigned char)((i)&0xff);
                table[j]        = float(i);
            }
            fprintf(pFile, "/*sfu table ex int*/\n");
            break;
        case EX:
            for (int i = -103; i <= 88; i++) {
                unsigned char j = (unsigned char)((i)&0xff);
                table[j]        = exp(i);
            }
            fprintf(pFile, "/*sfu table ex*/\n");
            break;
        case LNX:
            for (int i = 0; i < size; i++) {
                table[i] = log(2.0) * (((float)i) - 127.0);
            }
            fprintf(pFile, "/*sfu table lnx*/\n");
            break;
        case EX_TAILOR:
            for (level = 1; level <= SFU_TAILOR_L_TABLE_SIZE; level++) {
                coeff = 1.0f;
                for (idx = 1; idx <= level; idx++) {
                    coeff /= (float)idx;
                }
                table[level - 1] = coeff;
            }
            fprintf(pFile, "/*sfu tailor table ex*/\n");
            break;
        case LNX_TAILOR:
            for (level = 1; level <= SFU_TAILOR_L_TABLE_SIZE; level++) {
                if (level % 2 == 0)
                    coeff = -1.0f;
                else
                    coeff = 1.0f;
                coeff /= (float)level;
                table[level - 1] = coeff;
            }
            fprintf(pFile, "/*sfu tailor table lnx*/\n");
            break;
        case ARCSINX_TAILOR:
            for (level = 1; level <= SFU_TAILOR_L_TABLE_SIZE; level++) {
                float de = 1.0;
                float nu = 1.0;
                if (level % 2 == 0) {
                    coeff = 0.0;
                } else {
                    for (idx = 1; idx <= level; idx++) {
                        if (idx == level) break;
                        if (idx % 2 == 0) de *= float(idx);
                        if (idx % 2 == 1) nu *= float(idx);
                    }
                    coeff = nu / de / float(level);
                }
                table[level - 1] = coeff;
            }
            fprintf(pFile, "/*arcsin tailor table arcsin*/\n");
            break;
        case UINT2FLOAT:
            for (int i = 0; i < size; i++) {
                table[i] = (float)i;
            }
            fprintf(pFile, "/*table uint2float*/\n");
            break;
        case INT2FLOAT:
            for (int i = 0; i < size; i++) {
                signed char j = (signed char)((unsigned char)((i)&0xff));
                table[i]      = (float)j;
            }
            fprintf(pFile, "/*table int2float*/\n");
            break;
        case SINX:
            for (level = 1; level <= SFU_TAILOR_L_TABLE_SIZE; level++) {
                if (level % 2 == 0) {
                    coeff = 0.0f;
                } else {
                    if (level % 4 == 1)
                        coeff = 1.0f;
                    else
                        coeff = -1.0f;

                    for (idx = 1; idx <= level; idx++) coeff /= (float)idx;
                }
                table[level - 1] = coeff;
            }
            fprintf(pFile, "/*tailor table sinx*/\n");
            break;
        case WARP: {
            unsigned char *table_x = (unsigned char *)table;
            unsigned char *table_y = table_x + (WARP_TABLE_SIZE >> 1);

            memset(table_x, 0, sizeof(unsigned char) * size);

            for (int y = 0; y < WARP_MAX_H; y++) {
                for (int x = 0; x < WARP_MAX_W; x++) {
                    table_x[y * WARP_MAX_W + x] = (unsigned char)x;
                    table_y[y * WARP_MAX_W + x] = (unsigned char)y;
                }
            }
            size /= (sizeof(float) / sizeof(unsigned char));
            fprintf(pFile, "/*warp table*/\n");
        } break;
        default:
            assert(0);
    }
    for (int i = 0; i < size; i++) {
        fprintf(pFile, "0x%x,\n", reinterpret_cast<unsigned int &>(table[i]));
    }
    return 0;
}

int main() {
    FILE *pFile;
    char filename[100];
    char *bm_top = getenv("BM_TOP");
    if (!bm_top) {
        printf("BM_TOP env is not set!\n");
        printf("source envsetup.sh first!\n");
        return -1;
    }
    // printf("bm_top : %s\n", bm_top);
    strcpy(filename, bm_top);
    strcat(filename, "/common/bm1684/include/l2_sram_table.h");
    printf("table file : %s\n", filename);

    pFile = fopen(filename, "w");

    fprintf(pFile, "unsigned int l2_sram_table[] = {\n");

    int lookup_table_num   = 256;
    float *table           = new float[lookup_table_num];
    unsigned char *table_x = new unsigned char[WARP_TABLE_SIZE];

    construct_table(pFile, table, SFU_TABLE_SIZE, EX_INT);

    construct_table(pFile, table, SFU_TABLE_SIZE, EX);

    construct_table(pFile, table, SFU_TABLE_SIZE, LNX);

    construct_table(pFile, table, SFU_TAILOR_L_TABLE_SIZE, EX_TAILOR);

    construct_table(pFile, table, SFU_TAILOR_L_TABLE_SIZE, LNX_TAILOR);

    construct_table(pFile, table, SFU_TAILOR_L_TABLE_SIZE, ARCSINX_TAILOR);

    construct_table(pFile, table, UINT2FLOAT_TABLE_SIZE, UINT2FLOAT);

    construct_table(pFile, table, INT2FLOAT_TABLE_SIZE, INT2FLOAT);

    construct_table(pFile, table, SFU_TAILOR_L_TABLE_SIZE, SINX);

    construct_table(pFile, (float *)table_x, WARP_TABLE_SIZE, WARP);

    fprintf(pFile, "};\n");
    fclose(pFile);

    delete[] table;
    delete[] table_x;
    return 0;
}
