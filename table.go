package main

import (
	"bytes"

	"github.com/cespare/xxhash"
)

type Value struct {
	key   []byte
	value string
}

/*
Bucket is the
*/
type Bucket []Value

func (buck *Bucket) add(value Value) {
	*buck = append(*buck, value)
}

/*
Table is the main hash table struct
*/
type Table struct {
	buckets []Bucket
	size    uint64
}

func NewHashTable(tableSize, bucketSize uint64) *Table {
	table := &Table{
		buckets: make([]Bucket, tableSize),
		size:    tableSize,
	}
	var x uint64
	for x = 0; x < tableSize; x++ {
		table.buckets[x] = make(Bucket, 0, bucketSize)
	}
	return table
}

func (table *Table) Set(value Value) {
	hashKey := xxhash.Sum64(value.key) % table.size
	//x := Value{key: key, value: value}
	table.buckets[hashKey].add(value)
}

func (table *Table) Get(key []byte) (Value, bool) {
	hashKey := xxhash.Sum64(key) % table.size
	list := table.buckets[hashKey]
	for _, element := range list {
		if bytes.Compare(element.key, key) == 0 {
			return element, true
		}
	}
	return Value{}, false
}
