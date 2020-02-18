
#include "random.h"

#include <stdlib.h>

int myRandom() {
  return random();
}

void (*callback)() = 0;

void registerCallback( void (*fnPtr)()) {
    callback = fnPtr;
}


void doCallback() {
     callback();
}
