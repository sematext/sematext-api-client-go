package api

import (
	"fmt"
	"math/rand"
	"time"
)

// Plan TODO Doc Comment
type Plan struct {
	AppType            string  `json:"appType"`
	Custom             bool    `json:"custom"`
	DataRetentionHours float32 `json:"dataRetentionHours"` // TODO Check datatype
	DefaultTrialPlan   bool    `json:"defaultTrialPlan"`
	Free               bool    `json:"free"`
	FreeTrialDays      int64   `json:"freeTrialDays"`
	ID                 int64   `json:"id"`
	MaxAlerts          int64   `json:"maxAlerts"`
	MaxDailyEvents     int64   `json:"maxDailyEvents"`
	Name               string  `json:"name"`
	PlanScheme         string  `json:"planScheme"`      // TODO Enum: [ SPM_1_0, SPM_2_0, SA_1_0, SEARCHENE_1_0, LOGSENE_1_0, LOGSENE_2_0, RUM_1_0, RUM_EA, SYNTHETICS_1_0 ]
	SematextService    string  `json:"sematextService"` // TODO Enum: [ LOGSENE, SPM, RUM, SYNTHETICS ]
	TrialPlan          bool    `json:"trialPlan"`
}

// LookupPlanID2Apptypes is a ookup of planId to appType.
var LookupPlanID2Apptypes = map[int]string{
	121: "Akka",
	122: "Akka",
	123: "Akka",
	124: "Akka",
	81:  "Apache",
	82:  "Apache",
	83:  "Apache",
	84:  "Apache",
	73:  "AWS EBS",
	74:  "AWS EBS",
	75:  "AWS EBS",
	76:  "AWS EBS",
	65:  "AWS EC2",
	66:  "AWS EC2",
	67:  "AWS EC2",
	68:  "AWS EC2",
	69:  "AWS ELB",
	70:  "AWS ELB",
	71:  "AWS ELB",
	72:  "AWS ELB",
	57:  "Cassandra",
	58:  "Cassandra",
	59:  "Cassandra",
	60:  "Cassandra",
	130: "ClickHouse",
	131: "ClickHouse",
	132: "ClickHouse",
	133: "ClickHouse",
	117: "Docker",
	118: "Docker",
	119: "Docker",
	120: "Docker",
	12:  "Elastic Search",
	13:  "Elastic Search",
	14:  "Elastic Search",
	29:  "Elastic Search",
	21:  "Hadoop-MRv1",
	22:  "Hadoop-MRv1",
	23:  "Hadoop-MRv1",
	32:  "Hadoop-MRv1",
	24:  "Hadoop-YARN",
	25:  "Hadoop-YARN",
	26:  "Hadoop-YARN",
	33:  "Hadoop-YARN",
	105: "HAProxy",
	106: "HAProxy",
	107: "HAProxy",
	108: "HAProxy",
	4:   "HBase",
	5:   "HBase",
	6:   "HBase",
	28:  "HBase",
	129: "Infra",
	18:  "JVM",
	19:  "JVM",
	20:  "JVM",
	31:  "JVM",
	93:  "Kafka",
	94:  "Kafka",
	95:  "Kafka",
	96:  "Kafka",
	125: "MongoDB",
	126: "MongoDB",
	127: "MongoDB",
	128: "MongoDB",
	77:  "MySQL",
	78:  "MySQL",
	79:  "MySQL",
	80:  "MySQL",
	85:  "Nginx",
	86:  "Nginx",
	87:  "Nginx",
	88:  "Nginx",
	89:  "Nginx-Plus",
	90:  "Nginx-Plus",
	91:  "Nginx-Plus",
	92:  "Nginx-Plus",
	109: "Node.js",
	110: "Node.js",
	111: "Node.js",
	112: "Node.js",
	49:  "Redis",
	50:  "Redis",
	51:  "Redis",
	52:  "Redis",
	15:  "Sensei",
	16:  "Sensei",
	17:  "Sensei",
	30:  "Sensei",
	1:   "Solr",
	2:   "Solr",
	3:   "Solr",
	27:  "Solr",
	45:  "SolrCloud",
	46:  "SolrCloud",
	47:  "SolrCloud",
	48:  "SolrCloud",
	97:  "Spark",
	98:  "Spark",
	99:  "Spark",
	100: "Spark",
	53:  "Storm",
	54:  "Storm",
	55:  "Storm",
	56:  "Storm",
	113: "Tomcat",
	114: "Tomcat",
	115: "Tomcat",
	116: "Tomcat",
	41:  "ZooKeeper",
	42:  "ZooKeeper",
	43:  "ZooKeeper",
	44:  "ZooKeeper",
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
	"Docker":         {117, 118, 119, 120},
	"Elastic Search": {12, 13, 14, 29},
	"Hadoop-MRv1":    {21, 22, 23, 32},
	"Hadoop-YARN":    {24, 25, 26, 33},
	"HAProxy":        {105, 106, 107, 108},
	"HBase":          {4, 5, 6, 28},
	"Infra":          {129},
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
}

// AssignPlanID provides a ramdom PlanID for testing.
func AssignPlanID(appType string) int {
	fmt.Println("------------------------------")
	fmt.Println(appType)
	fmt.Println("------------------------------")
	if plans, ok := LookupAppType2PlanID[appType]; ok {
		rand.Seed(time.Now().Unix())
		n := rand.Intn(len(plans))
		return LookupAppType2PlanID[appType][n]
	}
	panic("CODE ERROR - bad appType.")
}
