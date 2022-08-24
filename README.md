# timer

## Summary

A command-line timer utility to improve productivity.

## Installation

```
go get -u github.com/gokpm/timer
go build
mv timer $HOME/bin
```

## Example

```
$ timer 10m
+5m
```

```
$ timer 10m
-3m
```

```
$ timer 5m
status
29m53.849311561s
```

```
$ timer 30m
exit
```
