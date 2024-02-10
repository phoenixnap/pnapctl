## pnapctl download invoice

Download invoice.

### Synopsis

Download invoice.

```
pnapctl download invoice INVOICE_ID [flags]
```

### Examples

```
pnapctl download invoice <INVOICE_ID>
```

### Options

```
  -d, --destination string   Set the destination for downloading the invoice. (default "./invoice.pdf")
  -h, --help                 help for invoice
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl download](pnapctl_download.md)	 - Download an invoice.

