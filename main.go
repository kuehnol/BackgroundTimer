package main

import (
	"flag"
	"github.com/gen2brain/beeep"
	"strconv"
	"time"
)

func main() {
	var (
		timeValue   int
		timeUnit    string
		timerReason string
	)

	flag.IntVar(&timeValue, "t", 0, "Provide a time value")
	flag.StringVar(&timeUnit, "u", "sec", "Provide a time unit (sec, min or h)")
	flag.StringVar(&timerReason, "r", "Timer has finished", "Provide a reason for the timer - use quotation marks")

	flag.Parse()

	timer(timeValue, timeUnit, timerReason)
}

func timer(timeValue int, unit string, reason string) {

	if timeValue < 1 {
		panic("timeValue under 1 or no value given")
	}

	timeDuration := time.Duration(timeValue)

	switch unit {
	case "sec":
		timeDuration = timeDuration * time.Second
	case "min":
		timeDuration = timeDuration * time.Minute
	case "h":
		timeDuration = timeDuration * time.Hour
	default:
		panic("No valid time unit given")
	}

	time.Sleep(timeDuration)

	titleString := strconv.Itoa(timeValue) + " " + unit + " timer done"
	desktopNotifier(titleString, reason)
}

func desktopNotifier(title string, subtitle string) {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}

	err = beeep.Alert(title, subtitle, "./assets/timer-sand-symbolic.svg")
	if err != nil {
		panic(err)
	}
}
