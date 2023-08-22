[![Build Status](https://travis-ci.org/hpcloud/tail.svg)](https://travis-ci.org/hpcloud/tail)
[![Build status](https://ci.appveyor.com/api/projects/status/vrl3paf9md0a7bgk/branch/master?svg=true)](https://ci.appveyor.com/project/Nino-K/tail/branch/master)

# Go package for tail-ing files

A Go package striving to emulate the features of the BSD `tail` program.

**fork from github.com/hpcloud/tail**

## changelog

1. support go module and remove vendor dir
2. upgrade dependent lib to latest version and go version up to go 1.17
3. fix goroutine crash possibility

## go module config usage
```
require (
	github.com/hpcloud/tail v1.0.0
)

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)

replace github.com/hpcloud/tail v1.0.0 => github.com/liumingmin/tail v1.0.2
```


## go module  only use poll watch
```
require (
	github.com/hpcloud/tail v1.0.0
)

require gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect

replace github.com/hpcloud/tail v1.0.0 => github.com/liumingmin/tail v1.1.0
```

## code usage
```Go
t, err := tail.TailFile("/var/log/nginx.log", tail.Config{Follow: true, Poll: true})
for line := range t.Lines {
    fmt.Println(line.Text)
}
```

See [API documentation](http://godoc.org/github.com/hpcloud/tail).

## Log rotation

Tail comes with full support for truncation/move detection as it is
designed to work with log rotation tools.

## Installing

    go get github.com/hpcloud/tail/...

## Windows support

**windows should use poll mode instead of fsnotify**

```go
tail.TailFile("/var/log/nginx.log", tail.Config{Follow: true, Poll: true})
```


This package [needs assistance](https://github.com/hpcloud/tail/labels/Windows) for full Windows support.
