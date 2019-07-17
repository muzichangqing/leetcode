class Solution {
public:
    int uniqueMorseRepresentations(vector<string>& words) {
        string morse_map[26] = {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."};
        vector<string> morse_rs;
        for(vector<string>::iterator it = words.begin(); it != words.end(); it++) {
        	string words_morse = "";
        	for(int i = 0; i < it->size(); i++) {
        		char c = it->at(i);
        		string c_morse = morse_map[c - 97];
        		words_morse += c_morse;
        	}
        	vector<string>::iterator f_it;
        	f_it = find(morse_rs.begin(), morse_rs.end(), words_morse);
        	if (f_it == morse_rs.end()) {
        		morse_rs.push_back(words_morse);
        	}
        }
        return morse_rs.size();
    }
};