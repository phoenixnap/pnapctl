<h1 align="center">
  <br>
  <a href="https://phoenixnap.com/bare-metal-cloud"><img src="https://user-images.githubusercontent.com/78744488/109779287-16da8600-7c06-11eb-81a1-97bf44983d33.png" alt="phoenixnap Bare Metal Cloud" width="300"></a>
  <br>
  Bare Metal Cloud CLI
  <br>
</h1>

<p align="center">
This CLI allows you to interact with Bare Metal Cloud APIs to deploy new and manage existing servers directly from the terminal.
</p>

<p align="center">
  <a href="https://phoenixnap.com/bare-metal-cloud">Bare Metal Cloud</a> •
  <a href="https://developers.phoenixnap.com/apis">API</a> •
  <a href="https://developers.phoenixnap.com/">Developers Portal</a> •
  <a href="http://phoenixnap.com/kb">Knowledge Base</a> •
  <a href="https://developers.phoenixnap.com/support">Support</a>
</p>

## Requirements

- [Bare Metal Cloud](https://bmc.phoenixnap.com) account
- [Go](https://golang.org/dl/)

## Creating a Bare Metal Cloud account

1. Go to the [Bare Metal Cloud signup page](https://support.phoenixnap.com/wap-jpost3/bmcSignup).
2. Follow the prompts to set up your account.
3. Use your credentials to [log in to Bare Metal Cloud portal](https://bmc.phoenixnap.com).

:arrow_forward: **Video tutorial:** [How to Create a Bare Metal Cloud Account](https://www.youtube.com/watch?v=RLRQOisEB-k)
<br>

:arrow_forward: **Video tutorial:** [Introduction to Bare Metal Cloud](https://www.youtube.com/watch?v=8TLsqgLDMN4)

## CTL Installation

The CLI can be either used manually or as part of automation scripts.

You can use pnapctl on Linux, OS X, and Windows-based AMD64 systems. The binary is available for download through the following links:

* Linux: [pnapctl-linux-amd64.tar.gz](https://github.com/phoenixnap/pnapctl/releases/latest/download/pnapctl-linux-amd64.tar.gz)
* OS X: [pnapctl-darwin-amd64.tar.gz](https://github.com/phoenixnap/pnapctl/releases/latest/download/pnapctl-darwin-amd64.tar.gz)
* Windows: [pnapctl-windows-amd64.zip](https://github.com/phoenixnap/pnapctl/releases/latest/download/pnapctl-windows-amd64.zip)

:open_book: Detailed steps on how to install CLI are available on our developers portal: [https://developers.phoenixnap.com/cli](https://developers.phoenixnap.com/cli)

## CTL Setup Steps

1. Get [`go`](https://golang.org/) and install.
2. Install `make`<br> using the following command: 
    `sudo apt-get install build-essential`
3. Clone this repository.
4. Go into the `pnapctl` folder.
5. Move the `sample-config.yaml` file to `$HOME/.pnap/config.yaml` and add your client credentials (ID and secret)
6. Run `make build-simple` to build.

The executable will be generated in the `bin` folder. This is an example of command execution:

   `./bin/pnapctl get servers`

## Commands

You can view all the available commands with different options for each command on our [GitHub page](https://github.com/phoenixnap/pnapctl/blob/latest/docs/pnapctl.md).

These commands enable you to create, delete, modify or shut down a resource, as well as to perform actions such as submit a modification request, print version, reset, tag or update a resource. For a better understanding of what each action does, please consult the [API documentation](https://developers.phoenixnap.com/cli).

## Multi OS Build

We are using `gox` for the multi OS build. Note that unless otherwise specified, `gox` will build the pnapctl against an unnecessarily long list of OSs so it is suggested to define the desired OS architectures.

* `make build` -> build a version for all supported OS architectures.
* `make build BUILD_PLATFORMS="linux/amd64 windows/amd64"` -> build a version for linux and windows, 64 bit.

## Running Tests

Usage:

* `make test` to run tests
* `make test-verbose` to run tests in verbose mode
* `make test-race` for race tests
* `make test-coverage` for test coverage (will output report.xml in out/unit-tests/).
* `make test-coverage-show` for showing a GUI with coverage information (will output `cover.out`)
* `make test-tparse` to run tests in a pretty format (requires `tparse` to be installed)
* `make test PKG=./commands/create` to restrict test to a package

### Running Tests with TParse

TParse is a command line tool used to summarize a Go test output. It is also useful when analysing test coverage. 

#### Installation and Usage: 
1. Verify that the bin directory for GO is included in `$PATH`
2. Install [`tparse`](https://github.com/mfridman/tparse) by running `go install github.com/mfridman/tparse@latest`
3. Use `go test -json -cover ./... | tparse -all` to run all tests or `go test -json -cover ./commands/get | tparse -all` to run specific tests. 

:grey_exclamation: **NOTE:** You can also run `make test-tparse` as a shortcut, which also works with `PKG`.

### Showing test coverage information

TParse can show you the total coverage%, but there's a way to find the exact lines being covered.

1. Run `go test ./... -coverprofile cover.out`. This runs all tests and outputs coverage information into `cover.out`.
2. Run `go tool cover -html=cover.out`. This will use the `cover.out` you just generated to display coverage information in your browser.

:grey_exclamation: **NOTE:** You can also run `make test-coverage-show`, which will run both of the mentioned commands. It works with `PKG` as well.

## Running Component Tests

Component tests are executed using [`bats`](https://bats-core.readthedocs.io/en/stable/). To install bats using `npm` run:

```
$ npm install -g bats
```

Usage:
* `make component-tests` to fetch and verify versions of required libraries and run the component tests (will output junit report in out/component-tests/).

## Debugging 

Our preferred IDE for developemnt in Go is VSCode. To debug Go we make use of [Delve](https://github.com/go-delve/delve), which is a debugger for the Go programming language. 

Setup: 
1. Clone and install Delve
  ```
  $ git clone https://github.com/go-delve/delve
  $ cd delve
  $ go install github.com/go-delve/delve/cmd/dlv

  ```
2. In VSCode, create a new debug configuration 
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
            "program": "<reaplace_with_path_to_your_workspace>/pnapctl/src",
            "env": {},
            "args": ["get", "servers"] // replace args accordingly             
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

## Bare Metal Cloud community

Become part of the Bare Metal Cloud community to get updates on new features, help us improve the platform, and engage with developers and other users.

- Follow [@phoenixNAP on Twitter](https://twitter.com/phoenixnap)
- Join the [official Slack channel](https://phoenixnap.slack.com)
- Sign up for our [Developers Monthly newsletter](https://phoenixnap.com/developers-monthly-newsletter)

### Resources

- [Product page](https://phoenixnap.com/bare-metal-cloud)
- [Instance pricing](https://phoenixnap.com/bare-metal-cloud/instances)
- [YouTube tutorials](https://www.youtube.com/watch?v=8TLsqgLDMN4&list=PLWcrQnFWd54WwkHM0oPpR1BrAhxlsy1Rc&ab_channel=PhoenixNAPGlobalITServices)
- [Developers Portal](https://developers.phoenixnap.com)
- [Knowledge Base](https://phoenixnap.com/kb)
- [Blog](https:/phoenixnap.com/blog)

### Documentation

- [API documentation](https://developers.phoenixnap.com/apis)

### Contact phoenixNAP

Get in touch with us if you have questions or need help with Bare Metal Cloud.

<p align="left">
  <a href="https://twitter.com/phoenixNAP">Twitter</a> •
  <a href="https://www.facebook.com/phoenixnap">Facebook</a> •
  <a href="https://www.linkedin.com/company/phoenix-nap">LinkedIn</a> •
  <a href="https://www.instagram.com/phoenixnap">Instagram</a> •
  <a href="https://www.youtube.com/user/PhoenixNAPdatacenter">YouTube</a> •
  <a href="https://developers.phoenixnap.com/support">Email</a> 
</p>

<p align="center">
  <br>
  <a href="https://phoenixnap.com/bare-metal-cloud"><img src="https://user-images.githubusercontent.com/81640346/115243282-0c773b80-a123-11eb-9de7-59e3934a5712.jpg" alt="phoenixnap Bare Metal Cloud"></a>
</p>
