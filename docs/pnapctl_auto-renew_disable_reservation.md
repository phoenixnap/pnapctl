## pnapctl auto-renew disable reservation

Disable auto-renew for a reservation

### Synopsis

Disable auto-renew for a reservation.
	
Requires a file (yaml or json) containing the information needed to disable auto-renew.

```
pnapctl auto-renew disable reservation [RESERVATION_ID] [flags]
```

### Examples

```

# Disable auto-renew for a specific reservation
pnapctl auto-renew disable reservation <RESERVATION_ID> --filename=<FILENAME>

# reservationAutoRenewDisable.yaml
autoRenewDisableReasons: "disable reason"
```

### Options

```
  -f, --filename string   File containing required information for creation
      --full              Shows all reservation details
  -h, --help              help for reservation
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl auto-renew disable](pnapctl_auto-renew_disable.md)	 - Disable auto-renew for a resource.

