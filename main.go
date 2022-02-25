package main

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

func main() {
	log.Println("end")

	rand.Seed(time.Now().UTC().UnixNano())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	InitDataSet(5, dataSize)
	log.Println("init dataset", dataset)
	sort.StringSlice(dataset).Sort()
	log.Println("sorted dataset", dataset)

	for i := 0; i < 3; i++ {
		idx := i % len(dataset)
		k := dataset[idx]
		binIdx := sort.SearchStrings(dataset, k)
		if binIdx != idx {
			log.Printf("[%d] %s Not-Found at [%d]\n", idx, k, binIdx)
		}
	}

	log.Println("end")
}
