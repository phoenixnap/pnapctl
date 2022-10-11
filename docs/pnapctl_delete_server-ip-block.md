## pnapctl delete server-ip-block

Remove an ip-block from a server.

### Synopsis

Remove an ip-block from a server.

Requires two IDs passed as arguments and a file (yaml or json) containing the information needed. First one being the server id and second being the ip-block id. 

```
pnapctl delete server-ip-block SERVER_ID IP_BLOCK_ID [flags]
```

### Examples

```
# Remove an ip-block from a server. 
pnapctl delete server-ip-block <SERVER_ID> <IP_BLOCK_ID> --filename <FILE_PATH>

# serveripblockdelete.yaml
deleteIpBlocks: false
```

### Options

```
  -f, --filename string   File containing required information for creation
  -h, --help              help for server-ip-block
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl delete](pnapctl_delete.md)	 - Delete a resource.

