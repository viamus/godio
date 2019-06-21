package main

import (
	"log"
	"os"

	"github.com/viamus/godio/pkg/audioencoder"
)

//GetAudioMatrix returning audio matrix from file
func GetAudioMatrix(path string) (audiomatrix *audioencoder.AudioSpectrum, err error) {
	reader := &audioencoder.Reader{}

	log.Println("Reading audio data.")
	audiofile, err := reader.ReadAudioFile(path)

	if err != nil {
		return nil, err
	}

	log.Println("Converting audio data into audio matrix.")
	audiomatrix, err = reader.GetAudioMatrix(audiofile)

	log.Printf("Matrix converted with size of %d bytes.", len(audiomatrix.Matrix))

	return audiomatrix, err
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetPrefix("[GODIO] ")

	if len(os.Args) == 1 {
		log.Fatal("File path not informed.")
	}

	args := os.Args[1]

	_, err := GetAudioMatrix(args)

	handleError(err)
}
