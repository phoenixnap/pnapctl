## pnapctl get product

Retrieves all products.

### Synopsis

Retrieve all products.

Prints all information about products.
By default, the data is printed in table format.

```
pnapctl get product [flags]
```

### Options

```
      --category string       Product category to filter products by.
  -h, --help                  help for product
      --location string       Location to filter products by.
  -o, --output string         Define the output format. Possible values: table, json, yaml (default "table")
      --product-code string   Product code to filter products by.
      --sku-code string       Sku code to filter products by.
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Info (default) to Debug.
```

### SEE ALSO

* [pnapctl get](pnapctl_get.md)	 - Display one or many resources.

