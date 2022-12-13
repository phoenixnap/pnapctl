## pnapctl patch storage-network

Patch a storage network.

### Synopsis

Patch a storage network.
	
Requires a file (yaml or json) containing the information needed to patch the storage network.

```
pnapctl patch storage-network [ID] [flags]
```

### Examples

```
# Patch a storage network using the contents of storageNetworkPatch.yaml as request body.
pnapctl patch storage-network <ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE]

# storageNetworkPatch.yaml
name: "UpdatedSN"
description: "Description"
```

### Options

```
  -f, --filename string   File containing required information for updating
  -h, --help              help for storage-network
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl patch](pnapctl_patch.md)	 - Modify a resource.

