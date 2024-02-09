## pnapctl update ssh-key

Update an ssh-key.

### Synopsis

Update an ssh-key.

Requires a file (yaml or json) containing the information needed to modify the ssh-key.

```
pnapctl update ssh-key SSH_KEY_ID [flags]
```

### Examples

```
# Update an ssh-key as per sshKeyUpdate.yaml
pnapctl update ssh-key <SSH_KEY_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# sshKeyUpdate.yaml
default: true
name: default ssh key
```

### Options

```
  -f, --filename string   File containing required information for updating
      --full              Shows all ssh key details
  -h, --help              help for ssh-key
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl update](pnapctl_update.md)	 - Update a resource.

