package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/term"
)

var prevResults []float64

func handleInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for range c {
		fmt.Println("\nBye!")
		close(c)
		os.Exit(0)
	}
}

func readInput(input chan []byte) {
	// Configure terminal for raw input on Windows
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error configuring terminal:", err)
		os.Exit(1)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	var b []byte = make([]byte, 1)
	for {
		_, err := os.Stdin.Read(b)
		if err != nil {
			break
		}
		input <- b
	}
}

func calcBpm(durations []time.Duration) int64 {
	n := len(durations)
	sum := int64(0)
	for i := 0; i < n; i++ {
		sum += durations[i].Milliseconds()
	}
	average := float64(sum) / float64(n)
	minInMs := float64((time.Second * 60).Milliseconds())
	currentBpm := minInMs / average

	prevResults = append(prevResults, currentBpm)
	if len(prevResults) > 3 {
		prevResults = prevResults[len(prevResults)-3:]
	}
	n = len(prevResults)
	prevSum := float64(0.0)
	for i := 0; i < n; i++ {
		prevSum += prevResults[i]
	}
	bpm := int64(float64(prevSum) / float64(n))

	return bpm
}

func handleTap() {
	input := make(chan []byte)
	go readInput(input)

	s := time.Now()
	var durations []time.Duration

	fmt.Println("Tempo Tapper - Tap any key to the rhythm!")
	fmt.Println("Press Ctrl+C to quit")
	fmt.Print("BPM: ")

	for {
		select {
		case <-input:
			d := time.Since(s)
			durations = append(durations, d)
			if len(durations) > 4 {
				durations = durations[len(durations)-4:]
			}
			fmt.Printf("\r%3d", calcBpm(durations))
			s = time.Now()
		}
	}
}

func main() {
	go handleInterrupt()
	handleTap()
}
