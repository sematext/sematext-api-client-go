
# <img src="https://sematext.com/wp-content/uploads/2020/09/just-octi-blue.png" valign="bottom" width="60px"/>**&nbsp;&nbsp;sematext-api-client-go**

<br>

>*A [Sematext Cloud](https://sematext.com/cloud/) API client, for interaction with Sematext Cloud solution monitoring, alerting and log shipping.*

<br>

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

<br>
One of a family of clients in following flavours:
<br>
<br>

* [sematext-api-client-javascript](https://github.com/sematext/sematext-api-client-javascript "Javascript")
* [sematext-api-client-rust](https://github.com/sematext/sematext-api-client-rust "Rust")
* [sematext-api-client-ruby](https://github.com/sematext/sematext-api-client-ruby "Ruby")
* [sematext-api-client-python](https://github.com/sematext/sematext-api-client-python "Python")
* [sematext-api-client-php](https://github.com/sematext/sematext-api-client-php "PHP")
* [sematext-api-client-java](https://github.com/sematext/sematext-api-client-java "Java")
* [sematext-api-client-go](https://github.com/sematext/sematext-api-client-go "Go/Golang")

<br>
Refer to below link for deeper information on the API itself.
<br>
<br>

* [Sematext Cloud API Reference](https://sematext.com/docs/api/ "API Reference")

<br>

## Contents

- [<img src="https://sematext.com/wp-content/uploads/2020/09/just-octi-blue.png" valign="bottom" width="60px"/>**&nbsp;&nbsp;sematext-api-client-go**](#sematext-api-client-go)
  - [Contents](#contents)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
  - [Authentication](#authentication)
  - [Versioning](#versioning)
  - [License](#license)
  - [Acknowledgements](#acknowledgements)
  - [Reference](#reference)
  - [Documentation For Models](#documentation-for-models)

<br>

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

* [Go](https://golang.org/) - v1.13.8


### Installation

Put the package under your project folder and add the following in import:
```golang
import "./stcloud"
```

## Authentication

This client code requires a Sematext API Access token to function. You can find this by logging into your [Sematext Cloud Account](https://apps.sematext.com/ui/account/api)


## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).


## License

This project is licensed under the Apache License - see the [LICENSE](./LICENSE) file for details


## Acknowledgements

This API client was initially generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.

- API version: v3
- Package version: 1.0.0


## Reference

All URIs are relative to *https://localhost*

| Class                      | Method                                                                                                           | HTTP request                                                        | Description                                                                                             |
| -------------------------- | ---------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| *AlertNotificationsAPI*    | [**GetAlertNotificationsForAppUsingPOST**](docs/AlertNotificationsAPI.md#getalertnotificationsforappusingpost)   | **Post** /users-web/api/v3/apps/{appID}/notifications/alerts        | Get alert notifications for an app                                                                      |
| *AlertNotificationsAPI*    | [**GetAlertNotificationsForUserUsingPOST**](docs/AlertNotificationsAPI.md#getalertnotificationsforuserusingpost) | **Post** /users-web/api/v3/notifications/alerts                     | Get alert notifications for a user                                                                      |
| *AlertsAPI*                | [**CreateAlertUsingPOST**](docs/AlertsAPI.md#createalertusingpost)                                               | **Post** /users-web/api/v3/alerts                                   | Create alert rule                                                                                       |
| *AlertsAPI*                | [**DeleteAlertRuleUsingDELETE**](docs/AlertsAPI.md#deletealertruleusingdelete)                                   | **Delete** /users-web/api/v3/alerts/{updateableAlertID}             | Delete alert rule                                                                                       |
| *AlertsAPI*                | [**DisableAlertRuleUsingPUT**](docs/AlertsAPI.md#disablealertruleusingput)                                       | **Put** /users-web/api/v3/alerts/{updateableAlertID}/disable        | Disable alert rule                                                                                      |
| *AlertsAPI*                | [**EnableAlertRuleUsingPUT**](docs/AlertsAPI.md#enablealertruleusingput)                                         | **Put** /users-web/api/v3/alerts/{updateableAlertID}/enable         | Enable alert rule                                                                                       |
| *AlertsAPI*                | [**GetAlertRulesForAppUsingGET**](docs/AlertsAPI.md#getalertrulesforappusingget)                                 | **Get** /users-web/api/v3/apps/{appID}/alerts                       | Get alert rules for an app                                                                              |
| *AppsAPI*                  | [**GetAppTypesUsingGET**](docs/AppsAPI.md#getapptypesusingget)                                                   | **Get** /users-web/api/v3/apps/types                                | Get all App types supported for the account identified with apiKey                                      |
| *AppsAPI*                  | [**GetUsingGET**](docs/AppsAPI.md#getusingget)                                                                   | **Get** /users-web/api/v3/apps/{anyStateAppID}                      | Gets defails for one particular App                                                                     |
| *AppsAPI*                  | [**InviteAppGuestsUsingPOST**](docs/AppsAPI.md#inviteappguestsusingpost)                                         | **Post** /users-web/api/v3/apps/guests                              | Invite guests to an app                                                                                 |
| *AppsAPI*                  | [**ListAppsUsersUsingGET**](docs/AppsAPI.md#listappsusersusingget)                                               | **Get** /users-web/api/v3/apps/users                                | Get all users of apps accessible to this account                                                        |
| *AppsAPI*                  | [**ListUsingGET**](docs/AppsAPI.md#listusingget)                                                                 | **Get** /users-web/api/v3/apps                                      | Get all apps accessible by account identified with apiKey                                               |
| *AppsAPI*                  | [**UpdateDescriptionUsingPUT**](docs/AppsAPI.md#updatedescriptionusingput)                                       | **Put** /users-web/api/v3/apps/{anyStateAppID}/description          | Update description of the app                                                                           |
| *AppsAPI*                  | [**UpdateUsingPUT1**](docs/AppsAPI.md#updateusingput1)                                                           | **Put** /users-web/api/v3/apps/{anyStateAppID}                      | Update app                                                                                              |
| *AwsSettingsControllerAPI* | [**UpdateUsingPUT**](docs/AwsSettingsControllerAPI.md#updateusingput)                                            | **Put** /users-web/api/v3/apps/{appID}/aws                          | Update App&#39;s AWS CloudWatch settings                                                                |
| *BillingAPI*               | [**GetDetailedInvoiceUsingGET**](docs/BillingAPI.md#getdetailedinvoiceusingget)                                  | **Get** /users-web/api/v3/billing/invoice/{service}/{year}/{month}  | Get invoice details                                                                                     |
| *BillingAPI*               | [**ListAvailablePlansUsingGET**](docs/BillingAPI.md#listavailableplansusingget)                                  | **Get** /users-web/api/v3/billing/availablePlans                    | Get available plans                                                                                     |
| *BillingAPI*               | [**UpdatePlanUsingPUT**](docs/BillingAPI.md#updateplanusingput)                                                  | **Put** /users-web/api/v3/billing/info/{appID}                      | Update plan for an app                                                                                  |
| *LogsAppAPI*               | [**CreateLogseneApplication**](docs/LogsAppAPI.md#createlogseneapplication)                                      | **Post** /logsene-reports/api/v3/apps                               | Create Logs App                                                                                         |
| *MetricsAPI*               | [**ListDataSeriesUsingPOST**](docs/MetricsAPI.md#listdataseriesusingpost)                                        | **Post** /spm-reports/api/v3/apps/{appID}/metrics/data              | Get metrics data points for an app                                                                      |
| *MetricsAPI*               | [**ListFiltersUsingPOST1**](docs/MetricsAPI.md#listfiltersusingpost1)                                            | **Post** /spm-reports/api/v3/apps/{appID}/metrics/filters           | Get metrics filters and their values for an app                                                         |
| *MetricsAPI*               | [**ListMetricsKeysUsingGET**](docs/MetricsAPI.md#listmetricskeysusingget)                                        | **Get** /spm-reports/api/v3/apps/{appID}/metrics/keys               | Get metrics keys for an app                                                                             |
| *MetricsAPI*               | [**ListMetricsUsingGET**](docs/MetricsAPI.md#listmetricsusingget)                                                | **Get** /spm-reports/api/v3/apps/{appID}/metrics                    | Get metrics info for an app                                                                             |
| *MonitoringAppAPI*         | [**CreateSpmApplication1**](docs/MonitoringAppAPI.md#createspmapplication1)                                      | **Post** /spm-reports/api/v3/apps                                   | Create Monitoring App                                                                                   |
| *ResetPasswordAPI*         | [**ResetPasswordUsingPOST**](docs/ResetPasswordAPI.md#resetpasswordusingpost)                                    | **Post** /users-web/api/v3/account/password/reset                   | Reset Password                                                                                          |
| *SavedQueriesAPI*          | [**DeleteSavedQueryUsingDELETE**](docs/SavedQueriesAPI.md#deletesavedqueryusingdelete)                           | **Delete** /users-web/api/v3/savedQueries/{updateableQueryID}       | Delete saved query                                                                                      |
| *SavedQueriesAPI*          | [**GetSavedQueriesForAppUsingGET**](docs/SavedQueriesAPI.md#getsavedqueriesforappusingget)                       | **Get** /users-web/api/v3/apps/{appID}/savedQueries                 | Get saved queries for an app                                                                            |
| *SavedQueriesAPI*          | [**SaveQueryUsingPOST**](docs/SavedQueriesAPI.md#savequeryusingpost)                                             | **Post** /users-web/api/v3/savedQueries                             | Create saved query                                                                                      |
| *SavedQueriesAPI*          | [**SaveQueryUsingPUT**](docs/SavedQueriesAPI.md#savequeryusingput)                                               | **Put** /users-web/api/v3/savedQueries/{updateableQueryID}          | Update saved query                                                                                      |
| *SubscriptionsAPI*         | [**ListUsingGET1**](docs/SubscriptionsAPI.md#listusingget1)                                                      | **Get** /users-web/api/v3/apps/{appID}/subscriptions                | Get subscriptions for an app                                                                            |
| *SubscriptionsAPI*         | [**SendReportUsingPOST**](docs/SubscriptionsAPI.md#sendreportusingpost)                                          | **Post** /users-web/api/v3/apps/{appID}/report/send                 | Trigger emailing of report for an app                                                                   |
| *TagAPIControllerAPI*      | [**GetTagNamesUsingGET1**](docs/TagAPIControllerAPI.md#gettagnamesusingget1)                                     | **Get** /spm-reports/api/v3/apps/{addIDs}/tagNames                  | Gets tag names for the given application identifiers appearing in the given time frame.                 |
| *TagAPIControllerAPI*      | [**GetUsingGET1**](docs/TagAPIControllerAPI.md#getusingget1)                                                     | **Get** /spm-reports/api/v3/apps/{addIDs}/tags                      | Gets values for specified tags for the given application identifiers appearing in the given time frame. |
| *TagAPIControllerAPI*      | [**GetUsingGET2**](docs/TagAPIControllerAPI.md#getusingget2)                                                     | **Get** /spm-reports/api/v3/apps/{addIDs}/metrics/filters           | Gets values for specified tags for the given application identifiers appearing in the given time frame. |
| *TokensAPIControllerAPI*   | [**CreateAppToken**](docs/TokensAPIControllerAPI.md#createapptoken)                                              | **Post** /users-web/api/v3/apps/{appID}/tokens                      | Create new app token                                                                                    |
| *TokensAPIControllerAPI*   | [**DeleteAppToken1**](docs/TokensAPIControllerAPI.md#deleteapptoken1)                                            | **Delete** /users-web/api/v3/apps/{appID}/tokens/{tokenID}          | Delete app token                                                                                        |
| *TokensAPIControllerAPI*   | [**GetAppTokens1**](docs/TokensAPIControllerAPI.md#getapptokens1)                                                | **Get** /users-web/api/v3/apps/{appID}/tokens                       | Get app available tokens                                                                                |
| *TokensAPIControllerAPI*   | [**RegenerateAppToken**](docs/TokensAPIControllerAPI.md#regenerateapptoken)                                      | **Post** /users-web/api/v3/apps/{appID}/tokens/{tokenID}/regenerate | Regenerate app token)                                                                                   |
| *TokensAPIControllerAPI*   | [**UpdateAppToken**](docs/TokensAPIControllerAPI.md#updateapptoken)                                              | **Put** /users-web/api/v3/apps/{appID}/tokens/{tokenID}             | Update app token (enable/disable)                                                                       |



## Documentation For Models

 - [AlertNotificationRequest](docs/AlertNotificationRequest.md)
 - [AlertRule](docs/AlertRule.md)
 - [AlertRuleScheduleTimeRangeDto](docs/AlertRuleScheduleTimeRangeDto.md)
 - [AlertRuleScheduleWeekdayDto](docs/AlertRuleScheduleWeekdayDto.md)
 - [App](docs/App.md)
 - [AppDescription](docs/AppDescription.md)
 - [AppMetadata](docs/AppMetadata.md)
 - [BasicAuthMethodDto](docs/BasicAuthMethodDto.md)
 - [BasicOrganizationDto](docs/BasicOrganizationDto.md)
 - [BillingInfo](docs/BillingInfo.md)
 - [CloudWatchSettings](docs/CloudWatchSettings.md)
 - [CreateAppInfo](docs/CreateAppInfo.md)
 - [DataSeriesFilter](docs/DataSeriesFilter.md)
 - [DataSeriesRequest](docs/DataSeriesRequest.md)
 - [FilterValue](docs/FilterValue.md)
 - [GenericAPIResponse](docs/GenericAPIResponse.md)
 - [Invitation](docs/Invitation.md)
 - [ModelError](docs/ModelError.md)
 - [NotificationIntegration](docs/NotificationIntegration.md)
 - [Plan](docs/Plan.md)
 - [ReportInfo](docs/ReportInfo.md)
 - [SavedQuery](docs/SavedQuery.md)
 - [ServiceIntegration](docs/ServiceIntegration.md)
 - [UpdateAppInfo](docs/UpdateAppInfo.md)
 - [UserInfo](docs/UserInfo.md)
 - [UserPermissions](docs/UserPermissions.md)
 - [UserRole](docs/UserRole.md)
