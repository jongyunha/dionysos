package scheduler

import (
	"context"
	"dionysos/config"
	"dionysos/docker"
	"errors"
	"fmt"
	"sync"
	"time"
)

func ScheduleTask(dockerClient *docker.Client, config *config.Config) {
	var intervalDuration, timeoutDuration time.Duration

	// Determine the interval duration based on the interval unit
	switch config.Docker.IntervalUnit {
	case "s":
		intervalDuration = time.Duration(config.Docker.Interval) * time.Second
	case "m":
		intervalDuration = time.Duration(config.Docker.Interval) * time.Minute
	case "h":
		intervalDuration = time.Duration(config.Docker.Interval) * time.Hour
	default:
		fmt.Println("Invalid interval unit. Use 's' for seconds, 'm' for minutes, or 'h' for hours.")
		return
	}

	// Determine the timeout duration based on the timeout unit
	switch config.Docker.TimeoutUnit {
	case "s":
		timeoutDuration = time.Duration(config.Docker.Timeout) * time.Second
	case "m":
		timeoutDuration = time.Duration(config.Docker.Timeout) * time.Minute
	case "h":
		timeoutDuration = time.Duration(config.Docker.Timeout) * time.Hour
	default:
		fmt.Println("Invalid timeout unit. Use 's' for seconds, 'm' for minutes, or 'h' for hours.")
		return
	}

	ticker := time.NewTicker(intervalDuration)
	defer ticker.Stop()

	var wg sync.WaitGroup

	for range ticker.C {
		deadline := time.Now().Add(timeoutDuration)

		if config.Docker.Concurrent {
			wg.Add(1)
			go func() {
				defer wg.Done()
				executeWithDeadline(dockerClient, config, deadline)
			}()
		} else {
			executeWithDeadline(dockerClient, config, deadline)
		}
	}
}

func executeWithDeadline(dockerClient *docker.Client, config *config.Config, deadline time.Time) {
	// Create a context with deadline
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	if err := dockerClient.ExecuteDocker(ctx, config.Docker.Image, config.Docker.Tag); err != nil {
		// Check if the error was caused by context deadline
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			fmt.Printf("Error: Docker container execution exceeded deadline: %v\n", err)
		} else {
			fmt.Printf("Error executing Docker container: %v\n", err)
		}
	}
}
