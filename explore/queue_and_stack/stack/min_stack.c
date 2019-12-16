#include <stdlib.h>
#include <string.h>

typedef struct {
    int size;
    int* data;
    int top;

    int minSize;
    int* minData;
    int minTop;
} MinStack;

#define INIT_STACK_SIZE 500
#define PER_INC_STACK_SIZE 100
/** initialize your data structure here. */

MinStack* minStackCreate() {
    MinStack* obj = (MinStack*)malloc(sizeof(MinStack));
    
    obj->size = INIT_STACK_SIZE;
    obj->data = (int*)malloc(sizeof(int) * obj->size);
    obj->top = -1;

    obj->minSize = INIT_STACK_SIZE;
    obj->minData = (int*)malloc(sizeof(int) * obj->minSize);
    obj->minTop = -1;

    return obj;
}

MinStack* minStackIncSize(MinStack* obj, int type) {
    if (type == 1) {
        int newSize = obj->size + PER_INC_STACK_SIZE;
        int* newData = (int*)malloc(sizeof(int) * newSize);
        memcpy(newData, obj->data, obj->size * sizeof(int));
        free(obj->data);
        obj->data = newData;
        obj->size = newSize;
    } else {
        int newSize = obj->minSize + PER_INC_STACK_SIZE;
        int* newData = (int*)malloc(sizeof(int) * newSize);
        memcpy(newData, obj->minData, obj->minSize * sizeof(int));
        free(obj->minData);
        obj->minData = newData;
        obj->minSize = newSize;
    }
    return obj;
}

void minStackPush(MinStack* obj, int x) {
    if (obj->top == obj->size - 1) {
        minStackIncSize(obj, 1);
    }
    obj->data[++(obj->top)] = x;
    if (obj->minTop < 0 || obj->minData[obj->minTop] >= x) {
        if (obj->minTop == obj->minSize - 1) minStackIncSize(obj, 2);
        obj->minData[++(obj->minTop)] = x;
    }
}

void minStackPop(MinStack* obj) {
    if (obj->top < 0) return;
    int topValue = obj->data[(obj->top)--];
    if (obj->minTop < 0) return;
    if (obj->minData[obj->minTop] == topValue) (obj->minTop)--;
}

int minStackTop(MinStack* obj) {
    return obj->data[obj->top];
}

int minStackGetMin(MinStack* obj) {
    return obj->minData[obj->minTop];
}

void minStackFree(MinStack* obj) {
    free(obj->data);
    free(obj->minData);
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