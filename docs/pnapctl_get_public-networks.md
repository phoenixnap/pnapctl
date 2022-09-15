## pnapctl get public-networks

Retrieve one or all public networks.

### Synopsis

Retrieve one or all public networks.

Prints detailed information about the public networks.
By default, the data is printed in table format.

To print a specific public network, an ID needs to be passed as an argument.

```
pnapctl get public-networks [PUBLIC_NETWORK_ID] [flags]
```

### Examples

```

# List all public networks.
pnapctl get public-networks [--location <LOCATION>] [--output <OUTPUT_TYPE>]

# List all details of a specific public network.
pnapctl get public-networks <PUBLIC_NETWORK_ID> [--output <OUTPUT_TYPE>]
```

### Options

```
  -h, --help              help for public-networks
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

