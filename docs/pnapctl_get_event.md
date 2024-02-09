## pnapctl get event

Retrieve all events relating to your account.

### Synopsis

Retrieve all events relating to your account.
	
By default, the data is printed in table format.

```
pnapctl get event [flags]
```

### Examples

```

# List all events.
pnapctl get events [--from <FROM>] [--to <TO>] [--limit <LIMIT>] [--order <ORDER>] [--username <USERNAME>] [--verb <VERB>] [--uri <URI>] [--output <OUTPUT_TYPE>]
```

### Options

```
      --from string       A 'from' filter. Needs to be in the following format: '2021-04-27T16:24:57.123Z'
  -h, --help              help for event
      --limit int         Limit the number of records returned.
      --order string      Ordering of the event's time. Must be 'ASC' or 'DESC'
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
      --to string         A 'to' filter. Needs to be in the following format: '2021-04-27T16:24:57.123Z'
      --uri string        The request URI.
      --username string   The username that did the actions.
      --verb string       The HTTP verb corresponding to the action. Must be 'POST', 'PUT', 'PATCH', 'DELETE'
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl get](pnapctl_get.md)	 - Display one or many resources.

