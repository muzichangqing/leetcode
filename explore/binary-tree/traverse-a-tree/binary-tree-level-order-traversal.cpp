/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> result;
        if (root == NULL) return result;

        vector<TreeNode *> queue;
        queue.push_back(root);
        int current_level_count = 1;
        int next_level_count = 0;
        
        vector<int> current_level;
        
        while(queue.size() != 0) {
            TreeNode *current_node = queue.front();

            if (current_node->left != NULL) {
                queue.push_back(current_node->left);
                next_level_count++;
            }
            if (current_node->right != NULL) {
                queue.push_back(current_node->right);
                next_level_count++;
            }
            
            current_level.push_back(current_node->val);
            
            queue.erase(queue.begin());
            current_level_count--;
            
            if (current_level_count == 0) {
                result.push_back(current_level);
                current_level.clear();
                current_level_count = next_level_count;
                next_level_count = 0;
            }
        }
        
        return result;
    }
};