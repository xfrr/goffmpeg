package media_test

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"testing"
	"time"

	"github.com/xfrr/goffmpeg/v2/pkg/media"
)

const mockJSON = `{
	"streams": [
			{
					"index": 0,
					"codec_name": "h264",
					"codec_long_name": "H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10",
					"profile": "Constrained Baseline",
					"codec_type": "video",
					"codec_tag_string": "avc1",
					"codec_tag": "0x31637661",
					"width": 480,
					"height": 270,
					"coded_width": 480,
					"coded_height": 270,
					"closed_captions": 10,
					"initial_padding": 0,
					"has_b_frames": 0,
					"sample_aspect_ratio": "1:1",
					"display_aspect_ratio": "16:9",
					"pix_fmt": "yuv420p",
					"level": 30,
					"field_order": "progressive",
					"refs": 1,
					"is_avc": "true",
					"nal_length_size": "4",
					"id": "0x1",
					"r_frame_rate": "30/1",
					"avg_frame_rate": "30/1",
					"time_base": "1/30",
					"start_pts": 0,
					"start_time": "0.000000",
					"duration_ts": 901,
					"duration": "30.033333",
					"bit_rate": "301201",
					"bits_per_raw_sample": "8",
					"nb_frames": "901",
					"extradata_size": 40,
					"disposition": {
							"default": 1,
							"dub": 0,
							"original": 0,
							"comment": 0,
							"lyrics": 0,
							"karaoke": 0,
							"forced": 0,
							"hearing_impaired": 0,
							"visual_impaired": 0,
							"clean_effects": 0,
							"attached_pic": 0,
							"timed_thumbnails": 0,
							"captions": 0,
							"descriptions": 0,
							"metadata": 0,
							"dependent": 0,
							"still_image": 0
					},
					"tags": {
							"creation_time": "2015-08-07T09:13:02.000000Z",
							"language": "und",
							"handler_name": "L-SMASH Video Handler",
							"vendor_id": "[0][0][0][0]",
							"encoder": "AVC Coding"
					}
			},
			{
					"index": 1,
					"codec_name": "aac",
					"codec_long_name": "AAC (Advanced Audio Coding)",
					"profile": "LC",
					"codec_type": "audio",
					"codec_tag_string": "mp4a",
					"codec_tag": "0x6134706d",
					"sample_fmt": "fltp",
					"sample_rate": "48000",
					"channels": 2,
					"channel_layout": "stereo",
					"bits_per_sample": 0,
					"initial_padding": 0,
					"id": "0x2",
					"r_frame_rate": "0/0",
					"avg_frame_rate": "0/0",
					"time_base": "1/48000",
					"start_pts": 0,
					"start_time": "0.000000",
					"duration_ts": 1465280,
					"duration": "30.526667",
					"bit_rate": "112000",
					"nb_frames": "1431",
					"extradata_size": 2,
					"disposition": {
							"default": 1,
							"dub": 0,
							"original": 0,
							"comment": 0,
							"lyrics": 0,
							"karaoke": 0,
							"forced": 0,
							"hearing_impaired": 0,
							"visual_impaired": 0,
							"clean_effects": 0,
							"attached_pic": 0,
							"timed_thumbnails": 0,
							"captions": 0,
							"descriptions": 0,
							"metadata": 0,
							"dependent": 0,
							"still_image": 0
					},
					"tags": {
							"creation_time": "2015-08-07T09:13:02.000000Z",
							"language": "und",
							"handler_name": "L-SMASH Audio Handler",
							"vendor_id": "[0][0][0][0]"
					}
			}
	],
	"format": {
			"filename": "testdata/input.mp4",
			"nb_streams": 2,
			"nb_programs": 0,
			"format_name": "mov,mp4,m4a,3gp,3g2,mj2",
			"format_long_name": "QuickTime / MOV",
			"start_time": "0.000000",
			"duration": "30.526667",
			"size": "1570024",
			"bit_rate": "411449",
			"probe_score": 100,
			"tags": {
					"major_brand": "mp42",
					"minor_version": "0",
					"compatible_brands": "mp42mp41isomavc1",
					"creation_time": "2015-08-07T09:13:02.000000Z"
			}
	}
}`

