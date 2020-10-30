# Run test

## Bazel test

```shell
bazel test examples/internal/integration/... --test_arg=--skylb-endpoints="" --test_arg=--debug-svc-endpoint=custom-ease-gateway-test=localhost:9090
```

> 通过bazel test 执行integration test
>
> 因为引入skylb-api, 运行时需要指定skylb-endpoints