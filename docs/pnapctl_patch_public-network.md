## pnapctl patch public-network

Patch a public network.

### Synopsis

Patch a public network.

Requires a file (yaml or json) containing the information needed to patch the public network.

```
pnapctl patch public-network [ID] [flags]
```

### Examples

```
# Patch a public network using the contents of publicNetworkPatch.yaml as request body. 
pnapctl patch public-network <PUBLIC_NETWORK_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# publicNetworkPatch.yaml
name: Network From CLI (Yaml)
description: This network was updated from the CLI using YAML
```

### Options

```
  -f, --filename string   File containing required information for updating
  -h, --help              help for public-network
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl patch](pnapctl_patch.md)	 - Modify a resource.

