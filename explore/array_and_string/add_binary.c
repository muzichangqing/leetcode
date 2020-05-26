#include <stdlib.h>
#include <string.h>

char * addBinary(char * a, char * b){
    size_t a_size = strlen(a) + 1;
    size_t b_size = strlen(b) + 1;
    size_t ans_size;
    char *ans;
    char *tmp;
    char *cursor1;
    char *cursor2;

    if (a_size > b_size) {
        ans_size = a_size + 1;
        ans = (char *)malloc(sizeof(char) * ans_size);
        memcpy(ans+1, a, a_size * sizeof(char));
        tmp = b;
        cursor2 = b + b_size - 2;
    } else {
        ans_size = b_size + 1;
        ans = (char *)malloc(sizeof(char) * ans_size);
        memcpy(ans+1, b, b_size * sizeof(char));
        tmp = a;
        cursor2 = a + a_size - 2;
    }

    *ans = '0';
    char i = 0;
    cursor1 = ans + ans_size - 2; 
    while((tmp - 1) != cursor2) {
        *cursor1 = *cursor1 + *cursor2 - 96 + i;
        if (*cursor1 > 1) {
            *cursor1 = *cursor1 - 2 + 48;
            i = 1;
        } else {
            *cursor1 = *cursor1 + 48;
            i = 0;
        }
        cursor1--;
        cursor2--;
    }

    while((ans - 1) != cursor1 && i > 0) {
         *cursor1 = *cursor1 - 48 + i;
        if (*cursor1 > 1) {
            *cursor1 = '0';
            i = 1;
        } else {
            *cursor1 = *cursor1 + 48;
            i = 0;
        }
        cursor1--;
    }

    if (*ans == '0') {
        char *new_ans = (char *)malloc(sizeof(char) * (ans_size - 1));
        memcpy(new_ans, ans + 1, (ans_size - 1) * sizeof(char));
        free(ans);
        ans = new_ans;
    } 

    return ans;
}