package api

import "fmt"

// CloudWatchSettings TODO Doc Comment
type CloudWatchSettings struct {
	AwsCloudWatchAccessKey string `json:"accessKey"`
	AwsCloudWatchSecretKey string `json:"secretKey"`
	AwsFetchFrequency      string `json:"fetchFrequency"`   // TODO Enum:[ MINUTE, FIVE_MINUTES, FIFTEEN_MINUTES ]
	AwsRegion              string `json:"Region,omitempty"` // TODO Enum: [ US_EAST_1, US_WEST_1, EU_WEST_1, US_WEST_2, AP_SOUTHEAST_1, AP_SOUTHEAST_2, AP_NORTHEAST_1, SA_EAST_1, GovCloud, CN_NORTH_1, US_EAST_2, AP_SOUTH_1, AP_NORTHEAST_2, CA_CENTRAL_1, EU_CENTRAL_1, EU_WEST_2 ]
}

// Persist TODO Doc comment
func (cloudWatchSettings *CloudWatchSettings) Persist(id int, client *Client) (*CloudWatchSettings, error) {

	path := fmt.Sprintf("/users-web/api/v3/apps/%d/aws", id)
	genericAPIResponse, err := client.PutJSON(path, cloudWatchSettings)
	if err != nil {
		return nil, err
	}
	var cws *CloudWatchSettings
	if cws, err = genericAPIResponse.ExtractCloudWatchSettings(); err != nil {
		return nil, err
	}

	return cws, nil
}
