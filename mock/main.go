package main

import (
	countdown "mock/pkg"
	"os"
	"time"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleep := &DefaultSleeper{}
	countdown.Countdown(os.Stdout, sleep)
}
