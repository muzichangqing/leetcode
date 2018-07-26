class Solution {
public:
    string toLowerCase(string str) {
        string rs = "";
        for (int i = 0; i < str.length(); i++) {
        	char c = str[i];
        	if (c >= 'A' && c <= 'Z') {
        		rs.append(1, char(c + 32));
        	} else {
        		rs.append(1, c);
        	}
        }
        return rs;
    }
};