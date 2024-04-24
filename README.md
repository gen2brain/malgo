## malgo
[![Build Status](https://github.com/gen2brain/malgo/actions/workflows/build.yml/badge.svg)](https://github.com/gen2brain/malgo/actions)
[![GoDoc](https://godoc.org/github.com/gen2brain/malgo?status.svg)](https://godoc.org/github.com/gen2brain/malgo) 
[![Go Report Card](https://goreportcard.com/badge/github.com/gen2brain/malgo?branch=master)](https://goreportcard.com/report/github.com/gen2brain/malgo) 
<!--[![Go Cover](http://gocover.io/_badge/github.com/gen2brain/malgo)](http://gocover.io/github.com/gen2brain/malgo)-->

Go bindings for [miniaudio](https://github.com/dr-soft/miniaudio) library.

Requires `cgo` but does not require linking to anything on the Windows/macOS and it links only `-ldl` on Linux/BSDs.

### Installation

    go get -u github.com/gen2brain/malgo

### Documentation

Documentation on [GoDoc](https://godoc.org/github.com/gen2brain/malgo). Also check [examples](https://github.com/gen2brain/malgo/tree/master/_examples).

### Platforms

* Windows (WASAPI, DirectSound, WinMM)
* Linux (PulseAudio, ALSA, JACK)
* FreeBSD/NetBSD/OpenBSD (OSS/audio(4)/sndio)
* macOS/iOS (CoreAudio)
* Android (OpenSL|ES, AAudio)
