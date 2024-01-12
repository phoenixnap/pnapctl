## pnapctl get reservation

Retrieve one or all reservations

### Synopsis

Retrieve one or all reservations.

```
pnapctl get reservation [RESERVATION_ID] [flags]
```

### Examples

```

# Retrieve all reservations
pnapctl get reservations [--category=<CATEGORY>] [--full] [--output=<OUTPUT_TYPE>]

# Retrieve a specific reservation
pnapctl get reservation <RESERVATION_ID> [--full] [--output=<OUTPUT_TYPE>]
```

### Options

```
      --category string   Product category to filter reservations by.
      --full              Shows all reservation details
  -h, --help              help for reservation
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl get](pnapctl_get.md)	 - Display one or many resources.

