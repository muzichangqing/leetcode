int lengthOfLongestSubstring(char* s) {
    int set[255];
    for(int i = 0; i < 255; i++) {
        set[i] = 0;
    }

    int i, j, count;
    i = j = count = 0;
    while(s[j] != '\0') {
        if (!set[s[j]]) {
            set[s[j++]] = 1;
            if (j - i > count) {
                count = j - i;
            }
        } else {
            set[s[i++]] = 0;
        }
    }

    return count;
}