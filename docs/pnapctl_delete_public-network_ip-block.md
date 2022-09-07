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
pnapctl delete public-network ip-block <NETWORK_ID> <IP_BLOCK_ID> [--output <OUTPUT_TYPE>]
```

### Options

```
  -h, --help   help for ip-block
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl delete public-network](pnapctl_delete_public-network.md)	 - Delete a public network.

