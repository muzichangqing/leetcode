class Solution:
    def decodeString(self, s: str) -> str:
        decode_strs = [];
        stack = [];
       
        for c in s:
            if c.isdigit() or c == '[':
               stack.append(c)
            elif c == ']':
                p_c = stack.pop()
                repeat_chars = []
                while p_c != '[':
                    repeat_chars.append(p_c)
                    p_c = stack.pop()
                repeat_num = 0;
                repeat_num_b = 1;
          
                while len(stack) > 0 and stack[len(stack) - 1].isdigit():
                    repeat_num_c = stack.pop()
                    repeat_num += int(repeat_num_c) * repeat_num_b
                    repeat_num_b *= 10

                repeat_chars.reverse()
                repeat_str =  "".join(repeat_chars) * int(repeat_num)
                if len(stack) > 0:
                    stack.append(repeat_str)
                else:
                    decode_strs.append(repeat_str)
            else:
                if len(stack) > 0:
                    stack.append(c)
                else:
                    decode_strs.append(c)

        return "".join(decode_strs)
    

if __name__ == '__main__':
    obj = Solution()
    print(obj.decodeString("30[a2[c]]"))
