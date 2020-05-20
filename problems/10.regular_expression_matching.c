#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

bool isMatchDp(int i, int j, char *s, size_t ss, char *p, size_t ps, int **memo);

bool isMatch(char *s, char *p){
    size_t ss = strlen(s);
    size_t ps = strlen(p);
    int **memo = (int **)malloc(sizeof(int *) * (ss + 1));
    for (int i = 0; i <= ss; i++) {
        memo[i] = (int *)malloc(sizeof(int) * (ps + 1));
    }
    for (int i = 0; i <= ss; i++) {
        for (int j = 0; j <= ps; j++) {
            memo[i][j] = -1;
        }
    }
    bool is_match = isMatchDp(0, 0, s, ss, p, ps, memo);
    for (int i = 0; i <= ss; i++) {
        free(memo[i]);
    }
    free(memo);
    return is_match;
}

bool isMatchDp(int i, int j, char *s, size_t ss, char *p, size_t ps, int **memo) {
    if (memo[i][j] != -1) {
        return memo[i][j] == 1;
    }

    bool is_match = false;
    if (j == ps) {
        is_match = i == ss;
    } else {
        bool first_match = (i < ss && (s[i] == p[j] || p[j] == '.'));
        if (p[j+1] == '*') {
            is_match = (isMatchDp(i, j+2, s, ss, p, ps, memo) || (first_match && isMatchDp(i+1, j, s, ss, p, ps, memo)));
        } else {
            is_match = first_match && isMatchDp(i+1, j+1, s, ss, p, ps, memo);
        }
    }
    memo[i][j] = is_match ? 1 : 0;
    return is_match;
}