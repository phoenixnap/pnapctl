## pnapctl patch server

Patch a server.

### Synopsis

Patch a server.

Requires a file (yaml or json) containing the information needed to patch the server.

```
pnapctl patch server SERVER_ID [flags]
```

### Examples

```
# Patch a server using the contents of serverPatch.yaml as request body. 
pnapctl patch server <SERVER_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# serverPatch.yaml
hostname: patched-server
description: My custom server edit
```

### Options

```
  -f, --filename string   File containing required information for creation
      --full              Shows all server details
  -h, --help              help for server
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl patch](pnapctl_patch.md)	 - Modify a resource.

