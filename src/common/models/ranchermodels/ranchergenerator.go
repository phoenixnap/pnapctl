package ranchermodels

import (
	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"phoenixnap.com/pnapctl/tests/generators"
)

func GenerateClusters(n int) []ranchersdk.Cluster {
	var clusterlist []ranchersdk.Cluster
	for i := 0; i < n; i++ {
		clusterlist = append(clusterlist, GenerateCluster())
	}
	return clusterlist
}

func GenerateCluster() ranchersdk.Cluster {
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

func GenerateCLICluster() Cluster {
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

func GenerateNodePool() ranchersdk.NodePool {
	return ranchersdk.NodePool{
		Name:       generators.RandSeqPointer(10),
		NodeCount:  nil,
		ServerType: generators.RandSeqPointer(10),
		SshConfig:  nil,
		Nodes:      nil,
	}
}

func GenerateCLINodePool() NodePool {
	return NodePool{
		Name:       generators.RandSeqPointer(10),
		NodeCount:  nil,
		ServerType: generators.RandSeqPointer(10),
		SshConfig:  nil,
		Nodes:      nil,
	}
}

func GenerateNode() ranchersdk.Node {
	return ranchersdk.Node{
		ServerId: generators.RandSeqPointer(10),
	}
}

func GenerateCLINode() Node {
	return Node{
		ServerId: generators.RandSeqPointer(10),
	}
}

func GenerateRancherClusterCertificates() ranchersdk.RancherClusterCertificates {
	return ranchersdk.RancherClusterCertificates{
		CaCertificate:  generators.RandSeqPointer(10),
		Certificate:    generators.RandSeqPointer(10),
		CertificateKey: generators.RandSeqPointer(10),
	}
}

func GenerateCLIRancherClusterCertificates() RancherClusterCertificates {
	return RancherClusterCertificates{
		CaCertificate:  generators.RandSeqPointer(10),
		Certificate:    generators.RandSeqPointer(10),
		CertificateKey: generators.RandSeqPointer(10),
	}
}

func GenerateRancherClusterConfig() ranchersdk.RancherClusterConfig {
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

func GenerateCLIRancherClusterConfig() RancherClusterConfig {
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

func GenerateRancherServerMetadata() ranchersdk.RancherServerMetadata {
	return ranchersdk.RancherServerMetadata{
		Url:      generators.RandSeqPointer(10),
		Username: generators.RandSeqPointer(10),
		Password: generators.RandSeqPointer(10),
	}
}

func GenerateCLIRancherServerMetadata() RancherServerMetadata {
	return RancherServerMetadata{
		Url:      generators.RandSeqPointer(10),
		Username: generators.RandSeqPointer(10),
		Password: generators.RandSeqPointer(10),
	}
}

func GenerateSshConfig() ranchersdk.SshConfig {
	return ranchersdk.SshConfig{
		InstallDefaultKeys: nil,
		Keys:               generators.RandListStringPointer(3),
		KeyIds:             generators.RandListStringPointer(3),
	}
}

func GenerateCLISshConfig() SshConfig {
	return SshConfig{
		InstallDefaultKeys: nil,
		Keys:               generators.RandListStringPointer(3),
		KeyIds:             generators.RandListStringPointer(3),
	}
}

func GenerateRancherDeleteResult() ranchersdk.DeleteResult {
	return ranchersdk.DeleteResult{
		Result:    generators.RandSeq(10),
		ClusterId: generators.RandSeqPointer(10),
	}
}
