## pnapctl create cluster

Create a new cluster.

### Synopsis

Create a new cluster.
	
Requires a file (yaml or json) containing the information needed to create the cluster.

```
pnapctl create cluster [flags]
```

### Examples

```
# Create a new cluster as described in clusterCreate.yaml
pnapctl create cluster --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# clusterCreate.yaml
location: PHX
name: rancher-cluster-test
nodePools:
  - serverType: s1.c1.medium

```

### Options

```
  -f, --filename string   File containing required information for creation
  -h, --help              help for cluster
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl create](pnapctl_create.md)	 - Create a resource.

