# tick

`tick` writes some string to output for `n` time units (called "duration") every
`m` time units (called "interval").

## Usage

```
tick duration [options]
  -duration=0: total duration to run
  -exit-code=0: the exit code when the run finishes
  -interval=1s: interval between each tick
  -silent=false: do not output any text
  -stderr=false: output to stderr
  -text=".": the text to output on each tick
```

A duration or interval string is a sequence of decimal numbers, each with
optional fraction and a unit suffix, such as "300ms", "1.5h" or "2h45m". Valid
time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

## Building

Simply run `go build`.
