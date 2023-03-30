# Go Type Cast [![Build Status](https://github.com/xgfone/cast/actions/workflows/go.yml/badge.svg)](https://github.com/xgfone/cast/actions/workflows/go.yml) [![GoDoc](https://pkg.go.dev/badge/github.com/xgfone/cast)](https://pkg.go.dev/github.com/xgfone/cast) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](https://raw.githubusercontent.com/xgfone/cast/master/LICENSE)

Provide some functions, supporting `Go1.18+`, to convert the value between different types, such as `ToXXX`.

## Installation
```shell
$ go get -u github.com/xgfone/cast
```

## API

#### Convert Function
```go
func ToTime(any interface{}) (dst time.Time, err error)
func ToBool(any interface{}) (dst bool, err error)
func ToInt64(any interface{}) (dst int64, err error)
func ToUint64(any interface{}) (dst uint64, err error)
func ToFloat64(any interface{}) (dst float64, err error)
func ToString(any interface{}) (dst string, err error)
func ToDuration(any interface{}) (dst time.Duration, err error)

func ToTimeInLocation(any interface{}, loc *time.Location, layouts ...string) (time.Time, error)
func MustToTimeInLocation(any interface{}, loc *time.Location, layouts ...string) time.Time
func MustParseTime(value string, loc *time.Location, layouts ...string) time.Time
func TryParseTime(value string, loc *time.Location, layouts ...string) (time.Time, error)

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
