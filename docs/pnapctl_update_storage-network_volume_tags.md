## pnapctl update storage-network volume tags

Update the tags of a storage network volume.

### Synopsis

Update the tags of a storage network volume.
	
Requires a file (yaml or json) containing the information needed to update the tags of a storage network volume.

```
pnapctl update storage-network volume tags [flags]
```

### Examples

```
# Update the tags of a storage network volume as per storageNetworkVolumeTagsUpdate.yaml
pnapctl update storage-network volume tags <STORAGE_NETWORK_ID> <VOLUME_ID> --filename <FILENAME> [--output <OUTPUT_TYPE>]

# storageNetworkVolumeTagsUpdate.yaml
```

### Options

```
  -f, --filename string   File containing required information for updating
  -h, --help              help for tags
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl update storage-network volume](pnapctl_update_storage-network_volume.md)	 - Update a volume.

