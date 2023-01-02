package main

import (
	"mocking"
	"os"
	"time"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 + time.Second, SleepFunction: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
