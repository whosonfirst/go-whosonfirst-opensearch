// Code generated by endpoint/awsrulesfn/internal/partition. DO NOT EDIT.

package awsrulesfn

// GetPartition returns an AWS [Partition] for the region provided. If the
// partition cannot be determined nil will be returned.
func GetPartition(region string) *PartitionConfig {
	return getPartition(partitions, region)
}

var partitions = []Partition{
	{
		ID:          "aws",
		RegionRegex: "^(us|eu|ap|sa|ca|me|af)\\-\\w+\\-\\d+$",
		DefaultConfig: PartitionConfig{
			Name:               "aws",
			DnsSuffix:          "amazonaws.com",
			DualStackDnsSuffix: "api.aws",
			SupportsFIPS:       true,
			SupportsDualStack:  true,
		},
		Regions: map[string]RegionOverrides{
			"af-south-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"ap-east-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"ap-northeast-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"ap-northeast-2": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"ap-northeast-3": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"ap-south-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"ap-south-2": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"ap-southeast-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"ap-southeast-2": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"ap-southeast-3": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"aws-global": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"ca-central-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"eu-central-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"eu-central-2": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"eu-north-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"eu-south-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"eu-south-2": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"eu-west-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"eu-west-2": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"eu-west-3": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"me-central-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"me-south-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"sa-east-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"us-east-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"us-east-2": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"us-west-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"us-west-2": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
		},
	},
	{
		ID:          "aws-cn",
		RegionRegex: "^cn\\-\\w+\\-\\d+$",
		DefaultConfig: PartitionConfig{
			Name:               "aws-cn",
			DnsSuffix:          "amazonaws.com.cn",
			DualStackDnsSuffix: "api.amazonwebservices.com.cn",
			SupportsFIPS:       true,
			SupportsDualStack:  true,
		},
		Regions: map[string]RegionOverrides{
			"aws-cn-global": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"cn-north-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"cn-northwest-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
		},
	},
	{
		ID:          "aws-us-gov",
		RegionRegex: "^us\\-gov\\-\\w+\\-\\d+$",
		DefaultConfig: PartitionConfig{
			Name:               "aws-us-gov",
			DnsSuffix:          "amazonaws.com",
			DualStackDnsSuffix: "api.aws",
			SupportsFIPS:       true,
			SupportsDualStack:  true,
		},
		Regions: map[string]RegionOverrides{
			"aws-us-gov-global": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"us-gov-east-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"us-gov-west-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
		},
	},
	{
		ID:          "aws-iso",
		RegionRegex: "^us\\-iso\\-\\w+\\-\\d+$",
		DefaultConfig: PartitionConfig{
			Name:               "aws-iso",
			DnsSuffix:          "c2s.ic.gov",
			DualStackDnsSuffix: "c2s.ic.gov",
			SupportsFIPS:       true,
			SupportsDualStack:  false,
		},
		Regions: map[string]RegionOverrides{
			"aws-iso-global": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"us-iso-east-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"us-iso-west-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
		},
	},
	{
		ID:          "aws-iso-b",
		RegionRegex: "^us\\-isob\\-\\w+\\-\\d+$",
		DefaultConfig: PartitionConfig{
			Name:               "aws-iso-b",
			DnsSuffix:          "sc2s.sgov.gov",
			DualStackDnsSuffix: "sc2s.sgov.gov",
			SupportsFIPS:       true,
			SupportsDualStack:  false,
		},
		Regions: map[string]RegionOverrides{
			"aws-iso-b-global": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
			"us-isob-east-1": {
				Name:               nil,
				DnsSuffix:          nil,
				DualStackDnsSuffix: nil,
				SupportsFIPS:       nil,
				SupportsDualStack:  nil,
			},
		},
	},
}
