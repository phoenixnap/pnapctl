## pnapctl reserve server

Reserve a specific server.

### Synopsis

Reserve a specific server for future use.

Requires a file (yaml or json) containing the information needed to reserve the specific server.

```
pnapctl reserve server SERVER_ID [flags]
```

### Examples

```
# Reserve a specific server with pricing model described in serverReserve.yaml
pnapctl reserve server <SERVER_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# serverReserve.yaml
pricingModel: ONE_MONTH_RESERVATION
```

### Options

```
  -f, --filename string   File containing required information for reservation
      --full              Shows all server details
  -h, --help              help for server
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl reserve](pnapctl_reserve.md)	 - Reserve the resource for future use.

