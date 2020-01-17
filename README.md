# pnap-cli

## Setup

1. Get [`go`](https://golang.org/) and install
2. Install `make`

    `sudo apt-get install build-essential`

3. Clone this repository.
4. Go into the `pnap-cli` folder.
5. Move the `sample-config.yaml` file to `$HOME/.pnap/config.yaml` and add your client credentials (ID and secret)
6. Run `make build-simple` to build.

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
make build ENVIRONMENT_NAME=dev
```

## Running Tests

Usage:

* `make test` to run tests
* `make test-verbose` to run tests in verbose mode
* `make test-race` for race tests
* `make test-coverage` for test coverage (will output report.xml in test/coverage/).
* `make test PKG=./commands/create` to restrict test to a package

## Mocks

We are using [`mockgen`](https://github.com/golang/mock), the mock generation tool by `gomock`.

Note that *only interfaces can be mocked.* An example can be seen in [`Client`](./pnapctl/client/client.go), and [`Printer`](./pnapctl/printer/printer.go).

Mocks must be placed in the `tests/mocks` directory. They also need to have a package name of `mocks`. The following is an example command to generate the `Printer`'s mocks:

`make generate-mock MOCK_SOURCE=common/printer/printer.go MOCK_DESTINATION=tests/mocks/mock_printer.go`

## Configuration
Details can be passed using a config file. This file can be passed as an argument, as environment variable `PNAPCTL_HOME`, or can be read if placed in `~/.pnap/config.yaml`. An example of this file is in `sample-config.yaml`. In order to currently test the application, this `yaml` file can be used by using the following command: `pnapctl --config=sample-config.yaml ...` or simply copying/symlinking the file to your home directory.

## Current folder structure

Every command is its own folder, having a `.go` file that represents it. So, to check `pnapctl get servers`, the directory structure would be `./pnapctl/commands/get/servers`.
