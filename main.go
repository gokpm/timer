package main

import (
	"fmt"
	"os"
	"time"
)

var stdin = make(chan string, 1)

func main() {
	/*
		Validate the arguments to the program.
	*/
	if len(os.Args) < 2 {
		return
	}
	/*
		Parse the duration entered by the user.
	*/
	duration, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		Start the stop watch.
	*/
	deadLine := time.Now().Add(duration)
	/*
		Read stdin stream.
	*/
	go Scan()
	go func() {
		for input := range stdin {
			if len(input) <= 0 {
				continue
			}
			switch input[0] {
			case '+':
				if len(input) < 2 {
					continue
				}
				duration, err := time.ParseDuration(input[1:])
				if err != nil {
					fmt.Println(err)
					continue
				}
				deadLine = deadLine.Add(duration)
			case '-':
				if len(input) < 2 {
					continue
				}
				duration, err := time.ParseDuration(input[1:])
				if err != nil {
					fmt.Println(err)
					continue
				}
				deadLine = deadLine.Add(-duration)
			default:
				switch input {
				case "status":
					fmt.Println(time.Until(deadLine).String())
				case "exit":
					deadLine = time.Now()
				default:
					err = ErrInvalidInput
					fmt.Println(err)
					continue
				}
			}
		}
	}()
	for time.Now().Before(deadLine) {
	}
	fmt.Println("Time up!")
	_, err = Exec("notify-send", "-i", "bash", "-u", "critical", "timer", "Done!")
	if err != nil {
		fmt.Println(err)
		return
	}
}
