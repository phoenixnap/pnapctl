## pnapctl delete public-network ip-block

Delete an ip-block on a public network.

### Synopsis

Delete an ip-block on a public network.

```
pnapctl delete public-network ip-block [ID] [flags]
```

### Examples

```
# Delete an ip-block on a public network.
pnapctl delete public-network ip-block <NETWORK_ID> <IP_BLOCK_ID> [--output <OUTPUT_TYPE>] 	[--force=false] 

```

### Options

```
      --force           Controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. Defaults to false.
  -h, --help            help for ip-block
  -o, --output string   Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl delete public-network](pnapctl_delete_public-network.md)	 - Deletes a public network.

