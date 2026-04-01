## pnapctl transfer-reservation server

Transfer reservation of server elsewhere.

### Synopsis

Transfer reservation of server elsewhere.
	
Requires a file (yaml or json) containing the information needed to transfer a server's reservation.

```
pnapctl transfer-reservation server SERVER_ID [flags]
```

### Examples

```
# Transfer a server's reservations using the contents of serverTransferReservation.yaml as request body. 
pnapctl transfer-reservation server <SERVER_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# serverTransferReservation.yaml
targetServerId: "<SERVER ID>"
```

### Options

```
  -f, --filename string   File containing required information for transfer reservation
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

* [pnapctl transfer-reservation](pnapctl_transfer-reservation.md)	 - Transfer a reservation from one point to another(??)

