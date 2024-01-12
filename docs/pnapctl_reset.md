## pnapctl reset

Reset the resource to original state.

### Synopsis

Reset the resource to the same state as it was originally created.
NOTE: Any data on the resource will be lost.

```
pnapctl reset [flags]
```

### Options

```
  -h, --help   help for reset
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl](pnapctl.md)	 - pnapctl creates new and manages existing bare metal servers.
* [pnapctl reset server](pnapctl_reset_server.md)	 - Resets a specific server.

