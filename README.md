# 测试结果
```
[gaofei@thinkpad perfmapbin]$ go test -run=. -bench=.  -count 5 -benchmem
goos: linux
goarch: amd64
pkg: perfmapbin
cpu: Intel(R) Core(TM) i3-2310M CPU @ 2.10GHz
BenchmarkBinSeach-4      6806210               187.3 ns/op             0 B/op          0 allocs/op
BenchmarkBinSeach-4      5973517               202.7 ns/op             0 B/op          0 allocs/op
BenchmarkBinSeach-4      5750232               218.2 ns/op             0 B/op          0 allocs/op
BenchmarkBinSeach-4      5424562               221.8 ns/op             0 B/op          0 allocs/op
BenchmarkBinSeach-4      5424704               227.1 ns/op             0 B/op          0 allocs/op
BenchmarkMapSeach-4     16204354               138.6 ns/op             0 B/op          0 allocs/op
BenchmarkMapSeach-4      5550140               225.8 ns/op             0 B/op          0 allocs/op
BenchmarkMapSeach-4      6838604               180.6 ns/op             0 B/op          0 allocs/op
BenchmarkMapSeach-4      6416245               188.8 ns/op             0 B/op          0 allocs/op
BenchmarkMapSeach-4      6539950               183.9 ns/op             0 B/op          0 allocs/op
PASS
ok      perfmapbin      16.328s
```

# 结论：
- map 相较二分查找速度更快。cpu 消耗更多。 bin 消耗 cpu 24%, map 消耗cpu 33%