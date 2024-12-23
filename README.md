# Noop Image

A minimal Docker image designed to print a message and exit. This image is useful for overriding images in development versions of services, acting as a no-op (no operation).

The message printed is:

```
noop image used, exiting. (see https://github.com/rjocoleman/noop for more info)
```

To better illustrate the practical application and benefit of using the Noop Docker image in development environments, the README can include an example showing how to effectively use Docker Compose to merge and override service definitions. This approach allows you to replace specific services, such as the Datadog Agent, with the Noop image in a development context without altering the primary `docker-compose.yml` configuration. Below is how you can incorporate this into the "Why Use the Noop Image?" section of your README:

## Why Use the Noop Image?

The Noop Docker image serves as a minimal placeholder (around 2mb), perfect for scenarios where you want to override service images in configurations such as Docker Compose, without the overhead of running actual services. This approach is particularly useful in development environments where certain services (e.g., monitoring agents) are not necessary, thus saving resources and simplifying configurations.

### Example: Overriding Datadog Agent in Development

Consider a scenario where your application's production Docker setup includes a Datadog Agent for monitoring. In development, you might not need the Datadog Agent running, but you still want to keep your Docker Compose setup consistent. This is where the Noop image comes in handy.

You have a `docker-compose.yml` defining all your services, including Datadog Agent:

```yaml
# docker-compose.yml
version: "3.8"
services:
  app:
    image: myapp:latest
    ...
  datadog-agent:
    image: datadog/agent:latest
    ...
```

For development, you create a `docker-compose.dev.yml` file that overrides the Datadog Agent service to use the Noop image:

```yaml
# docker-compose.dev.yml
version: "3.8"
services:
  datadog-agent:
    image: rjocoleman/noop:latest
```

To merge these configurations and apply the override for development, use the following Docker Compose command:

```bash
docker compose -f docker-compose.yml -f docker-compose.dev.yml up
```

This command combines the two Compose files, effectively replacing the Datadog Agent service with the Noop image in your development environment. Note: the `convert` command prints the combined configuration to the console, letting you verify that the override has been applied correctly.

The merged compose file would look something like this:

```yaml
version: "3.8"
services:
  app:
    command: "true"
    image: alpine:latest

  datadog-agent:
    image: rjocoleman/noop:latest
```

Using the Noop image in this manner ensures that your application can run without unnecessary services in development, while maintaining a clean and consistent setup across different environments.

For more information about multiple Docker Compose files (merging and overriding): https://docs.docker.com/compose/compose-file/13-merge/


## Using the Noop Image

### Docker Run

To run the `noop` image directly with Docker:

```bash
docker run --rm rjocoleman/noop:latest
```

### Docker Compose

To use the `noop` image with Docker Compose, include it in your `docker-compose.yml` file like this:

```yaml
version: "3.8"
services:
  noop-service:
    image: rjocoleman/noop:latest
```

Then, run the service using:

```bash
docker-compose up
```

## Silent output

To run the noop image without any message printed to stdout (i.e. suppress `noop image used, exiting. (see https://github.com/rjocoleman/noop for more info)`) you can set the environment variable `SILENCE_OUTPUT=true`. Use this wisely as it can be confusing why services aren't running.

## Service Dependencies and Health Checks

For services that need to wait for a noop service with health checks (e.g., using `depends_on: condition: service_healthy`), the container needs to stay alive. Set `NOOP_INFINITY=true` to make the container stay running indefinitely (similar to `sleep infinity`):

```yaml
services:
  noop-service:
    image: rjocoleman/noop:latest
    environment:
      - NOOP_INFINITY=true
      - SILENCE_OUTPUT=true
    healthcheck:
      test: ["/noop"]
      interval: 1s
      timeout: 1s
      retries: 1
      start_period: 0s

  dependent-service:
    depends_on:
      noop-service:
        condition: service_healthy

## Building From Source

To build the `noop` image from source, clone the repository and use the provided Dockerfile.

```bash
git clone https://github.com/rjocoleman/noop.git
cd noop
docker build -t your-tag-name .
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
