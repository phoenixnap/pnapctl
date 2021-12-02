package tagmodels

import (
	"fmt"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
)

type ResourceAssignment struct {
	ResourceName string  `json:"resourceName" yaml:"resourceName"`
	Value        *string `json:"value" yaml:"value"`
}

func ResourceAssignmentFromSdk(assignment *tagapisdk.ResourceAssignment) *ResourceAssignment {
	if assignment == nil {
		return nil
	}

	return &ResourceAssignment{
		ResourceName: assignment.ResourceName,
		Value:        assignment.Value,
	}
}

func ResourceAssignmentToTableStrings(resourceAssignment tagapisdk.ResourceAssignment) string {
	var value string

	if resourceAssignment.Value == nil {
		value = "N/A"
	} else {
		value = *resourceAssignment.Value
	}

	return fmt.Sprintf("%s : %s", resourceAssignment.ResourceName, value)
}
