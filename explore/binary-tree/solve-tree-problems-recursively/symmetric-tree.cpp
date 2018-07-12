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
    bool isSymmetric(TreeNode* root) {
    	vector<int> *traverse_before_reverse = new vector<int>;
    	vector<int> *traverse_after_reverse = new vector<int>;
    	traverseTree(root, traverse_before_reverse);
    	reverseTree(root);
    	traverseTree(root, traverse_after_reverse);
    	if (traverse_before_reverse->size() != traverse_after_reverse->size()) {
    		return false;
    	}
    	for(int i = 0; i < traverse_after_reverse->size(); i++) {
    		if (traverse_after_reverse->at(i) != traverse_before_reverse->at(i)) {
    			return false;
    		}
    	}
    	delete traverse_before_reverse;
    	delete traverse_after_reverse;
    	return true;
    }

    TreeNode *reverseTree(TreeNode *root) {
    	if (root == NULL) return NULL;
    	root->right = reverseTree(root->left);
    	root->left = reverseTree(root->right);
    	return root;
    }

    void traverseTree(TreeNode *root, vector<int> *result) {
    	if (root == NULL) return;
    	result->push_back(root->val);
    	traverseTree(root->left, result);
    	traverseTree(root->right, result);
    	return;
    }
};