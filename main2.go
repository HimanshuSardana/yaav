package main

import (
    "fmt"
    "github.com/gordonklaus/portaudio"
    "math"
)

func main() {
    portaudio.Initialize()
    defer portaudio.Terminate()

    in := make([]float32, 64)
    stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
    if err != nil {
        panic(err)
    }
    defer stream.Close()

    err = stream.Start()
    if err != nil {
        panic(err)
    }
    defer stream.Stop()

    for {
        err = stream.Read()
        if err != nil {
            panic(err)
        }
        amplitude := calculateAmplitude(in)
        fmt.Printf("Amplitude: %f\n", amplitude)
    }
}

func calculateAmplitude(samples []float32) float32 {
    var sum float32
    for _, sample := range samples {
        sum += sample * sample
    }
    return float32(math.Sqrt(float64(sum / float32(len(samples)))))
}

