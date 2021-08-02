
#include "MurmurHash3.h"
#include "bridge.h"
#include <stdio.h>
#include "stdlib.h"

long long hash64(const char *key, int len)
{
  long long out[2];
  MurmurHash3_x64_128(key, len, 0, &out);
  return out[0];
}