# go mod sandbox

[![Makefile CI](https://github.com/keyno63/go-mod-sandbox/actions/workflows/makefile.yml/badge.svg)](https://github.com/keyno63/go-mod-sandbox/actions/workflows/makefile.yml)

## Purpose
go module の使い方を確認するようのもの  

以下を目指している
- Local でビルド・ Run できる
- Local の Docker でビルド・Run できる
- Remote の Docker でビルド・Run できる

## How to develop

this app (repository) is built by go command.  
Use [Makefile](./Makefile) with make command usually.  

### build

`build` is making binary.

run to build on local machine. 

```shell
make build
```

or, run to build on docker.

```shell
make builddockr
```

### run

`run` is app run.

run to run locally.

```shell
make run
```

run to run on docker.

```shell
make rundocker
```

### test

run to test locally.

```shell
make test
```

## License

This repository is MIT License.  
see [License](./LICENSE) file.
