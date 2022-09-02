package ranchermodels

import (
	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
)

type ClusterWorkloadConfiguration struct {
	Name        *string `json:"name,omitempty" yaml:"name,omitempty"`
	ServerCount *int32  `json:"serverCount,omitempty" yaml:"serverCount,omitempty"`
	ServerType  string  `json:"serverType" yaml:"serverType"`
	Location    string  `json:"location" yaml:"location"`
}

func (c ClusterWorkloadConfiguration) ToSdk() *ranchersdk.ClusterWorkloadConfiguration {
	return &ranchersdk.ClusterWorkloadConfiguration{
		Name:        c.Name,
		ServerCount: c.ServerCount,
		ServerType:  c.ServerType,
		Location:    c.Location,
	}
}

func ClusterWorkloadConfigurationFromSdk(configuration *ranchersdk.ClusterWorkloadConfiguration) *ClusterWorkloadConfiguration {
	if configuration == nil {
		return nil
	}

	return &ClusterWorkloadConfiguration{
		Name:        configuration.Name,
		ServerCount: configuration.ServerCount,
		ServerType:  configuration.ServerType,
		Location:    configuration.Location,
	}
}
