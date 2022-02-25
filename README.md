# 测试结果
```
$ go test -run=. -bench=.  -count 5 -benchmem
goos: linux
goarch: amd64
pkg: perfmapbin
cpu: Intel(R) Core(TM) i3-2310M CPU @ 2.10GHz
BenchmarkBinSeach-4      7996130               152.0 ns/op             0 B/op          0 allocs/op
BenchmarkBinSeach-4      7800070               183.3 ns/op             0 B/op          0 allocs/op
BenchmarkBinSeach-4      2131078               477.1 ns/op             0 B/op          0 allocs/op
BenchmarkBinSeach-4      2653921               532.5 ns/op             0 B/op          0 allocs/op
BenchmarkBinSeach-4      2625172               503.6 ns/op             0 B/op          0 allocs/op
BenchmarkMapSeach-4      9320173               163.6 ns/op             0 B/op          0 allocs/op
BenchmarkMapSeach-4      8732587               164.9 ns/op             0 B/op          0 allocs/op
BenchmarkMapSeach-4      8046327               145.8 ns/op             0 B/op          0 allocs/op
BenchmarkMapSeach-4      8146843               162.7 ns/op             0 B/op          0 allocs/op
BenchmarkMapSeach-4      7346401               159.2 ns/op             0 B/op          0 allocs/op
PASS
ok      perfmapbin      17.620s
```

# 结论：
- map 相较二分查找速度略快。cpu 消耗更多。 bin 消耗 cpu 24%, map 消耗cpu 33%
- 二分查找每次都需要比较两个key值，key值比较复杂度也影响测试结果
- 根据场景选择合适做法