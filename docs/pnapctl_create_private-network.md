## pnapctl create private-network

Create a new private network.

### Synopsis

Create a new private-network.

Requires a file (yaml or json) containing the information needed to create the private network.

```
pnapctl create private-network [flags]
```

### Examples

```
# Create a new private network as per privateNetworkCreate.yaml
pnapctl create private-network --filename <FILE_PATH> [--output <OUTPUT_TYPE>] [--force]

# privateNetworkCreate.yaml
name: Example CLI Network,
location: PHX,
locationDefault: false,
description: Example CLI Network,
cidr: 10.0.0.0/24
```

### Options

```
  -f, --filename string   File containing required information for creation
      --force             Controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. Defaults to false.
  -h, --help              help for private-network
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl create](pnapctl_create.md)	 - Create a resource.

