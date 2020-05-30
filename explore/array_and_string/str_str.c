#include <string.h>

int strStr(char *haystack, char *needle) {
    if (*haystack == '\0' && *needle == '\0') {
        return 0;
    }
    char *cursor = haystack;
    size_t needle_len = strlen(needle);
    size_t needle_size = needle_len * sizeof(char);
    size_t haystack_len = strlen(haystack);
    int index = 0;
    while(*cursor != '\0') {
        if (haystack_len - index < needle_len) {
            break;
        }
        if (memcmp(cursor, needle, needle_size) == 0) {
            return index;
        }
        cursor++;
        index++;
    }
    return -1;
}