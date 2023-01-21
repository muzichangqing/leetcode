package interviewtop

func isValid(s string) bool {
    stack := make([]byte, len(s))
    top := -1
    for _, b := range []byte(s) {
        if b == '(' || b == '[' || b == '{' {
            top++
            stack[top] = b
        } else {
            if b == ')' {
                if top > -1 && stack[top] == '(' {
                    top--
                } else {
                    return false
                }
            } else if b == ']' {
                if top > -1 && stack[top] == '[' {
                    top--
                } else {
                    return false
                }
            } else {
                if top > -1 && stack[top] == '{' {
                    top--
                } else {
                    return false
                }
            }
        }
    }
    return top == -1
}
