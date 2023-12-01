package main

import (
	"fmt"
	"hash/fnv"
)

// HashMap はハッシュテーブルを表します。
type HashMap struct {
	buckets [][]KeyValuePair
	size    int
}

// KeyValuePair はキーと値のペアを表します。
type KeyValuePair struct {
	Key   string
	Value interface{}
}

// NewHashMap は新しいHashMapを作成します。
func NewHashMap(size int) *HashMap {
	return &HashMap{
		buckets: make([][]KeyValuePair, size),
		size:    size,
	}
}

// hashFunction はハッシュ関数です。
func hashFunction(key string, size int) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % size
}

// Put はハッシュテーブルにキーと値のペアを追加します。
func (hm *HashMap) Put(key string, value interface{}) {
	index := hashFunction(key, hm.size)
	bucket := hm.buckets[index]

	for i, pair := range bucket {
		if pair.Key == key {
			hm.buckets[index][i].Value = value
			return
		}
	}

	hm.buckets[index] = append(bucket, KeyValuePair{Key: key, Value: value})
}

// Get はハッシュテーブルからキーに対応する値を取得します。
func (hm *HashMap) Get(key string) (interface{}, bool) {
	index := hashFunction(key, hm.size)
	bucket := hm.buckets[index]

	for _, pair := range bucket {
		if pair.Key == key {
			return pair.Value, true
		}
	}

	return nil, false
}

func main() {
	// ハッシュマップの作成。
	hm := NewHashMap(100)

	// ハッシュマップにデータを追加。
	hm.Put("key1", "value1")
	hm.Put("key2", "value2")

	// ハッシュマップからデータを取得。
	if value, found := hm.Get("key1"); found {
		fmt.Println("key1:", value)
	} else {
		fmt.Println("key1 not found")
	}

	if value, found := hm.Get("key2"); found {
		fmt.Println("key2:", value)
	} else {
		fmt.Println("key2 not found")
	}
}
