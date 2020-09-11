package hashtable

const myHashSetDataSize = 769

// MyHashSet a simple hashSet
type MyHashSet struct {
	data [][]int
}

// Constructor initialize your data structure here.
func Constructor() MyHashSet {
	data := make([][]int, myHashSetDataSize)
	return MyHashSet{data}
}

// Add add a element to set
func (set *MyHashSet) Add(key int) {
	if set.Contains(key) {
		return
	}
	keyHash := hash(key)
	if set.data[keyHash] == nil {
		set.data[keyHash] = []int{key}
	} else {
		keys := set.data[keyHash]
		keys = append(keys, key)
		set.data[keyHash] = keys
	}
}

// Remove remove a value from set. if not exists, do nothing.
func (set *MyHashSet) Remove(key int) {
	if !set.Contains(key) {
		return
	}
	keyHash := hash(key)
	keys := set.data[keyHash]
	keyIndex := -1
	for index, element := range keys {
		if element == key {
			keyIndex = index
			break
		}
	}
	set.data[keyHash] = append(keys[:keyIndex], keys[keyIndex+1:]...)
}

// Contains returns true if this set contains the specified element
func (set *MyHashSet) Contains(key int) bool {
	keyHash := hash(key)
	keys := set.data[keyHash]
	if keys == nil {
		return false
	}
	for _, element := range keys {
		if element == key {
			return true
		}
	}
	return false
}

func hash(key int) int {
	return key % myHashSetDataSize
}

/**
 * 所有的值都在 [0, 1000000]的范围内。
 * 操作的总数目在[1, 10000]范围内。
 * 不要使用内建的哈希集合库。
 */

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
