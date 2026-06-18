## pnapctl update server os-configuration ipxe

Update the iPXE OS configuration of a server.

### Synopsis

Update the iPXE OS configuration of a server.

Requires a file (yaml or json) containing the information needed to modify the server's iPXE configuration.

```
pnapctl update server os-configuration ipxe SERVER_ID [flags]
```

### Examples

```
# Update the iPXE OS configuration of a server as per serverOsConfigurationIpxeUpdate.yaml
pnapctl update server os-configuration ipxe <SERVER_ID> --filename <FILENAME> [--output <OUTPUT_TYPE>]

# serverOsConfigurationIpxeUpdate.yaml
url: https://example.com/boot.ipxe
nativeVlanConfiguration:
  vlanId: 10
  staticDhcpAddressV4: 185.74.213.56
```

### Options

```
  -f, --filename string   File containing required information for updating
  -h, --help              help for ipxe
  -o, --output string     Define the output format. Possible values: table, json, yaml (default "table")
```

### Options inherited from parent commands

```
      --config string   config file defaults to the environment variable "PNAPCTL_HOME" or "pnap.yaml" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.
```

### SEE ALSO

* [pnapctl update server os-configuration](pnapctl_update_server_os-configuration.md)	 - Update a server's OS configuration.

