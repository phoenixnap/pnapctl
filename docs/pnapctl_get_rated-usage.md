## pnapctl get rated-usage

Retrieve all rated-usages for the given time period.

### Synopsis

Retrieve all rated-usages for the given time period.

Prints all information about the rated-usages for the given time period.
By default, the data is printed in table format.

Every record corresponds to a charge. All date & times are in UTC.
Note: "from" and "to" are required and need to be in a valid YYYY/MM format.

```
pnapctl get rated-usage [flags]
```

### Examples

```

# List all rated usages.
pnapctl get rated-usages --from=2020/10 --to=2021/11 [--category <CATEGORY>] [--output <OUTPUT_TYPE>]

```

### Options

```
      --category string   The product category to filter by.
      --from string       From year month (inclusive) to filter rated usage records by.
      --full              Shows all rated usage details
  -h, --help              help for rated-usage
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
      --to string         To year month (inclusive) to filter rated usage records by.
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl get](pnapctl_get.md)	 - Display one or many resources.
* [pnapctl get rated-usage month-to-date](pnapctl_get_rated-usage_month-to-date.md)	 - Retrieve all rated-usages for the current calendar month.

