package device

import (
	"errors"
	"github.com/pidurentry/buttplug-go"
	"time"
)

type FleshlightLaunch struct {
	launch   buttplug.FleshlightLaunch
	current  int
	movement map[int]time.Duration
}

func NewFleshlightLaunch(launch buttplug.FleshlightLaunch) (*FleshlightLaunch, error) {
	device := &FleshlightLaunch{
		launch:  launch,
		current: 99,
		movement: map[int]time.Duration{
			10: 20 * time.Millisecond,
			20: 9 * time.Millisecond,
			30: 7 * time.Millisecond,
			40: 5 * time.Millisecond,
			50: 4 * time.Millisecond,
			60: 3500 * time.Microsecond,
		},
	}

	if err := device.Move(0, 20); err != nil {
		return nil, err
	}

	return device, nil
}

func (device *FleshlightLaunch) Move(Value, Speed uint64) error {
	position, speed := device.position(Value), device.speed(Speed)

	err := device.launch.FleshlightCmd(position, speed)
	if err != nil {
		return err
	}

	movement, ok := device.movement[speed]
	if ok {
		time.Sleep(movement * time.Duration(device.diff(position)))
	}
	device.current = position

	return nil
}

func (device *FleshlightLaunch) position(position uint64) int {
	if position < 0 {
		return 0
	}

	if position > 99 {
		return 0
	}

	return int(position)
}

func (device *FleshlightLaunch) speed(speed uint64) int {
	if speed < 10 {
		return 10
	}

	if speed > 80 {
		return 80
	}

	return int(speed)
}

func (device *FleshlightLaunch) diff(position int) int {
	if position > device.current {
		return position - device.current
	} else {
		return device.current - position
	}
}

func (device *FleshlightLaunch) Rotate(Speed uint64, Clockwise bool) error {
	return errors.New("unsupported command - rotate")
}

func (device *FleshlightLaunch) Vibrate(Speed uint64) error {
	return errors.New("unsupported command - vibrate")
}
