package api

// TestDiscountCodeMetrics is used in auto test mocks.
const TestDiscountCodeMetrics = "TERRAFORM__METRICS_00"

// TestDiscountCodeLogs is used in auto test mocks.
const TestDiscountCodeLogs = "TERRAFORM__LOGS_00"

// TestDiscountCodeRum is used in auto test mocks.
const TestDiscountCodeRum = "TERRAFORM__RUM_00"

// STRegion2AWSRegion is a lookup of ST AWS_REGION => aws_region
var STRegion2AWSRegion = map[string]string{
	"US_EAST_1":      "us-east-1",
	"US_WEST_1":      "us-west-1",
	"EU_WEST_1":      "eu-west-1",
	"US_WEST_2":      "us-west-2",
	"AP_SOUTHEAST_1": "ap-southeast-1",
	"AP_SOUTHEAST_2": "ap-southeast-2",
	"AP_NORTHEAST_1": "ap-northeast-1",
	"SA_EAST_1":      "sa-east-1",
	"GovCloud":       "us-gov-west-1",
	"CN_NORTH_1":     "cn-north-1",
	"US_EAST_2":      "us-east-2",
	"AP_SOUTH_1":     "ap-south-1",
	"AP_NORTHEAST_2": "ap-northeast-2",
	"CA_CENTRAL_1":   "ca-central-1",
	"EU_CENTRAL_1":   "eu-central-1",
	"EU_WEST_2":      "eu-west-2",
}

// AWSRegion2STRegion is a lookup of aws_region => ST AWS_REGION
var AWSRegion2STRegion = map[string]string{
	"us-east-1":      "US_EAST_1",
	"us-west-1":      "US_WEST_1",
	"eu-west-1":      "EU_WEST_1",
	"us-west-2":      "US_WEST_2",
	"ap-southeast-1": "AP_SOUTHEAST_1",
	"ap-southeast-2": "AP_SOUTHEAST_2",
	"ap-northeast-1": "AP_NORTHEAST_1",
	"sa-east-1":      "SA_EAST_1",
	"us-gov-west-1":  "GovCloud",
	"cn-north-1":     "CN_NORTH_1",
	"us-east-2":      "US_EAST_2",
	"ap-south-1":     "AP_SOUTH_1",
	"ap-northeast-2": "AP_NORTHEAST_2",
	"ca-central-1":   "CA_CENTRAL_1",
	"eu-central-1":   "EU_CENTRAL_1",
	"eu-west-2":      "EU_WEST_2",
}
