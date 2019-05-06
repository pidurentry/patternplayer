package main

import (
	"bufio"
	"fmt"
	"github.com/pidurentry/buttplug-go"
	"github.com/pidurentry/pattern"
	"github.com/pidurentry/pattern/tools"
	"github.com/pidurentry/patternplayer/device"
	"os"
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

	player := NewPlayer(pattern, newDevice())
	if err := player.Start(); err != nil {
		panic(err)
	}
}

func newDevice() tools.Device {
	conn, err := buttplug.Dial("ws://localhost:12345/buttplug")
	if err != nil {
		panic(fmt.Sprintf("%#v", err))
	}
	handler := buttplug.NewHandler(conn)

	info, err := handler.Handshake("buttplug-example")
	if err != nil {
		panic(fmt.Sprintf("%#v", err))
	}
	fmt.Printf("%#v\n", info)

	go func() {
		ticker := info.MaxPingTime().Ticker()
		for {
			<-ticker.C
			if !handler.Ping() {
				panic("failed to ping server")
			}
		}
	}()

	deviceManager := buttplug.NewDeviceManager(handler)
	deviceManager.Scan(15 * time.Second).Wait()

	device, err := device.NewFleshlightLaunch(deviceManager.FleshlightLaunches()[0])
	if err != nil {
		panic(err)
	}
	return device
}
