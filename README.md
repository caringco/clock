# clock

*Clock* is a golang library having features such as static wall clock

Docs are present at [link](https://godoc.org/github.com/prakashpandey/clock)

## Features included

- **Wall Clock**: wallclock is static clock

## Future features

- **Stop Watch** : see branch [dev-stopWatch-impl](https://github.com/prakashpandey/clock/tree/dev-stopWatch-impl) for implementation

- **Alarm**: sends ticks on the given go-channel at a given time. Implementation branch [dev-alarmClock-impl](https://github.com/prakashpandey/clock/tree/dev-alarmClock-impl)

## Example

```
    // create new wallclock instance
    wc := NewWallClock()

    // move wallclock at the given time
    wc.MoveClockAt(time.Now())

    days := 7
    // move clock by years, days, hours, min, sec
    wc.MoveClockBy(0, days, 0, 0, 0)

    // time api's
    
    // GetTime is the current time of wallclock
    t := wc.GetTime()

    // GetUnixNano is the current time of wallclock returns as a Unix time
    n := wc.GetUnixNano()
    
    // Now is the current time of the system
    t := wc.Now()
    
```

## Test coverage

```
    $go test -coverprofile cover.out
    
    PASS
    coverage: 75.0% of statements
    ok      github.com/prakashpandey/clock  0.004s
```