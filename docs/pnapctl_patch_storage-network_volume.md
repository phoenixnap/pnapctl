## pnapctl patch storage-network volume

Patch a storage network's volume details.

### Synopsis

Patch a storage network's volume details.
	
Requires a file (yaml or json) containing the information needed to patch the storage network's volume.

```
pnapctl patch storage-network volume [storageNetworkID] [volumeID] [flags]
```

### Examples

```
# Patch a storage network's volume using the contents of storagenetworkvolumeupdate.yaml as request body.
pnapctl patch storage-network volume <storageNetworkID> <volumeID> --filename <FILE_PATH> [--output <OUTPUT_TYPE]

# storagenetworkvolumeupdate.yaml
capacityInGb: 2000
```

### Options

```
  -f, --filename string   File containing required information for updating
  -h, --help              help for volume
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -o, --output string   Define the output format. Possible values: table, json, yaml (default "table")
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl patch storage-network](pnapctl_patch_storage-network.md)	 - Patch a storage network.

