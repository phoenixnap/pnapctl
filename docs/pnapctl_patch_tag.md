## pnapctl patch tag

Patch/Update a tag.

### Synopsis

Patch/Update a tag.

Requires a file (yaml or json) containing the information needed to patch the tag.

```
pnapctl patch tag TAG_ID [flags]
```

### Examples

```
# Modify an existing tag as per tagPatch.yaml
pnapctl patch tag <TAG_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# tagPatch.yaml
name: Tag Name
description: The description of the tag.
isBillingTag: false
```

### Options

```
  -f, --filename string   File containing required information for updating
  -h, --help              help for tag
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl patch](pnapctl_patch.md)	 - Modify a resource.

