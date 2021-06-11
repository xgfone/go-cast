# Go Type Cast [![Build Status](https://api.travis-ci.com/xgfone/cast.svg?branch=master)](https://travis-ci.com/github/xgfone/cast) [![GoDoc](https://pkg.go.dev/badge/github.com/xgfone/cast)](https://pkg.go.dev/github.com/xgfone/cast) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](https://raw.githubusercontent.com/xgfone/cast/master/LICENSE)

The package supporting `Go1.7+` supplies some functions to convert the data between types, such as `ToXXX` and `MustToXXX`.

## Installation
```shell
$ go get -u github.com/xgfone/cast
```

## Usage

```go
// Bool
boolval, _ := cast.ToBool("1")
boolval, _ := cast.ToBool("0")
boolval, _ := cast.ToBool("t")
boolval, _ := cast.ToBool("f")
boolval, _ := cast.ToBool("on")
boolval, _ := cast.ToBool("off")
boolval, _ := cast.ToBool("true")
boolval, _ := cast.ToBool("false")

// Int
intval, _ := cast.ToInt("123")
intval, _ := cast.ToInt(123)
intval, _ := cast.ToInt(1.23)

// Int64
int64val, _ := cast.ToInt64("123")
int64val, _ := cast.ToInt64(123)
int64val, _ := cast.ToInt64(1.23)

// Uint
uintval, _ := cast.ToUint("123")
uintval, _ := cast.ToUint(123)
uintval, _ := cast.ToUint(1.23)

// Uint64
uint64val, _ := cast.ToUint64("123")
uint64val, _ := cast.ToUint64(123)
uint64val, _ := cast.ToUint64(1.23)

// Float64
float64val, _ := cast.ToFloat64("123")
float64val, _ := cast.ToFloat64("1.23")
float64val, _ := cast.ToFloat64(123)
float64val, _ := cast.ToFloat64(1.23)

// String
stringval, _ := cast.ToString("abc")
stringval, _ := cast.ToString(123)
stringval, _ := cast.ToString(1.23)
stringval, _ := cast.ToString(true)
stringval, _ := cast.ToString(false)

// Time
timeval, _ := cast.ToTime("2020-05-12 18:40:00")
timeval, _ := cast.ToTime("2020-05-12T18:40:00Z")
timeval, _ := cast.ToTime("2020-05-12T18:40:00+08:00")

// ......
```
