## pnapctl tag server

Tag a server.

### Synopsis

Tag a server.

Requires a file (yaml or json) containing the information needed to tag the server.

```
pnapctl tag server SERVER_ID [flags]
```

### Examples

```
# Tag a server as per serverTag.yaml. 
pnapctl tag server --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# serverTag.yaml
- name: tagName
  value: tagValue
- name: tagName2

```

### Options

```
  -f, --filename string   File containing required information for tagging
      --full              Shows all server details
  -h, --help              help for server
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl tag](pnapctl_tag.md)	 - Tag a resource.

