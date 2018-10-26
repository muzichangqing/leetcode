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
    bool hasPathSum(TreeNode* root, int sum) {
        int currentSum = 0;
        return hasPathSumRecursively(root, currentSum, sum);   
    }

    bool hasPathSumRecursively(TreeNode *root, int currentSum, int sum) {
        if (root == NULL) {
            return false;
        }
        currentSum += root->val;
        if (currentSum == sum && root->left == NULL && root->right == NULL) {
            return true;
        }
        return hasPathSumRecursively(root->left, currentSum, sum) 
            || hasPathSumRecursively(root->right, currentSum, sum);
    }
};
