public class Solution {    
    public bool IsValid(string s) {
        Stack<Parenthese> stack = new Stack<Parenthese>();
        foreach(char symbol in s) {
            Parenthese parenthese = new Parenthese();
            parenthese.Init(symbol);

            if (stack.Count == 0) {
                stack.Push(parenthese);
            } else {
                Parenthese last_p = stack.Peek();
                if (last_p.type == parenthese.type && last_p.toward + parenthese.toward == 0) {
                    stack.Pop();
                } else {
                    stack.Push(parenthese);
                }
            }
        }
        return stack.Count == 0;
    }

    struct Parenthese
    {
        public char symbol;

        public int toward;

        public int type;

        public void Init(char symbol) {
            this.symbol = symbol;
            switch (symbol) { 
                case '(':
                    this.type = 1;
                    this.toward = -1;
                    break;
                case '[':
                    this.type = 2;
                    this.toward = -1;
                    break;
                case '{':
                    this.type = 3;
                    this.toward = -1;
                    break;
                case ')':
                    this.type = 1;
                    this.toward = 1;
                    break;
                case ']':
                    this.type = 2;
                    this.toward = 1;
                    break;
                case '}':
                    this.type = 3;
                    this.toward = 1;
                    break;
                default:
                    this.toward = 0;
                    this.type = 0;
                    break;
            }
        }
    }
}