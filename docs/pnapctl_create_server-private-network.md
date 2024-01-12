## pnapctl create server-private-network

Create a new private network for server.

### Synopsis

Create a new private network for server.

Requires a file (yaml or json) containing the information needed to create the server private network.

```
pnapctl create server-private-network SERVER_ID [flags]
```

### Examples

```
# Add a server to a private network as defined in serverCreatePrivateNetwork.yaml
pnapctl create server-private-network <SERVER_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>] [--force=false]

# serverCreatePrivateNetwork.yaml
id: 5ff5cc9bc1acf144d9106233
ips: 
  - 10.0.0.1
  - 10.0.0.2
dhcp: false
statusDescription: in-progress

```

### Options

```
  -f, --filename string   File containing required information for creation
      --force             Controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. Defaults to false.
  -h, --help              help for server-private-network
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl create](pnapctl_create.md)	 - Create a resource.

