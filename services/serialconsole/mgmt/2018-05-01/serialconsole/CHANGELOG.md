# Unreleased Content

## Additive Changes

### New Constants

1. SerialPortState.Disabled
1. SerialPortState.Enabled

### New Funcs

1. *SerialPort.UnmarshalJSON([]byte) error
1. NewSerialPortsClient(string) SerialPortsClient
1. NewSerialPortsClientWithBaseURI(string, string) SerialPortsClient
1. PossibleSerialPortStateValues() []SerialPortState
1. SerialPort.MarshalJSON() ([]byte, error)
1. SerialPortsClient.Connect(context.Context, string, string, string, string, string) (SerialPortConnectResult, error)
1. SerialPortsClient.ConnectPreparer(context.Context, string, string, string, string, string) (*http.Request, error)
1. SerialPortsClient.ConnectResponder(*http.Response) (SerialPortConnectResult, error)
1. SerialPortsClient.ConnectSender(*http.Request) (*http.Response, error)
1. SerialPortsClient.Create(context.Context, string, string, string, string, string, SerialPort) (SerialPort, error)
1. SerialPortsClient.CreatePreparer(context.Context, string, string, string, string, string, SerialPort) (*http.Request, error)
1. SerialPortsClient.CreateResponder(*http.Response) (SerialPort, error)
1. SerialPortsClient.CreateSender(*http.Request) (*http.Response, error)
1. SerialPortsClient.Delete(context.Context, string, string, string, string, string) (autorest.Response, error)
1. SerialPortsClient.DeletePreparer(context.Context, string, string, string, string, string) (*http.Request, error)
1. SerialPortsClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. SerialPortsClient.DeleteSender(*http.Request) (*http.Response, error)
1. SerialPortsClient.Get(context.Context, string, string, string, string, string) (SerialPort, error)
1. SerialPortsClient.GetPreparer(context.Context, string, string, string, string, string) (*http.Request, error)
1. SerialPortsClient.GetResponder(*http.Response) (SerialPort, error)
1. SerialPortsClient.GetSender(*http.Request) (*http.Response, error)
1. SerialPortsClient.List(context.Context, string, string, string, string) (SerialPortListResult, error)
1. SerialPortsClient.ListBySubscriptions(context.Context) (SerialPortListResult, error)
1. SerialPortsClient.ListBySubscriptionsPreparer(context.Context) (*http.Request, error)
1. SerialPortsClient.ListBySubscriptionsResponder(*http.Response) (SerialPortListResult, error)
1. SerialPortsClient.ListBySubscriptionsSender(*http.Request) (*http.Response, error)
1. SerialPortsClient.ListPreparer(context.Context, string, string, string, string) (*http.Request, error)
1. SerialPortsClient.ListResponder(*http.Response) (SerialPortListResult, error)
1. SerialPortsClient.ListSender(*http.Request) (*http.Response, error)

### Struct Changes

#### New Structs

1. CloudError
1. CloudErrorBody
1. ProxyResource
1. Resource
1. SerialPort
1. SerialPortConnectResult
1. SerialPortListResult
1. SerialPortProperties
1. SerialPortsClient
