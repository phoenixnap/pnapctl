## pnapctl get tag

Retrieve one or all tags.

### Synopsis

Retrieve one or all tags.
	
Prints information about the tags.
By default, the data is printed in table format.

To print a specific tag, an ID needs to be passed as an argument.

```
pnapctl get tag [TAG_ID] [flags]
```

### Examples

```

# List all tags.
pnapctl get tags [--output <OUTPUT_TYPE>]

# List a specific tag.
pnapctl get tag <TAG_ID> [--output <OUTPUT_TYPE>]
```

### Options

```
  -h, --help            help for tag
      --name string     Name of the tag
  -o, --output string   Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl get](pnapctl_get.md)	 - Display one or many resources.

