# Goffmpeg
FFMPEG wrapper written in GO


# Dependencies
- [FFmpeg](https://www.ffmpeg.org/)

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
Manipulating media file
> Building
