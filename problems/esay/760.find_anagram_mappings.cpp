class Solution {
public:
    std::vector<int> anagramMappings(std::vector<int>& A, std::<int>& B) {
        int size = A.size();
        std::vector<int> result(size);
        for (size_t index_a = 0; index_a < size; index_a++) {
            for (size_t index_b = 0; index_b < size; index_b++) {
                if (A[index_a] == B[index_b]) {
                    result[index_a] = index_b;
                }
            }
        }
        return result;
    }
};
