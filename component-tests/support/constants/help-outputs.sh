outputHelpLong="pnapctl creates new and manages existing bare metal servers provided by the phoenixNAP Bare Metal Cloud service.
	
	Find More information at: https://developers.phoenixnap.com/cli

Usage:
  pnapctl [flags]
  pnapctl [command]

Available Commands:
  auto-renew   Modify auto-renew for a resource.
  completion   Generate the autocompletion script for the specified shell
  convert      Convert a resource.
  create       Create a resource.
  delete       Delete a resource.
  deprovision  Deprovision a resource.
  download     Download a resource.
  get          Display one or many resources.
  help         Help about any command
  patch        Modify a resource.
  pay          Pay a resource.
  power-off    Perform a hard shutdown on the resource.
  power-on     Power on a resource.
  reboot       Perform a soft reboot on a resource.
  request-edit Submit a modification request on a resource.
  reserve      Reserve the resource for future use.
  reset        Reset the resource to original state.
  shutdown     Perform a soft shutdown on the resource.
  tag          Tag a resource.
  update       Update a resource.
  version      Print version

Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -h, --help            help for pnapctl
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl [command] --help\" for more information about a command."

outputHelpGet="Display one or many resources.

Usage:
  pnapctl get [flags]
  pnapctl get [command]

Available Commands:
  account-billing-configuration Retrieve your account billing configuration
  bgp-peer-group                Retrieve one or all BGP peer groups.
  cluster                       Retrieve one or all clusters.
  event                         Retrieve all events relating to your account.
  invoice                       Retrieve one or all invoices for your account.
  ip-block                      Retrieve one or all ip-blocks for your account.
  location                      Retrieves all locations.
  private-network               Retrieve one or all private networks.
  product                       Retrieves all products.
  product-availabilities        Retrieve product availabilities
  public-network                Retrieve one or all public networks.
  quota                         Retrieve one or all quotas for your account.
  rated-usage                   Retrieve all rated-usages for the given time period.
  reservation                   Retrieve one or all reservations
  server                        Retrieve one or all servers.
  ssh-key                       Retrieve one or all ssh-keys for your account.
  storage-network               Retrieve one or all storage networks.
  tag                           Retrieve one or all tags.
  transaction                   Retrieve one or all transactions for your account.

Flags:
  -h, --help   help for get

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl get [command] --help\" for more information about a command."

outputHelpCreate="Create a resource.

Usage:
  pnapctl create [flags]
  pnapctl create [command]

Available Commands:
  bgp-peer-group         Create a new BGP peer group.
  cluster                Create a new cluster.
  ip-block               Create a new ip-block.
  private-network        Create a new private network.
  public-network         Create a new public network.
  reservation            Create a new reservation.
  server                 Create a new server.
  server-ip-block        Create a new ip-block for server.
  server-private-network Create a new private network for server.
  server-public-network  Create a new public network for server.
  ssh-key                Create a new ssh-key.
  storage-network        Create a new storage network.
  tag                    Create a new tag.

Flags:
  -h, --help   help for create

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl create [command] --help\" for more information about a command."

outputHelpDelete="Delete a resource.

Usage:
  pnapctl delete [flags]
  pnapctl delete [command]

Available Commands:
  bgp-peer-group         Deletes a BGP peer group.
  cluster                Deletes a specific cluster.
  ip-block               Deletes a specific ip-block.
  private-network        Deletes a specific private network.
  public-network         Deletes a public network.
  server-ip-block        Remove an ip-block from a server.
  server-private-network Remove a server from a private network.
  server-public-network  Remove a server from a public network.
  ssh-key                Deletes a specific SSH Key.
  storage-network        Deletes a specific storage network.
  tag                    Deletes a specific tag.

Flags:
  -h, --help   help for delete

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl delete [command] --help\" for more information about a command."

outputHelpDeprovision="Deprovision a resource

Usage:
  pnapctl deprovision [flags]
  pnapctl deprovision [command]

Available Commands:
  server      Deprovision a server.

Flags:
  -h, --help   help for deprovision

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl deprovision [command] --help\" for more information about a command."

outputHelpPatch="Modify a resource.

Usage:
  pnapctl patch [flags]
  pnapctl patch [command]

