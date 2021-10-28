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

### Running Tests with TParse

TParse is a command line tool used to summarise a go test output. It is also useful when analysing test coverage. 

Installation and Usage: 
1. Verify that the bin directory for GO is included in `$PATH`
2. Install [`tparse`](https://github.com/mfridman/tparse) by running `go install github.com/mfridman/tparse@latest`
3. Use `go test -json -cover ./... | tparse -all` to run all tests or `go test -json -cover ./commands/get | tparse -all` to run specific tests. 

## Debugging 

Our preferred IDE for developemnt in GO is VS Code. To debug GO we make use of a [Delve](https://github.com/go-delve/delve), which is a debugger for the Go programming language. 

Setup: 
1. Clone and install Delve
  ```
  $ git clone https://github.com/go-delve/delve
  $ cd delve
  $ go install github.com/go-delve/delve/cmd/dlv

  ```
2. In VS Code, create a new debug configuration 
  * Select the `Run and Debug` button from the Run view or hit the `F5` button to start the debugging mode. 
  * From the drop down menu next to `Run and Debug` select `Add Configuration...`. A new configuration `.vscode/launch.json` will be created.
  * Paste the following configuration in `launch.json`:
  ```
  {
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "<reaplace_with_path_to_your_workspace>/pnap-cli/",
            "env": {},
            "args": ["get", "servers"], // replace args accordingly 
            "buildFlags": "-tags dev"            
        }
    ]
  }
  ```


## Mocks

We are using [`mockgen`](https://github.com/golang/mock), the mock generation tool by `gomock`.

Note that *only interfaces can be mocked.* An example can be seen in [`Client`](./pnapctl/client/client.go), and [`Printer`](./pnapctl/printer/printer.go).

Mocks must be placed in the `tests/mocks` directory. They also need to have a package name of `mocks`. The following is an example command to generate the `Printer`'s mocks:

`make generate-mock MOCK_SOURCE=common/printer/printer.go MOCK_DESTINATION=tests/mocks/mock_printer.go`

## Configuration
Details can be passed using a config file. This file can be passed as an argument, as environment variable `PNAPCTL_HOME`, or can be read if placed in `~/.pnap/config.yaml`. An example of this file is in `sample-config.yaml`. In order to currently test the application, this `yaml` file can be used by using the following command: `pnapctl --config=sample-config.yaml ...` or simply copying/symlinking the file to your home directory.

## Current folder structure

Every command is its own folder, having a `.go` file that represents it. So, to check `pnapctl get servers`, the directory structure would be `./pnapctl/commands/get/servers`.

# Local SDK

While there is no public version of the SDK - a local one can be used to facilitate development. The following are instructions to set up the local SDK for an example sub-folder (`bmc-api`) in the SDK.

1. Create a new folder: `~/go/src/phoenixnap.com`
2. Put `pnap-cli` in above folder.
3. Put SDK subfolders (like `bmc-api`) in above folder.
4. Change SDK modules (in `go.mod`) to `phoenixnap.com/bmc-api`
5. Add following lines to pnap-cli's `go.mod`.

```
require phoenixnap.com/bmc-api v1.0.0
replace phoenixnap.com/bmc-api v1.0.0 => "../bmc-api"
```

6) Where used, add this import:
```
bmcApi "phoenixnap.com/bmc-api"
```

For any other sub-folder (`bmc-billing`, etc.) the above instructions work - simply replace `bmc-api` with the name of your sub-folder.

# Gitlab SDK

The private version of the SDK on Gitlab can be used as well, however for this to work locally there needs to be some setup.

1. Create a new gitlab access token.
2. Add the following into your `.gitconfig` file (it should be in your home directory):
```
[url "https://oauth2:{YOUR ACCESS TOKEN}@gitlab.com"]
	insteadOf = https://gitlab.com
```
3. Run the following command: `go env -w GOPRIVATE=gitlab.com/*`
  - This is so that the `go` compiler knows what dependencies are *private* - since it runs additional checks for public dependencies that won't work for private.