## pnapctl auto-renew reservation enable

Enable auto-renew for a reservation

### Synopsis

Enable auto-renew for a reservation.

```
pnapctl auto-renew reservation enable [RESERVATION_ID] [flags]
```

### Examples

```

# Enable auto-renew for a specific reservation
pnapctl auto-renew reservation enable <RESERVATION_ID>
```

### Options

```
      --full            Shows all reservation details
  -h, --help            help for enable
  -o, --output string   Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl auto-renew reservation](pnapctl_auto-renew_reservation.md)	 - autorenew for a resource.

