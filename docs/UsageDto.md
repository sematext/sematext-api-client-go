# UsageDto

## Properties

| Name                   | Type                                                | Description | Notes                        |
| ---------------------- | --------------------------------------------------- | ----------- | ---------------------------- |
| **DailyUsage**         | [**[]DailyDto**](DailyDto.md)                       |             | [optional] [default to null] |
| **DailyVolumeMB**      | **int64**                                           |             | [optional] [default to null] |
| **End**                | [**time.Time**](time.Time.md)                       |             | [optional] [default to null] |
| **FailedCount**        | **int64**                                           |             | [optional] [default to null] |
| **IngestedCount**      | **int64**                                           |             | [optional] [default to null] |
| **IngestedVolume**     | **int64**                                           |             | [optional] [default to null] |
| **LimitChangeEvents**  | [**[]LimitChangeEventDto**](LimitChangeEventDTO.md) |             | [optional] [default to null] |
| **MaxAllowedMB**       | **int64**                                           |             | [optional] [default to null] |
| **MaxLimitMB**         | **int64**                                           |             | [optional] [default to null] |
| **Start**              | [**time.Time**](time.Time.md)                       |             | [optional] [default to null] |
| **StoredCount**        | **int64**                                           |             | [optional] [default to null] |
| **StoredVolume**       | **int64**                                           |             | [optional] [default to null] |
| **VolumeChangeEvents** | [**[]LimitChangeEventDto**](LimitChangeEventDTO.md) |             | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
