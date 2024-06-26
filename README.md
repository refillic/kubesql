# KubeSQL

A simple CLI app that allows you to query Kubernetes with SQL.

## Prerequisites

* [Just](https://github.com/casey/just) for build scripts
* [goyacc](https://pkg.go.dev/golang.org/x/tools/cmd/goyacc?utm_source=godoc) for generating our SQL parser

## SQL Parser

We use [`goyacc`](https://pkg.go.dev/golang.org/x/tools/cmd/goyacc?utm_source=godoc) for generating our SQL parser. Our grammar file is taken from https://github.com/dolthub/vitess/blob/main/go/vt/sqlparser/sql.y

To generate the parser just run

```bash
just generate-parser
```

Make sure you have [goyacc](https://pkg.go.dev/golang.org/x/tools/cmd/goyacc?utm_source=godoc) installed. Instructions on how to install it can be found [here](https://cs.opensource.google/go/x/tools).

## Testing

To run tests just run

```bash
just test
```