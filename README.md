# TOPO

Custom Golang framework to build web applications fast and easy.

### Install TOPO cli

    ```go install "github.com/ManuelGarreRos/topo@latest"

### Create a new project

    ```topo new

### Run the project

    ```topo start

### Ports

- 8000: Web server
- 40000: Delve debugger
- 5432: Postgres

### Requirements

- Go 1.18
- Docker
- Docker Compose

### Features

- Full Docker support
- Hot reload using [air](https://github.com/cosmtrek/air)
- Debugging using [delve](https://github.com/go-delve/delve)