Available Commands:
  ip-block        Updates a specific ip-block.
  public-network  Patch a public network.
  server          Patch a server.
  storage-network Patch a storage network.
  tag             Patch/Update a tag.

Flags:
  -h, --help   help for patch

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl patch [command] --help\" for more information about a command."

outputHelpCompletion="Generate the autocompletion script for pnapctl for the specified shell.
See each sub-command's help for details on how to use the generated script.

Usage:
  pnapctl completion [command]

Available Commands:
  bash        Generate the autocompletion script for bash
  fish        Generate the autocompletion script for fish
  powershell  Generate the autocompletion script for powershell
  zsh         Generate the autocompletion script for zsh

Flags:
  -h, --help   help for completion

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl completion [command] --help\" for more information about a command."

outputHelpReset="Reset the resource to the same state as it was originally created.
NOTE: Any data on the resource will be lost.

Usage:
  pnapctl reset [flags]
  pnapctl reset [command]

Available Commands:
  server      Resets a specific server.

Flags:
  -h, --help   help for reset

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl reset [command] --help\" for more information about a command."

outputHelpPowerOff="Perform a hard shutdown on the resource.

Usage:
  pnapctl power-off [flags]
  pnapctl power-off [command]

Available Commands:
  server      Perform a hard shutdown on a specific server.

Flags:
  -h, --help   help for power-off

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl power-off [command] --help\" for more information about a command."

outputHelpPowerOn="Power on a resource.

Usage:
  pnapctl power-on [flags]
  pnapctl power-on [command]

Available Commands:
  server      Powers on a specific server.

Flags:
  -h, --help   help for power-on

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl power-on [command] --help\" for more information about a command."

outputHelpReboot="Perform a soft reboot on a resource.

Usage:
  pnapctl reboot [flags]
  pnapctl reboot [command]

Available Commands:
  server      Perform a soft reboot on a specific server.

Flags:
  -h, --help   help for reboot

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl reboot [command] --help\" for more information about a command."

outputHelpTag="Tag a resource.

Usage:
  pnapctl tag [flags]
  pnapctl tag [command]

Available Commands:
  server      Tag a server.

Flags:
  -h, --help   help for tag

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl tag [command] --help\" for more information about a command."

outputHelpRequestEdit="Submit a modification request on a resource.

Usage:
  pnapctl request-edit [flags]
  pnapctl request-edit [command]

Available Commands:
  quota       Submit a quota modification request.

Flags:
  -h, --help   help for request-edit

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl request-edit [command] --help\" for more information about a command."

outputHelpReserve="Reserve the resource to be used later on.

Usage:
  pnapctl reserve [flags]
  pnapctl reserve [command]

Available Commands:
  server      Reserve a specific server.

Flags:
  -h, --help   help for reserve

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl reserve [command] --help\" for more information about a command."

outputHelpShutdown="Perform a soft shutdown on the resource.

Usage:
  pnapctl shutdown [flags]
  pnapctl shutdown [command]

Available Commands:
  server      Perform a soft shutdown on a specific server.

Flags:
  -h, --help   help for shutdown

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl shutdown [command] --help\" for more information about a command."

outputHelpUpdate="Update a resource.

Usage:
  pnapctl update [flags]
  pnapctl update [command]

Available Commands:
  ip-block        Update an ip-block.
  private-network Update a private network.
  ssh-key         Update an ssh-key.
  storage-network Update a storage network.

Flags:
  -h, --help   help for update

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl update [command] --help\" for more information about a command."

outputHelpDownload="Download a resource.

Usage:
  pnapctl download [flags]
  pnapctl download [command]

Available Commands:
  invoice         Download a resource.

Flags:
  -h, --help         help for download
  -d, --destination  destination for the downloaded invoice

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl update [command] --help\" for more information about a command."

outputHelpPay="Pay a resource.

Usage:
  pnapctl pay [flags]
  pnapctl pay [command]

Available Commands:
  invoice         Pay a resource.

Flags:
  -h, --help         help for download

Global Flags:
      --config string   config file defaults to the environment variable \"PNAPCTL_HOME\" or \"pnap.yaml\" in the home directory.
  -v, --verbose         change log level from Warn (default) to Debug.

Use \"pnapctl update [command] --help\" for more information about a command."