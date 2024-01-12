## pnapctl deprovision server

Deprovision a server.

### Synopsis

Deprovision a server.

Requires a file (yaml or json) containing the information needed to deprovision a server.

```
pnapctl deprovision server SERVER_ID [flags]
```

### Examples

```
# Deprovision a server as per serverdeprovision.yaml
pnapctl deprovision server <SERVER_ID> --filename <FILE_PATH>

# serverdeprovision.yaml
deleteIpBlocks: false
```

### Options

```
  -f, --filename string   File containing required information for deprovisioning
  -h, --help              help for server
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl deprovision](pnapctl_deprovision.md)	 - Deprovision a resource.

