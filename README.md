# murmur2-go
Go implementation of MurmurHash2, based on the work by [Austin Appleby](https://code.google.com/p/smhasher/). This repo was forked to add go modules, a more modern test suite (though still lacking in rigour), and a ci pipeline to ensure build + test passes.

## Usage
Specified minimum version of go is v1.22. 

```bash
$ go get github.com/yarefs/murmur2-go
```

## Performance
Benchmarked as of 2024:

```bash
➜  murmur2-go git:(main) ✗ go test -bench=. ./...
goos: darwin
goarch: amd64
pkg: github.com/yarefs/murmur2-go/murmur
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMurmurHash2-12         61032399                19.41 ns/op
BenchmarkMurmurHash2A-12        56362617                21.29 ns/op
BenchmarkMurmurHash64A-12       81483355                14.70 ns/op
BenchmarkHash32_Murmur2-12      16060260                77.04 ns/op
BenchmarkHash32_FNV1-12         12585670                95.21 ns/op
BenchmarkHash32_FNV1a-12        12452218                95.75 ns/op
PASS
ok      github.com/yarefs/murmur2-go/murmur     8.125s
```

## License

[MIT](LICENSE)
