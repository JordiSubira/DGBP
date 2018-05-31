#include "api_gbp.h"

static const int FAIL_COMMAND = -1;


int checkPolicy( char* eid_src, char* eid_dst )
{

  FILE *fp;
  char path[1035];
  char cmd[64];

  //cmd = sprintf(cmd, "node query.js %s %s",eid_src,eid_dst)
  sprintf(cmd, "python server_app/server.py %s %s",eid_src,eid_dst);

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
