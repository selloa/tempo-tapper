package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	scanner := bufio.NewScanner(os.Stdin)
	s := time.Now()
	var durations []time.Duration

	fmt.Println("Tempo Tapper - Press Enter to tap!")
	fmt.Println("Press Ctrl+C to quit")
	fmt.Print("BPM: ")

	for scanner.Scan() {
		d := time.Since(s)
		durations = append(durations, d)
		if len(durations) > 4 {
			durations = durations[len(durations)-4:]
		}
		fmt.Printf("\r%3d", calcBpm(durations))
		s = time.Now()
	}
}

func main() {
	go handleInterrupt()
	handleTap()
}
