package main

import (
	"os"
	"bufio"
	"github.com/pidurentry/pattern"
	"fmt"
	"time"
)

func main() {
	file, err := os.Open(config.pattern)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pattern, err := (&pattern.Parser{}).Load(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	player := NewPlayer(pattern, &debug{})
	if err := player.Start(); err != nil {
		panic(err)
	}
}

type debug struct {}

func (*debug) Move(Value, Speed uint64) error {
	fmt.Printf("Move to %d with speed %d\n", Value, Speed)
	time.Sleep(250 * time.Millisecond)
	return nil
}

func (*debug) Rotate(Speed uint64, Clockwise bool) error {
	direction := "clockwise"
	if !Clockwise {
		direction = "anticlockwise"
	}
	fmt.Printf("Rotate %s with speed %d\n", direction, Speed)
	time.Sleep(250 * time.Millisecond)
	return nil
}

func (*debug) Vibrate(Speed uint64) error {
	fmt.Printf("Vibrate at speed %d\n", Speed)
	time.Sleep(250 * time.Millisecond)
	return nil
}