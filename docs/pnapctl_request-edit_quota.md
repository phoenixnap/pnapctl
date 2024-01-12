## pnapctl request-edit quota

Submit a quota modification request.

### Synopsis

Submit a quota modification request.

Requires a file (yaml or json) containing the information needed to submit a quota edit request.

```
pnapctl request-edit quota QUOTA_ID [flags]
```

### Examples

```
# Submit an edit request on an existing quota as per requestEditQuota.yaml
pnapctl request-edit quota <QUOTA_ID> --filename <FILE_PATH>

# requestEditQuota.yaml
limit: 75
reason: My current limit is not enough.
```

### Options

```
  -f, --filename string   File containing required information for submission
  -h, --help              help for quota
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl request-edit](pnapctl_request-edit.md)	 - Submit a modification request on a resource.

