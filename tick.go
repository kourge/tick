package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

const (
	durationUsage = "total duration to run"
	intervalUsage = "interval between each tick"
	textUsage     = "the text to output on each tick"
	exitCodeUsage = "the exit code when the run finishes"
	silentUsage   = "do not output any text"
	stderrUsage   = "output to stderr"
)

var duration time.Duration
var interval time.Duration
var text string
var exitCode int
var silent bool

var shouldUseStderr bool
var out io.Writer

func init() {
	flag.DurationVar(&duration, "duration", 0*time.Second, durationUsage)
	flag.DurationVar(&interval, "interval", 1*time.Second, intervalUsage)
	flag.StringVar(&text, "text", ".", textUsage)
	flag.IntVar(&exitCode, "exit-code", 0, exitCodeUsage)
	flag.BoolVar(&silent, "silent", false, silentUsage)
	flag.BoolVar(&shouldUseStderr, "stderr", false, stderrUsage)
}

func usage() {
	fmt.Println("usage: tick duration [options]")
	flag.PrintDefaults()
	fmt.Println(`A duration or interval string is a sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".`)
	os.Exit(0)
}

func isDurationSpecified() bool {
	isSpecified := false
	flag.Visit(func(flag *flag.Flag) {
		if flag.Name == "duration" {
			isSpecified = true
		}
	})
	return isSpecified
}

func write(text string) {
	if _, err := io.WriteString(out, text); err != nil {
		panic(err.Error())
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		usage()
	}

	if !isDurationSpecified() {
		if dur, err := time.ParseDuration(flag.Arg(0)); err != nil {
			panic(err.Error())
		} else {
			duration = dur
			args = args[1:]
		}
	}

	switch {
	case silent:
		out = ioutil.Discard
	case shouldUseStderr:
		out = os.Stderr
	default:
		out = os.Stdout
	}

	end := time.After(duration)
	tick := time.Tick(interval)

	for {
		select {
		case <-tick:
			write(text)
		case <-end:
			os.Exit(exitCode)
		}
	}
}
