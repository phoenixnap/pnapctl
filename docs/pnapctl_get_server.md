## pnapctl get server

Retrieve one or all servers.

### Synopsis

Retrieve one or all servers.

Prints brief or detailed information about the servers.
By default, the data is printed in table format.

To print a specific server, an ID needs to be passed as an argument.

```
pnapctl get server [SERVER_ID] [flags]
```

### Examples

```

# List all servers.
pnapctl get servers [--tag <TagName>.<TagValue>] [--tag <TagName>] [--full] [--output <OUTPUT_TYPE>]

# List all specific server.
pnapctl get servers <SERVER_ID> [--full] [--output <OUTPUT_TYPE>]
```

### Options

```
      --full              Shows all server details
  -h, --help              help for server
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
      --tag stringArray   Filter by tag
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl get](pnapctl_get.md)	 - Display one or many resources.

