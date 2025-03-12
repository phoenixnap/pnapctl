## pnapctl create bgp-peer-group

Create a new BGP peer group.

### Synopsis

Create a BGP peer group.

Requires a file (yaml or json) containing the information needed to create the BGP peer group.

```
pnapctl create bgp-peer-group [flags]
```

### Examples

```
# Create a public network using the contents of bgpPeerGroupCreate.yaml as request body. 
pnapctl create bgp-peer-group --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# bgpPeerGroupCreate.yaml
location: "PHX"
asn: 98239
password: "password"
advertisedRoutes: "DEFAULT"
```

### Options

```
  -f, --filename string   File containing required information for creation
  -h, --help              help for bgp-peer-group
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl create](pnapctl_create.md)	 - Create a resource.

