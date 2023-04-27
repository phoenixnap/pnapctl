## pnapctl patch server private-network

Patch a server's private network.

### Synopsis

Patch a server's private network.
	
Requires a file (yaml or json) containing the information needed to patch the server.

```
pnapctl patch server private-network SERVER_ID NETWORK_ID [flags]
```

### Examples

```
# Patch a server using the contents of serverPrivateNetworkPatch.yaml as the request body.
pnapctl patch server private-network <SERVER_ID> <NETWORK_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# serverPrivateNetworkPatch.yaml
hostname: patched-server
description: My custom server edit
```

### Options

```
  -f, --filename string   File containing required information for updating
      --force             Controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. Defaults to false.
  -h, --help              help for private-network
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl patch server](pnapctl_patch_server.md)	 - Patch a server.

