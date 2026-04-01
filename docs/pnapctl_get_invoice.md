## pnapctl get invoice

Retrieve one or all invoices for your account.

### Synopsis

Retrieve one or all invoices for your account.
Prints all information about the invoices assigned to your account.
By default, the data is printed in json format.
Table format isn't supported for this command.
To print a specific invoice, an invoice ID needs to be passed as an argument.

```
pnapctl get invoice [INVOICE_ID] [flags]
```

### Examples

```

# List all invoices.
pnapctl get invoices [--number <NUMBER>] [--status <STATUS>] [--sentOnFrom <SENT_ON_FROM>] [--sentOnTo <SENT_ON_TO>] [--limit <LIMIT>] [--offset <OFFSET>] [--sortField <SORT_FIELD>] [--sortDirection <SORT_DIRECTION>] [--output <OUTPUT_TYPE>]

# List a specific invoice.
pnapctl get invoice <INVOICE_ID> [--output <OUTPUT_TYPE>]
```

### Options

```
  -h, --help                   help for invoice
      --limit int              The limit of the number of results returned. The number of records returned may be smaller than the limit.
      --number string          A user-friendly reference number assigned to the invoice.
      --offset int             The number of items to skip in the results.
  -o, --output string          Define the output format. Possible values: table, json, yaml (default "table")
      --sentOnFrom string      Minimum value to filter invoices by sent on date.
      --sentOnTo string        Maximum value to filter invoices by sent on date.
      --sortDirection string   Sort Given Field depending on the desired direction. Default sorting is descending.
      --sortField string       If a sortField is requested, pagination will be done after sorting. Default sorting is by number.
      --status string          Payment status of the invoice.
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl get](pnapctl_get.md)	 - Display one or many resources.

