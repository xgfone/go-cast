# Go Type Cast [![Build Status](https://github.com/xgfone/go-cast/actions/workflows/go.yml/badge.svg)](https://github.com/xgfone/go-cast/actions/workflows/go.yml) [![GoDoc](https://pkg.go.dev/badge/github.com/xgfone/go-cast)](https://pkg.go.dev/github.com/xgfone/go-cast) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](https://raw.githubusercontent.com/xgfone/go-cast/master/LICENSE)

Provide some functions, supporting `Go1.21+`, to convert the value between different types, such as `ToXXX`.

## Installation

```shell
$ go get -u github.com/xgfone/go-cast
```

## API

#### Convert Function

```go
func ToTime(any any) (dst time.Time, err error)
func ToBool(any any) (dst bool, err error)
func ToInt64(any any) (dst int64, err error)
func ToUint64(any any) (dst uint64, err error)
func ToFloat64(any any) (dst float64, err error)
func ToString(any any) (dst string, err error)
func ToDuration(any any) (dst time.Duration, err error)

func ToTimeInLocation(any any, loc *time.Location, layouts ...string) (time.Time, error)
func MustToTimeInLocation(any any, loc *time.Location, layouts ...string) time.Time
func MustParseTime(value string, loc *time.Location, layouts ...string) time.Time
func TryParseTime(value string, loc *time.Location, layouts ...string) (time.Time, error)
func Set(dst, src any) (err error)

// Must is the generic function and used by associating with ToXXX. For example,
//   Must(ToBool(any))
//   Must(ToInt64(any))
//   Must(ToUint64(any))
//   Must(ToFloat64(any))
//   Must(ToString(any))
//   Must(ToDuration(any))
//   Must(ToTime(any))
func Must[T any](value T, err error) T
```
