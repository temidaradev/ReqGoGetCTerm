#include "libfetch.h"
#include <stdio.h>
#include <unistd.h>

int main() {
  while (1) {
    char *data = FetchData();
    printf("Received: %s\n", data);
    sleep(1);
  }
  return 0;
}
