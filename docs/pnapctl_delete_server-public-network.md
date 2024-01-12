## pnapctl delete server-public-network

Remove a server from a public network.

### Synopsis

Remove a server from a public network.

Requires two IDs passed as arguments. First one being the server id and second being the public network id. 

```
pnapctl delete server-public-network SERVER_ID PUBLIC_NETWORK_ID [flags]
```

### Examples

```
# remove a server from a public network 
pnapctl delete server-public-network <SERVER_ID> <PUBLIC_NETWORK_ID>

```

### Options

```
  -h, --help   help for server-public-network
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl delete](pnapctl_delete.md)	 - Delete a resource.

