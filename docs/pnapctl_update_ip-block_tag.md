## pnapctl update ip-block tag

Updates an ip block's tags.

### Synopsis

Update an existing ip-block's tag.

Requires a file (yaml or json) containing the information needed to update the ip-block's tags.
	

```
pnapctl update ip-block tag IP_BLOCK_ID [flags]
```

### Examples

```
# Update a tag on an existing ip-block with request body as described in ipblockputtag.yaml
pnapctl update ip-block tag <IP_BLOCK_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# ipblockputtag.yaml
---
- name: ip block tag name
  value: ip block tag value
```

### Options

```
  -f, --filename string   File containing required information for updating
      --full              Shows all ip-block details
  -h, --help              help for tag
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl update ip-block](pnapctl_update_ip-block.md)	 - Update an ip-block.

