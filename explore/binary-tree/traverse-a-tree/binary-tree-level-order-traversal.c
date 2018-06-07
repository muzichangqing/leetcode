/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *columnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
struct QueueNode {
    struct TreeNode *element;
    struct QueueNode *next;
};

struct Queue {
    struct QueueNode *head;
    struct QueueNode *end;
};
struct Queue *init_queue();
void add_queue(struct Queue *queue, struct TreeNode* element);
struct TreeNode *remove_queue(struct Queue *queue);
int is_queue_empty(struct Queue *queue);
void delete_queue(struct Queue *queue);

int** levelOrder(struct TreeNode* root, int** columnSizes, int* returnSize) {
    struct Queue *queue = init_queue();

    add_queue(queue, root);
    struct TreeNode *currentNode;
    int currentLevelCount = 1;
    int nextLevelCount = 0;
    
    while (!is_queue_empty(queue)) {
        currentNode = remove_queue(queue);
        currentLevelCount--;
        printf("%d\t", currentNode->val);
        if (currentNode->left != NULL) {
            add_queue(queue, currentNode->left);
            nextLevelCount++;
        }
        if (currentNode->right != NULL) {
            add_queue(queue, currentNode->right);
            nextLevelCount++;
        }
        if (currentLevelCount == 0) {
            currentLevelCount = nextLevelCount;
            nextLevelCount = 0;
        }
    }
    printf("\n");
    return NULL;  
}

struct Queue *init_queue() {
    struct Queue *queue = (struct Queue *) malloc (sizeof(struct Queue));
    queue->head = queue->end = NULL;
    return queue;
}

void add_queue(struct Queue *queue, struct TreeNode* element) {
    struct QueueNode *newNode = (struct QueueNode *) malloc (sizeof(struct QueueNode));
    newNode->element = element;
    newNode->next = NULL;
    if (queue->end == NULL) {
        queue->head = queue->end = newNode;
    } else {
        queue->end->next = newNode;
    }
}

struct TreeNode *remove_queue(struct Queue *queue) {
    struct TreeNode *element = queue->end->element;
    free(queue->end);
    if (queue->head == queue->end) {
        queue->head = NULL;
    }
    queue->end = NULL;
    return element;
}

int is_queue_empty(struct Queue *queue) {
    return queue->end == NULL;
}

void delete_queue(struct Queue *queue) {
    while (queue->end != NULL) {
        remove_queue(queue);
    }
    free(queue);
}

