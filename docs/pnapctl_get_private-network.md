## pnapctl get private-network

Retrieve one or all private networks.

### Synopsis

Retrieve one or all private networks.

Prints detailed information about the private networks.
By default, the data is printed in table format.

To print a specific private network, an ID needs to be passed as an argument.

```
pnapctl get private-network [PRIVATE_NETWORK_ID] [flags]
```

### Examples

```

# List all private networks.
pnapctl get private-networks [--location <LOCATION>] [--output <OUTPUT_TYPE>]

# List all details of a specific private network.
pnapctl get private-networks <PRIVATE_NETWORK_ID> [--output <OUTPUT_TYPE>]
```

### Options

```
  -h, --help              help for private-network
      --location string   Filter by location
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl get](pnapctl_get.md)	 - Display one or many resources.

