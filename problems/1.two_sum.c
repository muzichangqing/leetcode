#include <stdlib.h>
#ifndef NULL
    #define NULL ((void *)0)
#endif

typedef struct dictEntry {
	int key;
	int value;
	struct dictEntry *next;
} dictEntry;

typedef struct dict {
	struct dictEntry **table;
	unsigned long size;
} dict;

unsigned int dictHash(int key);
struct dict *createDict(int size);
int dictAdd(struct dict *d, int key, int value);
int dictFind(struct dict *d, int key);
void deleteDict(struct dict *d);

int dict_find_rs;

#define DICT_OK 1
#define DICT_ERR 0

int* twoSum(int* nums, int numsSize, int target) {
	dict *d = createDict(numsSize);
	for (int i = 0; i < numsSize; i++) {
		dictAdd(d, nums[i], i);
	}
	for (int i = 0; i < numsSize; i++) {
		int consume = target - nums[i];
		int index = dictFind(d, consume);
		if (dict_find_rs == DICT_ERR) continue;
        if (index == i) continue;
		deleteDict(d);
		int *result = (int *)malloc(sizeof(int) * 2);
		result[0] = i;
		result[1] = index;
		return result;
	}
	return 0;
}

unsigned int dictHash(int key) {
	unsigned int seed = 131;
	unsigned int hash = 0;

	int mod = key % 10;
	while (mod) {
		hash = hash * seed + mod;
		key = key / 10;
		mod = key % 10;
	}
	return (hash & 0x7FFFFFFF);
}


dict *createDict(int size) {
	dictEntry **table = (dictEntry **)malloc(sizeof(dictEntry *) * size);
	for (int i = 0; i < size; i++) {
		table[i] = NULL;
	}
	dict *d = (dict *)malloc(sizeof(dict));
	d->table = table;
	d->size = size;
	return d;
}

void deleteDict(dict *d) {
	for (int i = 0; i < d->size; i++) {
		dictEntry *de = d->table[i];
		while (de) {
			dictEntry *next = de->next;
			free(de);
			de = next;
		}
	}
	free(d);
}

int dictAdd(dict *d, int key, int value) {
	dictEntry *nde = (dictEntry *)malloc(sizeof(dictEntry));
	nde->key = key;
	nde->value = value;
	nde->next = NULL;

	int index = dictHash(key) % d->size;
	dictEntry *de = d->table[index];
	if (de == NULL) {
		d->table[index] = nde;
		return 1;
	}
	while (de != NULL) {
		if (de->key == key) return 0;
		if (!de->next) break;
		de = de->next;
	}
	de->next = nde;
	return 1;
}

int dictFind(dict *d, int key) {
	int index = dictHash(key) % d->size;
	dictEntry *de = d->table[index];
	while (de && de->key != key) {
		de = de->next;
	}
	if (de) {
		dict_find_rs = DICT_OK;
		return de->value;
	}
	dict_find_rs = DICT_ERR;
	return 0;
}