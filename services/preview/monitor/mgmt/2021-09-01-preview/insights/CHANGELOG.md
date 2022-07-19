# Unreleased

## Breaking Changes

### Removed Funcs

1. *ActionGroupsPostTestNotificationsFuture.UnmarshalJSON([]byte) error
1. ActionGroupsClient.GetTestNotifications(context.Context, string) (TestNotificationDetailsResponse, error)
1. ActionGroupsClient.GetTestNotificationsPreparer(context.Context, string) (*http.Request, error)
1. ActionGroupsClient.GetTestNotificationsResponder(*http.Response) (TestNotificationDetailsResponse, error)
1. ActionGroupsClient.GetTestNotificationsSender(*http.Request) (*http.Response, error)
1. ActionGroupsClient.PostTestNotifications(context.Context, NotificationRequestBody) (ActionGroupsPostTestNotificationsFuture, error)
1. ActionGroupsClient.PostTestNotificationsPreparer(context.Context, NotificationRequestBody) (*http.Request, error)
1. ActionGroupsClient.PostTestNotificationsResponder(*http.Response) (TestNotificationResponse, error)
1. ActionGroupsClient.PostTestNotificationsSender(*http.Request) (ActionGroupsPostTestNotificationsFuture, error)

### Struct Changes

#### Removed Structs

1. ActionDetail
1. ActionGroupsPostTestNotificationsFuture
1. Context
1. TestNotificationDetailsResponse
1. TestNotificationResponse

#### Removed Struct Fields

1. ActionGroupResource.Identity
1. ActivityLogAlertResource.Identity
1. AzureResource.Identity
