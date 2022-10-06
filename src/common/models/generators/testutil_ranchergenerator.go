package generators

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
		NodePools:             []ranchersdk.NodePool{},
		Configuration:         nil,
		Metadata:              nil,
		WorkloadConfiguration: nil,
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

func GenerateNodeSdk() ranchersdk.Node {
	return ranchersdk.Node{
		ServerId: testutil.RandSeqPointer(10),
	}
}

func GenerateRancherClusterCertificatesSdk() ranchersdk.RancherClusterConfigCertificates {
	return ranchersdk.RancherClusterConfigCertificates{
		CaCertificate:  testutil.RandSeqPointer(10),
		Certificate:    testutil.RandSeqPointer(10),
		CertificateKey: testutil.RandSeqPointer(10),
	}
}

func GenerateRancherClusterConfigSdk() ranchersdk.ClusterConfiguration {
	return ranchersdk.ClusterConfiguration{
		Token:                    testutil.RandSeqPointer(10),
		TlsSan:                   testutil.RandSeqPointer(10),
		EtcdSnapshotScheduleCron: testutil.RandSeqPointer(10),
		EtcdSnapshotRetention:    nil,
		NodeTaint:                testutil.RandSeqPointer(10),
		ClusterDomain:            testutil.RandSeqPointer(10),
		Certificates:             nil,
	}
}

func GenerateRancherServerMetadataSdk() ranchersdk.ClusterMetadata {
	return ranchersdk.ClusterMetadata{
		Url:      testutil.RandSeqPointer(10),
		Username: testutil.RandSeqPointer(10),
		Password: testutil.RandSeqPointer(10),
	}
}

func GenerateSshConfigSdk() ranchersdk.NodePoolSshConfig {
	return ranchersdk.NodePoolSshConfig{
		InstallDefaultKeys: nil,
		Keys:               testutil.RandListStringPointer(3),
		KeyIds:             testutil.RandListStringPointer(3),
	}
}

func GenerateRancherDeleteResultSdk() *ranchersdk.DeleteResult {
	return &ranchersdk.DeleteResult{
		Result:    testutil.RandSeq(10),
		ClusterId: testutil.RandSeq(10),
	}
}
