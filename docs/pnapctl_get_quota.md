## pnapctl get quota

Retrieve one or all quotas for your account.

### Synopsis

Retrieve one or all quotas for your account.

Prints all information about the quotas assigned to your account.
By default, the data is printed in table format.

To print a specific quota, a quota ID needs to be passed as an argument.

```
pnapctl get quota [QUOTA_ID] [flags]
```

### Examples

```

# List all quotas in.
pnapctl get quotas [--output <OUTPUT_TYPE>]

# List a specific quota.
pnapctl get quota <QUOTA_ID> [--output <OUTPUT_TYPE>]
```

### Options

```
  -h, --help            help for quota
  -o, --output string   Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl get](pnapctl_get.md)	 - Display one or many resources.

