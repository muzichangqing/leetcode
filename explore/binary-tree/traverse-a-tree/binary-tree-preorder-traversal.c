/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
/**
 * Return an array of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* preorderTraversal(struct TreeNode* root, int* returnSize) {
    if (root == NULL) {
        returnSize = 0;
        return NULL;
    }
    
    const int STACK_ICR_SIZE = 100;
    const int RESULT_ICR_SIZE = 100;

    struct TreeNode **stack = (struct TreeNode **)malloc(sizeof(struct TreeNode *) * STACK_ICR_SIZE);
    int *result = (int *)malloc(sizeof(int) * RESULT_ICR_SIZE);
    int stack_top = 0;
    *returnSize = 0;

    stack[stack_top++] = root;
    struct TreeNode *cursorNode;

    while (stack_top > 0) {
        cursorNode = stack[stack_top--];
        result[(*returnSize)++] = cursorNode->val;
        if (cursorNode->right != NULL) {
            stack[stack_top++] = cursorNode->right;
        }
        if (cursorNode->left != NULL) {
            stack[stack_top++] = cursorNode->left;
        }
    }

    free(stack);

    return result;
}
