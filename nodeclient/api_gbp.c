#include "api_gbp.h"

static const int FAIL_COMMAND = -1;


int checkPolicy( char* eid_dst, char* eid_src )
{

  FILE *fp;
  char path[1035];
  char cmd[64];

  //cmd = sprintf(cmd, "node query.js %s %s",eid_src,eid_dst)
  sprintf(cmd, "node queryPolicy.js %s %s",eid_dst,eid_src);

  // Open the command for reading.
  fp = popen(cmd, "r");
  if (fp == NULL) {
    printf("Failed to run command\n" );
    return FAIL_COMMAND;
  }

  // Read the output a line at a time - output it. 
  while (fgets(path, sizeof(path)-1, fp) != NULL) {
    printf("%s", path);
  }
  // close 
  return WEXITSTATUS(pclose(fp));
  
}
