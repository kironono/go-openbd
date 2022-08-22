# `go-openbd`

[![Go Reference](https://pkg.go.dev/badge/github.com/kironono/go-openbd.svg)](https://pkg.go.dev/github.com/kironono/go-openbd)
[![ci](https://github.com/kironono/go-openbd/actions/workflows/ci.yml/badge.svg)](https://github.com/kironono/go-openbd/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

This is a go library for [OpenBD](https://openbd.jp/) API client.

## Installation
```
$ go get github.com/kironono/go-openbd
```

## Usage

Import package:

```go
import "github.com/kironono/openbd"
```

Fetch OpenDB book information:

```go
books, err := openbd.DefaultClient().Books(context.Background(), []string{"9784814400041", "9784873115658"})
```

## License

`go-openbd` is distributed under the terms of the MIT license.

See the [LICENSE](LICENSE) files in this repository for more information.
