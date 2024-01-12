## pnapctl get location

Retrieves all locations.

### Synopsis

Retrieve all locations.
	
Prints all information about locations.
By default, the data is printed in table format.

```
pnapctl get location [flags]
```

### Options

```
  -h, --help                      help for location
      --location string           Location to filter by.
  -o, --output string             Define the output format. Possible values: table, json, yaml (default "table")
      --product-category string   Product category to filter locations by.
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl get](pnapctl_get.md)	 - Display one or many resources.

