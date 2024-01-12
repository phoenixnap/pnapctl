## pnapctl get storage-network

Retrieve one or all storage networks.

### Synopsis

Retrieve one or all storage networks.
	
Prints information about the storage networks.
By default, the data is printed in table format.

To print a specific storage network, an ID needs to be passed as argument.

```
pnapctl get storage-network [ID] [flags]
```

### Examples

```

# List all storage networks.
pnapctl get storage-networks [--output <OUTPUT_TYPE>]

# List a specific storage network.
pnapctl get storage-network <ID> [--output <OUTPUT_TYPE>]
```

### Options

```
  -h, --help            help for storage-network
  -o, --output string   Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl get](pnapctl_get.md)	 - Display one or many resources.
* [pnapctl get storage-network volume](pnapctl_get_storage-network_volume.md)	 - Retrieve one or all volumes.

