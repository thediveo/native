# native

[![Go Reference](https://pkg.go.dev/badge/github.com/thediveo/native.svg)](https://pkg.go.dev/github.com/thediveo/native)
[![GitHub](https://img.shields.io/github/license/thediveo/native)](https://img.shields.io/github/license/thediveo/native)
![build and test](https://github.com/thediveo/native/workflows/build%20and%20test/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/native)](https://goreportcard.com/report/github.com/thediveo/native)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

`native` translates between numbers in fixed sizes and byte sequences in (host)
native endianess. Think of it as the missing convenience combination of
`bytes.Buffer` and `encoding/binary` for fixed size numbers (`uint8`, `uint16`,
`uint32`, and `uint64`).

Especially useful when dealing with
[netlink](https://en.wikipedia.org/wiki/Netlink). And no need to manually unroll
`encoding/binary` interface calls depending on endianess at runtime.

## Supported Go Versions

`native` supports versions of Go that are noted by the [Go release
policy](https://golang.org/doc/devel/release.html#policy), that is, major
versions _N_ and _N_-1 (where _N_ is the current major version).

## Copyright and License

`native` is Copyright 2023 Harald Albrecht, and licensed under the Apache
License, Version 2.0.
