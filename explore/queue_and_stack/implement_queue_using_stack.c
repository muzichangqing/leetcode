#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

#define STACK_INIT_SIZE 100
#define STACK_INCR_SIZE 50

typedef struct {
    int* m_stack;
    int m_top;
    int m_size;
    int *s_stack;
    int s_size;
} MyQueue;

/** Initialize your data structure here. */

MyQueue* myQueueCreate() {
    MyQueue* obj = (MyQueue *)malloc(sizeof(MyQueue));
    if (!obj) {
        return NULL;
    }
    obj->m_stack = (int *)malloc(sizeof(int) * STACK_INIT_SIZE);
    obj->m_top = -1;
    obj->m_size = STACK_INIT_SIZE;
    obj->s_stack = (int *)malloc(sizeof(int) * STACK_INIT_SIZE);
    obj->s_size = STACK_INIT_SIZE;
    return obj;
}

/** Push element x to the back of queue. */
void myQueuePush(MyQueue* obj, int x) {
    if (obj->m_top == obj->m_size - 1) {
        int new_m_size = obj->m_size + STACK_INCR_SIZE;
        int *new_m_stack = (int *)malloc(sizeof(int) * new_m_size);
        memcpy(new_m_stack, obj->m_stack, sizeof(int) * obj->m_size);
        free(obj->m_stack);
        obj->m_size = new_m_size;
        obj->m_stack = new_m_stack;
    }
    obj->m_stack[++(obj->m_top)] = x;
}

/** Removes the element from in front of queue and returns that element. */
int myQueuePop(MyQueue* obj) {
    if (obj->s_size <= obj->m_top) {
        int new_s_size = obj->s_size + STACK_INCR_SIZE;
        int *new_s_stack = (int *)malloc(sizeof(int) * new_s_size);
        free(obj->s_stack);
        obj->s_size = new_s_size;
        obj->s_stack = new_s_stack;
    }
    int s_top = -1;
    while(obj->m_top != -1) {
        obj->s_stack[++s_top] = obj->m_stack[obj->m_top--];
    }
    int element = obj->s_stack[s_top--];
    while(s_top != -1) {
        obj->m_stack[++(obj->m_top)] = obj->s_stack[s_top--];
    }
    return element;
}

/** Get the front element. */
int myQueuePeek(MyQueue* obj) {
    if (obj->s_size <= obj->m_top) {
        int new_s_size = obj->s_size + STACK_INCR_SIZE;
        int *new_s_stack = (int *)malloc(sizeof(int) * new_s_size);
        free(obj->s_stack);
        obj->s_size = new_s_size;
        obj->s_stack = new_s_stack;
    }
    int s_top = -1;
    while(obj->m_top != -1) {
        obj->s_stack[++s_top] = obj->m_stack[obj->m_top--];
    }
    int element = obj->s_stack[s_top];
    while(s_top != -1) {
        obj->m_stack[++(obj->m_top)] = obj->s_stack[s_top--];
    }
    return element;
}

/** Returns whether the queue is empty. */
bool myQueueEmpty(MyQueue* obj) {
    return obj->m_top == -1;
}

void myQueueFree(MyQueue* obj) {
    free(obj->m_stack);
    free(obj->s_stack);
    free(obj);
}

/**
 * Your MyQueue struct will be instantiated and called as such:
 * MyQueue* obj = myQueueCreate();
 * myQueuePush(obj, x);
 
 * int param_2 = myQueuePop(obj);
 
 * int param_3 = myQueuePeek(obj);
 
 * bool param_4 = myQueueEmpty(obj);
 
 * myQueueFree(obj);
*/