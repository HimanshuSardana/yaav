package main

import (
	"log"
	"os"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"time"
)

func main() {
	speed := 1.5

	f, err := os.Open("Always.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	NewSampleRate := beep.SampleRate(int(float64(format.SampleRate) * speed))

	speaker.Init(NewSampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)
	select {}
}
