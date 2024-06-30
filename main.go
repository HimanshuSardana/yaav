package main

import (
	"log"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"github.com/nsf/termbox-go"
)

func main() {
	speed := 1.0

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

	resampled := beep.ResampleRatio(4, speed, streamer)
	speaker.Play(resampled)

	var mu sync.Mutex

	go func() {
		err := termbox.Init()
		if err != nil {
			log.Fatal(err)
		}
		defer termbox.Close()

		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				mu.Lock()
				if ev.Key == termbox.KeyArrowRight || ev.Ch == ']' {
					speed += 0.1
					resampled.SetRatio(speed)
					fmt.Println(speed)
				} else if ev.Key == termbox.KeyArrowLeft || ev.Ch == '[' {
					speed -= 0.1
					resampled.SetRatio(speed)
					fmt.Println(speed)
				} else if ev.Ch == 'q' {
					os.Exit(0)
				}
				mu.Unlock()
			case termbox.EventInterrupt, termbox.EventError, termbox.EventResize:
				return
			}
		}
	}()

	select {}
}
