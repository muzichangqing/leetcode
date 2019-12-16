#include <string.h>
#include <stdlib.h>

int isValid(char * s){
    int strSize = strlen(s);
    size_t strMemSize = strSize * sizeof(char);
    if (strMemSize == 0) return 1;
    char* stack = (char*)malloc(strMemSize);
    int top = -1;
    for(int i = 0; i < strSize; i++) {
        char c = s[i];
        if (top != -1 && stack[top] != c && abs(c - stack[top]) < 3) top--;
        else stack[++top] = c;
    }
    free(stack);
    return top == -1;
}
