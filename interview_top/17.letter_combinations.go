package interviewtop

import "strings"

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }
    result := make([]string, 1)
    for _, b := range []byte(digits) {
        letters := letterMap[int(b-'0')] 
        lettersLen := len(letters)
        newResult := make([]string, len(result) * lettersLen)
        for i, str := range result {
            for j, letter := range letters {
                var builder strings.Builder
                builder.Grow(len(str)+1)
                builder.WriteString(str)
                builder.WriteByte(letter)
                newResult[i*lettersLen+j] = builder.String()
            }
        }
        result = newResult
    }
    return result 
}

var letterMap = [][]byte{
    nil,
    nil,
    {'a', 'b', 'c'},
    {'d', 'e', 'f'},
    {'g', 'h', 'i'},
    {'j', 'k', 'l'},
    {'m', 'n', 'o'},
    {'p', 'q', 'r', 's'},
    {'t', 'u', 'v'},
    {'w', 'x', 'y', 'z'},
}

