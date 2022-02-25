package main

import "math/rand"

type MyConn struct {
	Tag  string
	Info string
}

type ByTag []MyConn

func (a ByTag) Len() int           { return len(a) }
func (a ByTag) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTag) Less(i, j int) bool { return a[i].Tag < a[j].Tag }

const (
	dataSetSize = 100
	dataSize    = 1
)

var (
	dataset []string = make([]string, 0, dataSetSize)
)

func GenActivateCode(data []byte) {
	src := []byte("3456789abcdefghjkmnpqrstuvwxyABCDEFGHJKLMNPQRSTUVWXY")
	// src := []byte("0123456789")
	srcLen := len(src)
	for i := 0; i < len(data); i++ {
		data[i] = src[rand.Intn(srcLen)]
	}
}

func InitDataSet(num, size int) {
	for i := 0; i < num; i++ {
		data := make([]byte, size)
		GenActivateCode(data)
		dataset = append(dataset, string(data))
	}
}

func InitMap(dataMap map[string]int, dataset []string) {
	for k, v := range dataset {
		dataMap[v] = k
	}
}
