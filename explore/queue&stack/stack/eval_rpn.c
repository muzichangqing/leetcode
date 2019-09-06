#include <stdlib.h>
#include <string.h>

int evalRPN(char ** tokens, int tokensSize){
    int *stack = (int *)malloc(tokensSize * sizeof(int));
    int result = 0;
    int stackTopP = -1;
    for (int i = 0; i < tokensSize; i++) {
        char *str = tokens[i];
        if (strcmp(str, "+") == 0) {
            int opNum2 = stack[stackTopP--];
            int opNum1 = stack[stackTopP--];
            stack[++stackTopP] = opNum1 + opNum2;
        } else if (strcmp(str, "-") == 0) {
            int opNum2 = stack[stackTopP--];
            int opNum1 = stack[stackTopP--];
            stack[++stackTopP] = opNum1 - opNum2;
        } else if (strcmp(str, "*") == 0) {
            int opNum2 = stack[stackTopP--];
            int opNum1 = stack[stackTopP--];
            stack[++stackTopP] = opNum1 * opNum2;
        } else if (strcmp(str, "/") == 0) {
            int opNum2 = stack[stackTopP--];
            int opNum1 = stack[stackTopP--];
            stack[++stackTopP] = opNum1 / opNum2;
        } else {
            stack[++stackTopP] = atoi(str);
        }
    }
    result = stack[stackTopP];
    free(stack);
    return result;
}