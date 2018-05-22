class Solution {
public:
    bool isToeplitzMatrix(vector<vector<int>>& matrix) {
        int row = matrix.size();
        int col = matrix[0].size();
        for (int i = 0; i < col; i++) {
            int value = matrix[0][i];
            for (int ci = i + 1, ri = 1; ci < col && ri < row; ci++, ri++) {
                if (matrix[ri][ci] != value) return false;
            }
        }
        for (int i = 0; i < row; i++) {
            int value = matrix[i][0];
            for (int ci = 1, ri = i + 1; ci < col && ri < row; ci++, ri++ ) {
                if (matrix[ri][ci] != value) return false;
            }
        }
        return true;
    }
};
