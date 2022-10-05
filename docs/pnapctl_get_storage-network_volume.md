## pnapctl get storage-network volume

Retrieve one or all volumes.

### Synopsis

Retrieve one or all volumes.
	
Prints information about the volumes.
By default, the data is printed in table format.

To print a specific volume, an ID needs to be passed as argument.

```
pnapctl get storage-network volume [ID] [flags]
```

### Examples

```

# List all volumes.
pnapctl get volumes [--full] [--output <OUTPUT_TYPE>]

# List a specific volume.
pnapctl get volume <ID> [--full] [--output <OUTPUT_TYPE>]
```

### Options

```
      --full   Shows all volume details
  -h, --help   help for volume
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -o, --output string   Define the output format. Possible values: table, json, yaml (default "table")
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl get storage-network](pnapctl_get_storage-network.md)	 - Retrieve one or all storage networks.

