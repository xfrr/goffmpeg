# Goffmpeg
[![Build & Test](https://github.com/xfrr/goffmpeg/actions/workflows/build_and_test.yml/badge.svg)](https://github.com/xfrr/goffmpeg/actions/workflows/build_and_test.yml)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/93e018e5008b4439acbb30d715b22e7f)](https://www.codacy.com/app/francisco.romero/goffmpeg?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=xfrr/goffmpeg&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/xfrr/goffmpeg)](https://goreportcard.com/report/github.com/xfrr/goffmpeg)
[![GoDoc](https://godoc.org/github.com/xfrr/goffmpeg?status.svg)](https://godoc.org/github.com/xfrr/goffmpeg)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)

FFMPEG wrapper written in GO

> New implementation with an easy-to-use API and interfaces to extend the transcoding capabilities.
> https://github.com/floostack/transcoder

## Features

- [x] Transcoding
- [x] Streaming
- [x] Progress
- [x] Filters
- [x] Thumbnails
- [x] Watermark
- [ ] Concatenation
- [ ] Subtitles

## Dependencies
- [FFmpeg](https://www.ffmpeg.org/)
- [FFProbe](https://www.ffmpeg.org/ffprobe.html)

## Supported platforms

 - Linux
 - OS X
 - Windows

## Installation
Install the package with the following command:
```shell
go get github.com/xfrr/goffmpeg
```

## Usage
Check the [examples](./examples)