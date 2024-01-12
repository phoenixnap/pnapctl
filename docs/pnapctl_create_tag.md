## pnapctl create tag

Create a new tag.

### Synopsis

Create a new tag.

Requires a file (yaml or json) containing the information needed to create the tag.

```
pnapctl create tag [flags]
```

### Examples

```
# Create a new tag as described in tagCreate.yaml
pnapctl create tag --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# tagCreate.yaml
name: TagName
description: The description of the tag.
isBillingTag: false

```

### Options

```
  -f, --filename string   File containing required information for creation
  -h, --help              help for tag
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl create](pnapctl_create.md)	 - Create a resource.

