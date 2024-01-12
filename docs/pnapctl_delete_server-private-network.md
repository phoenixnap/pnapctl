## pnapctl delete server-private-network

Remove a server from a private network.

### Synopsis

Remove a server from a private network.

Requires two IDs passed as arguments. First one being the server id and second being the private network id. 

```
pnapctl delete server-private-network SERVER_ID PRIVATE_NETWORK_ID [flags]
```

### Examples

```
# remove a server from a private network 
pnapctl delete server-private-network <SERVER_ID> <PRIVATE_NETWORK_ID>

```

### Options

```
  -h, --help   help for server-private-network
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl delete](pnapctl_delete.md)	 - Delete a resource.

