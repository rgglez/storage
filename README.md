# storage

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![GitHub all releases](https://img.shields.io/github/downloads/rgglez/storage/total)
![GitHub issues](https://img.shields.io/github/issues/rgglez/storage)
![GitHub commit activity](https://img.shields.io/github/commit-activity/y/rgglez/storage)
[![Go Report Card](https://goreportcard.com/badge/github.com/rgglez/storage)](https://goreportcard.com/report/github.com/rgglez/storage)
[![GitHub release](https://img.shields.io/github/release/rgglez/storage.svg)](https://github.com/rgglez/storage/releases/)

This Go module encapsulates the [go-storage](https://github.com/rgglez/go-storage) library, 
which makes transparent the storage and retrieval of files to and from a number of both 
cloud services and local supports. A full list of supported backends can be
found in the go-storage [README](https://github.com/rgglez/go-storage/blob/master/README.md).

It provides 2 functions, one for writing and one for reading. It takes care of the boilerplate.

You need to provide the constructor ***NewStorage(cnn string)*** a connection string in the
go-storage format:

```go
cnn := "oss://bucket/?credential=env&endpoint=http://127.0.0.1:9090&name=bucket"
```

## License

Apache-2.0. Please read the [LICENSE](LICENSE) file.

Copyright (c) 2024 Rodolfo González González.
