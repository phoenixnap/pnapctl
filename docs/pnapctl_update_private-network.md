## pnapctl update private-network

Update a private network.

### Synopsis

Update a private network.

Requires a file (yaml or json) containing the information needed to modify the private-network.

```
pnapctl update private-network PRIVATE_NETWORK_ID [flags]
```

### Examples

```
# Update a private network as per privateNetworkUpdate.yaml
pnapctl update private-network <PRIVATE_NETWORK_ID> --filename <FILENAME> [--output <OUTPUT_TYPE>]

# privateNetworkUpdate.yaml
name: Example CLI Network Updated,
description: Example CLI Network (Updated Description),
locationDefault: true
```

### Options

```
  -f, --filename string   File containing required information for creation
  -h, --help              help for private-network
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl update](pnapctl_update.md)	 - Update a resource.

