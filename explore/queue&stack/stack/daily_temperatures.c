#include <stdlib.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* dailyTemperatures(int* T, int TSize, int* returnSize){
    int* stack = (int*)malloc(TSize * sizeof(int));
    int top = -1;
    int* result = (int*)malloc(TSize * sizeof(int));
    *returnSize = TSize;
    if (TSize == 1) {
        result[0] = 0;
        return result;
    }
    result[TSize - 1] = 0;
    for(int i = 0; i < TSize - 1; i++) {
        int tomT = T[i+1];
        stack[++top] = i;
        result[i] = 0;
        while(top > -1) {
            int topI = stack[top];
            if (T[topI] < tomT) {
                result[topI] = (i+1) - topI;
                top--;
            } else {
                break;
            }
        }
    }
    free(stack);
    return result;
}

