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
int* postorderTraversal(struct TreeNode* root, int* returnSize) {
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
    
    struct TreeNode *cursorNode = root;
    
    while (stack_top > 0 || cursorNode != NULL) {
        // TODO: postorderTraversal
        while(cursorNode != NULL) {
            if (cursorNode->right != NULL) stack[stack_top++] = cursorNode->right;
            stack[stack_top++] = cursorNode;
            cursorNode = cursorNode->left;
        }
        
        cursorNode = stack[--stack_top];
        
        if (cursorNode->right == stack[stack_top - 1]) {
            stack_top--;
            stack[stack_top++] = cursorNode;
            cursorNode = cursorNode->right;
        } else {
            result[(*returnSize)++] = cursorNode->val;
            cursorNode = NULL;
        }
        
    }
    
    free(stack);
    return result;
}