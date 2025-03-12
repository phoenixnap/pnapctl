## pnapctl get bgp-peer-group

Retrieve one or all BGP peer groups.

### Synopsis

Retrieve one or all BGP peer groups.

Prints detailed information about the BGP peer groups.
By default, the data is printed in table format.

To print a specific BGP peer group, an ID needs to be passed as an argument.

```
pnapctl get bgp-peer-group [PUBLIC_NETWORK_ID] [flags]
```

### Examples

```

# List all BGP peer groups.
pnapctl get bgp-peer-groups [--location <LOCATION>] [--output <OUTPUT_TYPE>]

# List all details of a specific public network.
pnapctl get bgp-peer-groups <BGP_PEER_GROUP_ID> [--output <OUTPUT_TYPE>]
```

### Options

```
  -h, --help              help for bgp-peer-group
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

