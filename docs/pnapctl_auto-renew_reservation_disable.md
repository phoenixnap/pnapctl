## pnapctl auto-renew reservation disable

Disable auto-renew for a reservation

### Synopsis

Disable auto-renew for a reservation.
	
Requires a file (yaml or json) containing the information needed to disable auto-renew.

```
pnapctl auto-renew reservation disable [RESERVATION_ID] [flags]
```

### Examples

```

# Disable auto-renew for a specific reservation
pnapctl auto-renew reservation disable <RESERVATION_ID> --filename=<FILENAME>

# reservationAutoRenewDisable.yaml
autoRenewDisableReasons: "disable reason"
```

### Options

```
  -f, --filename string   File containing required information for creation
      --full              Shows all reservation details
  -h, --help              help for disable
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl auto-renew reservation](pnapctl_auto-renew_reservation.md)	 - autorenew for a resource.

