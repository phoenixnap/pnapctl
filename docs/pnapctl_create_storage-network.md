## pnapctl create storage-network

Create a new storage network.

### Synopsis

Create a storage network.
	
Requires a file (yaml or json) containing the information needed to create the storage network.

```
pnapctl create storage-network [flags]
```

### Examples

```
# Create a storage network using the contents of storageNetworkCreate.yaml as request body.
pnapctl create storage-network --filename <FILE_PATH> [--output <OUTPUT_TYPE]

# storageNetworkCreate.yaml
name: "CreatedSN"
description: "Description"
location: "PHX"
volumes:
  - name: "VolumeName"
    description: "VDescription"
    pathSuffix: "/cliyaml"
    capacityInGb: 1000
```

### Options

```
  -f, --filename string   File containing required information for creation
  -h, --help              help for storage-network
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl create](pnapctl_create.md)	 - Create a resource.

