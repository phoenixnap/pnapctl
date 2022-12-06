## pnapctl get rated-usage month-to-date

Retrieve all rated-usages for the current calendar month.

### Synopsis

Retrieve all rated-usages for the current calendar month.
	
Prints all information about the rated-usages for the current month.
By default, the data is printed in a table format.

Every record corresponds to a charge. All dates & times are in UTC.

```
pnapctl get rated-usage month-to-date [flags]
```

### Examples

```

# List all rated-usages	

```

### Options

```
      --category string   The product category to filter by.
      --full              Shows all rated-usage details
  -h, --help              help for month-to-date
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl get rated-usage](pnapctl_get_rated-usage.md)	 - Retrieve all rated-usages for the given time period.

