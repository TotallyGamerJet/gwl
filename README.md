GWL
======
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[godocs]: http://godoc.org/github.com/totallygamerjet/gwl
GWL is a crossplatform window and input api written in pure Go inspired by [glfw](https://github.com/glfw/glfw).

## Goals
1. Allow for concurrent calls to all functions
2. Provide an idiomatic Go experience
3. No CGO

##Note
This project is a WIP and only currently builds for windows. I have not found
an easy way to interface with MacOS without CGO. If you have an idea please
create an issue. The Linux version will use [XGB](https://github.com/BurntSushi/xgb)
but I do not have a system to develop on.

## Usage
1. Get the package using `go get github.com/totallygamerjet/gwl`
3. Import the package `import "github.com/totallygamerjet/gwl"`

## Example
Examples can be found in the [examples](https://github.com/TotallyGamerJet/gwl/tree/master/examples) folder.

## Contribute
All contributions are welcome: bug reports, pull requests, documentation etc.
