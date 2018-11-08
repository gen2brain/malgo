## malgo
[![TravisCI Build Status](https://travis-ci.org/gen2brain/malgo.svg?branch=master)](https://travis-ci.org/gen2brain/malgo) 
[![AppVeyor Build Status](https://ci.appveyor.com/api/projects/status/eofqkk271yjd3s3g?svg=true)](https://ci.appveyor.com/project/gen2brain/malgo)
[![GoDoc](https://godoc.org/github.com/gen2brain/malgo?status.svg)](https://godoc.org/github.com/gen2brain/malgo) 
[![Go Report Card](https://goreportcard.com/badge/github.com/gen2brain/malgo?branch=master)](https://goreportcard.com/report/github.com/gen2brain/malgo) 
<!--[![Go Cover](http://gocover.io/_badge/github.com/gen2brain/malgo)](http://gocover.io/github.com/gen2brain/malgo)-->

Go bindings for [mini_al](https://github.com/dr-soft/mini_al), mini audio library.

Requires `cgo` but does not require linking to anything on the Windows/macOS and it links only `-ldl` on Linux.

### Installation

    go get -u github.com/gen2brain/malgo

### Documentation

Documentation on [GoDoc](https://godoc.org/github.com/gen2brain/malgo). Also check [examples](https://github.com/gen2brain/malgo/tree/master/_examples).

### Platforms

* Windows (WASAPI, DirectSound, WinMM)
* Linux (PulseAudio, ALSA, JACK)
* FreeBSD/NetBSD/OpenBSD (OSS)
* macOS (CoreAudio)
* Android (OpenSL|ES)
