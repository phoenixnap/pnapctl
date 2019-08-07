# pnap-cli

To generate mocks, get [`mockgen`](https://github.com/golang/mock). The usual command to write is as follows: `mockgen --source=path/to/source.go --destination path/to/mock.go --package=mocks`

Clone this repository in `~/go/src/phoenixnap.com/`

To test everything, run `go test ./tests/...`. To colorize the output, download `gotest` from `go get -u github.com/rakyll/gotest`.

## Current folder structure

Every command is its own folder, having a `.go` file that represents it. So, to check `pnapctl bmc get servers`, the directory structure would be `./pnapctl/bmc/get/servers`.