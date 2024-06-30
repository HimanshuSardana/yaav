package main

import (
    "log"
    "os"
    "github.com/gopxl/beep"
    "github.com/gopxl/beep/mp3"
    "github.com/gopxl/beep/speaker"
    "time"
)

type Reverb struct {
    streamer beep.Streamer
    delay    int
    feedback float64
    buffer   []float64
    index    int
}

func NewReverb(streamer beep.Streamer, delay time.Duration, feedback float64, sampleRate beep.SampleRate) *Reverb {
    delaySamples := int(sampleRate.N(delay))
    return &Reverb{
        streamer: streamer,
        delay:    delaySamples,
        feedback: feedback,
        buffer:   make([]float64, delaySamples),
    }
}

func (r *Reverb) Stream(samples [][2]float64) (n int, ok bool) {
    n, ok = r.streamer.Stream(samples)
    for i := range samples[:n] {
        delayedSample := r.buffer[r.index]
        for ch := 0; ch < len(samples[i]); ch++ {
            samples[i][ch] += delayedSample * r.feedback
        }
        r.buffer[r.index] = samples[i][0]
        r.index = (r.index + 1) % r.delay
    }
    return n, ok
}

func (r *Reverb) Err() error {
    return r.streamer.Err()
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

    sampleRate := format.SampleRate

    speaker.Init(sampleRate, sampleRate.N(time.Second/10))

    reverb := NewReverb(streamer, 50*time.Millisecond, 0.5, sampleRate)

    speaker.Play(beep.Seq(reverb, beep.Callback(func() {
        log.Println("Playback finished")
    })))

    select {}
}

