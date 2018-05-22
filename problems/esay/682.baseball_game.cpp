class Solution {
public:
    int strtoint(string str) {
        int sign = 1;
        int length = str.length();
        int result = 0;
        for (int i = 0; i < length; i++) {
            char c = str[i];
            if (c == '-') sign = -1;
            else {
                result = 10 * result + c - 48;
            }
        }
        return result * sign;
    }
    int calPoints(vector<string>& ops) {
        vector<int> points;
        string plu("+");
        string dou("D");
        string cle("C");
        for (auto it = ops.cbegin(); it != ops.cend(); it++) {
            if (*it == plu) {
                int size = points.size();
                points.push_back(points[size - 1] + points[size - 2]);
            } else if (*it == dou) {
                points.push_back(points.back() * 2);
            } else if (*it == cle) {
                points.pop_back();
            } else {
                points.push_back(strtoint(*it));
            }
        }
        int result = 0;
        for (auto it = points.cbegin(); it != points.cend(); it++){
            result += *it;
        }
        return result;
    }
};
