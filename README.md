# pnap-cli

## Setup

1. Get [`go`](https://golang.org/) and install
2. Clone this repository in `$GOPATH/src/phoenixnap.com`. *(`$GOPATH` is normally `~/go/src`)*
3. Go into the `pnap-cli` folder.
4. Run `go get` to get all dependencies.
5. Get [`gomock`](https://github.com/golang/mock)
6. Move/symlink the `sample-config.yaml` file to `$HOME/pnap.yaml`
7. Run `go build` to build, or `go run main.go` to run.

The executable generated will have the same name as the folder. This means that `go build` in this repository with its default name will produce an executable called `pnap-cli`. To change its name, use `go build -o <name>`

**Note:** To run all tests, run `go test ./tests/...`. If you'd like a colourised output *(success as green, fail as red)* get `gotest` using `go get -u github.com/rakyll/gotest` and run `gotest ./tests/...` instead.

## Mocks

To generate mocks, get [`mockgen`](https://github.com/golang/mock), the mock generation tool by `gomock`.

Note that *only interfaces can be mocked.* An example can be seen in [`Client`](./pnapctl/client/client.go), and [`Printer`](./pnapctl/printer/printer.go).

Mocks must be placed in the `pnapctl/mocks` directory. They also need to have a package name of `mocks`. The following is an example command to generate the `Printer`'s mocks:

`mockgen --source=pnapctl/printer/printer.go --destination=pnapctl/mocks/mock_printer.go --package=mocks`

## Configuration
Details can be passed using a config file. This file can be passed as an argument, or can be read if placed in `~/pnap.yaml`. An example of this file is in `sample-config.yaml`. In order to currently test the application, this `yaml` file can be used by using the following command: `pnapctl bmc --config=sample-config.yaml ...` or simply copying/symlinking the file to your home directory.

## Current folder structure

Every command is its own folder, having a `.go` file that represents it. So, to check `pnapctl bmc get servers`, the directory structure would be `./pnapctl/bmc/get/servers`.


