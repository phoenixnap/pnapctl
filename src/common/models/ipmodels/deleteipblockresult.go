package ipmodels

import ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi"

type DeleteIpBlockResult struct {
	Result    string `yaml:"result" json:"result"`
	IpBlockId string `yaml:"ipBlockId" json:"ipBlockId"`
}

func DeleteIpBlockResultFromSdk(result ipapisdk.DeleteIpBlockResult) DeleteIpBlockResult {
	return DeleteIpBlockResult{
		Result:    result.Result,
		IpBlockId: result.IpBlockId,
	}
}
