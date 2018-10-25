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
    	vector<TreeNode *> *traverse_before_reverse = new vector<TreeNode *>;
    	vector<TreeNode *> *traverse_after_reverse = new vector<TreeNode *>;
    	traverseTree(root, traverse_before_reverse);
    	reverseTree(root);
    	traverseTree(root, traverse_after_reverse);
    	if (traverse_before_reverse->size() != traverse_after_reverse->size()) {
    		return false;
    	}
    	for(int i = 0; i < traverse_after_reverse->size(); i++) {
			TreeNode *node1 = traverse_before_reverse->at(i);
			TreeNode *node2 = traverse_after_reverse->at(i);
			if (node1 == NULL && node2 != NULL) {
				return false;
			}
			if (node1 != NULL && node2 == NULL) {
				return false;
			}
			if (node1 != NULL && node1->val != node2->val) {
				return false;
			}
    	}
    	delete traverse_before_reverse;
    	delete traverse_after_reverse;
    	return true;
    }

    TreeNode *reverseTree(TreeNode *root) {
    	if (root == NULL) return NULL;
		TreeNode *tmp = root->right;
    	root->right = reverseTree(root->left);
    	root->left = reverseTree(tmp);
    	return root;
    }

    void traverseTree(TreeNode *root, vector<TreeNode *> *result) {
		result->push_back(root);
    	if (root != NULL) {
			traverseTree(root->left, result);
    		traverseTree(root->right, result);
		}
    	return;
    }
};