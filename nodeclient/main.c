#include <stdio.h>
#include "api_gbp.h"

int main(int argc, char* argv[]) {

char eid[32];

for (int i=0; i<500;i++){

 	sprintf(eid, "0.0.%d",i);

	int r = createUser("PKI",eid,"Dep1");

	//printf("The return code is %d \n", r);
	printf("The return code is %d \n", r);
}



return 0;
}