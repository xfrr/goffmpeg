# Goffmpeg
[![Build & Test](https://github.com/xfrr/goffmpeg/actions/workflows/build_and_test.yml/badge.svg?branch=v2)](https://github.com/xfrr/goffmpeg/actions/workflows/build_and_test.yml)
[![Code Quality](https://api.codacy.com/project/badge/Grade/93e018e5008b4439acbb30d715b22e7f)](https://www.codacy.com/app/francisco.romero/goffmpeg?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=xfrr/goffmpeg&amp;utm_campaign=Badge_Grade)
[![Coverage](https://codecov.io/gh/xfrr/goffmpeg/graph/badge.svg?token=LjqrgDKO69)](https://codecov.io/gh/xfrr/goffmpeg)
[![Go Report Card](https://goreportcard.com/badge/github.com/xfrr/goffmpeg)](https://goreportcard.com/report/github.com/xfrr/goffmpeg)
[![GoDoc](https://godoc.org/github.com/xfrr/goffmpeg/v2?status.svg)](https://godoc.org/github.com/xfrr/goffmpeg/v2)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)

FFMPEG wrapper written in the Go

## Supported features

- [x] Execute ffmpeg command and get progress events
- [x] Read file and get metadata with ffprobe
- [x] Pipe Options
- [x] Global Options
- [x] Video Options
- [x] Audio Options

## Requirements
- [FFmpeg](https://www.ffmpeg.org/) - Supported versions: 4.0 ~ 6.0
- [FFProbe](https://www.ffmpeg.org/ffprobe.html)

## Supported platforms

 - Linux
 - OS X
 - Windows

## Installation
Install the package with the following command:
```shell
go get github.com/xfrr/goffmpeg/v2
```

## Usage
Check the [examples](./examples)
