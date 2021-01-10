package domain

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/taciomcosta/kronos/internal/domain/ticker"
)

type Job struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Tick    string `json:"tick"`
	ticker  ticker.Ticker
}

func NewJob(name string, command string, tick string) (Job, error) {
	ticker, err := ticker.NewTicker(tick)
	if err != nil {
		return Job{}, err
	}
	return Job{name, command, tick, ticker}, nil
}

func (j *Job) Run(t time.Time) {
	if !j.ticker.IsTimeSet(t) {
		return
	}
	fmt.Printf("Executing %s\n", j.Name)
	j.ExecCommand()
}

func (j *Job) ExecCommand() error {
	cmd := exec.Command(j.Command)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
