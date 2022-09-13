## pnapctl create server-public-network

Create a new public network for server.

### Synopsis

Create a new public network for server.

Requires a file (yaml or json) containing the information needed to create the server public network.

```
pnapctl create server-public-network SERVER_ID [flags]
```

### Examples

```
# Add a server to a public network as defined in serverCreatePublicNetwork.yaml
pnapctl create server-public-network <SERVER_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# serverCreatePublicNetwork.yaml
id: 5ff5cc9bc1acf144d9106233
ips: 
  - 10.0.0.1
  - 10.0.0.2
statusDescription: in-progress

```

### Options

```
  -f, --filename string   File containing required information for creation
  -h, --help              help for server-public-network
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl create](pnapctl_create.md)	 - Create a resource.

