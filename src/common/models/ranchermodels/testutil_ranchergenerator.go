package ranchermodels

import (
	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"phoenixnap.com/pnapctl/testsupport/generators"
)

func GenerateClusterListSdk(n int) []ranchersdk.Cluster {
	var clusterlist []ranchersdk.Cluster
	for i := 0; i < n; i++ {
		clusterlist = append(clusterlist, GenerateClusterSdk())
	}
	return clusterlist
}

func GenerateClusterSdk() ranchersdk.Cluster {
	return ranchersdk.Cluster{
		Id:                    generators.RandSeqPointer(10),
		Name:                  generators.RandSeqPointer(10),
		Description:           generators.RandSeqPointer(10),
		Location:              generators.RandSeq(10),
		InitialClusterVersion: generators.RandSeqPointer(10),
		NodePools:             nil,
		Configuration:         nil,
		Metadata:              nil,
		StatusDescription:     generators.RandSeqPointer(10),
	}
}

func GenerateClusterCli() Cluster {
	return Cluster{
		Id:                    generators.RandSeqPointer(10),
		Name:                  generators.RandSeqPointer(10),
		Description:           generators.RandSeqPointer(10),
		Location:              generators.RandSeq(10),
		InitialClusterVersion: generators.RandSeqPointer(10),
		NodePools:             nil,
		Configuration:         nil,
		Metadata:              nil,
		StatusDescription:     generators.RandSeqPointer(10),
	}
}

func GenerateNodePoolSdk() ranchersdk.NodePool {
	return ranchersdk.NodePool{
		Name:       generators.RandSeqPointer(10),
		NodeCount:  generators.RanNumberPointer(),
		ServerType: generators.RandSeqPointer(10),
		SshConfig:  nil,
		Nodes:      nil,
	}
}

func GenerateNodePoolCli() NodePool {
	return NodePool{
		Name:       generators.RandSeqPointer(10),
		NodeCount:  nil,
		ServerType: generators.RandSeqPointer(10),
		SshConfig:  nil,
		Nodes:      nil,
	}
}

func GenerateNodeSdk() ranchersdk.Node {
	return ranchersdk.Node{
		ServerId: generators.RandSeqPointer(10),
	}
}

func GenerateNodeCli() Node {
	return Node{
		ServerId: generators.RandSeqPointer(10),
	}
}

func GenerateRancherClusterCertificatesSdk() ranchersdk.RancherClusterCertificates {
	return ranchersdk.RancherClusterCertificates{
		CaCertificate:  generators.RandSeqPointer(10),
		Certificate:    generators.RandSeqPointer(10),
		CertificateKey: generators.RandSeqPointer(10),
	}
}

func GenerateRancherClusterCertificatesCli() RancherClusterCertificates {
	return RancherClusterCertificates{
		CaCertificate:  generators.RandSeqPointer(10),
		Certificate:    generators.RandSeqPointer(10),
		CertificateKey: generators.RandSeqPointer(10),
	}
}

func GenerateRancherClusterConfigSdk() ranchersdk.RancherClusterConfig {
	return ranchersdk.RancherClusterConfig{
		Token:                    generators.RandSeqPointer(10),
		TlsSan:                   generators.RandSeqPointer(10),
		EtcdSnapshotScheduleCron: generators.RandSeqPointer(10),
		EtcdSnapshotRetention:    nil,
		NodeTaint:                generators.RandSeqPointer(10),
		ClusterDomain:            generators.RandSeqPointer(10),
		Certificates:             nil,
	}
}

func GenerateRancherClusterConfigCli() RancherClusterConfig {
	return RancherClusterConfig{
		Token:                    generators.RandSeqPointer(10),
		TlsSan:                   generators.RandSeqPointer(10),
		EtcdSnapshotScheduleCron: generators.RandSeqPointer(10),
		EtcdSnapshotRetention:    nil,
		NodeTaint:                generators.RandSeqPointer(10),
		ClusterDomain:            generators.RandSeqPointer(10),
		Certificates:             nil,
	}
}

func GenerateRancherServerMetadataSdk() ranchersdk.RancherServerMetadata {
	return ranchersdk.RancherServerMetadata{
		Url:      generators.RandSeqPointer(10),
		Username: generators.RandSeqPointer(10),
		Password: generators.RandSeqPointer(10),
	}
}

func GenerateRancherServerMetadataCli() RancherServerMetadata {
	return RancherServerMetadata{
		Url:      generators.RandSeqPointer(10),
		Username: generators.RandSeqPointer(10),
		Password: generators.RandSeqPointer(10),
	}
}

func GenerateSshConfigSdk() ranchersdk.SshConfig {
	return ranchersdk.SshConfig{
		InstallDefaultKeys: nil,
		Keys:               generators.RandListStringPointer(3),
		KeyIds:             generators.RandListStringPointer(3),
	}
}

func GenerateSshConfigCli() SshConfig {
	return SshConfig{
		InstallDefaultKeys: nil,
		Keys:               generators.RandListStringPointer(3),
		KeyIds:             generators.RandListStringPointer(3),
	}
}

func GenerateRancherDeleteResultSdk() ranchersdk.DeleteResult {
	return ranchersdk.DeleteResult{
		Result:    generators.RandSeq(10),
		ClusterId: generators.RandSeqPointer(10),
	}
}
