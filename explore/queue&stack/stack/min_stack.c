#include <stdlib.h>
#include <string.h>


typedef struct {
    int *data;
    int top_p;
    int size;
    MinStack* minDataStack;
} MinStack;

/** initialize your data structure here. */

#define INIT_STACK_SIZE 1000
#define PER_StACK_INC_SIZE 500

MinStack* minStackCreate() {
    MinStack* minDataStack = (MinStack *)malloc(sizeof(MinStack));
    minDataStack->minDataStack = NULL;
    minDataStack->top_p = -1;
    minDataStack->size = INIT_STACK_SIZE;
    minDataStack->data = (int*)malloc(sizeof(int) * INIT_STACK_SIZE);

    MinStack* minStack = (MinStack *)malloc(sizeof(MinStack));
    minStack->data = (int*)malloc(sizeof(int) * INIT_STACK_SIZE);
    minStack->top_p = -1;
    minStack->size = INIT_STACK_SIZE;
    minStack->minDataStack = minDataStack;
    return minStack;
}

MinStack* increaseMinStackSize(MinStack* obj) {
    int* newData = (int*)malloc(sizeof(int) * (obj->size + PER_StACK_INC_SIZE));
    memcpy(newData, obj->data, sizeof(int) * obj->size);
    free(obj->data);
    obj->data = newData;
    return obj;
}

void minStackPush(MinStack* obj, int x) {
    if (obj->top_p >= obj->size - 1) {
        obj = increaseMinStackSize(obj);
    }
    obj->data[obj->top_p++] = x;
    if (obj->minDataStack->top_p < 0 || obj->minDataStack->data[obj->minDataStack->top_p] > x) {
        if (obj->minDataStack->top_p >= obj->minDataStack->size - 1) {
            obj->minDataStack = increaseMinStackSize(obj->minDataStack);
        }
        obj->minDataStack->data[obj->minDataStack->top_p++] = x;
    }

}

void minStackPop(MinStack* obj) {
    if (obj->top_p < 0) {
        return;
    }
    int popNum = obj->data[obj->top_p--];
    if (popNum == obj->minDataStack->data[obj->minDataStack->top_p]) {
        obj->minDataStack->top_p--;
    }
}

int minStackTop(MinStack* obj) {
    return obj->data[obj->top_p];
}

int minStackGetMin(MinStack* obj) {
    return obj->minDataStack[obj->minDataStack->top_p];
}

void minStackFree(MinStack* obj) {
    free(obj->minDataStack->data);
    free(obj->minDataStack);
    free(obj->data);
    free(obj);
}

/**
 * Your MinStack struct will be instantiated and called as such:
 * MinStack* obj = minStackCreate();
 * minStackPush(obj, x);
 
 * minStackPop(obj);
 
 * int param_3 = minStackTop(obj);
 
 * int param_4 = minStackGetMin(obj);
 
 * minStackFree(obj);
*/