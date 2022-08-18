# `go-openbd`

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

This is a go library for [OpenBD](https://openbd.jp/) API client.

## Installation
```
$ go get github.com/kironono/go-openbd
```

## Usage

Import package:

```go
import "github.com/kironono/go-openbd"
```

Fetch OpenDB book information:

```go
books, err := openbd.DefaultClient().Books([]string{"9784814400041", "9784873115658"})
```

## License

`twee` is distributed under the terms of the MIT license.

See the [LICENSE](LICENSE) files in this repository for more information.
