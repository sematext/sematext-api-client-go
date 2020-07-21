
# <img src="./assets/octi-blue.png" valign="bottom" width="60px"/>**&nbsp;&nbsp;sematext-api-client-go**

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

- [&nbsp;&nbsp;sematext-api-client-go](#sematext-api-client-go)
  - [Contents](#contents)
  - [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Authentication](#authentication)
  - [Reference](#reference)
  - [Documentation for Models](#documentation-for-models)

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

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AlertNotificationsApi* | [**GetAlertNotificationsForAppUsingPOST**](docs/AlertNotificationsApi.md#getalertnotificationsforappusingpost) | **Post** /users-web/api/v3/apps/{appId}/notifications/alerts | Get alert notifications for an app
*AlertNotificationsApi* | [**GetAlertNotificationsForUserUsingPOST**](docs/AlertNotificationsApi.md#getalertnotificationsforuserusingpost) | **Post** /users-web/api/v3/notifications/alerts | Get alert notifications for a user
*AlertsApi* | [**CreateAlertUsingPOST**](docs/AlertsApi.md#createalertusingpost) | **Post** /users-web/api/v3/alerts | Create alert rule
*AlertsApi* | [**DeleteAlertRuleUsingDELETE**](docs/AlertsApi.md#deletealertruleusingdelete) | **Delete** /users-web/api/v3/alerts/{updateableAlertId} | Delete alert rule
*AlertsApi* | [**DisableAlertRuleUsingPUT**](docs/AlertsApi.md#disablealertruleusingput) | **Put** /users-web/api/v3/alerts/{updateableAlertId}/disable | Disable alert rule
*AlertsApi* | [**EnableAlertRuleUsingPUT**](docs/AlertsApi.md#enablealertruleusingput) | **Put** /users-web/api/v3/alerts/{updateableAlertId}/enable | Enable alert rule
*AlertsApi* | [**GetAlertRulesForAppUsingGET**](docs/AlertsApi.md#getalertrulesforappusingget) | **Get** /users-web/api/v3/apps/{appId}/alerts | Get alert rules for an app
*AppsApi* | [**GetAppTypesUsingGET**](docs/AppsApi.md#getapptypesusingget) | **Get** /users-web/api/v3/apps/types | Get all App types supported for the account identified with apiKey
*AppsApi* | [**GetUsingGET**](docs/AppsApi.md#getusingget) | **Get** /users-web/api/v3/apps/{anyStateAppId} | Gets defails for one particular App
*AppsApi* | [**InviteAppGuestsUsingPOST**](docs/AppsApi.md#inviteappguestsusingpost) | **Post** /users-web/api/v3/apps/guests | Invite guests to an app
*AppsApi* | [**ListAppsUsersUsingGET**](docs/AppsApi.md#listappsusersusingget) | **Get** /users-web/api/v3/apps/users | Get all users of apps accessible to this account
*AppsApi* | [**ListUsingGET**](docs/AppsApi.md#listusingget) | **Get** /users-web/api/v3/apps | Get all apps accessible by account identified with apiKey
*AppsApi* | [**UpdateDescriptionUsingPUT**](docs/AppsApi.md#updatedescriptionusingput) | **Put** /users-web/api/v3/apps/{anyStateAppId}/description | Update description of the app
*AppsApi* | [**UpdateUsingPUT1**](docs/AppsApi.md#updateusingput1) | **Put** /users-web/api/v3/apps/{anyStateAppId} | Update app
*AwsSettingsControllerApi* | [**UpdateUsingPUT**](docs/AwsSettingsControllerApi.md#updateusingput) | **Put** /users-web/api/v3/apps/{appId}/aws | Update App&#39;s AWS CloudWatch settings
*BillingApi* | [**GetDetailedInvoiceUsingGET**](docs/BillingApi.md#getdetailedinvoiceusingget) | **Get** /users-web/api/v3/billing/invoice/{service}/{year}/{month} | Get invoice details
*BillingApi* | [**ListAvailablePlansUsingGET**](docs/BillingApi.md#listavailableplansusingget) | **Get** /users-web/api/v3/billing/availablePlans | Get available plans
*BillingApi* | [**UpdatePlanUsingPUT**](docs/BillingApi.md#updateplanusingput) | **Put** /users-web/api/v3/billing/info/{appId} | Update plan for an app
*LogsAppApi* | [**CreateLogseneApplication**](docs/LogsAppApi.md#createlogseneapplication) | **Post** /logsene-reports/api/v3/apps | Create Logs App
*MetricsApi* | [**ListDataSeriesUsingPOST**](docs/MetricsApi.md#listdataseriesusingpost) | **Post** /spm-reports/api/v3/apps/{appId}/metrics/data | Get metrics data points for an app
*MetricsApi* | [**ListFiltersUsingPOST1**](docs/MetricsApi.md#listfiltersusingpost1) | **Post** /spm-reports/api/v3/apps/{appId}/metrics/filters | Get metrics filters and their values for an app
*MetricsApi* | [**ListMetricsKeysUsingGET**](docs/MetricsApi.md#listmetricskeysusingget) | **Get** /spm-reports/api/v3/apps/{appId}/metrics/keys | Get metrics keys for an app
*MetricsApi* | [**ListMetricsUsingGET**](docs/MetricsApi.md#listmetricsusingget) | **Get** /spm-reports/api/v3/apps/{appId}/metrics | Get metrics info for an app
*MonitoringAppApi* | [**CreateSpmApplication1**](docs/MonitoringAppApi.md#createspmapplication1) | **Post** /spm-reports/api/v3/apps | Create Monitoring App
*ResetPasswordApi* | [**ResetPasswordUsingPOST**](docs/ResetPasswordApi.md#resetpasswordusingpost) | **Post** /users-web/api/v3/account/password/reset | Reset Password
*SavedQueriesApi* | [**DeleteSavedQueryUsingDELETE**](docs/SavedQueriesApi.md#deletesavedqueryusingdelete) | **Delete** /users-web/api/v3/savedQueries/{updateableQueryId} | Delete saved query
*SavedQueriesApi* | [**GetSavedQueriesForAppUsingGET**](docs/SavedQueriesApi.md#getsavedqueriesforappusingget) | **Get** /users-web/api/v3/apps/{appId}/savedQueries | Get saved queries for an app
*SavedQueriesApi* | [**SaveQueryUsingPOST**](docs/SavedQueriesApi.md#savequeryusingpost) | **Post** /users-web/api/v3/savedQueries | Create saved query
*SavedQueriesApi* | [**SaveQueryUsingPUT**](docs/SavedQueriesApi.md#savequeryusingput) | **Put** /users-web/api/v3/savedQueries/{updateableQueryId} | Update saved query
*SubscriptionsApi* | [**ListUsingGET1**](docs/SubscriptionsApi.md#listusingget1) | **Get** /users-web/api/v3/apps/{appId}/subscriptions | Get subscriptions for an app
*SubscriptionsApi* | [**SendReportUsingPOST**](docs/SubscriptionsApi.md#sendreportusingpost) | **Post** /users-web/api/v3/apps/{appId}/report/send | Trigger emailing of report for an app
*TagApiControllerApi* | [**GetTagNamesUsingGET1**](docs/TagApiControllerApi.md#gettagnamesusingget1) | **Get** /spm-reports/api/v3/apps/{appIds}/tagNames | Gets tag names for the given application identifiers appearing in the given time frame.
*TagApiControllerApi* | [**GetUsingGET1**](docs/TagApiControllerApi.md#getusingget1) | **Get** /spm-reports/api/v3/apps/{appIds}/tags | Gets values for specified tags for the given application identifiers appearing in the given time frame.
*TagApiControllerApi* | [**GetUsingGET2**](docs/TagApiControllerApi.md#getusingget2) | **Get** /spm-reports/api/v3/apps/{appIds}/metrics/filters | Gets values for specified tags for the given application identifiers appearing in the given time frame.


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
 - [GenericApiResponse](docs/GenericApiResponse.md)
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


