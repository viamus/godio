package audioencoder

import "os"

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
	Matrix [][2]float64
	Format AudioFormat
}

//AudioFormat is  a audio converted metada struct
type AudioFormat struct {
	SampleRate  int
	NumChannels int
	Precision   int
}
