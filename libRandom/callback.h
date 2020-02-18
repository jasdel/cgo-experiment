#ifndef CALLBACK_H
#define CALLBACK_H

typedef int (*callback_t)(int id, int i);

void* init_work(int id, int calls, callback_t cb);
void cleanup(void* thread_id);

#endif // CALLBACK_H
