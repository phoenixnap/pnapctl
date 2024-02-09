## pnapctl create ip-block

Create a new ip-block.

### Synopsis

Create a new ip-block.

Requires a file (yaml or json) containing the information needed to create the ip-block.

```
pnapctl create ip-block [flags]
```

### Examples

```
# Create a new ip-block as described in ipblockcreate.yaml
pnapctl create ip-block --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# ipblockcreate.yaml
cidrBlockSize: /28
location: PHX
```

### Options

```
  -f, --filename string   File containing required information for creation
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

* [pnapctl create](pnapctl_create.md)	 - Create a resource.

