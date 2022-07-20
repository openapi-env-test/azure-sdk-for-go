# Unreleased

## Breaking Changes

### Signature Changes

#### Funcs

1. ProtectedItemsClient.CreateOrUpdate
	- Returns
		- From: ProtectedItemResource, error
		- To: ProtectedItemsCreateOrUpdateFuture, error
1. ProtectedItemsClient.CreateOrUpdateSender
	- Returns
		- From: *http.Response, error
		- To: ProtectedItemsCreateOrUpdateFuture, error
1. ProtectedItemsClient.Delete
	- Returns
		- From: autorest.Response, error
		- To: ProtectedItemsDeleteFuture, error
1. ProtectedItemsClient.DeleteSender
	- Returns
		- From: *http.Response, error
		- To: ProtectedItemsDeleteFuture, error

## Additive Changes

### New Funcs

1. *ProtectedItemsCreateOrUpdateFuture.UnmarshalJSON([]byte) error
1. *ProtectedItemsDeleteFuture.UnmarshalJSON([]byte) error

### Struct Changes

#### New Structs

1. ProtectedItemsCreateOrUpdateFuture
1. ProtectedItemsDeleteFuture
