package printer

import (
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

// Storage Network
func PrintStorageNetworkResponse(storageNetwork *networkstorageapi.StorageNetwork, commandName string) error {
	networkStorageToPrint := PrepareNetworkStorageForPrinting(*storageNetwork)
	return MainPrinter.PrintOutput(networkStorageToPrint, commandName)
}

func PrintStorageNetworkListResponse(storageNetworks []networkstorageapi.StorageNetwork, commandName string) error {
	networkStoragesToPrint := iterutils.Map(storageNetworks, PrepareNetworkStorageForPrinting)
	return MainPrinter.PrintOutput(networkStoragesToPrint, commandName)
}

func PrepareNetworkStorageForPrinting(storageNetwork networkstorageapi.StorageNetwork) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.StorageNetworkTableFromSdk(storageNetwork)
	default:
		return storageNetwork
	}
}

// Volume
func PrintVolumeResponse(volume *networkstorageapi.Volume, full bool, commandName string) error {
	volumeToPrint := PrepareVolumeForPrinting(*volume, full)
	return MainPrinter.PrintOutput(volumeToPrint, commandName)
}

func PrintVolumeListResponse(volumes []networkstorageapi.Volume, full bool, commandName string) error {
	volumesToPrint := iterutils.Map(volumes, withFull(full, PrepareVolumeForPrinting))
	return MainPrinter.PrintOutput(volumesToPrint, commandName)
}

func PrepareVolumeForPrinting(volume networkstorageapi.Volume, full bool) interface{} {
	table := OutputIsTable()

	switch {
	case table && full:
		return tables.VolumeTableFromSdk(volume)
	case table:
		return tables.ShortVolumeTableFromSdk(volume)
	default:
		return volume
	}
}
