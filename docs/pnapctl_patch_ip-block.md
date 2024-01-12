## pnapctl patch ip-block

Updates a specific ip-block.

### Synopsis

Patch an existing ip-block.

Requires a file (yaml or json) containing the information needed to update the ip-block.

```
pnapctl patch ip-block IP_BLOCK_ID [flags]
```

### Examples

```
# Update an existing ip-block with request body as described in ipblockpatch.yaml
	pnapctl patch ip-block <IP_BLOCK_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]
	
	# ipblockpatch.yaml
	description: ip block description
```

### Options

```
  -f, --filename string   File containing required information for updating
      --full              Shows all ip-block details
  -h, --help              help for ip-block
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl patch](pnapctl_patch.md)	 - Modify a resource.

