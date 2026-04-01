## pnapctl delete bgp-peer-group

Deletes a BGP peer group.

### Synopsis

Delete a BGP peer group.

```
pnapctl delete bgp-peer-group [ID] [flags]
```

### Examples

```
# Delete a BGP peer group
pnapctl delete BGP peer group <ID>
```

### Options

```
  -h, --help            help for bgp-peer-group
  -o, --output string   Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl delete](pnapctl_delete.md)	 - Delete a resource.

