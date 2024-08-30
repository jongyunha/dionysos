package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"io"
	"os"
)

// DockerClient wraps the Docker API client
type DockerClient struct {
	cli *client.Client
}

// NewDockerClient initializes a new Docker client
func NewDockerClient() (*DockerClient, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("error creating Docker client: %v", err)
	}
	return &DockerClient{cli: cli}, nil
}

func (d *DockerClient) ExecuteDocker(ctx context.Context, dockerImage string, tag string) error {
	// Pull the Docker image if not available locally
	out, err := d.cli.ImagePull(ctx, fmt.Sprintf("%s:%s", dockerImage, tag), image.PullOptions{})
	if err != nil {
		return fmt.Errorf("error pulling Docker image: %v", err)
	}
	defer func(out io.ReadCloser) {
		_ = out.Close()
	}(out)

	_, _ = io.Copy(os.Stdout, out)

	// Create the Docker container
	resp, err := d.cli.ContainerCreate(ctx, &container.Config{
		Image: fmt.Sprintf("%s:%s", dockerImage, tag),
	}, nil, nil, nil, "")
	if err != nil {
		return fmt.Errorf("error creating Docker container: %v", err)
	}

	// Start the Docker container
	if err := d.cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return fmt.Errorf("error starting Docker container: %v", err)
	}

	// Wait for the container to finish and get logs
	statusCh, errCh := d.cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("error waiting for Docker container: %v", err)
		}
	case <-statusCh:
	}

	logs, err := d.cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true})
	if err != nil {
		return fmt.Errorf("error getting logs from Docker container: %v", err)
	}
	defer func(logs io.ReadCloser) {
		_ = logs.Close()
	}(logs)

	_, _ = io.Copy(os.Stdout, logs)

	// Remove the container after execution
	err = d.cli.ContainerRemove(ctx, resp.ID, container.RemoveOptions{})
	if err != nil {
		return fmt.Errorf("error removing Docker container: %v", err)
	}

	return nil
}
