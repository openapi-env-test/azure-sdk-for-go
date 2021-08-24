# Unreleased

## Breaking Changes

### Removed Funcs

1. Client.ListUpgradeNotifications(context.Context, string, string, float64) (NotificationListResponse, error)
1. Client.ListUpgradeNotificationsPreparer(context.Context, string, string, float64) (*http.Request, error)
1. Client.ListUpgradeNotificationsResponder(*http.Response) (NotificationListResponse, error)
1. Client.ListUpgradeNotificationsSender(*http.Request) (*http.Response, error)

### Struct Changes

#### Removed Struct Fields

1. NotificationListResponse.autorest.Response
