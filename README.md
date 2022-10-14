# structmapper

## Conversion...
DTO와 Entity간의 변환을 조금 더 쉽게 하기 위해 만들어 보았다. 

## Benchmark
```
goos: windows
goarch: amd64
pkg: github.com/wonksing/structmapper
cpu: 12th Gen Intel(R) Core(TM) i7-12700
BenchmarkMapper-20                494835              4215 ns/op            3128 B/op        115 allocs/op
BenchmarkMapperCached-20         2389815               986.1 ns/op           200 B/op          9 allocs/op
PASS
ok      github.com/wonksing/structmapper        5.540s
```

- 타겟 구조체 생성 포함
```
goos: windows
goarch: amd64
pkg: github.com/wonksing/structmapper
cpu: 12th Gen Intel(R) Core(TM) i7-12700
BenchmarkMapper-20                553504              4170 ns/op            3128 B/op        115 allocs/op
BenchmarkMapperCached-20         2129167              1129 ns/op             376 B/op         13 allocs/op
PASS
ok      github.com/wonksing/structmapper        5.931s
```