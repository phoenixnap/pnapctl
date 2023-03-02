## pnapctl create storage-network volume

Create a storage network's volume.

### Synopsis

Create a storage network's volume.
	
Requires a file (yaml or json) containing the information needed to create a storage network's volume.

```
pnapctl create storage-network volume [storageNetworkID] [flags]
```

### Examples

```
# Create a storage network's volume using the contents of storagenetworkvolumecreate.yaml as request body.
pnapctl create storage-network volume <storageNetworkID> --filename <FILE_PATH> [--output <OUTPUT_TYPE]

# storagenetworkvolumecreate.yaml
name: name
description:description
capacityInGb: 2000
pathSuffix: /pathSuffix
```

### Options

```
  -f, --filename string   File containing required information for updating
  -h, --help              help for volume
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl create storage-network](pnapctl_create_storage-network.md)	 - Create a new storage network.

