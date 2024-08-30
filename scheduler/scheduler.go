package scheduler

import (
	"context"
	"dionysos/docker"
	"fmt"
	"sync"
	"time"
)

func ScheduleTask(dockerClient *docker.DockerClient, image string, tag string, interval int, intervalUnit string, allowConcurrency bool) {
	var duration time.Duration

	// Determine the duration based on the interval unit
	switch intervalUnit {
	case "s":
		duration = time.Duration(interval) * time.Second
	case "m":
		duration = time.Duration(interval) * time.Minute
	case "h":
		duration = time.Duration(interval) * time.Hour
	default:
		fmt.Println("Invalid interval unit. Use 's' for seconds, 'm' for minutes, or 'h' for hours.")
		return
	}

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	var wg sync.WaitGroup

	for {
		select {
		case <-ticker.C:
			if allowConcurrency {
				wg.Add(1)
				go func() {
					defer wg.Done()
					if err := dockerClient.ExecuteDocker(context.Background(), image, tag); err != nil {
						fmt.Printf("Error executing Docker container: %v\n", err)
					}
				}()
			} else {
				if err := dockerClient.ExecuteDocker(context.Background(), image, tag); err != nil {
					fmt.Printf("Error executing Docker container: %v\n", err)
				}
			}
		}
	}
}
