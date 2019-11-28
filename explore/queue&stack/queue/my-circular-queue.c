#include <stdlib.h>
#include <stdbool.h>

typedef struct MyCircularQueue
{
    int *data;
    int head;
    int tail;
    int size;
} MyCircularQueue;

/** Initialize your data structure here. Set the size of the queue to be k. */
MyCircularQueue *myCircularQueueCreate(int k)
{
    MyCircularQueue *obj;
    if ((obj = malloc(sizeof(*obj))) == NULL)
        return NULL;
    if ((obj->data = malloc(sizeof(int) * k)) == NULL)
    {
        free(obj);
        return NULL;
    }
    obj->head = obj->tail = -1;
    obj->size = k;
    return obj;
}

/** Checks whether the circular queue is empty or not. */
bool myCircularQueueIsEmpty(MyCircularQueue *obj)
{
    return obj->head == -1;
}

/** Checks whether the circular queue is full or not. */
bool myCircularQueueIsFull(MyCircularQueue *obj)
{
    return (obj->tail + 1) % obj->size == obj->head;
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
bool myCircularQueueEnQueue(MyCircularQueue *obj, int value)
{
    if (myCircularQueueIsFull(obj))
        return false;
    if (myCircularQueueIsEmpty(obj))
        obj->head = 0;
    obj->tail = (obj->tail + 1) % obj->size;
    obj->data[obj->tail] = value;
    return true;
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
bool myCircularQueueDeQueue(MyCircularQueue *obj)
{
	if (myCircularQueueIsEmpty(obj))
		return false;
	if (obj->head == obj->tail)
		obj->head = obj->tail = -1;
	else 
		obj->head = (obj->head + 1) % obj->size;
	return true;
}

/** Get the front item from the queue. */
int myCircularQueueFront(MyCircularQueue *obj)
{
    if (myCircularQueueIsEmpty(obj))
        return -1;
    return obj->data[obj->head];
}

/** Get the last item from the queue. */
int myCircularQueueRear(MyCircularQueue *obj)
{
    if (myCircularQueueIsEmpty(obj))
        return -1;
    return obj->data[obj->tail];
}

void myCircularQueueFree(MyCircularQueue *obj)
{
    free(obj->data);
    free(obj);
}