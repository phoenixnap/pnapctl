## pnapctl get ssh-key

Retrieve one or all ssh-keys for your account.

### Synopsis

Retrieve one or all ssh-keys for your account.

Prints one or all ssh-keys assigned to your account.
By default, the data is printed in table format.

To print a specific ssh-key, an ID linked to the resource needs to be passed as an argument.

```
pnapctl get ssh-key [SSH_KEY_ID] [flags]
```

### Examples

```

# List all ssh-keys.
pnapctl get ssh-keys [--full] [--output <OUTPUT_TYPE>]

# List a specific ssh-key.
pnapctl get ssh-key <SSH_KEY_ID> [--full] [--output <OUTPUT_TYPE>]
```

### Options

```
      --full            Shows all ssh key details
  -h, --help            help for ssh-key
  -o, --output string   Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl get](pnapctl_get.md)	 - Display one or many resources.

