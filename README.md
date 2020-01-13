# pnap-cli

## Setup

1. Get [`go`](https://golang.org/) and install
2. Install `make`

    `sudo apt-get install build-essential`

3. Clone this repository.
4. Go into the `pnap-cli` folder.
5. Install [`gomock`](https://github.com/golang/mock):
    - go get github.com/golang/mock/mockgen
    - go get github.com/golang/mock/gomock
    - go install github.com/golang/mock/mockgen
6. Move the `sample-config.yaml` file to `$HOME/.pnap/config.yaml` and add your client credentials (ID and secret)
7. Run `make build-simple` to build.

The executable will be generated in `bin` folder. This is an example of command execution:

   `./bin/pnapctl get servers`

## Multi OS Build

We are using `gox` for multi os build. Note that unless otherwise specified `gox` will build the pnapctl against an unnecessarily long list of OS's so it is suggested to define the desire OS architectures.

* `make build` -> build a version for all supported OS architectures.
* `make build BUILD_PLATFORMS="linux/amd64 windows/amd64"` -> build a version for linux and windows, 64 bit.

## Multi Environment Build

We are building two versions of `pnapctl` executable. Default one for `prod` environment and one for `dev` environment. They are uploaded to `prod/dev` environment of Apigee portal. 
You can change `dev` environment specific properties by editing `./pnapctl/configuration/properties_dev.go` file. By
default build is executed for `dev` environment but you can specify different environment by overriding build variable `ENVIRONMENT_NAME`, e.g:

```
make build-simple ENVIRONMENT_NAME=prod
```

or

```
make build ENVIRONMENT_NAME=prod
```

## Running Tests

If you'd like a colourised output *(success as green, fail as red)* get `gotest` using `go get -u github.com/rakyll/gotest` and run `gotest ./tests/...` instead.

* `go test ./tests/...` -> run all tests
* `go test ./tests/... -v` -> run all tests with verbose output
* `go test ./tests/create_test.go` -> run all tests in `create_test.go` file
* `go test ./tests/create_test.go -run TestCreateServerSuccessYAML` -> run all tests in`create_test.go` that match the regex `TestCreateServerSuccessYAML`


## Mocks

To generate mocks, get [`mockgen`](https://github.com/golang/mock), the mock generation tool by `gomock`.

Note that *only interfaces can be mocked.* An example can be seen in [`Client`](./pnapctl/client/client.go), and [`Printer`](./pnapctl/printer/printer.go).

Mocks must be placed in the `pnapctl/mocks` directory. They also need to have a package name of `mocks`. The following is an example command to generate the `Printer`'s mocks:

`mockgen --source=pnapctl/printer/printer.go --destination=pnapctl/mocks/mock_printer.go --package=mocks`

## Configuration
Details can be passed using a config file. This file can be passed as an argument, as environment variable `PNAPCTL_HOME`, or can be read if placed in `~/.pnap/config.yaml`. An example of this file is in `sample-config.yaml`. In order to currently test the application, this `yaml` file can be used by using the following command: `pnapctl --config=sample-config.yaml ...` or simply copying/symlinking the file to your home directory.

## Current folder structure

Every command is its own folder, having a `.go` file that represents it. So, to check `pnapctl get servers`, the directory structure would be `./pnapctl/commands/get/servers`.
