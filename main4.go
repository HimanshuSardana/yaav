package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

func extractSamples(streamer beep.Streamer) {
	bufferSize := 1024

	samples := make([][2]float64, bufferSize)

	for {
		n, ok := streamer.Stream(samples)
		if !ok {
			break
		}

		for i := 0; i < n; i++ {
			println(samples[i][0], samples[i][1]) 
		}
	}
}

func main() {
	f, err := os.Open("Always.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	extractSamples(streamer)
}
