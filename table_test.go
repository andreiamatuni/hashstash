package main

import (
	"fmt"
	"strconv"
	"testing"
)

type Pair struct {
	key   []byte
	value string
}

var dataSet10 [10]Value
var dataSet100 [100]Value
var dataSet1000 [1000]Value
var dataSet10000 [10000]Value
var dataSet100000 [100000]Value
var dataSet1000000 [1000000]Value

func init10DataSet() [10]Value {
	data := 0
	for x := 0; x < len(dataSet10); x++ {
		stringdata := strconv.Itoa(data)
		dataSet10[x] = Value{key: []byte(stringdata), value: stringdata}
		data++
	}
	return dataSet10
}

func init100DataSet() [100]Value {
	data := 0
	for x := 0; x < len(dataSet100); x++ {
		stringdata := strconv.Itoa(data)
		dataSet100[x] = Value{key: []byte(stringdata), value: stringdata}
		data++
	}
	return dataSet100
}

func init1000DataSet() [1000]Value {
	data := 0
	for x := 0; x < len(dataSet1000); x++ {
		stringdata := strconv.Itoa(data)
		dataSet1000[x] = Value{key: []byte(stringdata), value: stringdata}
		data++
	}
	return dataSet1000
}

func init10000DataSet() [10000]Value {
	data := 0
	for x := 0; x < len(dataSet10000); x++ {
		stringdata := strconv.Itoa(data)
		dataSet10000[x] = Value{key: []byte(stringdata), value: stringdata}
		data++
	}
	return dataSet10000
}

func init100000DataSet() [100000]Value {
	data := 0
	for x := 0; x < len(dataSet100000); x++ {
		stringdata := strconv.Itoa(data)
		dataSet10000[x] = Value{key: []byte(stringdata), value: stringdata}
		data++
	}
	return dataSet100000
}

func init1000000DataSet() [1000000]Value {
	data := 0
	for x := 0; x < len(dataSet1000000); x++ {
		stringdata := strconv.Itoa(data)
		dataSet1000000[x] = Value{key: []byte(stringdata), value: stringdata}
		data++
	}
	return dataSet1000000
}

func TestHashTable(t *testing.T) {

	data10 := init10DataSet()

	table := NewHashTable(7, 3)

	for _, pair := range data10 {
		table.Set(pair)
	}

	fmt.Printf("\n\n")
	fmt.Println("\n\nsize of the table: ", len(table.buckets))
	fmt.Println()

	result, _ := table.Get([]byte(strconv.Itoa(42)))
	fmt.Println("\n\nthe result: ", result)

}

func BenchmarkHashTable10(b *testing.B) {
	data10 := init10DataSet()

	table := NewHashTable(30, 1)
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, pair := range data10 {
			table.Set(pair)
		}
	}
}

// func BenchmarkGoHashTable10(b *testing.B) {
// 	data10 := init10DataSet()
//
// 	table := make(map[string]string)
// 	b.ResetTimer()
// 	b.ReportAllocs()
// 	for n := 0; n < b.N; n++ {
// 		for _, pair := range data10 {
// 			table[pair.key] = pair.value
// 		}
// 	}
// }

func BenchmarkHashTable1000(b *testing.B) {
	data1000 := init1000DataSet()

	table := NewHashTable(3000, 3)
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, pair := range data1000 {
			table.Set(pair)
		}
	}
}

func BenchmarkHashTable10000(b *testing.B) {
	data10000 := init10000DataSet()

	table := NewHashTable(15000, 6)
	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		for _, pair := range data10000 {
			table.Set(pair)
		}
	}
}

// func BenchmarkGoHashTable1000(b *testing.B) {
// 	data1000 := init1000DataSet()
//
// 	table := make(map[string]string)
// 	b.ResetTimer()
// 	b.ReportAllocs()
// 	for n := 0; n < b.N; n++ {
// 		for _, pair := range data1000 {
// 			table[pair.key] = pair.value
// 		}
// 	}
// }

func BenchmarkHashTable1000000(b *testing.B) {
	data1000000 := init1000000DataSet()

	table := NewHashTable(2000000, 4)
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, pair := range data1000000 {
			table.Set(pair)
		}
	}
}
