## pnapctl create reservation

Create a new reservation.

### Synopsis

Create a new reservation.
	
Requires a file (yaml or json) containing the information needed to create the reservation.

```
pnapctl create reservation [RESERVATION_ID] [flags]
```

### Examples

```

# Create a specific reservation
pnapctl create reservation <RESERVATION_ID> --filename=<FILENAME>

# reservationCreate.yaml
sku: "skuCode"
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

* [pnapctl create](pnapctl_create.md)	 - Create a resource.

