/*
 * Sematext Cloud API
 *
 * API Explorer provides access and documentation for Sematext REST API. The REST API requires the API Key to be sent as part of `Authorization` header. E.g.: `Authorization : apiKey e5f18450-205a-48eb-8589-7d49edaea813`.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stcloud

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes a oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// Configuration TODO godoc commenting
type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
}

// NewConfiguration TODO godoc commenting
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "https://localhost",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	return cfg
}

// AddDefaultHeader TODO godoc commenting
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

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

// IsValidSematextRegion checks sematext api region is valid.
func IsValidSematextRegion(region string) bool {
	switch region {
	case "EU", "US":
		return true
	}
	return false
}

// IsValidUUID checks a string is UUIDv4
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

// LookupPlanID2Apptypes is a ookup of planId to appType.
var LookupPlanID2Apptypes = map[int]string{
	121:   "Akka",
	122:   "Akka",
	123:   "Akka",
	124:   "Akka",
	81:    "Apache",
	82:    "Apache",
	83:    "Apache",
	84:    "Apache",
	73:    "AWS EBS",
	74:    "AWS EBS",
	75:    "AWS EBS",
	76:    "AWS EBS",
	65:    "AWS EC2",
	66:    "AWS EC2",
	67:    "AWS EC2",
	68:    "AWS EC2",
	69:    "AWS ELB",
	70:    "AWS ELB",
	71:    "AWS ELB",
	72:    "AWS ELB",
	57:    "Cassandra",
	58:    "Cassandra",
	59:    "Cassandra",
	60:    "Cassandra",
	130:   "ClickHouse",
	131:   "ClickHouse",
	132:   "ClickHouse",
	133:   "ClickHouse",
	12:    "Elastic Search",
	13:    "Elastic Search",
	14:    "Elastic Search",
	29:    "Elastic Search",
	21:    "Hadoop-MRv1",
	22:    "Hadoop-MRv1",
	23:    "Hadoop-MRv1",
	32:    "Hadoop-MRv1",
	24:    "Hadoop-YARN",
	25:    "Hadoop-YARN",
	26:    "Hadoop-YARN",
	33:    "Hadoop-YARN",
	105:   "HAProxy",
	106:   "HAProxy",
	107:   "HAProxy",
	108:   "HAProxy",
	4:     "HBase",
	5:     "HBase",
	6:     "HBase",
	28:    "HBase",
	129:   "Infra",
	300:   "Infra",
	301:   "Infra",
	302:   "Infra",
	18:    "JVM",
	19:    "JVM",
	20:    "JVM",
	31:    "JVM",
	93:    "Kafka",
	94:    "Kafka",
	95:    "Kafka",
	96:    "Kafka",
	125:   "MongoDB",
	126:   "MongoDB",
	127:   "MongoDB",
	128:   "MongoDB",
	77:    "MySQL",
	78:    "MySQL",
	79:    "MySQL",
	80:    "MySQL",
	85:    "Nginx",
	86:    "Nginx",
	87:    "Nginx",
	88:    "Nginx",
	89:    "Nginx-Plus",
	90:    "Nginx-Plus",
	91:    "Nginx-Plus",
	92:    "Nginx-Plus",
	109:   "Node.js",
	110:   "Node.js",
	111:   "Node.js",
	112:   "Node.js",
	49:    "Redis",
	50:    "Redis",
	51:    "Redis",
	52:    "Redis",
	15:    "Sensei",
	16:    "Sensei",
	17:    "Sensei",
	30:    "Sensei",
	1:     "Solr",
	2:     "Solr",
	3:     "Solr",
	27:    "Solr",
	45:    "SolrCloud",
	46:    "SolrCloud",
	47:    "SolrCloud",
	48:    "SolrCloud",
	97:    "Spark",
	98:    "Spark",
	99:    "Spark",
	100:   "Spark",
	53:    "Storm",
	54:    "Storm",
	55:    "Storm",
	56:    "Storm",
	113:   "Tomcat",
	114:   "Tomcat",
	115:   "Tomcat",
	116:   "Tomcat",
	41:    "ZooKeeper",
	42:    "ZooKeeper",
	43:    "ZooKeeper",
	44:    "ZooKeeper",
	10100: "Logsene",
	10101: "Logsene",
	10102: "Logsene",
	10103: "Logsene",
	10104: "Logsene",
	10105: "Logsene",
	10106: "Logsene",
	10107: "Logsene",
	10108: "Logsene",
	10109: "Logsene",
	10110: "Logsene",
	10111: "Logsene",
	10112: "Logsene",
	10113: "Logsene",
	10114: "Logsene",
	10115: "Logsene",
	10116: "Logsene",
	10117: "Logsene",
	10118: "Logsene",
	10119: "Logsene",
	10120: "Logsene",
	10121: "Logsene",
	10122: "Logsene",
	10123: "Logsene",
	10124: "Logsene",
	10125: "Logsene",
	10126: "Logsene",
	10127: "Logsene",
	10128: "Logsene",
	10129: "Logsene",
	10130: "Logsene",
	10131: "Logsene",
	10132: "Logsene",
	10133: "Logsene",
	10134: "Logsene",
	10135: "Logsene",
	10136: "Logsene",
	10137: "Logsene",
	10138: "Logsene",
	10139: "Logsene",
	10140: "Logsene",
	10141: "Logsene",
	10142: "Logsene",
	10143: "Logsene",
	10144: "Logsene",
	10145: "Logsene",
	10146: "Logsene",
	10147: "Logsene",
	10148: "Logsene",
	10149: "Logsene",
	10150: "Logsene",
	10151: "Logsene",
	10152: "Logsene",
	10153: "Logsene",
	10154: "Logsene",
	10155: "Logsene",
	10156: "Logsene",
	10157: "Logsene",
	10158: "Logsene",
	10159: "Logsene",
	10160: "Logsene",
	10161: "Logsene",
	10162: "Logsene",
	10163: "Logsene",
	10164: "Logsene",
	10165: "Logsene",
	10166: "Logsene",
	10167: "Logsene",
	10168: "Logsene",
	10169: "Logsene",
	10170: "Logsene",
	10171: "Logsene",
	10172: "Logsene",
	10173: "Logsene",
	10174: "Logsene",
	10175: "Logsene",
	10176: "Logsene",
	5004:  "postgresql",
	5005:  "postgresql",
	5006:  "postgresql",
	5007:  "postgresql",
	5008:  "rabbitmq",
	5009:  "rabbitmq",
	5010:  "rabbitmq",
	5011:  "rabbitmq",
	10533: "mobile-logs",
	10534: "mobile-logs",
	10535: "mobile-logs",
	10536: "mobile-logs",
	10537: "mobile-logs",
	10538: "mobile-logs",
	10539: "mobile-logs",
	10540: "mobile-logs",
	10541: "mobile-logs",
	10542: "mobile-logs",
	10543: "mobile-logs",
	10544: "mobile-logs",
	10545: "mobile-logs",
	10546: "mobile-logs",
	10547: "mobile-logs",
	10548: "mobile-logs",
	10549: "mobile-logs",
	10550: "mobile-logs",
	10551: "mobile-logs",
	10552: "mobile-logs",
	10553: "mobile-logs",
	10554: "mobile-logs",
	10555: "mobile-logs",
	10556: "mobile-logs",
	10557: "mobile-logs",
	10558: "mobile-logs",
	10559: "mobile-logs",
	10560: "mobile-logs",
	10561: "mobile-logs",
	10562: "mobile-logs",
	10563: "mobile-logs",
	10564: "mobile-logs",
	10565: "mobile-logs",
	10566: "mobile-logs",
	10567: "mobile-logs",
	10568: "mobile-logs",
	10569: "mobile-logs",
	10570: "mobile-logs",
	10571: "mobile-logs",
	10572: "mobile-logs",
	10573: "mobile-logs",
	10574: "mobile-logs",
	10575: "mobile-logs",
	10576: "mobile-logs",
	10577: "mobile-logs",
	10578: "mobile-logs",
	10579: "mobile-logs",
	10580: "mobile-logs",
	10581: "mobile-logs",
	10582: "mobile-logs",
	10583: "mobile-logs",
	10584: "mobile-logs",
	10585: "mobile-logs",
	10586: "mobile-logs",
	10587: "mobile-logs",
	10588: "mobile-logs",
	10589: "mobile-logs",
	10590: "mobile-logs",
	10591: "mobile-logs",
	10592: "mobile-logs",
	10593: "mobile-logs",
	10594: "mobile-logs",
	10595: "mobile-logs",
	10596: "mobile-logs",
	10597: "mobile-logs",
	10598: "mobile-logs",
	10599: "mobile-logs",
	10600: "mobile-logs",
	10601: "mobile-logs",
	10602: "mobile-logs",
	10603: "mobile-logs",
	10604: "mobile-logs",
	10605: "mobile-logs",
	10606: "mobile-logs",
	10607: "mobile-logs",
	10608: "mobile-logs",
	10609: "mobile-logs",
}

// LookupAppType2PlanID is a lookup of appType to [planId]
var LookupAppType2PlanID = map[string][]int{
	"Akka":           {121, 122, 123, 124},
	"Apache":         {81, 82, 83, 84},
	"AWS EBS":        {73, 74, 75, 76},
	"AWS EC2":        {65, 66, 67, 68},
	"AWS ELB":        {69, 70, 71, 72},
	"Cassandra":      {57, 58, 59, 60},
	"ClickHouse":     {130, 131, 132, 133},
	"Elastic Search": {12, 13, 14, 29},
	"Hadoop-MRv1":    {21, 22, 23, 32},
	"Hadoop-YARN":    {24, 25, 26, 33},
	"HAProxy":        {105, 106, 107, 108},
	"HBase":          {4, 5, 6, 28},
	"Infra":          {129, 300, 301, 302},
	"JVM":            {18, 19, 20, 31},
	"Kafka":          {93, 94, 95, 96},
	"MongoDB":        {125, 126, 127, 128},
	"MySQL":          {77, 78, 79, 80},
	"Nginx":          {85, 86, 87, 88},
	"Nginx-Plus":     {89, 90, 91, 92},
	"Node.js":        {109, 110, 111, 112},
	"Redis":          {49, 50, 51, 52},
	"Sensei":         {15, 16, 17, 30},
	"Solr":           {1, 2, 3, 27},
	"SolrCloud":      {45, 46, 47, 48},
	"Spark":          {97, 98, 99, 100},
	"Storm":          {53, 54, 55, 56},
	"Tomcat":         {113, 114, 115, 116},
	"ZooKeeper":      {41, 42, 43, 44},
	"Logsene": {
		10100,
		10101,
		10102,
		10103,
		10104,
		10105,
		10106,
		10107,
		10108,
		10109,
		10110,
		10111,
		10112,
		10113,
		10114,
		10115,
		10116,
		10117,
		10118,
		10119,
		10120,
		10121,
		10122,
		10123,
		10124,
		10125,
		10126,
		10127,
		10128,
		10129,
		10130,
		10131,
		10132,
		10133,
		10134,
		10135,
		10136,
		10137,
		10138,
		10139,
		10140,
		10141,
		10142,
		10143,
		10144,
		10145,
		10146,
		10147,
		10148,
		10149,
		10150,
		10151,
		10152,
		10153,
		10154,
		10155,
		10156,
		10157,
		10158,
		10159,
		10160,
		10161,
		10162,
		10163,
		10164,
		10165,
		10166,
		10167,
		10168,
		10169,
		10170,
		10171,
		10172,
		10173,
		10174,
		10175,
		10176,
	},
	"postgresql": {
		5004,
		5005,
		5006,
		5007,
	},
	"rabbitmq": {
		5008,
		5009,
		5010,
		5011,
	},
	"mobile-logs": {
		10533,
		10534,
		10535,
		10536,
		10537,
		10538,
		10539,
		10540,
		10541,
		10542,
		10543,
		10544,
		10545,
		10546,
		10547,
		10548,
		10549,
		10550,
		10551,
		10552,
		10553,
		10554,
		10555,
		10556,
		10557,
		10558,
		10559,
		10560,
		10561,
		10562,
		10563,
		10564,
		10565,
		10566,
		10567,
		10568,
		10569,
		10570,
		10571,
		10572,
		10573,
		10574,
		10575,
		10576,
		10577,
		10578,
		10579,
		10580,
		10581,
		10582,
		10583,
		10584,
		10585,
		10586,
		10587,
		10588,
		10589,
		10590,
		10591,
		10592,
		10593,
		10594,
		10595,
		10596,
		10597,
		10598,
		10599,
		10600,
		10601,
		10602,
		10603,
		10604,
		10605,
		10606,
		10607,
		10608,
		10609,
	},
}

// AssignPlanID provides a ramdom PlanID for testing. - TODO rest to only use free plans
func AssignPlanID(appType string) int {
	if plans, ok := LookupAppType2PlanID[appType]; ok {
		if appType == "Logsene" {
			return LookupAppType2PlanID[appType][0]
		}
		rand.Seed(time.Now().Unix())
		n := rand.Intn(len(plans))
		return LookupAppType2PlanID[appType][n]
	}
	panic("CODE ERROR - bad appType.")
}
