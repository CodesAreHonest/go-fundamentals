## Module Links

Adapt `fundamentals/CallOtherModules` module so it can find `fundamentals/CreateModule`.

```bash
$ go mod edit -replace fundamentals/CreateModule=../2-create-module
$ go mod tidy
go: found fundamentals/CreateModule in fundamentals/CreateModule v0.0.0-00010101000000-000000000000
```

## References

1. https://go.dev/doc/tutorial/handle-errors
2. https://go.dev/doc/tutorial/random-greeting
