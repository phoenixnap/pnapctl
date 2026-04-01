## pnapctl convert reservation

Convert a reservation

### Synopsis

Convert a reservation.
	
Requires a file (yaml or json) containing the information needed to convert the reservation

```
pnapctl convert reservation [RESERVATION_ID] [flags]
```

### Examples

```

# Convert a specific reservation
pnapctl convert reservation <RESERVATION_ID> --filename=[FILENAME]

# convertReservation.yaml
sku: "SKU_CODE"
```

### Options

```
  -f, --filename string   File containing required information for conversion
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

* [pnapctl convert](pnapctl_convert.md)	 - Convert a resource.

