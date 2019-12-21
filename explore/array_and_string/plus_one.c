#include <stdlib.h>
#include <string.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* plusOne(int* digits, int digitsSize, int* returnSize){
    int *result = (int *)malloc(sizeof(int) * digitsSize);
    memcpy(result, digits, digitsSize * sizeof(int));
    *returnSize = digitsSize;
    int carry = 1;
    for (int i = digitsSize - 1; i >= 0; i--) {
        int sum = digits[i] + carry;
        result[i] = sum > 9 ? 0 : sum;
        carry = sum > 9 ? 1 : 0;
        if (carry == 0) {
            break;
        }
    }
    if (carry == 1) {
        int *newResult = (int *)malloc(sizeof(int) * (digitsSize + 1));
        *returnSize = digitsSize + 1;
        int *cpyPtr = newResult;
        memcpy(++cpyPtr, result, sizeof(int) * digitsSize);
        newResult[0] = carry;
        free(result);
        result = newResult;
    }

    return result;
}