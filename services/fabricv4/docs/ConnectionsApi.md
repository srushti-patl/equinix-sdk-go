# \ConnectionsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnection**](ConnectionsApi.md#CreateConnection) | **Post** /fabric/v4/connections | Create Connection
[**CreateConnectionAction**](ConnectionsApi.md#CreateConnectionAction) | **Post** /fabric/v4/connections/{connectionId}/actions | Connection Actions
[**CreateConnectionsInBulk**](ConnectionsApi.md#CreateConnectionsInBulk) | **Post** /fabric/v4/connections/bulk | Bulk Connections
[**DeleteConnectionByUuid**](ConnectionsApi.md#DeleteConnectionByUuid) | **Delete** /fabric/v4/connections/{connectionId} | Delete by ID
[**GetConnectionByUuid**](ConnectionsApi.md#GetConnectionByUuid) | **Get** /fabric/v4/connections/{connectionId} | Get Connection by ID
[**PutConnectionByUuid**](ConnectionsApi.md#PutConnectionByUuid) | **Put** /fabric/v4/connections/{connectionId} | Replace by ID
[**SearchConnections**](ConnectionsApi.md#SearchConnections) | **Post** /fabric/v4/connections/search | Search connections
[**UpdateConnectionByUuid**](ConnectionsApi.md#UpdateConnectionByUuid) | **Patch** /fabric/v4/connections/{connectionId} | Update by ID
[**UpdateConnectionsByUuids**](ConnectionsApi.md#UpdateConnectionsByUuids) | **Patch** /fabric/v4/connections/bulk | Update Connections
[**ValidateConnections**](ConnectionsApi.md#ValidateConnections) | **Post** /fabric/v4/connections/validate | Validate Connection



## CreateConnection

> Connection CreateConnection(ctx).ConnectionPostRequest(connectionPostRequest).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).XSOURCE(xSOURCE).AccountSubCustomerUcmId(accountSubCustomerUcmId).DryRun(dryRun).Execute()

Create Connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	connectionPostRequest := *openapiclient.NewConnectionPostRequest(openapiclient.ConnectionType("EVPL_VC"), "Name_example", []openapiclient.SimplifiedNotification{*openapiclient.NewSimplifiedNotification(openapiclient.SimplifiedNotification_type("NOTIFICATION"), []string{"Emails_example"})}, int32(123), *openapiclient.NewConnectionSide(), *openapiclient.NewConnectionSide()) // ConnectionPostRequest | 
	xCORRELATIONID := "xCORRELATIONID_example" // string | Correlation identifier (optional)
	xAUTHUSERNAME := "xAUTHUSERNAME_example" // string | User name (optional)
	xSOURCE := "xSOURCE_example" // string | source (optional)
	accountSubCustomerUcmId := "accountSubCustomerUcmId_example" // string | subCustomerUcmId (optional)
	dryRun := true // bool | option to verify that API calls will succeed (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsApi.CreateConnection(context.Background()).ConnectionPostRequest(connectionPostRequest).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).XSOURCE(xSOURCE).AccountSubCustomerUcmId(accountSubCustomerUcmId).DryRun(dryRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.CreateConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnection`: Connection
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.CreateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionPostRequest** | [**ConnectionPostRequest**](ConnectionPostRequest.md) |  | 
 **xCORRELATIONID** | **string** | Correlation identifier | 
 **xAUTHUSERNAME** | **string** | User name | 
 **xSOURCE** | **string** | source | 
 **accountSubCustomerUcmId** | **string** | subCustomerUcmId | 
 **dryRun** | **bool** | option to verify that API calls will succeed | [default to false]

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnectionAction

> ConnectionAction CreateConnectionAction(ctx, connectionId).ConnectionActionRequest(connectionActionRequest).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).XSOURCE(xSOURCE).Execute()

Connection Actions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	connectionId := "connectionId_example" // string | Connection Id
	connectionActionRequest := *openapiclient.NewConnectionActionRequest(openapiclient.Actions("CONNECTION_CREATION_ACCEPTANCE")) // ConnectionActionRequest | 
	xCORRELATIONID := "xCORRELATIONID_example" // string | Correlation identifier (optional)
	xAUTHUSERNAME := "xAUTHUSERNAME_example" // string | User name (optional)
	xSOURCE := "xSOURCE_example" // string | source (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsApi.CreateConnectionAction(context.Background(), connectionId).ConnectionActionRequest(connectionActionRequest).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).XSOURCE(xSOURCE).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.CreateConnectionAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectionAction`: ConnectionAction
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.CreateConnectionAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionActionRequest** | [**ConnectionActionRequest**](ConnectionActionRequest.md) |  | 
 **xCORRELATIONID** | **string** | Correlation identifier | 
 **xAUTHUSERNAME** | **string** | User name | 
 **xSOURCE** | **string** | source | 

### Return type

[**ConnectionAction**](ConnectionAction.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnectionsInBulk

> ConnectionBulk CreateConnectionsInBulk(ctx).ConnectionBulkPostRequest(connectionBulkPostRequest).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).XSOURCE(xSOURCE).AccountSubCustomerUcmId(accountSubCustomerUcmId).Execute()

Bulk Connections



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	connectionBulkPostRequest := *openapiclient.NewConnectionBulkPostRequest() // ConnectionBulkPostRequest | 
	xCORRELATIONID := "xCORRELATIONID_example" // string | Correlation identifier (optional)
	xAUTHUSERNAME := "xAUTHUSERNAME_example" // string | User name (optional)
	xSOURCE := "xSOURCE_example" // string | source (optional)
	accountSubCustomerUcmId := "accountSubCustomerUcmId_example" // string | subCustomerUcmId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsApi.CreateConnectionsInBulk(context.Background()).ConnectionBulkPostRequest(connectionBulkPostRequest).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).XSOURCE(xSOURCE).AccountSubCustomerUcmId(accountSubCustomerUcmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.CreateConnectionsInBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectionsInBulk`: ConnectionBulk
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.CreateConnectionsInBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionsInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionBulkPostRequest** | [**ConnectionBulkPostRequest**](ConnectionBulkPostRequest.md) |  | 
 **xCORRELATIONID** | **string** | Correlation identifier | 
 **xAUTHUSERNAME** | **string** | User name | 
 **xSOURCE** | **string** | source | 
 **accountSubCustomerUcmId** | **string** | subCustomerUcmId | 

### Return type

[**ConnectionBulk**](ConnectionBulk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectionByUuid

> Connection DeleteConnectionByUuid(ctx, connectionId).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).XSOURCE(xSOURCE).AccountSubCustomerUcmId(accountSubCustomerUcmId).Execute()

Delete by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	connectionId := "connectionId_example" // string | Connection UUID
	xCORRELATIONID := "12345-6789-10123" // string | Correlation identifier (optional)
	xAUTHUSERNAME := "alice" // string | User name (optional)
	xSOURCE := "xSOURCE_example" // string | source (optional)
	accountSubCustomerUcmId := "accountSubCustomerUcmId_example" // string | subCustomerUcmId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsApi.DeleteConnectionByUuid(context.Background(), connectionId).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).XSOURCE(xSOURCE).AccountSubCustomerUcmId(accountSubCustomerUcmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.DeleteConnectionByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteConnectionByUuid`: Connection
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.DeleteConnectionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **string** | Correlation identifier | 
 **xAUTHUSERNAME** | **string** | User name | 
 **xSOURCE** | **string** | source | 
 **accountSubCustomerUcmId** | **string** | subCustomerUcmId | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionByUuid

> Connection GetConnectionByUuid(ctx, connectionId).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).Direction(direction).Execute()

Get Connection by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	connectionId := "connectionId_example" // string | Connection Id
	xCORRELATIONID := "xCORRELATIONID_example" // string | Correlation identifier (optional)
	xAUTHUSERNAME := "xAUTHUSERNAME_example" // string | User name (optional)
	direction := openapiclient.ConnectionDirection("INTERNAL") // ConnectionDirection | Connection Direction (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsApi.GetConnectionByUuid(context.Background(), connectionId).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.GetConnectionByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionByUuid`: Connection
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.GetConnectionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **string** | Correlation identifier | 
 **xAUTHUSERNAME** | **string** | User name | 
 **direction** | [**ConnectionDirection**](ConnectionDirection.md) | Connection Direction | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutConnectionByUuid

> Connection PutConnectionByUuid(ctx, connectionId).ConnectionPutRequest(connectionPutRequest).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).AccountSubCustomerUcmId(accountSubCustomerUcmId).XSOURCE(xSOURCE).Execute()

Replace by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	connectionId := "connectionId_example" // string | Connection UUID
	connectionPutRequest := *openapiclient.NewConnectionPutRequest() // ConnectionPutRequest | 
	xCORRELATIONID := "xCORRELATIONID_example" // string | Correlation identifier (optional)
	xAUTHUSERNAME := "xAUTHUSERNAME_example" // string | User name (optional)
	accountSubCustomerUcmId := "accountSubCustomerUcmId_example" // string | subCustomerUcmId (optional)
	xSOURCE := "xSOURCE_example" // string | source (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsApi.PutConnectionByUuid(context.Background(), connectionId).ConnectionPutRequest(connectionPutRequest).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).AccountSubCustomerUcmId(accountSubCustomerUcmId).XSOURCE(xSOURCE).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.PutConnectionByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutConnectionByUuid`: Connection
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.PutConnectionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutConnectionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionPutRequest** | [**ConnectionPutRequest**](ConnectionPutRequest.md) |  | 
 **xCORRELATIONID** | **string** | Correlation identifier | 
 **xAUTHUSERNAME** | **string** | User name | 
 **accountSubCustomerUcmId** | **string** | subCustomerUcmId | 
 **xSOURCE** | **string** | source | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchConnections

> ConnectionSearchResponse SearchConnections(ctx).SearchRequest(searchRequest).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).Execute()

Search connections



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	searchRequest := *openapiclient.NewSearchRequest() // SearchRequest | 
	xCORRELATIONID := "xCORRELATIONID_example" // string | Correlation identifier (optional)
	xAUTHUSERNAME := "xAUTHUSERNAME_example" // string | User name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsApi.SearchConnections(context.Background()).SearchRequest(searchRequest).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.SearchConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchConnections`: ConnectionSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.SearchConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchRequest** | [**SearchRequest**](SearchRequest.md) |  | 
 **xCORRELATIONID** | **string** | Correlation identifier | 
 **xAUTHUSERNAME** | **string** | User name | 

### Return type

[**ConnectionSearchResponse**](ConnectionSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectionByUuid

> Connection UpdateConnectionByUuid(ctx, connectionId).ConnectionChangeOperation(connectionChangeOperation).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).XSOURCE(xSOURCE).AccountSubCustomerUcmId(accountSubCustomerUcmId).Execute()

Update by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	connectionId := "connectionId_example" // string | Connection Id
	connectionChangeOperation := []openapiclient.ConnectionChangeOperation{*openapiclient.NewConnectionChangeOperation("add", "/ipv6", interface{}(123))} // []ConnectionChangeOperation | 
	xCORRELATIONID := "xCORRELATIONID_example" // string | Correlation identifier (optional)
	xAUTHUSERNAME := "xAUTHUSERNAME_example" // string | User name (optional)
	xSOURCE := "xSOURCE_example" // string | source (optional)
	accountSubCustomerUcmId := "accountSubCustomerUcmId_example" // string | subCustomerUcmId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsApi.UpdateConnectionByUuid(context.Background(), connectionId).ConnectionChangeOperation(connectionChangeOperation).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).XSOURCE(xSOURCE).AccountSubCustomerUcmId(accountSubCustomerUcmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.UpdateConnectionByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConnectionByUuid`: Connection
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.UpdateConnectionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Connection Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionChangeOperation** | [**[]ConnectionChangeOperation**](ConnectionChangeOperation.md) |  | 
 **xCORRELATIONID** | **string** | Correlation identifier | 
 **xAUTHUSERNAME** | **string** | User name | 
 **xSOURCE** | **string** | source | 
 **accountSubCustomerUcmId** | **string** | subCustomerUcmId | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectionsByUuids

> Connection UpdateConnectionsByUuids(ctx).Uuid(uuid).ConnectionChangeOperation(connectionChangeOperation).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).XSOURCE(xSOURCE).Execute()

Update Connections



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	uuid := "uuid_example" // string | Connection UUID
	connectionChangeOperation := []openapiclient.ConnectionChangeOperation{*openapiclient.NewConnectionChangeOperation("add", "/ipv6", interface{}(123))} // []ConnectionChangeOperation | 
	xCORRELATIONID := "xCORRELATIONID_example" // string | Correlation identifier (optional)
	xAUTHUSERNAME := "xAUTHUSERNAME_example" // string | User name (optional)
	xSOURCE := "xSOURCE_example" // string | source (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsApi.UpdateConnectionsByUuids(context.Background()).Uuid(uuid).ConnectionChangeOperation(connectionChangeOperation).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).XSOURCE(xSOURCE).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.UpdateConnectionsByUuids``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConnectionsByUuids`: Connection
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.UpdateConnectionsByUuids`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionsByUuidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** | Connection UUID | 
 **connectionChangeOperation** | [**[]ConnectionChangeOperation**](ConnectionChangeOperation.md) |  | 
 **xCORRELATIONID** | **string** | Correlation identifier | 
 **xAUTHUSERNAME** | **string** | User name | 
 **xSOURCE** | **string** | source | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateConnections

> ConnectionResponse ValidateConnections(ctx).ValidateRequest(validateRequest).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).Execute()

Validate Connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	validateRequest := *openapiclient.NewValidateRequest() // ValidateRequest | 
	xCORRELATIONID := "12345-6789-10123" // string | Correlation identifier (optional)
	xAUTHUSERNAME := "alice" // string | User name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsApi.ValidateConnections(context.Background()).ValidateRequest(validateRequest).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.ValidateConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateConnections`: ConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.ValidateConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateRequest** | [**ValidateRequest**](ValidateRequest.md) |  | 
 **xCORRELATIONID** | **string** | Correlation identifier | 
 **xAUTHUSERNAME** | **string** | User name | 

### Return type

[**ConnectionResponse**](ConnectionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

