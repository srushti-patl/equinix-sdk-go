# \MetrosApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIBXsByMetroCode**](MetrosApi.md#GetIBXsByMetroCode) | **Get** /fabric/v4/metros/{metroCode}/ibxs | Get IBXs by Metro
[**GetMetroByCode**](MetrosApi.md#GetMetroByCode) | **Get** /fabric/v4/metros/{metroCode} | Get Metro by Code
[**GetMetros**](MetrosApi.md#GetMetros) | **Get** /fabric/v4/metros | Get all Metros



## GetIBXsByMetroCode

> IBXResponse GetIBXsByMetroCode(ctx, metroCode).IsTimeServiceEnabled(isTimeServiceEnabled).CorrelationId(correlationId).Offset(offset).Limit(limit).Execute()

Get IBXs by Metro



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
	metroCode := "metroCode_example" // string | Metro Code
	isTimeServiceEnabled := true // bool | Precision Time Service eligibility flag. (optional) (default to true)
	correlationId := "correlationId_example" // string | Correlation identifier (optional)
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetrosApi.GetIBXsByMetroCode(context.Background(), metroCode).IsTimeServiceEnabled(isTimeServiceEnabled).CorrelationId(correlationId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetrosApi.GetIBXsByMetroCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIBXsByMetroCode`: IBXResponse
	fmt.Fprintf(os.Stdout, "Response from `MetrosApi.GetIBXsByMetroCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metroCode** | **string** | Metro Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIBXsByMetroCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isTimeServiceEnabled** | **bool** | Precision Time Service eligibility flag. | [default to true]
 **correlationId** | **string** | Correlation identifier | 
 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**IBXResponse**](IBXResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetroByCode

> Metro GetMetroByCode(ctx, metroCode).CorrelationId(correlationId).Execute()

Get Metro by Code



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
	metroCode := "metroCode_example" // string | Metro Code
	correlationId := "correlationId_example" // string | Correlation identifier (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetrosApi.GetMetroByCode(context.Background(), metroCode).CorrelationId(correlationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetrosApi.GetMetroByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetroByCode`: Metro
	fmt.Fprintf(os.Stdout, "Response from `MetrosApi.GetMetroByCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metroCode** | **string** | Metro Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetroByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **correlationId** | **string** | Correlation identifier | 

### Return type

[**Metro**](Metro.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetros

> MetroResponse GetMetros(ctx).Presence(presence).Offset(offset).Limit(limit).CorrelationId(correlationId).Execute()

Get all Metros



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
	presence := openapiclient.Presence("MY_PORTS") // Presence | User On Boarded Metros based on Fabric resource availability (optional)
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)
	correlationId := "correlationId_example" // string | Correlation identifier (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetrosApi.GetMetros(context.Background()).Presence(presence).Offset(offset).Limit(limit).CorrelationId(correlationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetrosApi.GetMetros``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetros`: MetroResponse
	fmt.Fprintf(os.Stdout, "Response from `MetrosApi.GetMetros`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetrosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **presence** | [**Presence**](Presence.md) | User On Boarded Metros based on Fabric resource availability | 
 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 
 **correlationId** | **string** | Correlation identifier | 

### Return type

[**MetroResponse**](MetroResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

