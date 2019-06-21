package audioencoder

import (
	"encoding/gob"
	"os"
)

const (
	//AudioTypeUNKNOWN AudioType enum definition for unkown files
	AudioTypeUNKNOWN AudioType = 0
	//AudioTypeMP3 AudioType enum definition for mp3 files
	AudioTypeMP3 AudioType = 1
	//AudioTypeWAV AudioType enum definition for wav files
	AudioTypeWAV AudioType = 2
)

//AudioFile is a audio sample file
type AudioFile struct {
	Type AudioType
	File *os.File
}

//AudioType is a constant of a audiofile type
type AudioType int

//AudioSpectrum is a audio frequency vector
type AudioSpectrum struct {
	Matrix [][2]float64 `json:"matrix"`
	Format AudioFormat  `json:"format"`
}

//AudioFormat is  a audio converted metada struct
type AudioFormat struct {
	SampleRate  int `json:"samplerate"`
	NumChannels int `json:"numchannels"`
	Precision   int `json:"precision"`
}

//Save a binary file of the audio matrix structure
func (s AudioSpectrum) Save(path string) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(s)
	}
	file.Close()
	return err
}

//Load a binary file of the audio matrix strucutre
func (s *AudioSpectrum) Load(path string) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(s)
	}
	file.Close()
	return err
}
