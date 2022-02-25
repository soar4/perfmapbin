package main

import (
	"sort"
	"testing"
)

func BenchmarkBinSeach(b *testing.B) {
	InitDataSet(dataSetSize, dataSize)
	sort.StringSlice(dataset).Sort()

	for i := 0; i < b.N; i++ {
		idx := i % len(dataset)
		k := dataset[idx]
		binIdx := sort.SearchStrings(dataset, k)
		if dataset[binIdx] != dataset[idx] {
			b.Errorf("[%d][%s] Not-Found [%d][%s]\n", idx, k, binIdx, dataset[binIdx])
		}
	}
}

func BenchmarkMapSeach(b *testing.B) {
	var dataMap map[string]int = make(map[string]int)
	InitDataSet(dataSetSize, dataSize)

	InitMap(dataMap, dataset)

	for i := 0; i < b.N; i++ {
		idx := i % len(dataset)
		k := dataset[idx]
		if _, ok := dataMap[k]; !ok {
			b.Errorf("[%d][%s] Not-Found\n", idx, k)
		}
	}
}
