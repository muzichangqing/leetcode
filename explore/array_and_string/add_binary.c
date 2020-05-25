#include <stdlib.h>
#include <string.h>

char * addBinary(char * a, char * b){
    size_t a_size = strlen(a);
    size_t b_size = strlen(b);
    size_t ans_size = a_size > b_size ? a_size + 1 : b_size + 1;
    char *ans = (char *)malloc(sizeof(char) * (ans_size + 2));

    


    return ans;
}
