# native

[![Go Reference](https://pkg.go.dev/badge/github.com/thediveo/native.svg)](https://pkg.go.dev/github.com/thediveo/native)
[![GitHub](https://img.shields.io/github/license/thediveo/native)](https://img.shields.io/github/license/thediveo/native)
![build and test](https://github.com/thediveo/native/actions/workflows/buildandtest.yaml/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/native)](https://goreportcard.com/report/github.com/thediveo/native)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

`native` translates between numbers in fixed sizes and byte sequences in (host)
native endianess. Think of it as the missing convenience combination of
`bytes.Buffer` and `encoding/binary` for fixed size numbers (`uint8`, `uint16`,
`uint32`, and `uint64`).

Especially useful when dealing with
[netlink](https://en.wikipedia.org/wiki/Netlink). And no need to manually unroll
`encoding/binary` interface calls depending on endianess at runtime.

## DevContainer

> [!CAUTION]
>
> Do **not** use VSCode's "~~Dev Containers: Clone Repository in Container
> Volume~~" command, as it is utterly broken by design, ignoring
> `.devcontainer/devcontainer.json`.

1. `git clone https://github.com/thediveo/enumflag`
2. in VSCode: Ctrl+Shift+P, "Dev Containers: Open Workspace in Container..."
3. select `enumflag.code-workspace` and off you go...

## Supported Go Versions

`native` supports versions of Go that are noted by the [Go release
policy](https://golang.org/doc/devel/release.html#policy), that is, major
versions _N_ and _N_-1 (where _N_ is the current major version).

## Copyright and License

`native` is Copyright 2023, 2025 Harald Albrecht, and licensed under the Apache
License, Version 2.0.
