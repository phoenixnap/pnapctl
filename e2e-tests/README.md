# Using the `pnapctl` Docker Image

[![Release](https://img.shields.io/github/v/release/phoenixnap/pnapctl)](https://github.com/phoenixnap/pnapctl/releases)
[![Container Image](https://img.shields.io/badge/GHCR-pnapctl-blue)](https://github.com/phoenixnap/pnapctl/pkgs/container/pnapctl)

Run `pnapctl` without installing the binary locally by using the published Docker image.

**Resources**

* Container Image: [`ghcr.io/phoenixnap/pnapctl`](https://github.com/phoenixnap/pnapctl/pkgs/container/pnapctl)
* Documentation: [`github.com/phoenixnap/pnapctl/blob/master/docs/pnapctl.md`](https://github.com/phoenixnap/pnapctl/blob/master/docs/pnapctl.md)
* Source Code: [`github.com/phoenixnap/pnapctl`](https://github.com/phoenixnap/pnapctl)

## 1. Pull the image

```bash
docker pull ghcr.io/phoenixnap/pnapctl:latest
```

## 2. Configure a shell alias

```bash
alias pnapctl='docker run --rm -it \
  -e PNAP_CLIENT_ID="YOUR_CLIENT_ID" \
  -e PNAP_CLIENT_SECRET="YOUR_CLIENT_SECRET" \
  ghcr.io/phoenixnap/pnapctl:latest'
```
> Note: This alias is available only for the current shell session. Add it to your shell profile (for example, `.bashrc` or `.zshrc`) to make it persistent.
> Replace `YOUR_CLIENT_ID` and `YOUR_CLIENT_SECRET` with your PhoenixNAP API credentials.

## 3. Verify the setup

```bash
pnapctl version
```

## 4. Run `pnapctl` commands

```bash
pnapctl get servers
```

For additional examples and supported commands, refer to the official documentation.

## Running E2E Tests

E2E tests validate the published `pnapctl` Docker image with GO SDK against real PhoenixNAP APIs.

### Prerequisites

* Java 25+
* Maven 3.9+
* Docker
* Valid `PNAP_CLIENT_ID` and `PNAP_CLIENT_SECRET` credentials

### Usage

Export the required environment variables:

```bash
export PNAP_CLIENT_ID=<your-client-id>
export PNAP_CLIENT_SECRET=<your-client-secret>
export PNAPCTL_IMAGE=ghcr.io/phoenixnap/pnapctl:latest
export PNAPCTL_VERSION=1.13.0
```

Note: `PNAPCTL_VERSION` should be updated accordingly over time as new versions are released.

Run the E2E test suite:

```bash
cd e2e-tests
mvn clean verify 
```

To run specific test group/s use `mvn clean verify -Dgroups=testGroup1,testGroup2` command.

### Notes

* Tests interact with real PhoenixNAP resources and APIs.
    * Anything that is created will be tied to the account owner of the defined PNAP client credentials.
* The Docker image specified by `PNAPCTL_IMAGE` will be used for all CLI operations.
* Ensure the configured credentials have sufficient permissions to perform the tested operations.
