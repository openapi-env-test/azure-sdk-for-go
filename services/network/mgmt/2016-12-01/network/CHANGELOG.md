# Unreleased

## Breaking Changes

### Removed Funcs

1. *InterfacesDeleteFuture.UnmarshalJSON([]byte) error
1. InterfacesClient.Delete(context.Context, string, string) (InterfacesDeleteFuture, error)
1. InterfacesClient.DeletePreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteSender(*http.Request) (InterfacesDeleteFuture, error)
1. InterfacesClient.ListVirtualMachineScaleSetVMNetworkInterfaces(context.Context, string, string, string) (InterfaceListResultPage, error)
1. InterfacesClient.ListVirtualMachineScaleSetVMNetworkInterfacesComplete(context.Context, string, string, string) (InterfaceListResultIterator, error)
1. InterfacesClient.ListVirtualMachineScaleSetVMNetworkInterfacesPreparer(context.Context, string, string, string) (*http.Request, error)
1. InterfacesClient.ListVirtualMachineScaleSetVMNetworkInterfacesResponder(*http.Response) (InterfaceListResult, error)
1. InterfacesClient.ListVirtualMachineScaleSetVMNetworkInterfacesSender(*http.Request) (*http.Response, error)

### Struct Changes

#### Removed Structs

1. InterfacesDeleteFuture

## Additive Changes

### New Funcs

1. *InterfacesDeleteABCDFuture.UnmarshalJSON([]byte) error
1. InterfacesClient.DeleteABCD(context.Context, string, string) (InterfacesDeleteABCDFuture, error)
1. InterfacesClient.DeleteABCDPreparer(context.Context, string, string) (*http.Request, error)
1. InterfacesClient.DeleteABCDResponder(*http.Response) (autorest.Response, error)
1. InterfacesClient.DeleteABCDSender(*http.Request) (InterfacesDeleteABCDFuture, error)
1. InterfacesClient.ListVirtualMachineScaleSetVMNetworkInterfacesABC(context.Context, string, string, string) (InterfaceListResultPage, error)
1. InterfacesClient.ListVirtualMachineScaleSetVMNetworkInterfacesABCComplete(context.Context, string, string, string) (InterfaceListResultIterator, error)
1. InterfacesClient.ListVirtualMachineScaleSetVMNetworkInterfacesABCPreparer(context.Context, string, string, string) (*http.Request, error)
1. InterfacesClient.ListVirtualMachineScaleSetVMNetworkInterfacesABCResponder(*http.Response) (InterfaceListResult, error)
1. InterfacesClient.ListVirtualMachineScaleSetVMNetworkInterfacesABCSender(*http.Request) (*http.Response, error)

### Struct Changes

#### New Structs

1. InterfacesDeleteABCDFuture