func TestFile_UnmarshalJSON(t *testing.T) {
	var mockFile media.File
	if err := json.Unmarshal([]byte(mockJSON), &mockFile); err != nil {
		t.Errorf("File.UnmarshalJSON() error = %+v", err)
		return
	}

	t.Run("should unmarshal json to File struct", func(t *testing.T) {
		var file media.File
		if err := file.UnmarshalJSON([]byte(mockJSON)); err != nil {
			t.Errorf("File.UnmarshalJSON() error = %+v", err)
			return
		}
		if file.Filename() != "testdata/input.mp4" {
			t.Errorf("File.UnmarshalJSON() = %+v, WANT %+v", file.Filename(), "testdata/input.mp4")
		}
		if file.FormatName() != "mov,mp4,m4a,3gp,3g2,mj2" {
			t.Errorf("File.UnmarshalJSON() = %+v, WANT %+v", file.FormatName(), "mov,mp4,m4a,3gp,3g2,mj2")
		}
		if file.FormatLongName() != "QuickTime / MOV" {
			t.Errorf("File.UnmarshalJSON() = %+v, WANT %+v", file.FormatLongName(), "QuickTime / MOV")
		}
		if file.BitRate() != "411449" {
			t.Errorf("File.UnmarshalJSON() = %+v, WANT %+v", file.BitRate(), "411449")
		}
		if file.Size() != "1570024" {
			t.Errorf("File.UnmarshalJSON() = %+v, WANT %+v", file.Size(), "1570024")
		}
		if file.Encoder() != "" {
			t.Errorf("File.UnmarshalJSON() = %+v, WANT %+v", file.Encoder(), "")
		}
		if file.NbStreams() != 2 {
			t.Errorf("File.UnmarshalJSON() = %+v, WANT %+v", file.NbStreams(), 2)
		}
		if file.NbPrograms() != 0 {
			t.Errorf("File.UnmarshalJSON() = %+v, WANT %+v", file.NbPrograms(), 0)
		}
		if file.Duration() != 30*time.Second {
			t.Errorf("File.UnmarshalJSON() = %+v, WANT %+v", file.Duration(), 30*time.Second)
		}
		if file.ProbeScore() != 100 {
			t.Errorf("File.UnmarshalJSON() = %+v, WANT %+v", file.ProbeScore(), 100)
		}

		etags := media.Tags{
			MajorBrand:       "mp42",
			MinorVersion:     "0",
			CompatibleBrands: "mp42mp41isomavc1",
			CreationTime:     "2015-08-07T09:13:02.000000Z",
		}
		if !reflect.DeepEqual(file.Tags(), etags) {
			t.Errorf("File.UnmarshalJSON() = %+v, WANT %+v", file.Tags(), etags)
		}

		streams := file.Streams()
		if len(streams) != 2 {
			t.Errorf("File.UnmarshalJSON() = %+v, WANT %+v", len(streams), 2)
		}

		stream := file.StreamByIndex(0)
		if stream.Index != 0 {
			t.Errorf("File.UnmarshalJSON() = %+v, WANT %+v", stream.Index, 0)
		}

		estream := media.Stream{
			Index:              0,
			ID:                 "0x1",
			CodecName:          "h264",
			CodecLongName:      "H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10",
			Profile:            "Constrained Baseline",
			CodecType:          "video",
			CodecTimeBase:      "",
			CodecTagString:     "avc1",
			CodecTag:           "0x31637661",
			Width:              480,
			Height:             270,
			CodedWidth:         480,
			CodedHeight:        270,
			HasBFrames:         0,
			SampleAspectRatio:  "1:1",
			DisplayAspectRatio: "16:9",
			PixFmt:             "yuv420p",
			Level:              30,
			ChromaLocation:     "",
			Refs:               1,
			QuarterSample:      "",
			DivxPacked:         "",
			RFrameRrate:        "30/1",
			AvgFrameRate:       "30/1",
			TimeBase:           "1/30",
			DurationTs:         901,
			Duration:           "30.033333",
			BitRate:            "301201",
			NbFrames:           "901",
			ExtraDataSize:      40,
			BitsPerRawSample:   "8",
			BitsPerSample:      0,
			SampleFmt:          "",
			SampleRate:         "",
			Channels:           0,
			ClosedCaptions:     10,
			Disposition: media.Disposition{
				Default:         1,
				Dub:             0,
				Original:        0,
				Comment:         0,
				Lyrics:          0,
				Karaoke:         0,
				Forced:          0,
				HearingImpaired: 0,
				VisualImpaired:  0,
				CleanEffects:    0,
				AttachedPic:     0,
				TimedThumbnails: 0,
				Captions:        0,
				Descriptions:    0,
				Metadata:        0,
				Dependent:       0,
				StillImage:      0,
			},
			Tags: media.Tags{
				CreationTime: "2015-08-07T09:13:02.000000Z",
				Language:     "und",
				HandlerName:  "L-SMASH Video Handler",
				Encoder:      "AVC Coding",
				VendorID:     "[0][0][0][0]",
			},
			IsAvc:          "true",
			NalLengthSize:  "4",
			InitialPadding: 0,
			FieldOrder:     "progressive",
			StartTime:      "0.000000",
		}

		if !equalStream(*stream, estream) {
			t.Errorf("File.UnmarshalJSON() = %+v, WANT %+v", *stream, estream)
		}

		// marshal back to json
		sut, err := json.Marshal(file)
		if err != nil {
			t.Errorf("File.MarshalJSON() error = %+v", err)
			return
		}

		mockJSON, err := json.Marshal(mockFile)
		if err != nil {
			t.Errorf("File.MarshalJSON() error = %+v", err)
			return
		}

		equal, err := JSONEqual(bytes.NewReader(sut), bytes.NewReader([]byte(mockJSON)))
		if err != nil {
			t.Errorf("File.MarshalJSON() error = %+v", err)
			return
		}

		if !equal {
			t.Errorf("File.MarshalJSON() = %s, WANT %s", sut, mockJSON)
		}
	})
}

func equalStream(stream, estream media.Stream) bool {
	return reflect.DeepEqual(stream, estream)
}

// JSONEqual compares the JSON from two Readers.
func JSONEqual(a, b io.Reader) (bool, error) {
	var j, j2 map[string]interface{}
	d := json.NewDecoder(a)
	if err := d.Decode(&j); err != nil {
		return false, err
	}
	d = json.NewDecoder(b)
	if err := d.Decode(&j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}
