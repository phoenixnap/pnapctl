package ranchermodels

import (
	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"phoenixnap.com/pnapctl/testsupport/testutil"
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
		Id:                    testutil.RandSeqPointer(10),
		Name:                  testutil.RandSeqPointer(10),
		Description:           testutil.RandSeqPointer(10),
		Location:              testutil.RandSeq(10),
		InitialClusterVersion: testutil.RandSeqPointer(10),
		NodePools:             nil,
		Configuration:         nil,
		Metadata:              nil,
		StatusDescription:     testutil.RandSeqPointer(10),
	}
}

func GenerateClusterCli() Cluster {
	return Cluster{
		Id:                    testutil.RandSeqPointer(10),
		Name:                  testutil.RandSeqPointer(10),
		Description:           testutil.RandSeqPointer(10),
		Location:              testutil.RandSeq(10),
		InitialClusterVersion: testutil.RandSeqPointer(10),
		NodePools:             nil,
		Configuration:         nil,
		Metadata:              nil,
		StatusDescription:     testutil.RandSeqPointer(10),
	}
}

func GenerateNodePoolSdk() ranchersdk.NodePool {
	return ranchersdk.NodePool{
		Name:       testutil.RandSeqPointer(10),
		NodeCount:  testutil.RanNumberPointer(),
		ServerType: testutil.RandSeqPointer(10),
		SshConfig:  nil,
		Nodes:      nil,
	}
}

func GenerateNodePoolCli() NodePool {
	return NodePool{
		Name:       testutil.RandSeqPointer(10),
		NodeCount:  nil,
		ServerType: testutil.RandSeqPointer(10),
		SshConfig:  nil,
		Nodes:      nil,
	}
}

func GenerateNodeSdk() ranchersdk.Node {
	return ranchersdk.Node{
		ServerId: testutil.RandSeqPointer(10),
	}
}

func GenerateNodeCli() Node {
	return Node{
		ServerId: testutil.RandSeqPointer(10),
	}
}

func GenerateRancherClusterCertificatesSdk() ranchersdk.RancherClusterCertificates {
	return ranchersdk.RancherClusterCertificates{
		CaCertificate:  testutil.RandSeqPointer(10),
		Certificate:    testutil.RandSeqPointer(10),
		CertificateKey: testutil.RandSeqPointer(10),
	}
}

func GenerateRancherClusterCertificatesCli() RancherClusterCertificates {
	return RancherClusterCertificates{
		CaCertificate:  testutil.RandSeqPointer(10),
		Certificate:    testutil.RandSeqPointer(10),
		CertificateKey: testutil.RandSeqPointer(10),
	}
}

func GenerateRancherClusterConfigSdk() ranchersdk.RancherClusterConfig {
	return ranchersdk.RancherClusterConfig{
		Token:                    testutil.RandSeqPointer(10),
		TlsSan:                   testutil.RandSeqPointer(10),
		EtcdSnapshotScheduleCron: testutil.RandSeqPointer(10),
		EtcdSnapshotRetention:    nil,
		NodeTaint:                testutil.RandSeqPointer(10),
		ClusterDomain:            testutil.RandSeqPointer(10),
		Certificates:             nil,
	}
}

func GenerateRancherClusterConfigCli() RancherClusterConfig {
	return RancherClusterConfig{
		Token:                    testutil.RandSeqPointer(10),
		TlsSan:                   testutil.RandSeqPointer(10),
		EtcdSnapshotScheduleCron: testutil.RandSeqPointer(10),
		EtcdSnapshotRetention:    nil,
		NodeTaint:                testutil.RandSeqPointer(10),
		ClusterDomain:            testutil.RandSeqPointer(10),
		Certificates:             nil,
	}
}

func GenerateRancherServerMetadataSdk() ranchersdk.RancherServerMetadata {
	return ranchersdk.RancherServerMetadata{
		Url:      testutil.RandSeqPointer(10),
		Username: testutil.RandSeqPointer(10),
		Password: testutil.RandSeqPointer(10),
	}
}

func GenerateRancherServerMetadataCli() RancherServerMetadata {
	return RancherServerMetadata{
		Url:      testutil.RandSeqPointer(10),
		Username: testutil.RandSeqPointer(10),
		Password: testutil.RandSeqPointer(10),
	}
}

func GenerateSshConfigSdk() ranchersdk.SshConfig {
	return ranchersdk.SshConfig{
		InstallDefaultKeys: nil,
		Keys:               testutil.RandListStringPointer(3),
		KeyIds:             testutil.RandListStringPointer(3),
	}
}

func GenerateSshConfigCli() SshConfig {
	return SshConfig{
		InstallDefaultKeys: nil,
		Keys:               testutil.RandListStringPointer(3),
		KeyIds:             testutil.RandListStringPointer(3),
	}
}

func GenerateRancherDeleteResultSdk() ranchersdk.DeleteResult {
	return ranchersdk.DeleteResult{
		Result:    testutil.RandSeq(10),
		ClusterId: testutil.RandSeqPointer(10),
	}
}
