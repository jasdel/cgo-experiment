#include "callback.h"

#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

#include <time.h>

typedef struct user_data_t {
  int id;
  int calls;
  callback_t cb;
} user_data_t;


void *my_thread(void *user_data) {
  user_data_t* data = (user_data_t*)(user_data);

  double time_taken = 0;
  for (int i = 0; i < data->calls; i++) {
    struct timespec tstart={0,0}, tend={0,0};
    clock_gettime(CLOCK_MONOTONIC, &tstart);

    // clock is not accurate, want actual time.
    //clock_t start = clock();
    data->cb(data->id, i);

    //clock_t end = clock();
    clock_gettime(CLOCK_MONOTONIC, &tend);

    double taken = ((double)tend.tv_sec + 1.0e-9*tend.tv_nsec) -
           ((double)tstart.tv_sec + 1.0e-9*tstart.tv_nsec);
    //printf("some_long_computation took about %.6f seconds\n", taken);

    time_taken += taken;

    //time_taken += difftime(end, start)/(double)(CLOCKS_PER_SEC); // in seconds
  }

  printf("%d: average time over %d callback: %f (seconds)\n",
      data->id, data->calls, time_taken/(double)(data->calls));
  free(data);

  return NULL;
}

void* init_work(int id, int calls, callback_t cb) {
  pthread_t* thread_id = malloc(sizeof(pthread_t));
  user_data_t* user_data = malloc(sizeof(user_data_t));

  user_data->id = id;
  user_data->cb = cb;
  user_data->calls = calls;

  pthread_create(thread_id, NULL, my_thread, user_data);
  
  return thread_id;
}

void cleanup(void* thread_id) {
  pthread_t* id = (pthread_t*)(thread_id);
  int res = pthread_join(*id, NULL);
  if (res != 0) {
    printf("failed to wait for thread to stop\n");
  }
  free(id);
}
