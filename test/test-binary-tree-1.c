#include <stdio.h>
#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

int* preorderTraversal(struct TreeNode* root, int* returnSize) {
    if (root == NULL) {
        returnSize = 0;
        return NULL;
    }
    
    const int STACK_ICR_SIZE = 100;
    const int RESULT_ICR_SIZE = 100;

    struct TreeNode **stack = (struct TreeNode **)malloc(sizeof(struct TreeNode *) * STACK_ICR_SIZE);
    int *result = (int *)malloc(sizeof(int) * RESULT_ICR_SIZE);
    int stack_top = -1;
    *returnSize = 0;

    stack[++stack_top] = root;
    struct TreeNode *cursorNode;

    while (stack_top > -1) {
        cursorNode = stack[stack_top--];
        result[(*returnSize)++] = cursorNode->val;
        if (cursorNode->right != NULL) {
            stack[++stack_top] = cursorNode->right;
        }
        if (cursorNode->left != NULL) {
            stack[++stack_top] = cursorNode->left;
        }
    }

    free(stack);

    return result;
}

int main() {
    struct TreeNode *root = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    struct TreeNode *node2 = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    struct TreeNode *node3 = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    struct TreeNode *node4 = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    struct TreeNode *node5 = (struct TreeNode *)malloc(sizeof(struct TreeNode));

    root->val = 1;
    root->left = node2;
    root->right = node3;
    node2->val = 2;
    node2->left = NULL;
    node2->left = NULL;
    node3->val = 3;
    node3->left = node4;
    node3->right = NULL;
    node4->val = 4;
    node4->left = NULL;
    node4->right = node5;
    node5->val = 5;
    node5->left = NULL;
    node5->right = NULL;

    int *returnSize;
    int *result = preorderTraversal(root, returnSize);
    for(int i = 0; i < *returnSize; i++) {
        printf("%d\t", result[i]);
    }
    printf("\n");
    free(result);
    return 0;
}
