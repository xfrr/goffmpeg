# Goffmpeg
FFMPEG wrapper written in GO which allows to obtain the progress.

# Dependencies
- [FFmpeg](https://www.ffmpeg.org/)
- [FFProbe](https://www.ffmpeg.org/ffprobe.html)

# Supported platforms

 - Linux
 - OS X
 - Windows

# Getting started
How to transcode a media file
```go
package main

import (
    "goffmpeg/transcoder"
)

var inputPath = "/data/testmov"
var outputPath = "/data/testmp4.mp4"

func main() {

	// Create new instance of transcoder
    trans := new(transcoder.Transcoder)
	
	// Initialize transcoder passing the input file path and output file path
    err := trans.Initialize( inputPath, outputPath )
    // Handle error...

	// Start transcoder process
	done, err := trans.Run()
	
	// This channel is used to wait for the process to end
	<-done

}
```
How to get the transcoding progress
```go
...
func main() {

	// Create new instance of transcoder
    trans := new(transcoder.Transcoder)
	
	// Initialize transcoder passing the input file path and output file path
    err := trans.Initialize( inputPath, outputPath )
    // Handle error...

	// Start transcoder process
	done, err := trans.Run()
	
	// Returns a channel to get the transcoding progress
	progress, err := trans.Output()

	// Example of printing transcoding progress
	for msg := range progress {
		fmt.Println(msg)
	}
	
	// This channel is used to wait for the transcoding process to end
	<-done

}
```
# Progress properties
```golang
type Progress struct {
	FramesProcessed			string
	CurrentTime			string
	CurrentBitrate			string
	Progress			float64
}
```
# Media setters
Those options can be set before starting the transcoding.
```js
SetAspect
SetResolution
SetVideoBitRate
SetVideoBitRateTolerance
SetVideoMaxBitrate
SetVideoMinBitRate
SetVideoCodec
SetFrameRate
SetMaxKeyFrame
SetMinKeyFrame
SetKeyframeInterval
SetAudioCodec
SetAudioBitRate
SetAudioChannels
SetBufferSize
SetThreads
SetPreset
SetDuration
SetSeekTime
SetQuality
```

----
> Building...
