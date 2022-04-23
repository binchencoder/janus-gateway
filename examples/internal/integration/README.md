# Run test

## Bazel test

```shell
bazel run examples/internal/integration/... --test_arg=--skylb-endpoints="" --test_arg=--debug-svc-endpoint=custom-janus-gateway-test=localhost:9090
```

> 通过bazel run 执行integration test
>
> 因为引入skylb-api, 运行时需要指定skylb-endpoints