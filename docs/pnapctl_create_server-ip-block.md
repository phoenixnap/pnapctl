## pnapctl create server-ip-block

Create a new ip-block for server.

### Synopsis

Create a new ip-block for server.

Requires a file (yaml or json) containing the information needed to create the server ip-block.

```
pnapctl create server-ip-block SERVER_ID [flags]
```

### Examples

```
# Add an ip-block to a server defined in servercreateipblock.yaml
pnapctl create server-ip-block <SERVER_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# servercreateipblock.yaml
id: 5ff5cc9bc1acf144d9106233
vlanId: 11
```

### Options

```
  -f, --filename string   File containing required information for creation
  -h, --help              help for server-ip-block
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl create](pnapctl_create.md)	 - Create a resource.

