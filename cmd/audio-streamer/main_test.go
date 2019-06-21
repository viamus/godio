package main

import "testing"

func TestGetAudioMatrix(t *testing.T) {
	samples := []struct {
		path  string
		valid bool
	}{
		{"../../assets/audio-samples/sample-1.wav", false},
		{"../../assets/audio-samples/sample-2.wav", false},
		{"../../assets/audio-samples/sample-3.wav", false},
		{"../../assets/audio-samples/sample1.wav", true},
		{"../../assets/audio-samples/sample2.wav", true},
		{"../../assets/audio-samples/sample3.wav", true},
		{"../../assets/audio-samples/sample4.mp3", true},
		{"../../assets/audio-samples/sample5.mp3", true},
	}

	for _, sample := range samples {
		audiomatrix, err := GetAudioMatrix(sample.path)
		if sample.valid {
			if err != nil {
				t.Errorf("Sample:%s  Error:%s", sample.path, err.Error())
			} else if audiomatrix == nil {
				t.Errorf("Sample:%s  Error: AudioMatrix returned is nil, expected a valid value", sample.path)
			}
		} else {
			if err == nil {
				t.Errorf("Sample:%s  Error: AudioMatrix not valid but does not returned error", sample.path)
			}
		}
	}
}
