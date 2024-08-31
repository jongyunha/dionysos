package main

import (
	"dionysos/config"
	"dionysos/docker"
	"dionysos/scheduler"
	"fmt"
)

func main() {
	// Load the configuration from the YAML file
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		return
	}

	// Create a Docker client
	dockerClient, err := docker.NewDockerClient()
	if err != nil {
		fmt.Printf("Error creating Docker client: %v\n", err)
		return
	}

	// Schedule the Docker task using the configuration
	scheduler.ScheduleTask(dockerClient, cfg)
}
