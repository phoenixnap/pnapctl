## pnapctl create public-network

Create a new public network.

### Synopsis

Create a public network.

Requires a file (yaml or json) containing the information needed to create the public network.

```
pnapctl create public-network [flags]
```

### Examples

```
# Create a public network using the contents of publicNetworkCreate.yaml as request body. 
pnapctl create public-network --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# publicNetworkCreate.yaml
hostname: patched-server
description: My custom server edit
```

### Options

```
  -f, --filename string   File containing required information for creation.
  -h, --help              help for public-network
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl create](pnapctl_create.md)	 - Create a resource.
* [pnapctl create public-network ip-block](pnapctl_create_public-network_ip-block.md)	 - Create an ip-block on a public network.

