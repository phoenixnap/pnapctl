## pnapctl patch server public-network

Patch a server's public network.

### Synopsis

Patch a server's public network.
	
Requires a file (yaml or json) containing the information needed to patch the server.

```
pnapctl patch server public-network SERVER_ID NETWORK_ID [flags]
```

### Examples

```
# Patch a server using the contents of serverPublicNetworkPatch.yaml as the request body.
pnapctl patch server public-network <SERVER_ID> <NETWORK_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# serverPublicNetworkPatch.yaml
ips:
  - "10.0.0.0"
```

### Options

```
  -f, --filename string   File containing required information for updating
      --force             Controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. Defaults to false.
  -h, --help              help for public-network
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl patch server](pnapctl_patch_server.md)	 - Patch a server.

