# Benchmark Go

## How to run the benchmark

### Local

```shell
go test -benchmem -bench=$TARGET_FUNCTION_NAME_QUERY ./...
```

### Docker

```shell
docker run -i --rm -v $(pwd):/go/src/ golang go test -benchmem -bench=$TARGET_FUNCTION_NAME_QUERY ./...
```
