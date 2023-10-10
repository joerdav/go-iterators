# Go Iterators

[![xc compatible](https://xcfile.dev/badge.svg)](https://xcfile.dev)

Just messing around with Go function iterators.

## Tasks

### tools

gotip is required as this is an unreleased feature.

This installs gotip and installs the iterator branch.

```bash
go install golang.org/dl/gotip@latest
gotip download 510541   # download and build CL 510541              
gotip version  # should say "(w/ rangefunc)"                             
gotip run foo.go       
```

### run

Environment: GOEXPERIMENT=range

```bash
gotip run main.go
```

