# controller-model

[![Integration Test](https://github.com/instill-ai/controller-model/actions/workflows/integration-test.yml/badge.svg)](https://github.com/instill-ai/controller-model/actions/workflows/integration-test.yml)

`controller-model` service monitors the state of other services and resources within Instill Model.

## Local dev

On the local machine, clone `model` repository in your workspace, move to the repository folder, and launch all dependent microservices:
```bash
$ cd <your-workspace>
$ git clone https://github.com/instill-ai/model.git
$ cd model
$ make latest PROFILE=controller
```

Clone `controller-model` repository in your workspace and move to the repository folder:
```bash
$ cd <your-workspace>
$ git clone https://github.com/instill-ai/controller-model.git
$ cd controller-model
```

### Build the dev image

```bash
$ make build
```

### Run the dev container

```bash
$ make dev
```

Now, you have the Go project set up in the container, in which you can compile and run the binaries together with the integration test in each container shell.

### Run the server

```bash
$ docker exec -it controller-model /bin/bash
$ go run ./cmd/main
```

### Run the integration test

```bash
$ docker exec -it controller-model /bin/bash
$ make integration-test
```

### Stop the dev container

```bash
$ make stop
```

### CI/CD

The latest images will be published to Docker Hub [repository](https://hub.docker.com/r/instill/controller-model) at release.

## Contributing

Please refer to the [Contributing Guidelines](https://github.com/instill-ai/model/blob/main/.github/CONTRIBUTING.md) for more details.


## License

See the [LICENSE](./LICENSE) file for licensing information.
