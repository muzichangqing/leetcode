#include <stdlib.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* findDiagonalOrder(int** matrix, int matrixSize, int* matrixColSize, int* returnSize){
    if (matrixSize == 0) {
        *returnSize = 0;
        return NULL;
    }
    *returnSize = matrixSize * (*matrixColSize);
    int *returnLine = (int *)malloc(sizeof(int) * (*returnSize));
    int direction, x, y, index;
    // -1 左下， 1 右上
    direction = 1;
    x = y = index = 0;
    while (index < *returnSize) {
        if (x < matrixSize && y < *matrixColSize) {
            returnLine[index++] = matrix[x][y];
        }
        if (direction > 0) {
            // x - 1, y + 1
            if (x == 0) {
                direction = -1;
            } else {
                x -= 1;
            }
            y += 1;
        } else {
            if ( y == 0) {
                direction = 1;
            } else {
                y -= 1;
            }
            x += 1;
        }
    }
    return returnLine;
}