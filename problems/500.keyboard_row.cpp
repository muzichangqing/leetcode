class Solution {
public:
    vector<string> findWords(vector<string>& words) {
        std::vector<string> keyboard;
        keyboard.push_back(string("qwertyuiop"));
        keyboard.push_back(string("asdfghjkl"));
        keyboard.push_back(string("zxcvbnm"));
        int size = words.size();
        std::vector<string> result;
        for(int i = 0; i < size; i++) {
            string str = words[i];
            char first_c = str[0];
            if (first_c <= 90) first_c += 32;
            int line = keyboard[0].find(first_c) != string::npos ? 0 :
                (keyboard[1].find(first_c) != string::npos ? 1 : 2);
            int str_size = str.length();
            bool can_push = true;
            for (int j = 1; j < str_size; j++) {
                char c = str[j];
                if (c <= 90) c += 32;
                int line_c = keyboard[0].find(c) != string::npos ? 0 :
                    (keyboard[1].find(c) != string::npos ? 1 : 2);
                if (line != line_c) {
                    can_push = false;
                    break;
                }
            }
            if (can_push) result.push_back(str);
        }
        return result;
     }
};
