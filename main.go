package main

import (
	"io"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func main() {
	f, err := os.Open("audio.mp3")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder, err := mp3.NewDecoder(f)
	if err != nil {
		panic(err)
	}

	// SampleRate, NumChannels, BytesPerSample
	context, err := oto.NewContext(44100, 2, 2, 8192)
	if err != nil {
		panic(err)
	}
	player := context.NewPlayer()
	defer player.Close()

	if _, err := io.Copy(player, decoder); err != nil {
		panic(err)
	}
}
