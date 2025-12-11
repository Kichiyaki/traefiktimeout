# [TimeoutHandler](https://pkg.go.dev/net/http#TimeoutHandler) for Traefik

This middleware allows you to specify a timeout for each router
separately. [Traefik's built-in timeouts](https://doc.traefik.io/traefik/reference/install-configuration/entrypoints/#opt-transport-respondingTimeouts-readTimeout)
are specified for an entrypoint and there is no way to override them.

# Installation

1. Enable the plugin in your Traefik static configuration:
    ```yaml
    # traefik.yml
    experimental:
      plugins:
        traefiktimeout:
          moduleName: github.com/Kichiyaki/traefiktimeout
          version: v0.1.1  # Use the latest version
    ```
2. Configure the middleware in your dynamic configuration.

## Configuration options

| Parameter   | Is required | Description                                                                            | Default value |
|-------------|-------------|----------------------------------------------------------------------------------------|---------------|
| **timeout** | **YES**     | Time limit for requests in [Go duration format](https://pkg.go.dev/time#ParseDuration) |               |
| **message** |             | A message that will be sent to a user after a timeout                                  | Timeout       |

