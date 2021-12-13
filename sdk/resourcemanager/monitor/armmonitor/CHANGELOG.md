# Release History

## 0.2.0 (2021-12-13)
### Breaking Changes

- Field `SmsReceivers` of struct `ActionGroup` has been removed

### New Content

- New function `*ActionGroupsPostTestNotificationsPoller.Poll(context.Context) (*http.Response, error)`
- New function `ActionGroupsPostTestNotificationsPollerResponse.PollUntilDone(context.Context, time.Duration) (ActionGroupsPostTestNotificationsResponse, error)`
- New function `NotificationRequestBody.MarshalJSON() ([]byte, error)`
- New function `TestNotificationDetailsResponse.MarshalJSON() ([]byte, error)`
- New function `*ActionGroupsPostTestNotificationsPoller.FinalResponse(context.Context) (ActionGroupsPostTestNotificationsResponse, error)`
- New function `*ActionGroupsClient.BeginPostTestNotifications(context.Context, NotificationRequestBody, *ActionGroupsBeginPostTestNotificationsOptions) (ActionGroupsPostTestNotificationsPollerResponse, error)`
- New function `*ActionGroupsPostTestNotificationsPoller.Done() bool`
- New function `*ActionGroupsClient.GetTestNotifications(context.Context, string, *ActionGroupsGetTestNotificationsOptions) (ActionGroupsGetTestNotificationsResponse, error)`
- New function `*ActionGroupsPostTestNotificationsPoller.ResumeToken() (string, error)`
- New function `*ActionGroupsPostTestNotificationsPollerResponse.Resume(context.Context, *ActionGroupsClient, string) error`
- New struct `ActionDetail`
- New struct `ActionGroupsBeginPostTestNotificationsOptions`
- New struct `ActionGroupsGetTestNotificationsOptions`
- New struct `ActionGroupsGetTestNotificationsResponse`
- New struct `ActionGroupsGetTestNotificationsResult`
- New struct `ActionGroupsPostTestNotificationsPoller`
- New struct `ActionGroupsPostTestNotificationsPollerResponse`
- New struct `ActionGroupsPostTestNotificationsResponse`
- New struct `ActionGroupsPostTestNotificationsResult`
- New struct `Context`
- New struct `EventHubReceiver`
- New struct `NotificationRequestBody`
- New struct `TestNotificationDetailsResponse`
- New struct `TestNotificationResponse`
- New field `EventHubReceivers` in struct `ActionGroup`

Total 1 breaking change(s), 39 additive change(s).


## 0.1.0 (2021-10-08)
- To better align with the Azure SDK guidelines (https://azure.github.io/azure-sdk/general_introduction.html), we have decided to change the module path to "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor". Therefore, we are deprecating the old module path (which is "github.com/Azure/azure-sdk-for-go/sdk/monitor/armmonitor") to avoid confusion.