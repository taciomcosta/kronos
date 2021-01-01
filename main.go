package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s executable-path seconds", os.Args[0])
		return
	}

	program := os.Args[1]
	seconds, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("error on parsing ticker: %v\n", err)
	}

	for {
		cmd := exec.Command(program)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Running program error: %v\n", err)
			os.Exit(1)
		}
		time.Sleep(time.Duration(seconds) * time.Second)
	}
}
