/**
 * Return an array of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* shortestToChar(char* S, char C, int* returnSize) {
    *returnSize = strlen(S);
    int *result = (int *)malloc(*returnSize * sizeof(int));
    int last_c_offset;
    int current_c_offset;
    
    last_c_offset = current_c_offset = -1;

    for (int i = 0; i < *returnSize; i++) {
        if (S[i] == C) {
            result[i] = 0;
            last_c_offset = current_c_offset;
            current_c_offset = i;
            for (int j = last_c_offset + 1; j < current_c_offset; j++) {
                if (last_c_offset >= 0) {
                    int distance1 = j - last_c_offset;
                    int distance2 = current_c_offset - j;
                    result[j] = distance1 < distance2 ? distance1 : distance2;
                } else {
                    result[j] = current_c_offset - j;
                }
                
            }
        }
    }

    if (current_c_offset != *returnSize) {
        for (int i = current_c_offset + 1; i < *returnSize; i++) {
            result[i] = i - current_c_offset;
        }
    }

    return result;
}
