package audioencoder

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"
)

//Reader is a Strucuture definition for the audiofile reader
type Reader struct{}

//ReadAudioFile is a function that reads a audio file and gets its file type
func (r Reader) ReadAudioFile(path string) (file *AudioFile, err error) {
	result := &AudioFile{}

	result.File, err = os.Open(path)

	if err != nil {
		return result, err
	}

	result.Type, err = r.getFileContentType(result.File)

	return result, err
}

func (r Reader) getFileContentType(out *os.File) (AudioType, error) {

	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return AudioTypeUNKNOWN, err
	}

	defer out.Seek(0, io.SeekStart)

	contentType := http.DetectContentType(buffer)

	switch contentType {
	case "audio/wave":
		return AudioTypeWAV, nil
	case "audio/mpeg":
		return AudioTypeMP3, nil
	default:
		return AudioTypeUNKNOWN, nil
	}
}

//GetAudioMatrix return a audio matriz
func (r Reader) GetAudioMatrix(file *AudioFile) (spec *AudioSpectrum, err error) {
	if file == nil || file.File == nil {
		return nil, errors.New("File stream not found. Please load a valid audio file")
	}

	var streamer beep.StreamSeekCloser
	var format beep.Format

	switch file.Type {
	case AudioTypeMP3:
		streamer, format, err = mp3.Decode(file.File)
		break
	case AudioTypeWAV:
		streamer, format, err = wav.Decode(file.File)
		break
	default:
		return nil, errors.New("Audio Type not recognized. Please load a valid audio file")
	}

	if err != nil {
		return nil, err
	}

	defer streamer.Close()

	spec = &AudioSpectrum{}
	spec.Format.SampleRate = int(format.SampleRate)
	spec.Format.NumChannels = format.NumChannels
	spec.Format.Precision = format.Precision

	spec.Matrix = make([][2]float64, streamer.Len())

	_, ok := streamer.Stream(spec.Matrix)

	if !ok {
		return nil, errors.New("Could not read matrix audio from audio stream")
	}

	return spec, nil
}
