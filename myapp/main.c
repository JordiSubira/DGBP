#include <stdio.h>
#include "api_gbp.h"

int main(int argc, char* argv[]) {

int r = checkPolicy("0.0.0.1", "0.0.0.0");


//printf("The return code is %d \n", r);
printf("The return code is %d \n", r);

return 0;
}
