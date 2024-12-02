package datastructure_test

import (
	"fmt"
	"testing"
)

type HashTable struct {
	data [][]any
}

func NewHashTable(size ...int) *HashTable {
	if len(size) == 0 {
		size = append(size, 7)
	}
	return &HashTable{
		data: make([][]any, size[0]),
	}
}

func (table *HashTable) _hash(key string) int {
	defaultHash := 0

	for _, v := range key {
		defaultHash = (defaultHash + int(v)*17) % len(table.data)
	}
	// create key hash value
	// key unicode (python ord <- go int casting to rune)
	// 17 is prime number
	// % len is if size 7 we have 0 ~ 6 index
	// %7 is must range 0 ~ 6
	return defaultHash
}

func (table *HashTable) setItem(key string, value any) {
	index := table._hash(key)
	if table.data[index] == nil {
		table.data[index] = make([]any, 0)
	}
	table.data[index] = append(table.data[index], []any{key, value})
}

func (table *HashTable) getItem(key string) any {
	index := table._hash(key)
	if table.data[index] != nil {
		for i := range len(table.data[index]) {
			// pointing key []any{key, value} <- 0 (key), 1(value)
			if table.data[index][i].([]any)[0].(string) == key {
				return table.data[index][i].([]any)[1]
			}
		}
	}
	return nil
}

func (table *HashTable) keys() []any {
	allkeys := make([]any, 0)
	for i := range len(table.data) {
		// 0  :  [[fakeitem3 7150]]
		// 1  :  [[fakeitem1 4200] [fakeitem8 10832]]
		// 2  :  [[fakeitem6 4020]]
		// 3  :  [[fakeitem4 6640]]
		// 4  :  [[fakeitem2 61350] [fakeitem9 3952]]
		// 5  :  [[fakeitem7 9521]]
		// 6  :  [[fakeitem5 74820]]
		//step 1. iter index is not null
		//step 2. iterator all index keys and saved allkeys array
		// return array then user knows keys in hashtable
		if table.data[i] != nil {
			for j := range len(table.data[i]) {
				allkeys = append(allkeys, table.data[i][j].([]any)[0])
			}
		}
	}
	return allkeys
}

func (table *HashTable) printHashTable() {
	for i, val := range table.data {
		fmt.Println(i, " : ", val)
	}
}

func TestHashTable(t *testing.T) {
	table := NewHashTable()

	table.setItem("fakeitem1", 4200)
	table.setItem("fakeitem2", 61350)
	table.setItem("fakeitem3", 7150)
	table.setItem("fakeitem4", 6640)
	table.setItem("fakeitem5", 74820)
	table.setItem("fakeitem6", 4020)
	table.setItem("fakeitem7", 9521)
	table.setItem("fakeitem8", 10832)
	table.setItem("fakeitem9", 3952)
	table.printHashTable()

	fmt.Println(table.getItem("fakeitem7"))
	fmt.Println(table.getItem("fakeitem9"))
	fmt.Println(table.getItem("fakeitem12041"))
	fmt.Println(table.keys()...)
}
