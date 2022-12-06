## pnapctl create public-network ip-block

Create an ip-block on a public network.

### Synopsis

Create an ip-block on a public network.

Requires a file (yaml or json) containing the information needed to create an ip-block.

```
pnapctl create public-network ip-block [NETWORK_ID] [flags]
```

### Examples

```
# Create an ip-block using the contents of publicNetworkIpBlockCreate.yaml as request body. 
pnapctl create public-network ip-block <NETWORK_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# publicNetworkIpBlockCreate.yaml
hostname: patched-server
description: My custom server edit
```

### Options

```
  -f, --filename string   File containing required information for creation
  -h, --help              help for ip-block
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl create public-network](pnapctl_create_public-network.md)	 - Create a new public network.

