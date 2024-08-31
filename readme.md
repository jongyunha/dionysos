# Dionysus

Dionysus is a Go library that enables the scheduling and management of Docker container executions using a cron-like mechanism. It allows users to configure Docker tasks through a YAML configuration file, specifying options such as the Docker image, tag, execution interval, and whether tasks should be executed concurrently.

## Features

- **Easy Configuration**: Manage Docker containers using a simple YAML configuration file.
- **Flexible Scheduling**: Schedule tasks with customizable intervals, including seconds, minutes, and hours.
- **Concurrency Control**: Run Docker containers concurrently or sequentially based on user preferences.
- **Docker API Integration**: Directly interact with Docker through its API, ensuring efficient and reliable container management.

## Installation

To use Dionysus in your project, first ensure you have Go installed and your project is set up. Then, clone this repository and integrate Dionysus into your project:

```bash
git clone https://github.com/yourusername/dionysus.git
```

### Dependencies

Dionysus relies on the following Go packages:

- `github.com/docker/docker/client`
- `gopkg.in/yaml.v2`

You can install these dependencies using `go get`:

```bash
go get github.com/docker/docker/client
go get gopkg.in/yaml.v2
```

## Configuration

Dionysus is configured through a `config.yaml` file. Below is an example configuration:

```yaml
docker:
image: "your-image-name"
tag: "latest"
interval: 5          # Interval for execution
interval_unit: "m"    # Time unit for interval ("s" for seconds, "m" for minutes, "h" for hours)
concurrent: false
timeout: 10           # Timeout for Docker container execution
timeout_unit: "s"     # Time unit for timeout ("s" for seconds, "m" for minutes, "h" for hours)
```

### Configuration Parameters

- **image**: The Docker image to be executed. Example: `"nginx"`.
- **tag**: The tag of the Docker image. Example: `"latest"`. This allows you to specify different versions of the image.
- **interval**: The interval at which the Docker container should be executed. Example: `5`.
- **interval_unit**: The time unit for the interval. This can be:
    - `"s"` for seconds
    - `"m"` for minutes
    - `"h"` for hours
- **concurrent**: Whether to allow concurrent execution of Docker containers. Set to `true` to allow multiple instances to run simultaneously, or `false` to ensure only one instance runs at a time.
- **timeout**: The timeout for Docker container execution.
- **timeout_unit**: The time unit for the timeout.

### Example Configuration

```yaml
docker:
image: "nginx"
tag: "alpine"
interval: 10
interval_unit: "m"
concurrent: true
timeout: 10
timeout_unit: "m"
```

In this example, Dionysus will pull and run the `nginx:alpine` Docker image every 10 minutes, allowing concurrent execution of multiple instances.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have any ideas, improvements, or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.