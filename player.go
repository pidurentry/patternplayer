package main

import (
	"fmt"
	"github.com/pidurentry/pattern"
	"github.com/pidurentry/pattern/tools"
	"time"
)

type Player struct {
	pattern   pattern.Pattern
	device    tools.Device
	active    string
	queue     []tools.Action
	interrupt chan bool
}

func NewPlayer(pattern pattern.Pattern, device tools.Device) *Player {
	return &Player{
		pattern:   pattern,
		device:    device,
		active:    pattern.Pattern,
		interrupt: make(chan bool, 0),
	}
}

func (player *Player) Start() error {
	for {
		if player.queue == nil || len(player.queue) == 0 {
			queue, ok := player.pattern.Patterns[player.active]
			if !ok {
				return fmt.Errorf("Unknown pattern '%s'", player.active)
			}

			player.queue = make([]tools.Action, len(queue))
			copy(player.queue, queue)
		}

		if len(player.queue) == 0 {
			return fmt.Errorf("Active pattern '%s' has no actions", player.active)
		}

		action := player.queue[0]
		player.queue = player.queue[1:]
		action.Apply(player, player.pattern.Variables, player.device)

		select {
		case <-player.interrupt:
			<-player.interrupt
		default:
		}
	}
	return nil
}

func (player *Player) Pause() error {
	player.interrupt <- true
	return nil
}

func (player *Player) Goto(pattern string) error {
	if _, ok := player.pattern.Patterns[pattern]; !ok {
		return fmt.Errorf("Unknown pattern '%s'", pattern)
	}

	player.active = pattern
	player.queue = nil
	return nil
}

func (player *Player) QueueActions(actions []tools.Action) error {
	if player.queue == nil || len(player.queue) == 0 {
		player.queue = make([]tools.Action, len(actions))
		copy(player.queue, actions)
	} else {
		queue := make([]tools.Action, len(actions))
		copy(queue, actions)
		player.queue = append(queue, player.queue...)
	}
	return nil
}

func (player *Player) Sleep(duration time.Duration) error {
	time.Sleep(duration)
	return nil
}
