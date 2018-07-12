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
    int maxDepth(TreeNode* root) {
        return maxDepthR(root, 0);
    }

    int maxDepthR(TreeNode *root, int depth) {
    	if (root == NULL) {
    		return depth;
    	}
    	int left_depth = maxDepthR(root->left, depth + 1);
    	int right_depth = maxDepthR(root->right, depth + 1);
    	return left_depth > right_depth ? left_depth : right_depth;
    }
};