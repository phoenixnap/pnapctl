## pnapctl create ssh-key

Create a new ssh-key.

### Synopsis

Create a new ssh-key.

Requires a file (yaml or json) containing the information needed to create the ssh-key.

```
pnapctl create ssh-key [flags]
```

### Examples

```
# Create a new ssh-key as described in sshKeyCreate.yaml
pnapctl create ssh-key --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# sshKeyCreate.yaml
default: true
name: default ssh key
key: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCyVGaw1PuEl98f4/7Kq3O9ZIvDw2OFOSXAFVqilSFNkHlefm1iMtPeqsIBp2t9cbGUf55xNDULz/bD/4BCV43yZ5lh0cUYuXALg9NI29ui7PEGReXjSpNwUD6ceN/78YOK41KAcecq+SS0bJ4b4amKZIJG3JWmDKljtv1dmSBCrTmEAQaOorxqGGBYmZS7NQumRe4lav5r6wOs8OACMANE1ejkeZsGFzJFNqvr5DuHdDL5FAudW23me3BDmrM9ifUzzjl1Jwku3bnRaCcjaxH8oTumt1a00mWci/1qUlaVFft085yvVq7KZbF2OPPbl+erDW91+EZ2FgEi+v1/CSJ5 test2@test
```

### Options

```
  -f, --filename string   File containing required information for creation
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

* [pnapctl create](pnapctl_create.md)	 - Create a resource.

