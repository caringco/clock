# clock

*Clock* is a golang library having features such as static wall clock

Docs are present at [link](https://godoc.org/github.com/prakashpandey/clock)

## Features included

- **Wall Clock**: wallclock is static clock

## Future features

- **Stop Watch** : see branch [dev-stopWatch-impl](https://github.com/prakashpandey/clock/tree/dev-stopWatch-impl) for implementation

- **Alarm**: sends ticks on the given go-channel at a given time.  

        see branch [dev-stopWatch-impl](https://github.com/prakashpandey/clock/tree/dev-stopWatch-impl) for implementation

## Test coverage

```
    $go test -coverprofile cover.out
    
    PASS
    coverage: 78.6% of statements
    ok  	github.com/prakashpandey/clock	0.002s
```

## Example

```
    // create new wallclock instance
    wc := NewWallClock()

    // move wallclock at the given time
    wc.MoveClockAt(time.Now())

    days := 7
    // move clock by years, days, hours, min, sec
    wc2.MoveClockBy(0, days, 0, 0, 0)
```