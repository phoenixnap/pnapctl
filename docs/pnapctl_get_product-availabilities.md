## pnapctl get product-availabilities

Retrieve product availabilities

### Synopsis

Retrieve one or all reservations.

```
pnapctl get product-availabilities [flags]
```

### Examples

```

# Retrieve all product-availabilities
pnapctl get product-availabilities 
	[--output=<OUTPUT_TYPE>] 
	[--category=<CATEGORY>] 
	[--code=<CODE>] 
	[--showOnlyMinQuantityAvailable=false] 
	[--location=<LOCATION>] 
	[--solution=<SOLUTION>] 
	[--minQuantity=<MIN_QUANTITY>]
```

### Options

```
      --category stringArray           Category to filter product availabilities by.
      --code stringArray               Category to filter product availabilities by.
  -h, --help                           help for product-availabilities
      --location stringArray           Category to filter product availabilities by.
      --minQuantity float32            Category to filter product availabilities by.
  -o, --output string                  Define the output format. Possible values: table, json, yaml (default "table")
      --showOnlyMinQuantityAvailable   Category to filter product availabilities by. (default true)
      --solution stringArray           Category to filter product availabilities by.
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl get](pnapctl_get.md)	 - Display one or many resources.

