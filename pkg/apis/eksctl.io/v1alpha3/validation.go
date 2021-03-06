package v1alpha3

import (
	"fmt"
)

// ValidateNodeGroup checks compatible fileds of a given nodegroup
func ValidateNodeGroup(i int, ng *NodeGroup) error {
	path := fmt.Errorf("nodegroups[%d]", i)
	if ng.Name == "" {
		return fmt.Errorf("%s.name must be set", path)
	}

	if ng.IAM.InstanceRoleARN != "" {
		p := fmt.Sprintf("%s.iam.instanceRoleARN and %s.iam", path, path)
		if ng.IAM.InstanceRoleName != "" {
			return fmt.Errorf("%s.instanceRoleName cannot be set at the same time", p)
		}
		if len(ng.IAM.AttachPolicyARNs) != 0 {
			return fmt.Errorf("%s.attachPolicyARNs cannot be set at the same time", p)
		}
		if ng.IAM.WithAddonPolicies.AutoScaler {
			return fmt.Errorf("%s.withAddonPolicies.autoScaler cannot be set at the same time", p)
		}
		if ng.IAM.WithAddonPolicies.ExternalDNS {
			return fmt.Errorf("%s.withAddonPolicies.externalDNS cannot be set at the same time", p)
		}
		if ng.IAM.WithAddonPolicies.ImageBuilder {
			return fmt.Errorf("%s.imageBuilder cannot be set at the same time", p)
		}
	}

	return nil
}
