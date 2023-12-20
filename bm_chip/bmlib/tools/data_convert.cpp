#include "string.h"
#include "stdlib.h"
#include "stdio.h"
#include "common.h"
#include <sstream>
int main(int argc, char * argv[])
{
    IF_VAL if_val;
    if ( argc < 3 ){
        printf("usage 1 for flt2hex 0 for hex2flt string\n");
        return 0;
    }

    if ( strcmp(argv[1], "1") == 0 ){
        float a = atof(argv[2]);
        if_val.fval = a;
        printf("The integer is %08x\n", if_val.ival);
    }
    else{
        unsigned int a ;
        sscanf(argv[2], "%x", &a);
        if_val.ival = a;
        printf("The float is %8.8f\n", if_val.fval);
    }
    return 1;

}
