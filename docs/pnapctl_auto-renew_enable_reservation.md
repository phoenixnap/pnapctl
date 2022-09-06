## pnapctl auto-renew enable reservation

Enable auto-renew for a reservation

### Synopsis

Enable auto-renew for a reservation.

```
pnapctl auto-renew enable reservation [RESERVATION_ID] [flags]
```

### Examples

```

# Enable auto-renew for a specific reservation
pnapctl auto-renew enable reservation <RESERVATION_ID>
```

### Options

```
      --full            Shows all reservation details
  -h, --help            help for reservation
  -o, --output string   Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl auto-renew enable](pnapctl_auto-renew_enable.md)	 - Enable auto-renew for a resource.

