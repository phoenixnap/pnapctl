## pnapctl get ip-block

Retrieve one or all ip-blocks for your account.

### Synopsis

Retrieve one or all ip-blocks for your account.

Prints all information about the ip-blocks owned by your account.
By default, the data is printed in table format.

To print a specific ip-block, an ip-block ID needs to be passed as an argument.

```
pnapctl get ip-block [IP_BLOCK_ID] [flags]
```

### Examples

```

# List all ip-blocks.
pnapctl get ip-blocks [--output <OUTPUT_TYPE>]

# List a specific ip-block.
pnapctl get ip-block <IP_BLOCK_ID> [--output <OUTPUT_TYPE>]
```

### Options

```
  -h, --help            help for ip-block
  -o, --output string   Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl get](pnapctl_get.md)	 - Display one or many resources.

