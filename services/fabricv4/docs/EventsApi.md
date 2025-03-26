# \EventsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /fabric/v4/events/{eventId} | Get Event
[**GetEventsByAsset**](EventsApi.md#GetEventsByAsset) | **Get** /fabric/v4/events | Get Events of Asset



## GetEvent

> GetEvent(ctx, eventId).Type_(type_).ConnectionId(connectionId).PortId(portId).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).Execute()

Get Event



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
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Event UUID
	type_ := openapiclient.getEvent_type_parameter("TRANSACTION_STATUS_AUDIT") // GetEventTypeParameter | type of events
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Connection UUID (optional)
	portId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID (optional)
	xCORRELATIONID := "xCORRELATIONID_example" // string | Correlation identifier (optional)
	xAUTHUSERNAME := "xAUTHUSERNAME_example" // string | User name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsApi.GetEvent(context.Background(), eventId).Type_(type_).ConnectionId(connectionId).PortId(portId).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | Event UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**GetEventTypeParameter**](GetEventTypeParameter.md) | type of events | 
 **connectionId** | **string** | Connection UUID | 
 **portId** | **string** | Port UUID | 
 **xCORRELATIONID** | **string** | Correlation identifier | 
 **xAUTHUSERNAME** | **string** | User name | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsByAsset

> InlineResponse200 GetEventsByAsset(ctx).Type_(type_).ConnectionId(connectionId).PortId(portId).RouterId(routerId).RoutingProtocolId(routingProtocolId).StartDateTime(startDateTime).EndDateTime(endDateTime).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).Execute()

Get Events of Asset



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	type_ := openapiclient.getEvent_type_parameter("TRANSACTION_STATUS_AUDIT") // GetEventTypeParameter | type of events
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Connection UUID (optional)
	portId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID (optional)
	routerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cloud Router UUID (optional)
	routingProtocolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Routing Protocol UUID (optional)
	startDateTime := time.Now() // time.Time | Start date and time (optional)
	endDateTime := time.Now() // time.Time | End date and time (optional)
	xCORRELATIONID := "xCORRELATIONID_example" // string | Correlation identifier (optional)
	xAUTHUSERNAME := "xAUTHUSERNAME_example" // string | User name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsApi.GetEventsByAsset(context.Background()).Type_(type_).ConnectionId(connectionId).PortId(portId).RouterId(routerId).RoutingProtocolId(routingProtocolId).StartDateTime(startDateTime).EndDateTime(endDateTime).XCORRELATIONID(xCORRELATIONID).XAUTHUSERNAME(xAUTHUSERNAME).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEventsByAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventsByAsset`: InlineResponse200
	fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEventsByAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsByAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**GetEventTypeParameter**](GetEventTypeParameter.md) | type of events | 
 **connectionId** | **string** | Connection UUID | 
 **portId** | **string** | Port UUID | 
 **routerId** | **string** | Cloud Router UUID | 
 **routingProtocolId** | **string** | Routing Protocol UUID | 
 **startDateTime** | **time.Time** | Start date and time | 
 **endDateTime** | **time.Time** | End date and time | 
 **xCORRELATIONID** | **string** | Correlation identifier | 
 **xAUTHUSERNAME** | **string** | User name | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

