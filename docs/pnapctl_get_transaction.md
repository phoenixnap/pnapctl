## pnapctl get transaction

Retrieve one or all transactions for your account.

### Synopsis

Retrieve one or all transactions for your account.

Prints all information about the transactions assigned to your account.
By default, the data is printed in json format.

Table format isn't supported for this command.

To print a specific transaction, a transaction ID needs to be passed as an argument.

```
pnapctl get transaction [TRANSACTION_ID] [flags]
```

### Examples

```

# List all transactions in.
pnapctl get transactions [--limit <LIMIT>] [--offset <OFFSET>] [--sortdirection <SORTDIRECTION>] [--sortfield <SORTFIELD>] [--from <FROM>] [--to <TO>] [--output <OUTPUT_TYPE>]

# List a specific transaction.
pnapctl get transactions <TRANSACTION_ID> [--output <OUTPUT_TYPE>]
```

### Options

```
      --from string            A 'from' filter. Needs to be in the following format: '2021-04-27T16:24:57.123Z'
  -h, --help                   help for transaction
      --limit int              Limit the number of records returned.
      --offset int             The number of items to skip in the results.
  -o, --output string          Define the output format. Possible values: table, json, yaml (default "table")
      --sortDirection string   Ordering of the event's time. Must be 'ASC' or 'DESC'
      --sortField string       If a sortField is requested, pagination will be done after sorting. Default sorting is by date.
      --to string              A 'to' filter. Needs to be in the following format: '2021-04-27T16:24:57.123Z'
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl get](pnapctl_get.md)	 - Display one or many resources.

